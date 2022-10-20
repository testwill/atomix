// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"bytes"
	"github.com/atomix/runtime/sdk/pkg/errors"
	"github.com/atomix/runtime/sdk/pkg/protocol/statemachine"
	"sync"
)

const Service = "atomix.runtime.indexedmap.v1.IndexedMap"

func RegisterStateMachine(registry *statemachine.PrimitiveTypeRegistry) {
	statemachine.RegisterPrimitiveType[*IndexedMapInput, *IndexedMapOutput](registry)(PrimitiveType)
}

var PrimitiveType = statemachine.NewPrimitiveType[*IndexedMapInput, *IndexedMapOutput](Service, indexedMapCodec,
	func(context statemachine.PrimitiveContext[*IndexedMapInput, *IndexedMapOutput]) statemachine.Executor[*IndexedMapInput, *IndexedMapOutput] {
		return newExecutor(NewIndexedMapStateMachine(context))
	})

type IndexedMapContext interface {
	statemachine.PrimitiveContext[*IndexedMapInput, *IndexedMapOutput]
	Events() statemachine.Proposals[*EventsInput, *EventsOutput]
}

func newContext(context statemachine.PrimitiveContext[*IndexedMapInput, *IndexedMapOutput]) IndexedMapContext {
	return &indexedMapContext{
		PrimitiveContext: context,
		events: statemachine.NewProposals[*IndexedMapInput, *IndexedMapOutput, *EventsInput, *EventsOutput](context).
			Decoder(func(input *IndexedMapInput) (*EventsInput, bool) {
				if events, ok := input.Input.(*IndexedMapInput_Events); ok {
					return events.Events, true
				}
				return nil, false
			}).
			Encoder(func(output *EventsOutput) *IndexedMapOutput {
				return &IndexedMapOutput{
					Output: &IndexedMapOutput_Events{
						Events: output,
					},
				}
			}).
			Build(),
	}
}

type indexedMapContext struct {
	statemachine.PrimitiveContext[*IndexedMapInput, *IndexedMapOutput]
	events statemachine.Proposals[*EventsInput, *EventsOutput]
}

func (c *indexedMapContext) Events() statemachine.Proposals[*EventsInput, *EventsOutput] {
	return c.events
}

type IndexedMapStateMachine interface {
	statemachine.Context[*IndexedMapInput, *IndexedMapOutput]
	statemachine.Recoverable
	Append(statemachine.Proposal[*AppendInput, *AppendOutput])
	Update(statemachine.Proposal[*UpdateInput, *UpdateOutput])
	Remove(statemachine.Proposal[*RemoveInput, *RemoveOutput])
	Clear(statemachine.Proposal[*ClearInput, *ClearOutput])
	Events(statemachine.Proposal[*EventsInput, *EventsOutput])
	Size(statemachine.Query[*SizeInput, *SizeOutput])
	Get(statemachine.Query[*GetInput, *GetOutput])
	FirstEntry(statemachine.Query[*FirstEntryInput, *FirstEntryOutput])
	LastEntry(statemachine.Query[*LastEntryInput, *LastEntryOutput])
	NextEntry(statemachine.Query[*NextEntryInput, *NextEntryOutput])
	PrevEntry(statemachine.Query[*PrevEntryInput, *PrevEntryOutput])
	Entries(statemachine.Query[*EntriesInput, *EntriesOutput])
}

func NewIndexedMapStateMachine(context statemachine.PrimitiveContext[*IndexedMapInput, *IndexedMapOutput]) IndexedMapStateMachine {
	return &indexedMapStateMachine{
		IndexedMapContext: newContext(context),
	}
}

// LinkedMapEntryValue is a doubly linked MapEntryValue
type LinkedMapEntryValue struct {
	*IndexedMapEntry
	Prev *LinkedMapEntryValue
	Next *LinkedMapEntryValue
}

type indexedMapStateMachine struct {
	IndexedMapContext
	lastIndex  uint64
	keys       map[string]*LinkedMapEntryValue
	indexes    map[uint64]*LinkedMapEntryValue
	firstEntry *LinkedMapEntryValue
	lastEntry  *LinkedMapEntryValue
	streams    map[statemachine.ProposalID]*IndexedMapListener
	timers     map[string]statemachine.CancelFunc
	watchers   map[statemachine.QueryID]statemachine.Query[*EntriesInput, *EntriesOutput]
	mu         sync.RWMutex
}

func (s *indexedMapStateMachine) reset() {
	if s.timers != nil {
		for _, cancel := range s.timers {
			cancel()
		}
	}
	s.timers = make(map[string]statemachine.CancelFunc)
	if s.watchers != nil {
		for _, watcher := range s.watchers {
			watcher.Cancel()
		}
	}
	s.watchers = make(map[statemachine.QueryID]statemachine.Query[*EntriesInput, *EntriesOutput])
	s.streams = make(map[statemachine.ProposalID]*IndexedMapListener)
	s.keys = make(map[string]*LinkedMapEntryValue)
	s.indexes = make(map[uint64]*LinkedMapEntryValue)
	s.firstEntry = nil
	s.lastEntry = nil
}

func (s *indexedMapStateMachine) Snapshot(writer *statemachine.SnapshotWriter) error {
	s.Log().Infow("Persisting IndexedMap to snapshot")
	if err := s.snapshotEntries(writer); err != nil {
		return err
	}
	if err := s.snapshotStreams(writer); err != nil {
		return err
	}
	return nil
}

func (s *indexedMapStateMachine) snapshotEntries(writer *statemachine.SnapshotWriter) error {
	if err := writer.WriteVarInt(len(s.keys)); err != nil {
		return err
	}
	entry := s.firstEntry
	for entry != nil {
		if err := writer.WriteMessage(s.firstEntry); err != nil {
			return err
		}
		entry = entry.Next
	}
	if err := writer.WriteVarUint64(s.lastIndex); err != nil {
		return err
	}
	return nil
}

func (s *indexedMapStateMachine) snapshotStreams(writer *statemachine.SnapshotWriter) error {
	if err := writer.WriteVarInt(len(s.streams)); err != nil {
		return err
	}
	for proposalID, listener := range s.streams {
		if err := writer.WriteVarUint64(uint64(proposalID)); err != nil {
			return err
		}
		if err := writer.WriteMessage(listener); err != nil {
			return err
		}
	}
	return nil
}

func (s *indexedMapStateMachine) Recover(reader *statemachine.SnapshotReader) error {
	s.Log().Infow("Recovering IndexedMap from snapshot")
	s.reset()
	if err := s.recoverEntries(reader); err != nil {
		return err
	}
	if err := s.recoverStreams(reader); err != nil {
		return err
	}
	return nil
}

func (s *indexedMapStateMachine) recoverEntries(reader *statemachine.SnapshotReader) error {
	n, err := reader.ReadVarInt()
	if err != nil {
		return err
	}

	var prevEntry *LinkedMapEntryValue
	for i := 0; i < n; i++ {
		entry := &LinkedMapEntryValue{}
		if err := reader.ReadMessage(entry); err != nil {
			return err
		}
		s.keys[entry.Key] = entry
		s.indexes[entry.Index] = entry
		if s.firstEntry == nil {
			s.firstEntry = entry
		}
		if prevEntry != nil {
			prevEntry.Next = entry
			entry.Prev = prevEntry
		}
		prevEntry = entry
		s.lastEntry = entry
		s.scheduleTTL(entry)
	}

	i, err := reader.ReadVarUint64()
	if err != nil {
		return err
	}
	s.lastIndex = i
	return nil
}

func (s *indexedMapStateMachine) recoverStreams(reader *statemachine.SnapshotReader) error {
	n, err := reader.ReadVarInt()
	if err != nil {
		return err
	}
	for i := 0; i < n; i++ {
		proposalID, err := reader.ReadVarUint64()
		if err != nil {
			return err
		}
		proposal, ok := s.IndexedMapContext.Events().Get(statemachine.ProposalID(proposalID))
		if !ok {
			return errors.NewFault("cannot find proposal %d", proposalID)
		}
		listener := &IndexedMapListener{}
		if err := reader.ReadMessage(listener); err != nil {
			return err
		}
		s.streams[proposal.ID()] = listener
		proposal.Watch(func(state statemachine.ProposalState) {
			if state != statemachine.Running {
				delete(s.streams, proposal.ID())
			}
		})
	}
	return nil
}

func (s *indexedMapStateMachine) Append(proposal statemachine.Proposal[*AppendInput, *AppendOutput]) {
	defer proposal.Close()

	// Check that the key does not already exist in the map
	if entry, ok := s.keys[proposal.Input().Key]; ok {
		proposal.Error(errors.NewAlreadyExists("key %s already exists at index %d", proposal.Input().Key, entry.Index))
		return
	}

	// Increment the map index
	s.lastIndex++
	index := s.lastIndex

	// Create a new entry value and set it in the map.
	entry := &LinkedMapEntryValue{
		IndexedMapEntry: &IndexedMapEntry{
			Index: index,
			Key:   proposal.Input().Key,
			Value: &IndexedMapValue{
				Value:   proposal.Input().Value,
				Version: uint64(proposal.ID()),
			},
		},
	}
	if proposal.Input().TTL != nil {
		expire := s.Scheduler().Time().Add(*proposal.Input().TTL)
		entry.Value.Expire = &expire
	}
	s.keys[entry.Key] = entry
	s.indexes[entry.Index] = entry

	// Set the first entry if not set
	if s.firstEntry == nil {
		s.firstEntry = entry
	}

	// If the last entry is set, link it to the new entry
	if s.lastEntry != nil {
		s.lastEntry.Next = entry
		entry.Prev = s.lastEntry
	}

	// Update the last entry
	s.lastEntry = entry

	// Schedule the timeout for the value if necessary.
	s.scheduleTTL(entry)

	s.notify(entry, &Event{
		Key:   entry.Key,
		Index: entry.Index,
		Event: &Event_Inserted_{
			Inserted: &Event_Inserted{
				Value: *newStateMachineValue(entry.Value),
			},
		},
	})

	proposal.Output(&AppendOutput{
		Entry: newStateMachineEntry(entry.IndexedMapEntry),
	})
}

func (s *indexedMapStateMachine) Update(proposal statemachine.Proposal[*UpdateInput, *UpdateOutput]) {
	defer proposal.Close()

	// Get the current entry value by key, index, or both
	var oldEntry *LinkedMapEntryValue
	if proposal.Input().Index != 0 {
		if e, ok := s.indexes[proposal.Input().Index]; !ok {
			proposal.Error(errors.NewNotFound("index %d not found", proposal.Input().Index))
			return
		} else if proposal.Input().Key != "" && e.Key != proposal.Input().Key {
			proposal.Error(errors.NewFault("key at index %d does not match proposed Update key %s", proposal.Input().Index, proposal.Input().Key))
			return
		} else {
			oldEntry = e
		}
	} else if proposal.Input().Key != "" {
		if e, ok := s.keys[proposal.Input().Key]; !ok {
			proposal.Error(errors.NewNotFound("key %s not found", proposal.Input().Key))
			return
		} else {
			oldEntry = e
		}
	} else {
		proposal.Error(errors.NewInvalid("must specify either a key or index to update"))
		return
	}

	// If a prev_version was specified, check that the previous version matches
	if proposal.Input().PrevVersion != 0 && oldEntry.Value.Version != proposal.Input().PrevVersion {
		proposal.Error(errors.NewConflict("key %s version %d does not match prev_version %d", oldEntry.Key, oldEntry.Value.Version, proposal.Input().PrevVersion))
		return
	}

	// If the value is equal to the current value, return a no-op.
	if bytes.Equal(oldEntry.Value.Value, proposal.Input().Value) {
		proposal.Output(&UpdateOutput{
			Entry: newStateMachineEntry(oldEntry.IndexedMapEntry),
		})
		return
	}

	// Create a new entry value and set it in the map.
	entry := &LinkedMapEntryValue{
		IndexedMapEntry: &IndexedMapEntry{
			Index: oldEntry.Index,
			Key:   oldEntry.Key,
			Value: &IndexedMapValue{
				Value:   proposal.Input().Value,
				Version: uint64(proposal.ID()),
			},
		},
		Prev: oldEntry.Prev,
		Next: oldEntry.Next,
	}
	if proposal.Input().TTL != nil {
		expire := s.Scheduler().Time().Add(*proposal.Input().TTL)
		entry.Value.Expire = &expire
	}
	s.keys[entry.Key] = entry
	s.indexes[entry.Index] = entry

	// Update links for previous and next entries
	if oldEntry.Prev != nil {
		oldEntry.Prev.Next = entry
	} else {
		s.firstEntry = entry
	}
	if oldEntry.Next != nil {
		oldEntry.Next.Prev = entry
	} else {
		s.lastEntry = entry
	}

	// Schedule the timeout for the value if necessary.
	s.scheduleTTL(entry)

	s.notify(entry, &Event{
		Key:   entry.Key,
		Index: entry.Index,
		Event: &Event_Updated_{
			Updated: &Event_Updated{
				Value: *newStateMachineValue(entry.Value),
			},
		},
	})

	proposal.Output(&UpdateOutput{
		Entry: newStateMachineEntry(entry.IndexedMapEntry),
	})
}

func (s *indexedMapStateMachine) Remove(proposal statemachine.Proposal[*RemoveInput, *RemoveOutput]) {
	defer proposal.Close()

	var entry *LinkedMapEntryValue
	var ok bool
	if proposal.Input().Index != 0 {
		if entry, ok = s.indexes[proposal.Input().Index]; !ok {
			proposal.Error(errors.NewNotFound("no entry found at index %d", proposal.Input().Index))
			return
		}
	} else {
		if entry, ok = s.keys[proposal.Input().Key]; !ok {
			proposal.Error(errors.NewNotFound("no entry found at key %s", proposal.Input().Key))
			return
		}
	}

	if proposal.Input().PrevVersion != 0 && entry.Value.Version != proposal.Input().PrevVersion {
		proposal.Error(errors.NewConflict("key %s version %d does not match prev_version %d", entry.Key, entry.Value.Version, proposal.Input().PrevVersion))
		return
	}

	// Delete the entry from the map.
	delete(s.keys, entry.Key)
	delete(s.indexes, entry.Index)

	// Cancel any TTLs.
	s.cancelTTL(proposal.Input().Key)

	// Update links for previous and next entries
	if entry.Prev != nil {
		entry.Prev.Next = entry.Next
	} else {
		s.firstEntry = entry.Next
	}
	if entry.Next != nil {
		entry.Next.Prev = entry.Prev
	} else {
		s.lastEntry = entry.Prev
	}

	s.notify(entry, &Event{
		Key:   entry.Key,
		Index: entry.Index,
		Event: &Event_Removed_{
			Removed: &Event_Removed{
				Value: *newStateMachineValue(entry.IndexedMapEntry.Value),
			},
		},
	})

	proposal.Output(&RemoveOutput{
		Entry: newStateMachineEntry(entry.IndexedMapEntry),
	})
}

func (s *indexedMapStateMachine) Clear(proposal statemachine.Proposal[*ClearInput, *ClearOutput]) {
	defer proposal.Close()
	for key, entry := range s.keys {
		s.notify(entry, &Event{
			Key:   entry.Key,
			Index: entry.Index,
			Event: &Event_Removed_{
				Removed: &Event_Removed{
					Value: *newStateMachineValue(entry.IndexedMapEntry.Value),
				},
			},
		})
		s.cancelTTL(key)
	}
	s.keys = make(map[string]*LinkedMapEntryValue)
	s.indexes = make(map[uint64]*LinkedMapEntryValue)
	s.firstEntry = nil
	s.lastEntry = nil
	proposal.Output(&ClearOutput{})
}

func (s *indexedMapStateMachine) Events(proposal statemachine.Proposal[*EventsInput, *EventsOutput]) {
	listener := &IndexedMapListener{
		Key: proposal.Input().Key,
	}
	s.streams[proposal.ID()] = listener
	proposal.Watch(func(state statemachine.ProposalState) {
		if state != statemachine.Running {
			delete(s.streams, proposal.ID())
		}
	})
}

func (s *indexedMapStateMachine) Size(query statemachine.Query[*SizeInput, *SizeOutput]) {
	defer query.Close()
	query.Output(&SizeOutput{
		Size_: uint32(len(s.keys)),
	})
}

func (s *indexedMapStateMachine) Get(query statemachine.Query[*GetInput, *GetOutput]) {
	defer query.Close()

	var entry *LinkedMapEntryValue
	var ok bool
	if query.Input().Index > 0 {
		if entry, ok = s.indexes[query.Input().Index]; !ok {
			query.Error(errors.NewNotFound("no entry found at index %d", query.Input().Index))
			return
		}
	} else {
		if entry, ok = s.keys[query.Input().Key]; !ok {
			query.Error(errors.NewNotFound("no entry found at key %s", query.Input().Key))
			return
		}
	}

	query.Output(&GetOutput{
		Entry: newStateMachineEntry(entry.IndexedMapEntry),
	})
}

func (s *indexedMapStateMachine) FirstEntry(query statemachine.Query[*FirstEntryInput, *FirstEntryOutput]) {
	defer query.Close()
	if s.firstEntry == nil {
		query.Error(errors.NewNotFound("map is empty"))
	} else {
		query.Output(&FirstEntryOutput{
			Entry: newStateMachineEntry(s.firstEntry.IndexedMapEntry),
		})
	}
}

func (s *indexedMapStateMachine) LastEntry(query statemachine.Query[*LastEntryInput, *LastEntryOutput]) {
	defer query.Close()
	if s.lastEntry == nil {
		query.Error(errors.NewNotFound("map is empty"))
	} else {
		query.Output(&LastEntryOutput{
			Entry: newStateMachineEntry(s.lastEntry.IndexedMapEntry),
		})
	}
}

func (s *indexedMapStateMachine) NextEntry(query statemachine.Query[*NextEntryInput, *NextEntryOutput]) {
	defer query.Close()
	entry, ok := s.indexes[query.Input().Index]
	if !ok {
		entry = s.firstEntry
		if entry == nil {
			query.Error(errors.NewNotFound("map is empty"))
			return
		}
		for entry != nil && entry.Index >= query.Input().Index {
			entry = entry.Next
		}
	}
	entry = entry.Next
	if entry == nil {
		query.Error(errors.NewNotFound("no entry found after index %d", query.Input().Index))
	} else {
		query.Output(&NextEntryOutput{
			Entry: newStateMachineEntry(entry.IndexedMapEntry),
		})
	}
}

func (s *indexedMapStateMachine) PrevEntry(query statemachine.Query[*PrevEntryInput, *PrevEntryOutput]) {
	defer query.Close()
	entry, ok := s.indexes[query.Input().Index]
	if !ok {
		entry = s.lastEntry
		if entry == nil {
			query.Error(errors.NewNotFound("map is empty"))
			return
		}
		for entry != nil && entry.Index >= query.Input().Index {
			entry = entry.Prev
		}
	}
	entry = entry.Prev
	if entry == nil {
		query.Error(errors.NewNotFound("no entry found prior to index %d", query.Input().Index))
	} else {
		query.Output(&PrevEntryOutput{
			Entry: newStateMachineEntry(entry.IndexedMapEntry),
		})
	}
}

func (s *indexedMapStateMachine) Entries(query statemachine.Query[*EntriesInput, *EntriesOutput]) {
	defer query.Close()
	entry := s.firstEntry
	for entry != nil {
		query.Output(&EntriesOutput{
			Entry: *newStateMachineEntry(entry.IndexedMapEntry),
		})
		entry = entry.Next
	}

	if query.Input().Watch {
		s.mu.Lock()
		s.watchers[query.ID()] = query
		s.mu.Unlock()
		query.Watch(func(state statemachine.QueryState) {
			if state != statemachine.Running {
				s.mu.Lock()
				delete(s.watchers, query.ID())
				s.mu.Unlock()
			}
		})
	} else {
		query.Close()
	}
}

func (s *indexedMapStateMachine) notify(entry *LinkedMapEntryValue, event *Event) {
	for proposalID, listener := range s.streams {
		if listener.Key == "" || listener.Key == event.Key {
			proposal, ok := s.IndexedMapContext.Events().Get(proposalID)
			if ok {
				proposal.Output(&EventsOutput{
					Event: *event,
				})
			} else {
				delete(s.streams, proposalID)
			}
		}
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, watcher := range s.watchers {
		watcher.Output(&EntriesOutput{
			Entry: *newStateMachineEntry(entry.IndexedMapEntry),
		})
	}
}

func (s *indexedMapStateMachine) scheduleTTL(entry *LinkedMapEntryValue) {
	s.cancelTTL(entry.Key)
	if entry.Value.Expire != nil {
		s.timers[entry.Key] = s.Scheduler().Schedule(*entry.Value.Expire, func() {
			// Delete the entry from the key/index maps
			delete(s.keys, entry.Key)
			delete(s.indexes, entry.Index)

			// Update links for previous and next entries
			if entry.Prev != nil {
				entry.Prev.Next = entry.Next
			} else {
				s.firstEntry = entry.Next
			}
			if entry.Next != nil {
				entry.Next.Prev = entry.Prev
			} else {
				s.lastEntry = entry.Prev
			}

			// Notify watchers of the removal
			s.notify(entry, &Event{
				Key:   entry.Key,
				Index: entry.Index,
				Event: &Event_Removed_{
					Removed: &Event_Removed{
						Value: Value{
							Value:   entry.Value.Value,
							Version: entry.Value.Version,
						},
						Expired: true,
					},
				},
			})
		})
	}
}

func (s *indexedMapStateMachine) cancelTTL(key string) {
	ttlCancelFunc, ok := s.timers[key]
	if ok {
		ttlCancelFunc()
	}
}

func newStateMachineEntry(entry *IndexedMapEntry) *Entry {
	return &Entry{
		Key:   entry.Key,
		Index: entry.Index,
		Value: newStateMachineValue(entry.Value),
	}
}

func newStateMachineValue(value *IndexedMapValue) *Value {
	if value == nil {
		return nil
	}
	return &Value{
		Value:   value.Value,
		Version: value.Version,
	}
}