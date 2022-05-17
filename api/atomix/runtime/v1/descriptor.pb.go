// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: atomix/runtime/v1/descriptor.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ComponentType int32

const (
	ComponentType_NONE    ComponentType = 0
	ComponentType_ATOM    ComponentType = 1
	ComponentType_MANAGER ComponentType = 2
)

var ComponentType_name = map[int32]string{
	0: "NONE",
	1: "ATOM",
	2: "MANAGER",
}

var ComponentType_value = map[string]int32{
	"NONE":    0,
	"ATOM":    1,
	"MANAGER": 2,
}

func (x ComponentType) String() string {
	return proto.EnumName(ComponentType_name, int32(x))
}

func (ComponentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_397f97b76c9fa393, []int{0}
}

// OperationType is an enum for specifying the type of operation
type OperationType int32

const (
	OperationType_COMMAND OperationType = 0
	OperationType_QUERY   OperationType = 1
	OperationType_CREATE  OperationType = 2
	OperationType_CLOSE   OperationType = 3
)

var OperationType_name = map[int32]string{
	0: "COMMAND",
	1: "QUERY",
	2: "CREATE",
	3: "CLOSE",
}

var OperationType_value = map[string]int32{
	"COMMAND": 0,
	"QUERY":   1,
	"CREATE":  2,
	"CLOSE":   3,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}

func (OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_397f97b76c9fa393, []int{1}
}

var E_Name = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "atomix.runtime.v1.name",
	Tag:           "bytes,50000,opt,name=name",
	Filename:      "atomix/runtime/v1/descriptor.proto",
}

var E_Component = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
	ExtensionType: (*ComponentType)(nil),
	Field:         50001,
	Name:          "atomix.runtime.v1.component",
	Tag:           "varint,50001,opt,name=component,enum=atomix.runtime.v1.ComponentType",
	Filename:      "atomix/runtime/v1/descriptor.proto",
}

var E_Headers = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         51000,
	Name:          "atomix.runtime.v1.headers",
	Tag:           "varint,51000,opt,name=headers",
	Filename:      "atomix/runtime/v1/descriptor.proto",
}

var E_Input = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         51001,
	Name:          "atomix.runtime.v1.input",
	Tag:           "varint,51001,opt,name=input",
	Filename:      "atomix/runtime/v1/descriptor.proto",
}

var E_Output = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         51002,
	Name:          "atomix.runtime.v1.output",
	Tag:           "varint,51002,opt,name=output",
	Filename:      "atomix/runtime/v1/descriptor.proto",
}

var E_OperationId = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MethodOptions)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         52000,
	Name:          "atomix.runtime.v1.operation_id",
	Tag:           "varint,52000,opt,name=operation_id",
	Filename:      "atomix/runtime/v1/descriptor.proto",
}

var E_OperationType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MethodOptions)(nil),
	ExtensionType: (*OperationType)(nil),
	Field:         52001,
	Name:          "atomix.runtime.v1.operation_type",
	Tag:           "varint,52001,opt,name=operation_type,enum=atomix.runtime.v1.OperationType",
	Filename:      "atomix/runtime/v1/descriptor.proto",
}

func init() {
	proto.RegisterEnum("atomix.runtime.v1.ComponentType", ComponentType_name, ComponentType_value)
	proto.RegisterEnum("atomix.runtime.v1.OperationType", OperationType_name, OperationType_value)
	proto.RegisterExtension(E_Name)
	proto.RegisterExtension(E_Component)
	proto.RegisterExtension(E_Headers)
	proto.RegisterExtension(E_Input)
	proto.RegisterExtension(E_Output)
	proto.RegisterExtension(E_OperationId)
	proto.RegisterExtension(E_OperationType)
}

func init() {
	proto.RegisterFile("atomix/runtime/v1/descriptor.proto", fileDescriptor_397f97b76c9fa393)
}

var fileDescriptor_397f97b76c9fa393 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0xae, 0x93, 0x40,
	0x14, 0x87, 0xe1, 0xde, 0xdb, 0x7f, 0xa7, 0xd2, 0xe0, 0xac, 0x1a, 0x13, 0xb1, 0x71, 0x65, 0xba,
	0x18, 0xac, 0xa6, 0x31, 0xb2, 0x30, 0x41, 0x44, 0x63, 0x22, 0x10, 0xa7, 0x75, 0xe1, 0xca, 0xd0,
	0x32, 0xb6, 0x93, 0x14, 0x86, 0xd0, 0xa1, 0xd1, 0x17, 0x70, 0x69, 0x5c, 0x75, 0xad, 0x6f, 0xa1,
	0x3e, 0x81, 0xcb, 0xba, 0x73, 0x69, 0xda, 0x17, 0x31, 0x50, 0x68, 0x6d, 0x88, 0xa9, 0x3b, 0x98,
	0x7c, 0xdf, 0xef, 0x9c, 0x39, 0x67, 0xe0, 0xb6, 0x2f, 0x78, 0xc8, 0xde, 0xe9, 0x49, 0x1a, 0x09,
	0x16, 0x52, 0x7d, 0x35, 0xd0, 0x03, 0xba, 0x9c, 0x26, 0x2c, 0x16, 0x3c, 0xc1, 0x71, 0xc2, 0x05,
	0x47, 0xd7, 0xf7, 0x0c, 0x2e, 0x18, 0xbc, 0x1a, 0xdc, 0xe8, 0xcd, 0x38, 0x9f, 0x2d, 0xa8, 0x9e,
	0x03, 0x93, 0xf4, 0x6d, 0x45, 0xea, 0xdf, 0x05, 0xc5, 0xe2, 0x61, 0xcc, 0x23, 0x1a, 0x89, 0xf1,
	0xfb, 0x98, 0xa2, 0x26, 0x5c, 0xb9, 0x9e, 0x6b, 0xab, 0x52, 0xf6, 0x65, 0x8e, 0x3d, 0x47, 0x95,
	0x51, 0x1b, 0x1a, 0x8e, 0xe9, 0x9a, 0xcf, 0x6c, 0xa2, 0x5e, 0xf4, 0x1f, 0x81, 0xe2, 0xc5, 0x34,
	0xf1, 0x05, 0xe3, 0x51, 0x6e, 0xb4, 0xa1, 0x61, 0x79, 0x8e, 0x63, 0xba, 0x4f, 0x54, 0x09, 0xb5,
	0xa0, 0xf6, 0xf2, 0x95, 0x4d, 0x5e, 0xab, 0x32, 0x02, 0xa8, 0x5b, 0xc4, 0x36, 0xc7, 0xb6, 0x7a,
	0x91, 0x1d, 0x5b, 0x2f, 0xbc, 0x91, 0xad, 0x5e, 0x1a, 0x43, 0xb8, 0x8a, 0xfc, 0x90, 0xa2, 0x5b,
	0x78, 0xdf, 0x1c, 0x2e, 0x9b, 0xc3, 0x23, 0x9a, 0xac, 0xd8, 0x94, 0x7a, 0x71, 0x16, 0xbd, 0xec,
	0x6e, 0x3e, 0x5c, 0xf6, 0xe4, 0x3b, 0x2d, 0x92, 0xe3, 0x86, 0x0f, 0xad, 0x69, 0xd9, 0xe8, 0x79,
	0xf7, 0x67, 0xee, 0x76, 0xee, 0xf5, 0x70, 0x65, 0x28, 0xf8, 0xe4, 0xbe, 0xe4, 0x98, 0x6a, 0x3c,
	0x84, 0xc6, 0x9c, 0xfa, 0x01, 0x4d, 0x96, 0xe8, 0x66, 0xa5, 0xc0, 0x53, 0x46, 0x17, 0x41, 0x19,
	0xff, 0xf5, 0x63, 0x16, 0xdf, 0x24, 0x25, 0x6f, 0x0c, 0xa1, 0xc6, 0xa2, 0x38, 0x15, 0xe7, 0xc4,
	0x6f, 0x85, 0xb8, 0xa7, 0x8d, 0x07, 0x50, 0xe7, 0xa9, 0xf8, 0x0f, 0xef, 0x7b, 0xe1, 0x15, 0xb8,
	0x61, 0xc1, 0x35, 0x5e, 0x2e, 0xe1, 0x0d, 0x0b, 0x90, 0x56, 0xd1, 0x1d, 0x2a, 0xe6, 0xfc, 0xe0,
	0x7f, 0x5e, 0x67, 0xbe, 0x42, 0xda, 0x07, 0xeb, 0x79, 0x60, 0x30, 0xe8, 0x1c, 0x43, 0x44, 0xb6,
	0xca, 0x73, 0x31, 0x5f, 0xd6, 0xff, 0x1e, 0xeb, 0xc9, 0xa3, 0x20, 0x0a, 0xff, 0xfb, 0xf7, 0x71,
	0xf7, 0xc7, 0x56, 0x93, 0x37, 0x5b, 0x4d, 0xfe, 0xbd, 0xd5, 0xe4, 0x4f, 0x3b, 0x4d, 0xda, 0xec,
	0x34, 0xe9, 0xd7, 0x4e, 0x93, 0x26, 0xf5, 0xbc, 0xd4, 0xfd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x1d, 0xaf, 0x9c, 0xc9, 0xe2, 0x02, 0x00, 0x00,
}