// Code generated by protoc-gen-go.
// source: src/rgraphql.proto
// DO NOT EDIT!

/*
Package rgraphql is a generated protocol buffer package.

It is generated from these files:
	src/rgraphql.proto

It has these top-level messages:
	RGQLQueryFieldDirective
	RGQLQueryTreeNode
	FieldArgument
	ASTValue
	RGQLClientMessage
	RGQLTreeMutation
	RGQLFieldMutation
	RGQLServerMessage
	RGQLValueMutation
*/
package rgraphql

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ASTValue_ASTValueKind int32

const (
	ASTValue_AST_VALUE_NULL   ASTValue_ASTValueKind = 0
	ASTValue_AST_VALUE_STRING ASTValue_ASTValueKind = 1
	ASTValue_AST_VALUE_ENUM   ASTValue_ASTValueKind = 2
	ASTValue_AST_VALUE_INT    ASTValue_ASTValueKind = 3
	ASTValue_AST_VALUE_FLOAT  ASTValue_ASTValueKind = 4
	ASTValue_AST_VALUE_BOOL   ASTValue_ASTValueKind = 5
	ASTValue_AST_VALUE_LIST   ASTValue_ASTValueKind = 6
	ASTValue_AST_VALUE_OBJECT ASTValue_ASTValueKind = 7
)

var ASTValue_ASTValueKind_name = map[int32]string{
	0: "AST_VALUE_NULL",
	1: "AST_VALUE_STRING",
	2: "AST_VALUE_ENUM",
	3: "AST_VALUE_INT",
	4: "AST_VALUE_FLOAT",
	5: "AST_VALUE_BOOL",
	6: "AST_VALUE_LIST",
	7: "AST_VALUE_OBJECT",
}
var ASTValue_ASTValueKind_value = map[string]int32{
	"AST_VALUE_NULL":   0,
	"AST_VALUE_STRING": 1,
	"AST_VALUE_ENUM":   2,
	"AST_VALUE_INT":    3,
	"AST_VALUE_FLOAT":  4,
	"AST_VALUE_BOOL":   5,
	"AST_VALUE_LIST":   6,
	"AST_VALUE_OBJECT": 7,
}

func (x ASTValue_ASTValueKind) String() string {
	return proto.EnumName(ASTValue_ASTValueKind_name, int32(x))
}
func (ASTValue_ASTValueKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type RGQLTreeMutation_SubtreeOperation int32

const (
	// Add a child tree to the subtree.
	RGQLTreeMutation_SUBTREE_ADD_CHILD RGQLTreeMutation_SubtreeOperation = 0
	// Delete a tree node and all children.
	RGQLTreeMutation_SUBTREE_DELETE RGQLTreeMutation_SubtreeOperation = 1
)

var RGQLTreeMutation_SubtreeOperation_name = map[int32]string{
	0: "SUBTREE_ADD_CHILD",
	1: "SUBTREE_DELETE",
}
var RGQLTreeMutation_SubtreeOperation_value = map[string]int32{
	"SUBTREE_ADD_CHILD": 0,
	"SUBTREE_DELETE":    1,
}

func (x RGQLTreeMutation_SubtreeOperation) String() string {
	return proto.EnumName(RGQLTreeMutation_SubtreeOperation_name, int32(x))
}
func (RGQLTreeMutation_SubtreeOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

type RGQLValueMutation_ValueOperation int32

const (
	// Set the value of this resolver.
	RGQLValueMutation_VALUE_SET RGQLValueMutation_ValueOperation = 0
	// Push a child resolver of this array.
	// Value will be the current value of that resolver.
	RGQLValueMutation_VALUE_PUSH RGQLValueMutation_ValueOperation = 1
	// Error this resolver (the value will enter the errored state).
	RGQLValueMutation_VALUE_ERROR RGQLValueMutation_ValueOperation = 2
	// Delete this resolver altogether.
	RGQLValueMutation_VALUE_DELETE RGQLValueMutation_ValueOperation = 3
)

var RGQLValueMutation_ValueOperation_name = map[int32]string{
	0: "VALUE_SET",
	1: "VALUE_PUSH",
	2: "VALUE_ERROR",
	3: "VALUE_DELETE",
}
var RGQLValueMutation_ValueOperation_value = map[string]int32{
	"VALUE_SET":    0,
	"VALUE_PUSH":   1,
	"VALUE_ERROR":  2,
	"VALUE_DELETE": 3,
}

func (x RGQLValueMutation_ValueOperation) String() string {
	return proto.EnumName(RGQLValueMutation_ValueOperation_name, int32(x))
}
func (RGQLValueMutation_ValueOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

type RGQLQueryFieldDirective struct {
	// Directive name
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Optional arguments.
	Args []*FieldArgument `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
}

func (m *RGQLQueryFieldDirective) Reset()                    { *m = RGQLQueryFieldDirective{} }
func (m *RGQLQueryFieldDirective) String() string            { return proto.CompactTextString(m) }
func (*RGQLQueryFieldDirective) ProtoMessage()               {}
func (*RGQLQueryFieldDirective) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RGQLQueryFieldDirective) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RGQLQueryFieldDirective) GetArgs() []*FieldArgument {
	if m != nil {
		return m.Args
	}
	return nil
}

type RGQLQueryTreeNode struct {
	// Integer ID of the node.
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Name of the field this node represents.
	FieldName string `protobuf:"bytes,2,opt,name=field_name,json=fieldName" json:"field_name,omitempty"`
	// Arguments.
	Args []*FieldArgument `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
	// Directives
	Directive []*RGQLQueryFieldDirective `protobuf:"bytes,4,rep,name=directive" json:"directive,omitempty"`
	// Children
	Children []*RGQLQueryTreeNode `protobuf:"bytes,5,rep,name=children" json:"children,omitempty"`
}

func (m *RGQLQueryTreeNode) Reset()                    { *m = RGQLQueryTreeNode{} }
func (m *RGQLQueryTreeNode) String() string            { return proto.CompactTextString(m) }
func (*RGQLQueryTreeNode) ProtoMessage()               {}
func (*RGQLQueryTreeNode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RGQLQueryTreeNode) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RGQLQueryTreeNode) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *RGQLQueryTreeNode) GetArgs() []*FieldArgument {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *RGQLQueryTreeNode) GetDirective() []*RGQLQueryFieldDirective {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *RGQLQueryTreeNode) GetChildren() []*RGQLQueryTreeNode {
	if m != nil {
		return m.Children
	}
	return nil
}

type FieldArgument struct {
	Name  string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value *ASTValue `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *FieldArgument) Reset()                    { *m = FieldArgument{} }
func (m *FieldArgument) String() string            { return proto.CompactTextString(m) }
func (*FieldArgument) ProtoMessage()               {}
func (*FieldArgument) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FieldArgument) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldArgument) GetValue() *ASTValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type ASTValue struct {
	// String representation of ENUM, STRING, or OBJECT (json)
	StringValue  string                     `protobuf:"bytes,1,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	ListValue    []*ASTValue                `protobuf:"bytes,2,rep,name=list_value,json=listValue" json:"list_value,omitempty"`
	IntValue     int32                      `protobuf:"varint,3,opt,name=int_value,json=intValue" json:"int_value,omitempty"`
	FloatValue   float32                    `protobuf:"fixed32,4,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	BoolValue    bool                       `protobuf:"varint,5,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	ObjectFields []*ASTValue_ASTObjectField `protobuf:"bytes,6,rep,name=object_fields,json=objectFields" json:"object_fields,omitempty"`
	Kind         ASTValue_ASTValueKind      `protobuf:"varint,7,opt,name=kind,enum=rgraphql.ASTValue_ASTValueKind" json:"kind,omitempty"`
}

func (m *ASTValue) Reset()                    { *m = ASTValue{} }
func (m *ASTValue) String() string            { return proto.CompactTextString(m) }
func (*ASTValue) ProtoMessage()               {}
func (*ASTValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ASTValue) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *ASTValue) GetListValue() []*ASTValue {
	if m != nil {
		return m.ListValue
	}
	return nil
}

func (m *ASTValue) GetIntValue() int32 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *ASTValue) GetFloatValue() float32 {
	if m != nil {
		return m.FloatValue
	}
	return 0
}

func (m *ASTValue) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *ASTValue) GetObjectFields() []*ASTValue_ASTObjectField {
	if m != nil {
		return m.ObjectFields
	}
	return nil
}

func (m *ASTValue) GetKind() ASTValue_ASTValueKind {
	if m != nil {
		return m.Kind
	}
	return ASTValue_AST_VALUE_NULL
}

type ASTValue_ASTObjectField struct {
	Key   string    `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value *ASTValue `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *ASTValue_ASTObjectField) Reset()                    { *m = ASTValue_ASTObjectField{} }
func (m *ASTValue_ASTObjectField) String() string            { return proto.CompactTextString(m) }
func (*ASTValue_ASTObjectField) ProtoMessage()               {}
func (*ASTValue_ASTObjectField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *ASTValue_ASTObjectField) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ASTValue_ASTObjectField) GetValue() *ASTValue {
	if m != nil {
		return m.Value
	}
	return nil
}

// Messages
type RGQLClientMessage struct {
	MutateTree  *RGQLTreeMutation  `protobuf:"bytes,1,opt,name=mutate_tree,json=mutateTree" json:"mutate_tree,omitempty"`
	MutateField *RGQLFieldMutation `protobuf:"bytes,2,opt,name=mutate_field,json=mutateField" json:"mutate_field,omitempty"`
}

func (m *RGQLClientMessage) Reset()                    { *m = RGQLClientMessage{} }
func (m *RGQLClientMessage) String() string            { return proto.CompactTextString(m) }
func (*RGQLClientMessage) ProtoMessage()               {}
func (*RGQLClientMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RGQLClientMessage) GetMutateTree() *RGQLTreeMutation {
	if m != nil {
		return m.MutateTree
	}
	return nil
}

func (m *RGQLClientMessage) GetMutateField() *RGQLFieldMutation {
	if m != nil {
		return m.MutateField
	}
	return nil
}

type RGQLTreeMutation struct {
	// ID of the node we are operating on.
	NodeId    uint32                            `protobuf:"varint,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Operation RGQLTreeMutation_SubtreeOperation `protobuf:"varint,2,opt,name=operation,enum=rgraphql.RGQLTreeMutation_SubtreeOperation" json:"operation,omitempty"`
	// The new node tree to add, if we are adding a child.
	Node *RGQLQueryTreeNode `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
}

func (m *RGQLTreeMutation) Reset()                    { *m = RGQLTreeMutation{} }
func (m *RGQLTreeMutation) String() string            { return proto.CompactTextString(m) }
func (*RGQLTreeMutation) ProtoMessage()               {}
func (*RGQLTreeMutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RGQLTreeMutation) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *RGQLTreeMutation) GetOperation() RGQLTreeMutation_SubtreeOperation {
	if m != nil {
		return m.Operation
	}
	return RGQLTreeMutation_SUBTREE_ADD_CHILD
}

func (m *RGQLTreeMutation) GetNode() *RGQLQueryTreeNode {
	if m != nil {
		return m.Node
	}
	return nil
}

type RGQLFieldMutation struct {
	// The node we are operating on.
	NodeId uint32 `protobuf:"varint,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
}

func (m *RGQLFieldMutation) Reset()                    { *m = RGQLFieldMutation{} }
func (m *RGQLFieldMutation) String() string            { return proto.CompactTextString(m) }
func (*RGQLFieldMutation) ProtoMessage()               {}
func (*RGQLFieldMutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RGQLFieldMutation) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

type RGQLServerMessage struct {
	MutateValue *RGQLValueMutation `protobuf:"bytes,1,opt,name=mutate_value,json=mutateValue" json:"mutate_value,omitempty"`
}

func (m *RGQLServerMessage) Reset()                    { *m = RGQLServerMessage{} }
func (m *RGQLServerMessage) String() string            { return proto.CompactTextString(m) }
func (*RGQLServerMessage) ProtoMessage()               {}
func (*RGQLServerMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RGQLServerMessage) GetMutateValue() *RGQLValueMutation {
	if m != nil {
		return m.MutateValue
	}
	return nil
}

type RGQLValueMutation struct {
	// The ID of the resolver execution (the value identifier).
	ValueNodeId uint32 `protobuf:"varint,1,opt,name=value_node_id,json=valueNodeId" json:"value_node_id,omitempty"`
	// The ID of the parent resolver execution (the parent value identifier).
	ParentValueNodeId uint32 `protobuf:"varint,2,opt,name=parent_value_node_id,json=parentValueNodeId" json:"parent_value_node_id,omitempty"`
	// The ID of the query node (the query identifier).
	QueryNodeId uint32 `protobuf:"varint,3,opt,name=query_node_id,json=queryNodeId" json:"query_node_id,omitempty"`
	// The operation on the value.
	Operation RGQLValueMutation_ValueOperation `protobuf:"varint,4,opt,name=operation,enum=rgraphql.RGQLValueMutation_ValueOperation" json:"operation,omitempty"`
	// The actual value itself
	ValueJson string `protobuf:"bytes,5,opt,name=value_json,json=valueJson" json:"value_json,omitempty"`
}

func (m *RGQLValueMutation) Reset()                    { *m = RGQLValueMutation{} }
func (m *RGQLValueMutation) String() string            { return proto.CompactTextString(m) }
func (*RGQLValueMutation) ProtoMessage()               {}
func (*RGQLValueMutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RGQLValueMutation) GetValueNodeId() uint32 {
	if m != nil {
		return m.ValueNodeId
	}
	return 0
}

func (m *RGQLValueMutation) GetParentValueNodeId() uint32 {
	if m != nil {
		return m.ParentValueNodeId
	}
	return 0
}

func (m *RGQLValueMutation) GetQueryNodeId() uint32 {
	if m != nil {
		return m.QueryNodeId
	}
	return 0
}

func (m *RGQLValueMutation) GetOperation() RGQLValueMutation_ValueOperation {
	if m != nil {
		return m.Operation
	}
	return RGQLValueMutation_VALUE_SET
}

func (m *RGQLValueMutation) GetValueJson() string {
	if m != nil {
		return m.ValueJson
	}
	return ""
}

func init() {
	proto.RegisterType((*RGQLQueryFieldDirective)(nil), "rgraphql.RGQLQueryFieldDirective")
	proto.RegisterType((*RGQLQueryTreeNode)(nil), "rgraphql.RGQLQueryTreeNode")
	proto.RegisterType((*FieldArgument)(nil), "rgraphql.FieldArgument")
	proto.RegisterType((*ASTValue)(nil), "rgraphql.ASTValue")
	proto.RegisterType((*ASTValue_ASTObjectField)(nil), "rgraphql.ASTValue.ASTObjectField")
	proto.RegisterType((*RGQLClientMessage)(nil), "rgraphql.RGQLClientMessage")
	proto.RegisterType((*RGQLTreeMutation)(nil), "rgraphql.RGQLTreeMutation")
	proto.RegisterType((*RGQLFieldMutation)(nil), "rgraphql.RGQLFieldMutation")
	proto.RegisterType((*RGQLServerMessage)(nil), "rgraphql.RGQLServerMessage")
	proto.RegisterType((*RGQLValueMutation)(nil), "rgraphql.RGQLValueMutation")
	proto.RegisterEnum("rgraphql.ASTValue_ASTValueKind", ASTValue_ASTValueKind_name, ASTValue_ASTValueKind_value)
	proto.RegisterEnum("rgraphql.RGQLTreeMutation_SubtreeOperation", RGQLTreeMutation_SubtreeOperation_name, RGQLTreeMutation_SubtreeOperation_value)
	proto.RegisterEnum("rgraphql.RGQLValueMutation_ValueOperation", RGQLValueMutation_ValueOperation_name, RGQLValueMutation_ValueOperation_value)
}

func init() { proto.RegisterFile("src/rgraphql.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 818 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x7f, 0x64, 0x8b, 0x43, 0x49, 0x59, 0x6f, 0x53, 0x58, 0x48, 0x60, 0x44, 0xe1, 0x49,
	0x68, 0x0a, 0x1b, 0x55, 0x0e, 0x3d, 0x14, 0x6d, 0x21, 0x5b, 0x74, 0xac, 0x94, 0x96, 0x9a, 0x25,
	0xe5, 0x43, 0x2f, 0x04, 0x6d, 0x6e, 0x14, 0x26, 0x34, 0xe9, 0x2c, 0x29, 0x03, 0x79, 0x8b, 0x3e,
	0x4b, 0xef, 0x7d, 0x80, 0xbe, 0x48, 0x81, 0xbe, 0x45, 0xc1, 0x59, 0x52, 0x34, 0x55, 0xab, 0x6d,
	0x6e, 0xbb, 0xdf, 0x7e, 0xdf, 0x37, 0xb3, 0x33, 0x3b, 0x24, 0xd0, 0x4c, 0x5c, 0x1f, 0x8b, 0xa5,
	0x08, 0x6e, 0xdf, 0x7d, 0x8c, 0x8f, 0x6e, 0x45, 0x9a, 0xa7, 0xb4, 0x5d, 0xed, 0xad, 0x5f, 0xe0,
	0x80, 0xbd, 0x7a, 0xe3, 0xbc, 0x59, 0x71, 0xf1, 0xe9, 0x2c, 0xe2, 0x71, 0x38, 0x89, 0x04, 0xbf,
	0xce, 0xa3, 0x3b, 0x4e, 0x29, 0xe8, 0x49, 0x70, 0xc3, 0xfb, 0xca, 0x40, 0x19, 0x1a, 0x0c, 0xd7,
	0xf4, 0x05, 0xe8, 0x81, 0x58, 0x66, 0x7d, 0x75, 0xa0, 0x0d, 0xcd, 0xd1, 0xc1, 0xd1, 0xda, 0x17,
	0xb5, 0x63, 0xb1, 0x5c, 0xdd, 0xf0, 0x24, 0x67, 0x48, 0xb2, 0xfe, 0x52, 0x60, 0x7f, 0x6d, 0xee,
	0x09, 0xce, 0x67, 0x69, 0xc8, 0x69, 0x0f, 0xd4, 0x28, 0x44, 0xd3, 0x2e, 0x53, 0xa3, 0x90, 0x1e,
	0x02, 0xbc, 0x2d, 0xc4, 0x3e, 0x06, 0x53, 0x31, 0x98, 0x81, 0xc8, 0xec, 0x7e, 0x44, 0xed, 0x7f,
	0x44, 0xa4, 0x3f, 0x82, 0x11, 0x56, 0xf9, 0xf7, 0x75, 0x54, 0x3c, 0xaf, 0x15, 0x5b, 0x2e, 0xca,
	0x6a, 0x0d, 0xfd, 0x16, 0xda, 0xd7, 0xef, 0xa2, 0x38, 0x14, 0x3c, 0xe9, 0xb7, 0x50, 0xff, 0xf4,
	0x01, 0x7d, 0x75, 0x17, 0xb6, 0x26, 0x5b, 0x17, 0xd0, 0x6d, 0x24, 0xf4, 0x60, 0xf5, 0x86, 0xd0,
	0xba, 0x0b, 0xe2, 0x95, 0xbc, 0xa5, 0x39, 0xa2, 0xb5, 0xf5, 0xd8, 0xf5, 0x2e, 0x8b, 0x13, 0x26,
	0x09, 0xd6, 0xef, 0x3a, 0xb4, 0x2b, 0x8c, 0x3e, 0x87, 0x4e, 0x96, 0x8b, 0x28, 0x59, 0xfa, 0x52,
	0x2d, 0x2d, 0x4d, 0x89, 0x49, 0xca, 0x37, 0x00, 0x71, 0x94, 0xe5, 0x7e, 0x65, 0xaf, 0x6d, 0xb1,
	0x37, 0x0a, 0x96, 0x94, 0x3c, 0x05, 0x23, 0x4a, 0x2a, 0x85, 0x36, 0x50, 0x86, 0x2d, 0xd6, 0x8e,
	0x92, 0xf2, 0xf0, 0x19, 0x98, 0x6f, 0xe3, 0x34, 0xa8, 0x8e, 0xf5, 0x81, 0x32, 0x54, 0x19, 0x20,
	0x24, 0x09, 0x87, 0x00, 0x57, 0x69, 0x1a, 0x97, 0xe7, 0xad, 0x81, 0x32, 0x6c, 0x33, 0xa3, 0x40,
	0xe4, 0xf1, 0x19, 0x74, 0xd3, 0xab, 0xf7, 0xfc, 0x3a, 0xf7, 0xb1, 0x93, 0x59, 0x7f, 0x77, 0xb3,
	0x19, 0x55, 0x4a, 0xc5, 0x62, 0x8e, 0x54, 0xac, 0x1f, 0xeb, 0xa4, 0xf5, 0x26, 0xa3, 0x2f, 0x41,
	0xff, 0x10, 0x25, 0x61, 0x7f, 0x6f, 0xa0, 0x0c, 0x7b, 0xa3, 0x67, 0x0f, 0xcb, 0x71, 0xf1, 0x53,
	0x94, 0x84, 0x0c, 0xc9, 0x4f, 0x1c, 0xe8, 0x35, 0x4d, 0x29, 0x01, 0xed, 0x03, 0xff, 0x54, 0x16,
	0xae, 0x58, 0x7e, 0x46, 0x2b, 0x7e, 0x53, 0xa0, 0x73, 0x3f, 0x08, 0xa5, 0x68, 0xef, 0x5f, 0x8e,
	0x9d, 0x85, 0xed, 0xcf, 0x16, 0x8e, 0x43, 0x76, 0xe8, 0x63, 0x20, 0x35, 0xe6, 0x7a, 0x6c, 0x3a,
	0x7b, 0x45, 0x94, 0x26, 0xd3, 0x9e, 0x2d, 0x2e, 0x88, 0x4a, 0xf7, 0xa1, 0x5b, 0x63, 0xd3, 0x99,
	0x47, 0x34, 0xfa, 0x05, 0x3c, 0xaa, 0xa1, 0x33, 0x67, 0x3e, 0xf6, 0x88, 0xde, 0xd4, 0x9e, 0xcc,
	0xe7, 0x0e, 0x69, 0x35, 0x31, 0x67, 0xea, 0x7a, 0x64, 0xb7, 0x19, 0x79, 0x7e, 0xf2, 0xda, 0x3e,
	0xf5, 0xc8, 0x9e, 0xf5, 0x6b, 0x39, 0x7a, 0xa7, 0x71, 0xc4, 0x93, 0xfc, 0x82, 0x67, 0x59, 0xb0,
	0xe4, 0xf4, 0x3b, 0x30, 0x6f, 0x56, 0x79, 0x90, 0x73, 0x3f, 0x17, 0x5c, 0xbe, 0x23, 0x73, 0xf4,
	0xa4, 0xf9, 0xc0, 0x8b, 0xb7, 0x7d, 0x51, 0x90, 0xa2, 0x34, 0x61, 0x20, 0xe9, 0x05, 0x46, 0x7f,
	0x80, 0x4e, 0x29, 0xc6, 0x96, 0x96, 0x85, 0xdb, 0x18, 0x0f, 0x2c, 0xf7, 0x5a, 0x5e, 0x46, 0x43,
	0xd0, 0xfa, 0x53, 0x01, 0xb2, 0x19, 0x80, 0x1e, 0xc0, 0x5e, 0x92, 0x86, 0xdc, 0x5f, 0x7f, 0x11,
	0x76, 0x8b, 0xed, 0x34, 0xa4, 0x53, 0x30, 0xd2, 0x5b, 0x2e, 0x90, 0x85, 0xa1, 0x7a, 0xa3, 0x17,
	0xdb, 0x13, 0x3d, 0x72, 0x57, 0x57, 0xc5, 0x95, 0xe6, 0x95, 0x84, 0xd5, 0x6a, 0x7a, 0x0c, 0x7a,
	0x61, 0x8a, 0x6f, 0xfc, 0x3f, 0xe6, 0x19, 0x89, 0xd6, 0xf7, 0x40, 0x36, 0xfd, 0xe8, 0x97, 0xb0,
	0xef, 0x2e, 0x4e, 0x3c, 0x66, 0xdb, 0xfe, 0x78, 0x32, 0xf1, 0x4f, 0xcf, 0xa7, 0xce, 0x84, 0xec,
	0x14, 0x1d, 0xa9, 0xe0, 0x89, 0xed, 0xd8, 0x9e, 0x4d, 0x14, 0xeb, 0x6b, 0x59, 0xfa, 0x46, 0x29,
	0xb6, 0x5e, 0xd4, 0x72, 0x25, 0xdb, 0xe5, 0xe2, 0x8e, 0x8b, 0xaa, 0x51, 0x75, 0xad, 0xeb, 0x89,
	0xff, 0x47, 0xea, 0xf8, 0x22, 0x37, 0x6b, 0x8d, 0xa0, 0xf5, 0x87, 0x2a, 0x5d, 0x1b, 0x14, 0x6a,
	0x41, 0x17, 0xed, 0xfc, 0x66, 0x26, 0x26, 0x82, 0x33, 0x59, 0xf7, 0x63, 0x78, 0x7c, 0x1b, 0x08,
	0x5e, 0x7d, 0x18, 0xd6, 0x54, 0x15, 0xa9, 0xfb, 0xf2, 0xec, 0xf2, 0x9e, 0xc0, 0x82, 0xee, 0xc7,
	0xa2, 0x86, 0x6b, 0xa6, 0x26, 0x4d, 0x11, 0x2c, 0x39, 0xe7, 0xf7, 0x9b, 0xa9, 0x63, 0x33, 0xbf,
	0xfa, 0x97, 0xbb, 0x1c, 0xe1, 0xee, 0xc1, 0x5e, 0x1e, 0x02, 0xc8, 0xbc, 0xde, 0x67, 0x69, 0x82,
	0x9f, 0x1d, 0x83, 0x19, 0x88, 0xbc, 0xce, 0xd2, 0xc4, 0x62, 0xd0, 0x6b, 0x6a, 0x69, 0x17, 0x8c,
	0x72, 0x28, 0x6d, 0x8f, 0xec, 0xd0, 0x1e, 0x80, 0xdc, 0xfe, 0xbc, 0x70, 0xcf, 0x89, 0x42, 0x1f,
	0x81, 0x59, 0x4e, 0x27, 0x63, 0x73, 0x46, 0x54, 0x4a, 0xa0, 0x23, 0x81, 0xb2, 0x9d, 0xda, 0xd5,
	0x2e, 0xfe, 0x32, 0x5f, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x23, 0xbc, 0xe3, 0x48, 0x07,
	0x00, 0x00,
}
