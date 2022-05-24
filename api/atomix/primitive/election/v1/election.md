# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [atomix/primitive/election/v1/election.proto](#atomix_primitive_election_v1_election-proto)
    - [AnointInput](#atomix-primitive-election-v1-AnointInput)
    - [AnointOutput](#atomix-primitive-election-v1-AnointOutput)
    - [AnointRequest](#atomix-primitive-election-v1-AnointRequest)
    - [AnointResponse](#atomix-primitive-election-v1-AnointResponse)
    - [EnterInput](#atomix-primitive-election-v1-EnterInput)
    - [EnterOutput](#atomix-primitive-election-v1-EnterOutput)
    - [EnterRequest](#atomix-primitive-election-v1-EnterRequest)
    - [EnterResponse](#atomix-primitive-election-v1-EnterResponse)
    - [Event](#atomix-primitive-election-v1-Event)
    - [EventsInput](#atomix-primitive-election-v1-EventsInput)
    - [EventsOutput](#atomix-primitive-election-v1-EventsOutput)
    - [EventsRequest](#atomix-primitive-election-v1-EventsRequest)
    - [EventsResponse](#atomix-primitive-election-v1-EventsResponse)
    - [EvictInput](#atomix-primitive-election-v1-EvictInput)
    - [EvictOutput](#atomix-primitive-election-v1-EvictOutput)
    - [EvictRequest](#atomix-primitive-election-v1-EvictRequest)
    - [EvictResponse](#atomix-primitive-election-v1-EvictResponse)
    - [GetTermInput](#atomix-primitive-election-v1-GetTermInput)
    - [GetTermOutput](#atomix-primitive-election-v1-GetTermOutput)
    - [GetTermRequest](#atomix-primitive-election-v1-GetTermRequest)
    - [GetTermResponse](#atomix-primitive-election-v1-GetTermResponse)
    - [LeaderElectionConfig](#atomix-primitive-election-v1-LeaderElectionConfig)
    - [PromoteInput](#atomix-primitive-election-v1-PromoteInput)
    - [PromoteOutput](#atomix-primitive-election-v1-PromoteOutput)
    - [PromoteRequest](#atomix-primitive-election-v1-PromoteRequest)
    - [PromoteResponse](#atomix-primitive-election-v1-PromoteResponse)
    - [Term](#atomix-primitive-election-v1-Term)
    - [WithdrawInput](#atomix-primitive-election-v1-WithdrawInput)
    - [WithdrawOutput](#atomix-primitive-election-v1-WithdrawOutput)
    - [WithdrawRequest](#atomix-primitive-election-v1-WithdrawRequest)
    - [WithdrawResponse](#atomix-primitive-election-v1-WithdrawResponse)
  
    - [Event.Type](#atomix-primitive-election-v1-Event-Type)
  
    - [LeaderElection](#atomix-primitive-election-v1-LeaderElection)
  
- [Scalar Value Types](#scalar-value-types)



<a name="atomix_primitive_election_v1_election-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## atomix/primitive/election/v1/election.proto



<a name="atomix-primitive-election-v1-AnointInput"></a>

### AnointInput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| candidate_id | [string](#string) |  |  |






<a name="atomix-primitive-election-v1-AnointOutput"></a>

### AnointOutput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| term | [Term](#atomix-primitive-election-v1-Term) |  |  |






<a name="atomix-primitive-election-v1-AnointRequest"></a>

### AnointRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.RequestHeaders](#atomix-primitive-v1-RequestHeaders) |  |  |
| input | [AnointInput](#atomix-primitive-election-v1-AnointInput) |  |  |






<a name="atomix-primitive-election-v1-AnointResponse"></a>

### AnointResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.ResponseHeaders](#atomix-primitive-v1-ResponseHeaders) |  |  |
| output | [AnointOutput](#atomix-primitive-election-v1-AnointOutput) |  |  |






<a name="atomix-primitive-election-v1-EnterInput"></a>

### EnterInput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| candidate_id | [string](#string) |  |  |






<a name="atomix-primitive-election-v1-EnterOutput"></a>

### EnterOutput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| term | [Term](#atomix-primitive-election-v1-Term) |  |  |






<a name="atomix-primitive-election-v1-EnterRequest"></a>

### EnterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.RequestHeaders](#atomix-primitive-v1-RequestHeaders) |  |  |
| input | [EnterInput](#atomix-primitive-election-v1-EnterInput) |  |  |






<a name="atomix-primitive-election-v1-EnterResponse"></a>

### EnterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.ResponseHeaders](#atomix-primitive-v1-ResponseHeaders) |  |  |
| output | [EnterOutput](#atomix-primitive-election-v1-EnterOutput) |  |  |






<a name="atomix-primitive-election-v1-Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Event.Type](#atomix-primitive-election-v1-Event-Type) |  |  |
| term | [Term](#atomix-primitive-election-v1-Term) |  |  |






<a name="atomix-primitive-election-v1-EventsInput"></a>

### EventsInput







<a name="atomix-primitive-election-v1-EventsOutput"></a>

### EventsOutput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| event | [Event](#atomix-primitive-election-v1-Event) |  |  |






<a name="atomix-primitive-election-v1-EventsRequest"></a>

### EventsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.RequestHeaders](#atomix-primitive-v1-RequestHeaders) |  |  |
| input | [EventsInput](#atomix-primitive-election-v1-EventsInput) |  |  |






<a name="atomix-primitive-election-v1-EventsResponse"></a>

### EventsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.ResponseHeaders](#atomix-primitive-v1-ResponseHeaders) |  |  |
| output | [EventsOutput](#atomix-primitive-election-v1-EventsOutput) |  |  |






<a name="atomix-primitive-election-v1-EvictInput"></a>

### EvictInput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| candidate_id | [string](#string) |  |  |






<a name="atomix-primitive-election-v1-EvictOutput"></a>

### EvictOutput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| term | [Term](#atomix-primitive-election-v1-Term) |  |  |






<a name="atomix-primitive-election-v1-EvictRequest"></a>

### EvictRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.RequestHeaders](#atomix-primitive-v1-RequestHeaders) |  |  |
| input | [EvictInput](#atomix-primitive-election-v1-EvictInput) |  |  |






<a name="atomix-primitive-election-v1-EvictResponse"></a>

### EvictResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.ResponseHeaders](#atomix-primitive-v1-ResponseHeaders) |  |  |
| output | [EvictOutput](#atomix-primitive-election-v1-EvictOutput) |  |  |






<a name="atomix-primitive-election-v1-GetTermInput"></a>

### GetTermInput







<a name="atomix-primitive-election-v1-GetTermOutput"></a>

### GetTermOutput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| term | [Term](#atomix-primitive-election-v1-Term) |  |  |






<a name="atomix-primitive-election-v1-GetTermRequest"></a>

### GetTermRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.RequestHeaders](#atomix-primitive-v1-RequestHeaders) |  |  |
| input | [GetTermInput](#atomix-primitive-election-v1-GetTermInput) |  |  |






<a name="atomix-primitive-election-v1-GetTermResponse"></a>

### GetTermResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.ResponseHeaders](#atomix-primitive-v1-ResponseHeaders) |  |  |
| output | [GetTermOutput](#atomix-primitive-election-v1-GetTermOutput) |  |  |






<a name="atomix-primitive-election-v1-LeaderElectionConfig"></a>

### LeaderElectionConfig







<a name="atomix-primitive-election-v1-PromoteInput"></a>

### PromoteInput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| candidate_id | [string](#string) |  |  |






<a name="atomix-primitive-election-v1-PromoteOutput"></a>

### PromoteOutput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| term | [Term](#atomix-primitive-election-v1-Term) |  |  |






<a name="atomix-primitive-election-v1-PromoteRequest"></a>

### PromoteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.RequestHeaders](#atomix-primitive-v1-RequestHeaders) |  |  |
| input | [PromoteInput](#atomix-primitive-election-v1-PromoteInput) |  |  |






<a name="atomix-primitive-election-v1-PromoteResponse"></a>

### PromoteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.ResponseHeaders](#atomix-primitive-v1-ResponseHeaders) |  |  |
| output | [PromoteOutput](#atomix-primitive-election-v1-PromoteOutput) |  |  |






<a name="atomix-primitive-election-v1-Term"></a>

### Term



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| leader_id | [string](#string) |  |  |
| candidates | [string](#string) | repeated |  |
| timestamp | [atomix.primitive.meta.v1.Timestamp](#atomix-primitive-meta-v1-Timestamp) |  |  |






<a name="atomix-primitive-election-v1-WithdrawInput"></a>

### WithdrawInput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| candidate_id | [string](#string) |  |  |






<a name="atomix-primitive-election-v1-WithdrawOutput"></a>

### WithdrawOutput



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| term | [Term](#atomix-primitive-election-v1-Term) |  |  |






<a name="atomix-primitive-election-v1-WithdrawRequest"></a>

### WithdrawRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.RequestHeaders](#atomix-primitive-v1-RequestHeaders) |  |  |
| input | [WithdrawInput](#atomix-primitive-election-v1-WithdrawInput) |  |  |






<a name="atomix-primitive-election-v1-WithdrawResponse"></a>

### WithdrawResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [atomix.primitive.v1.ResponseHeaders](#atomix-primitive-v1-ResponseHeaders) |  |  |
| output | [WithdrawOutput](#atomix-primitive-election-v1-WithdrawOutput) |  |  |





 


<a name="atomix-primitive-election-v1-Event-Type"></a>

### Event.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| CHANGED | 1 |  |


 

 


<a name="atomix-primitive-election-v1-LeaderElection"></a>

### LeaderElection
LeaderElection is a service for a leader election primitive

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Enter | [EnterRequest](#atomix-primitive-election-v1-EnterRequest) | [EnterResponse](#atomix-primitive-election-v1-EnterResponse) | Enter enters the leader election |
| Withdraw | [WithdrawRequest](#atomix-primitive-election-v1-WithdrawRequest) | [WithdrawResponse](#atomix-primitive-election-v1-WithdrawResponse) | Withdraw withdraws a candidate from the leader election |
| Anoint | [AnointRequest](#atomix-primitive-election-v1-AnointRequest) | [AnointResponse](#atomix-primitive-election-v1-AnointResponse) | Anoint anoints a candidate leader |
| Promote | [PromoteRequest](#atomix-primitive-election-v1-PromoteRequest) | [PromoteResponse](#atomix-primitive-election-v1-PromoteResponse) | Promote promotes a candidate |
| Evict | [EvictRequest](#atomix-primitive-election-v1-EvictRequest) | [EvictResponse](#atomix-primitive-election-v1-EvictResponse) | Evict evicts a candidate from the election |
| GetTerm | [GetTermRequest](#atomix-primitive-election-v1-GetTermRequest) | [GetTermResponse](#atomix-primitive-election-v1-GetTermResponse) | GetTerm gets the current leadership term |
| Events | [EventsRequest](#atomix-primitive-election-v1-EventsRequest) | [EventsResponse](#atomix-primitive-election-v1-EventsResponse) stream | Events listens for leadership events |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
