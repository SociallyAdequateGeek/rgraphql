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
	ASTVariable
	RGQLClientMessage
	RGQLTreeMutation
	RGQLServerMessage
	RGQLQueryError
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
	// Set/create the value of this resolver.
	RGQLValueMutation_VALUE_SET RGQLValueMutation_ValueOperation = 0
	// Error this resolver (the value will enter the errored state).
	RGQLValueMutation_VALUE_ERROR RGQLValueMutation_ValueOperation = 1
	// Delete this resolver
	RGQLValueMutation_VALUE_DELETE RGQLValueMutation_ValueOperation = 2
)

var RGQLValueMutation_ValueOperation_name = map[int32]string{
	0: "VALUE_SET",
	1: "VALUE_ERROR",
	2: "VALUE_DELETE",
}
var RGQLValueMutation_ValueOperation_value = map[string]int32{
	"VALUE_SET":    0,
	"VALUE_ERROR":  1,
	"VALUE_DELETE": 2,
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
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	VariableId uint32 `protobuf:"varint,2,opt,name=variable_id,json=variableId" json:"variable_id,omitempty"`
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

func (m *FieldArgument) GetVariableId() uint32 {
	if m != nil {
		return m.VariableId
	}
	return 0
}

type ASTVariable struct {
	Id        uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	JsonValue string `protobuf:"bytes,2,opt,name=json_value,json=jsonValue" json:"json_value,omitempty"`
}

func (m *ASTVariable) Reset()                    { *m = ASTVariable{} }
func (m *ASTVariable) String() string            { return proto.CompactTextString(m) }
func (*ASTVariable) ProtoMessage()               {}
func (*ASTVariable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ASTVariable) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ASTVariable) GetJsonValue() string {
	if m != nil {
		return m.JsonValue
	}
	return ""
}

// Messages
type RGQLClientMessage struct {
	MutateTree *RGQLTreeMutation `protobuf:"bytes,1,opt,name=mutate_tree,json=mutateTree" json:"mutate_tree,omitempty"`
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

type RGQLTreeMutation struct {
	// All node mutations in this step.
	NodeMutation []*RGQLTreeMutation_NodeMutation `protobuf:"bytes,1,rep,name=node_mutation,json=nodeMutation" json:"node_mutation,omitempty"`
	// Any new variables.
	Variables []*ASTVariable `protobuf:"bytes,2,rep,name=variables" json:"variables,omitempty"`
}

func (m *RGQLTreeMutation) Reset()                    { *m = RGQLTreeMutation{} }
func (m *RGQLTreeMutation) String() string            { return proto.CompactTextString(m) }
func (*RGQLTreeMutation) ProtoMessage()               {}
func (*RGQLTreeMutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RGQLTreeMutation) GetNodeMutation() []*RGQLTreeMutation_NodeMutation {
	if m != nil {
		return m.NodeMutation
	}
	return nil
}

func (m *RGQLTreeMutation) GetVariables() []*ASTVariable {
	if m != nil {
		return m.Variables
	}
	return nil
}

type RGQLTreeMutation_NodeMutation struct {
	// ID of the node we are operating on.
	NodeId uint32 `protobuf:"varint,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	// Operation we are taking.
	Operation RGQLTreeMutation_SubtreeOperation `protobuf:"varint,2,opt,name=operation,enum=rgraphql.RGQLTreeMutation_SubtreeOperation" json:"operation,omitempty"`
	// The new node tree to add, if we are adding a child.
	Node *RGQLQueryTreeNode `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
}

func (m *RGQLTreeMutation_NodeMutation) Reset()         { *m = RGQLTreeMutation_NodeMutation{} }
func (m *RGQLTreeMutation_NodeMutation) String() string { return proto.CompactTextString(m) }
func (*RGQLTreeMutation_NodeMutation) ProtoMessage()    {}
func (*RGQLTreeMutation_NodeMutation) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

func (m *RGQLTreeMutation_NodeMutation) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *RGQLTreeMutation_NodeMutation) GetOperation() RGQLTreeMutation_SubtreeOperation {
	if m != nil {
		return m.Operation
	}
	return RGQLTreeMutation_SUBTREE_ADD_CHILD
}

func (m *RGQLTreeMutation_NodeMutation) GetNode() *RGQLQueryTreeNode {
	if m != nil {
		return m.Node
	}
	return nil
}

type RGQLServerMessage struct {
	MutateValue *RGQLValueMutation `protobuf:"bytes,1,opt,name=mutate_value,json=mutateValue" json:"mutate_value,omitempty"`
	QueryError  *RGQLQueryError    `protobuf:"bytes,2,opt,name=query_error,json=queryError" json:"query_error,omitempty"`
}

func (m *RGQLServerMessage) Reset()                    { *m = RGQLServerMessage{} }
func (m *RGQLServerMessage) String() string            { return proto.CompactTextString(m) }
func (*RGQLServerMessage) ProtoMessage()               {}
func (*RGQLServerMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RGQLServerMessage) GetMutateValue() *RGQLValueMutation {
	if m != nil {
		return m.MutateValue
	}
	return nil
}

func (m *RGQLServerMessage) GetQueryError() *RGQLQueryError {
	if m != nil {
		return m.QueryError
	}
	return nil
}

// Communicating a failure in the input query.
type RGQLQueryError struct {
	QueryNodeId uint32 `protobuf:"varint,1,opt,name=query_node_id,json=queryNodeId" json:"query_node_id,omitempty"`
	ErrorJson   string `protobuf:"bytes,2,opt,name=error_json,json=errorJson" json:"error_json,omitempty"`
}

func (m *RGQLQueryError) Reset()                    { *m = RGQLQueryError{} }
func (m *RGQLQueryError) String() string            { return proto.CompactTextString(m) }
func (*RGQLQueryError) ProtoMessage()               {}
func (*RGQLQueryError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RGQLQueryError) GetQueryNodeId() uint32 {
	if m != nil {
		return m.QueryNodeId
	}
	return 0
}

func (m *RGQLQueryError) GetErrorJson() string {
	if m != nil {
		return m.ErrorJson
	}
	return ""
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
	// Do we have a value to set?
	HasValue bool `protobuf:"varint,6,opt,name=has_value,json=hasValue" json:"has_value,omitempty"`
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

func (m *RGQLValueMutation) GetHasValue() bool {
	if m != nil {
		return m.HasValue
	}
	return false
}

func init() {
	proto.RegisterType((*RGQLQueryFieldDirective)(nil), "rgraphql.RGQLQueryFieldDirective")
	proto.RegisterType((*RGQLQueryTreeNode)(nil), "rgraphql.RGQLQueryTreeNode")
	proto.RegisterType((*FieldArgument)(nil), "rgraphql.FieldArgument")
	proto.RegisterType((*ASTVariable)(nil), "rgraphql.ASTVariable")
	proto.RegisterType((*RGQLClientMessage)(nil), "rgraphql.RGQLClientMessage")
	proto.RegisterType((*RGQLTreeMutation)(nil), "rgraphql.RGQLTreeMutation")
	proto.RegisterType((*RGQLTreeMutation_NodeMutation)(nil), "rgraphql.RGQLTreeMutation.NodeMutation")
	proto.RegisterType((*RGQLServerMessage)(nil), "rgraphql.RGQLServerMessage")
	proto.RegisterType((*RGQLQueryError)(nil), "rgraphql.RGQLQueryError")
	proto.RegisterType((*RGQLValueMutation)(nil), "rgraphql.RGQLValueMutation")
	proto.RegisterEnum("rgraphql.RGQLTreeMutation_SubtreeOperation", RGQLTreeMutation_SubtreeOperation_name, RGQLTreeMutation_SubtreeOperation_value)
	proto.RegisterEnum("rgraphql.RGQLValueMutation_ValueOperation", RGQLValueMutation_ValueOperation_name, RGQLValueMutation_ValueOperation_value)
}

func init() { proto.RegisterFile("src/rgraphql.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0xc5, 0x4e, 0xe0, 0x8b, 0xaf, 0x93, 0x7c, 0x66, 0x54, 0x84, 0x05, 0x42, 0xa5, 0xde, 0x14,
	0x15, 0x29, 0x48, 0x61, 0x51, 0x55, 0xfd, 0x53, 0x20, 0x6e, 0x49, 0x15, 0xa0, 0x4c, 0x42, 0x16,
	0xdd, 0x58, 0x06, 0x4f, 0x13, 0x57, 0x89, 0x1d, 0xc6, 0x4e, 0xa4, 0xbe, 0x44, 0x5f, 0xa2, 0xef,
	0xd4, 0x77, 0xe8, 0x2b, 0x74, 0x55, 0xcd, 0x1d, 0xdb, 0x89, 0xc1, 0x44, 0xdd, 0x65, 0x8e, 0xef,
	0x3d, 0xf7, 0xdc, 0x33, 0x67, 0x02, 0x24, 0xe2, 0xb7, 0x47, 0x7c, 0xc8, 0xdd, 0xe9, 0xe8, 0x6e,
	0xdc, 0x98, 0xf2, 0x30, 0x0e, 0x49, 0x25, 0x3d, 0x5b, 0x5f, 0x60, 0x9b, 0x7e, 0xbc, 0xea, 0x5e,
	0xcd, 0x18, 0xff, 0xfe, 0xc1, 0x67, 0x63, 0xaf, 0xed, 0x73, 0x76, 0x1b, 0xfb, 0x73, 0x46, 0x08,
	0x94, 0x03, 0x77, 0xc2, 0x4c, 0x65, 0x5f, 0x39, 0xd0, 0x28, 0xfe, 0x26, 0x87, 0x50, 0x76, 0xf9,
	0x30, 0x32, 0xd5, 0xfd, 0xd2, 0x81, 0xde, 0xdc, 0x6e, 0x64, 0xbc, 0xd8, 0xdb, 0xe2, 0xc3, 0xd9,
	0x84, 0x05, 0x31, 0xc5, 0x22, 0xeb, 0xb7, 0x02, 0x9b, 0x19, 0x79, 0x9f, 0x33, 0x76, 0x11, 0x7a,
	0x8c, 0xd4, 0x41, 0xf5, 0x3d, 0x24, 0xad, 0x51, 0xd5, 0xf7, 0xc8, 0x1e, 0xc0, 0x57, 0xd1, 0xec,
	0xe0, 0x30, 0x15, 0x87, 0x69, 0x88, 0x5c, 0x2c, 0x4f, 0x2c, 0xfd, 0xc3, 0x44, 0xf2, 0x1e, 0x34,
	0x2f, 0xd5, 0x6f, 0x96, 0xb1, 0xe3, 0xd9, 0xa2, 0xe3, 0x91, 0x45, 0xe9, 0xa2, 0x87, 0xbc, 0x84,
	0xca, 0xed, 0xc8, 0x1f, 0x7b, 0x9c, 0x05, 0xe6, 0x3a, 0xf6, 0xef, 0x16, 0xf4, 0xa7, 0xbb, 0xd0,
	0xac, 0xd8, 0x6a, 0x43, 0x2d, 0x27, 0xa8, 0xd0, 0xbd, 0xa7, 0xa0, 0xcf, 0x5d, 0xee, 0xbb, 0x37,
	0x63, 0xe6, 0xf8, 0x1e, 0xee, 0x5a, 0xa3, 0x90, 0x42, 0x1d, 0xcf, 0x7a, 0x03, 0x7a, 0xab, 0xd7,
	0x1f, 0x24, 0x40, 0x91, 0x55, 0xdf, 0xa2, 0x30, 0x70, 0xe6, 0xee, 0x78, 0x96, 0x59, 0x25, 0x90,
	0x81, 0x00, 0xac, 0xcf, 0xd2, 0xee, 0xd3, 0xb1, 0xcf, 0x82, 0xf8, 0x9c, 0x45, 0x91, 0x3b, 0x64,
	0xe4, 0x35, 0xe8, 0x93, 0x59, 0xec, 0xc6, 0xcc, 0x89, 0x39, 0x93, 0x72, 0xf4, 0xe6, 0x4e, 0x7e,
	0x29, 0xb1, 0xcf, 0xb9, 0x28, 0xf2, 0xc3, 0x80, 0x82, 0x2c, 0x17, 0x98, 0xf5, 0x47, 0x05, 0xe3,
	0x7e, 0x01, 0xe9, 0x42, 0x2d, 0x08, 0x3d, 0xe6, 0x4c, 0x12, 0xc0, 0x54, 0xd0, 0xa8, 0xe7, 0x8f,
	0x73, 0x36, 0x84, 0x59, 0xd9, 0x80, 0x6a, 0xb0, 0x74, 0x22, 0xc7, 0xa0, 0xa5, 0x06, 0xa4, 0xb1,
	0xda, 0x5a, 0x30, 0x2d, 0xb9, 0x41, 0x17, 0x75, 0x3b, 0x3f, 0x15, 0xa8, 0x2e, 0x73, 0x92, 0x6d,
	0xf8, 0x0f, 0x35, 0x65, 0x76, 0x6d, 0x88, 0x63, 0xc7, 0x23, 0x1d, 0xd0, 0xc2, 0x29, 0xe3, 0x52,
	0xa8, 0x70, 0xac, 0xde, 0x3c, 0x5c, 0x21, 0xb4, 0x37, 0xbb, 0x11, 0x36, 0x5d, 0xa6, 0x2d, 0x74,
	0xd1, 0x4d, 0x8e, 0xa0, 0x2c, 0x48, 0xcd, 0x12, 0x5a, 0xb8, 0x32, 0x17, 0x58, 0x68, 0xbd, 0x05,
	0xe3, 0x3e, 0x1f, 0xd9, 0x82, 0xcd, 0xde, 0xf5, 0x49, 0x9f, 0xda, 0xb6, 0xd3, 0x6a, 0xb7, 0x9d,
	0xd3, 0xb3, 0x4e, 0xb7, 0x6d, 0xac, 0x11, 0x02, 0xf5, 0x14, 0x6e, 0xdb, 0x5d, 0xbb, 0x6f, 0x1b,
	0x8a, 0xf5, 0x23, 0x79, 0x3e, 0x3d, 0xc6, 0xe7, 0x8c, 0xa7, 0xf7, 0xf9, 0x0e, 0xaa, 0xc9, 0x7d,
	0xca, 0x14, 0x28, 0x45, 0x6a, 0x30, 0x0f, 0x99, 0xe1, 0x49, 0x00, 0x10, 0x24, 0xaf, 0x40, 0xbf,
	0x13, 0x5a, 0x1d, 0xc6, 0x79, 0xc8, 0xd1, 0x12, 0xbd, 0x69, 0x16, 0x2c, 0x63, 0x8b, 0xef, 0x14,
	0xee, 0xb2, 0xdf, 0x56, 0x0f, 0xea, 0xf9, 0xaf, 0xc4, 0x82, 0x9a, 0x24, 0xcb, 0x9b, 0x2f, 0x27,
	0x5c, 0xc8, 0x1b, 0xd8, 0x03, 0xc0, 0x51, 0x8e, 0x08, 0x6a, 0x1a, 0x5a, 0x44, 0x3e, 0x45, 0x61,
	0x60, 0xfd, 0x52, 0xe5, 0x96, 0x39, 0xc9, 0x82, 0x18, 0xd7, 0xbb, 0x4f, 0x8c, 0x60, 0x42, 0x7c,
	0x04, 0x4f, 0xa6, 0x2e, 0x67, 0x41, 0xec, 0xe4, 0x4b, 0xe5, 0xb3, 0xda, 0x94, 0xdf, 0x06, 0x4b,
	0x0d, 0x0f, 0xd4, 0x96, 0x1e, 0xaa, 0x3d, 0x5b, 0xce, 0x4b, 0x19, 0xf3, 0xf2, 0x62, 0x85, 0xb7,
	0x0d, 0x3c, 0x15, 0xc6, 0x65, 0x0f, 0x40, 0xea, 0xc2, 0xbd, 0xd7, 0xe5, 0xde, 0x88, 0x88, 0xbd,
	0xc9, 0x2e, 0x68, 0x23, 0x37, 0x4a, 0x2e, 0x71, 0x63, 0x5f, 0x39, 0xa8, 0xd0, 0xca, 0xc8, 0x8d,
	0xe4, 0x4b, 0x3e, 0x81, 0x7a, 0x9e, 0x98, 0xd4, 0x40, 0x1b, 0xb4, 0xba, 0xd7, 0xb6, 0xd3, 0xb3,
	0xfb, 0xc6, 0x1a, 0xf9, 0x1f, 0x74, 0x79, 0xb4, 0x29, 0xbd, 0xa4, 0x86, 0x42, 0x0c, 0xa8, 0x4a,
	0x20, 0x89, 0x8f, 0x7a, 0xb3, 0x81, 0x7f, 0xf5, 0xc7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x98,
	0x46, 0x25, 0x0c, 0x00, 0x06, 0x00, 0x00,
}
