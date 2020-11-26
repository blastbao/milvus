// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_msg.proto

package servicepb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	commonpb "github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	schemapb "github.com/zilliztech/milvus-distributed/internal/proto/schemapb"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//*
// @brief Placeholder value types
type PlaceholderType int32

const (
	PlaceholderType_NONE          PlaceholderType = 0
	PlaceholderType_VECTOR_BINARY PlaceholderType = 100
	PlaceholderType_VECTOR_FLOAT  PlaceholderType = 101
)

var PlaceholderType_name = map[int32]string{
	0:   "NONE",
	100: "VECTOR_BINARY",
	101: "VECTOR_FLOAT",
}

var PlaceholderType_value = map[string]int32{
	"NONE":          0,
	"VECTOR_BINARY": 100,
	"VECTOR_FLOAT":  101,
}

func (x PlaceholderType) String() string {
	return proto.EnumName(PlaceholderType_name, int32(x))
}

func (PlaceholderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{0}
}

//*
// @brief Collection name
type CollectionName struct {
	CollectionName       string   `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionName) Reset()         { *m = CollectionName{} }
func (m *CollectionName) String() string { return proto.CompactTextString(m) }
func (*CollectionName) ProtoMessage()    {}
func (*CollectionName) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{0}
}

func (m *CollectionName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionName.Unmarshal(m, b)
}
func (m *CollectionName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionName.Marshal(b, m, deterministic)
}
func (m *CollectionName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionName.Merge(m, src)
}
func (m *CollectionName) XXX_Size() int {
	return xxx_messageInfo_CollectionName.Size(m)
}
func (m *CollectionName) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionName.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionName proto.InternalMessageInfo

func (m *CollectionName) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

//*
// @brief Partition name
type PartitionName struct {
	CollectionName       string   `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	Tag                  string   `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionName) Reset()         { *m = PartitionName{} }
func (m *PartitionName) String() string { return proto.CompactTextString(m) }
func (*PartitionName) ProtoMessage()    {}
func (*PartitionName) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{1}
}

func (m *PartitionName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionName.Unmarshal(m, b)
}
func (m *PartitionName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionName.Marshal(b, m, deterministic)
}
func (m *PartitionName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionName.Merge(m, src)
}
func (m *PartitionName) XXX_Size() int {
	return xxx_messageInfo_PartitionName.Size(m)
}
func (m *PartitionName) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionName.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionName proto.InternalMessageInfo

func (m *PartitionName) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *PartitionName) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

//*
// @brief Row batch for Insert call
type RowBatch struct {
	CollectionName       string           `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	PartitionTag         string           `protobuf:"bytes,2,opt,name=partition_tag,json=partitionTag,proto3" json:"partition_tag,omitempty"`
	RowData              []*commonpb.Blob `protobuf:"bytes,3,rep,name=row_data,json=rowData,proto3" json:"row_data,omitempty"`
	HashKeys             []int32          `protobuf:"varint,4,rep,packed,name=hash_keys,json=hashKeys,proto3" json:"hash_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RowBatch) Reset()         { *m = RowBatch{} }
func (m *RowBatch) String() string { return proto.CompactTextString(m) }
func (*RowBatch) ProtoMessage()    {}
func (*RowBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{2}
}

func (m *RowBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RowBatch.Unmarshal(m, b)
}
func (m *RowBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RowBatch.Marshal(b, m, deterministic)
}
func (m *RowBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RowBatch.Merge(m, src)
}
func (m *RowBatch) XXX_Size() int {
	return xxx_messageInfo_RowBatch.Size(m)
}
func (m *RowBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RowBatch.DiscardUnknown(m)
}

var xxx_messageInfo_RowBatch proto.InternalMessageInfo

func (m *RowBatch) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *RowBatch) GetPartitionTag() string {
	if m != nil {
		return m.PartitionTag
	}
	return ""
}

func (m *RowBatch) GetRowData() []*commonpb.Blob {
	if m != nil {
		return m.RowData
	}
	return nil
}

func (m *RowBatch) GetHashKeys() []int32 {
	if m != nil {
		return m.HashKeys
	}
	return nil
}

//*
// @brief Placeholder value in DSL
type PlaceholderValue struct {
	Tag  string          `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Type PlaceholderType `protobuf:"varint,2,opt,name=type,proto3,enum=milvus.proto.service.PlaceholderType" json:"type,omitempty"`
	// values is a 2d-array, every array contains a vector
	Values               [][]byte `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlaceholderValue) Reset()         { *m = PlaceholderValue{} }
func (m *PlaceholderValue) String() string { return proto.CompactTextString(m) }
func (*PlaceholderValue) ProtoMessage()    {}
func (*PlaceholderValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{3}
}

func (m *PlaceholderValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlaceholderValue.Unmarshal(m, b)
}
func (m *PlaceholderValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlaceholderValue.Marshal(b, m, deterministic)
}
func (m *PlaceholderValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaceholderValue.Merge(m, src)
}
func (m *PlaceholderValue) XXX_Size() int {
	return xxx_messageInfo_PlaceholderValue.Size(m)
}
func (m *PlaceholderValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaceholderValue.DiscardUnknown(m)
}

var xxx_messageInfo_PlaceholderValue proto.InternalMessageInfo

func (m *PlaceholderValue) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *PlaceholderValue) GetType() PlaceholderType {
	if m != nil {
		return m.Type
	}
	return PlaceholderType_NONE
}

func (m *PlaceholderValue) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

type PlaceholderGroup struct {
	Placeholders         []*PlaceholderValue `protobuf:"bytes,1,rep,name=placeholders,proto3" json:"placeholders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PlaceholderGroup) Reset()         { *m = PlaceholderGroup{} }
func (m *PlaceholderGroup) String() string { return proto.CompactTextString(m) }
func (*PlaceholderGroup) ProtoMessage()    {}
func (*PlaceholderGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{4}
}

func (m *PlaceholderGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlaceholderGroup.Unmarshal(m, b)
}
func (m *PlaceholderGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlaceholderGroup.Marshal(b, m, deterministic)
}
func (m *PlaceholderGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaceholderGroup.Merge(m, src)
}
func (m *PlaceholderGroup) XXX_Size() int {
	return xxx_messageInfo_PlaceholderGroup.Size(m)
}
func (m *PlaceholderGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaceholderGroup.DiscardUnknown(m)
}

var xxx_messageInfo_PlaceholderGroup proto.InternalMessageInfo

func (m *PlaceholderGroup) GetPlaceholders() []*PlaceholderValue {
	if m != nil {
		return m.Placeholders
	}
	return nil
}

//*
// @brief Query for Search call
type Query struct {
	CollectionName string   `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	PartitionTags  []string `protobuf:"bytes,2,rep,name=partition_tags,json=partitionTags,proto3" json:"partition_tags,omitempty"`
	Dsl            string   `protobuf:"bytes,3,opt,name=dsl,proto3" json:"dsl,omitempty"`
	// placeholder_group contains the serialized PlaceholderGroup
	PlaceholderGroup     []byte   `protobuf:"bytes,4,opt,name=placeholder_group,json=placeholderGroup,proto3" json:"placeholder_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{5}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *Query) GetPartitionTags() []string {
	if m != nil {
		return m.PartitionTags
	}
	return nil
}

func (m *Query) GetDsl() string {
	if m != nil {
		return m.Dsl
	}
	return ""
}

func (m *Query) GetPlaceholderGroup() []byte {
	if m != nil {
		return m.PlaceholderGroup
	}
	return nil
}

//*
// @brief String response
type StringResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Value                string           `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StringResponse) Reset()         { *m = StringResponse{} }
func (m *StringResponse) String() string { return proto.CompactTextString(m) }
func (*StringResponse) ProtoMessage()    {}
func (*StringResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{6}
}

func (m *StringResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringResponse.Unmarshal(m, b)
}
func (m *StringResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringResponse.Marshal(b, m, deterministic)
}
func (m *StringResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringResponse.Merge(m, src)
}
func (m *StringResponse) XXX_Size() int {
	return xxx_messageInfo_StringResponse.Size(m)
}
func (m *StringResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StringResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StringResponse proto.InternalMessageInfo

func (m *StringResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *StringResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

//*
// @brief Bool response
type BoolResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Value                bool             `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BoolResponse) Reset()         { *m = BoolResponse{} }
func (m *BoolResponse) String() string { return proto.CompactTextString(m) }
func (*BoolResponse) ProtoMessage()    {}
func (*BoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{7}
}

func (m *BoolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolResponse.Unmarshal(m, b)
}
func (m *BoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolResponse.Marshal(b, m, deterministic)
}
func (m *BoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolResponse.Merge(m, src)
}
func (m *BoolResponse) XXX_Size() int {
	return xxx_messageInfo_BoolResponse.Size(m)
}
func (m *BoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BoolResponse proto.InternalMessageInfo

func (m *BoolResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BoolResponse) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

//*
// @brief String list response
type StringListResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Values               []string         `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StringListResponse) Reset()         { *m = StringListResponse{} }
func (m *StringListResponse) String() string { return proto.CompactTextString(m) }
func (*StringListResponse) ProtoMessage()    {}
func (*StringListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{8}
}

func (m *StringListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringListResponse.Unmarshal(m, b)
}
func (m *StringListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringListResponse.Marshal(b, m, deterministic)
}
func (m *StringListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringListResponse.Merge(m, src)
}
func (m *StringListResponse) XXX_Size() int {
	return xxx_messageInfo_StringListResponse.Size(m)
}
func (m *StringListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StringListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StringListResponse proto.InternalMessageInfo

func (m *StringListResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *StringListResponse) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

//*
// @brief Integer list response
type IntegerListResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Values               []int64          `protobuf:"varint,2,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IntegerListResponse) Reset()         { *m = IntegerListResponse{} }
func (m *IntegerListResponse) String() string { return proto.CompactTextString(m) }
func (*IntegerListResponse) ProtoMessage()    {}
func (*IntegerListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{9}
}

func (m *IntegerListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegerListResponse.Unmarshal(m, b)
}
func (m *IntegerListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegerListResponse.Marshal(b, m, deterministic)
}
func (m *IntegerListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegerListResponse.Merge(m, src)
}
func (m *IntegerListResponse) XXX_Size() int {
	return xxx_messageInfo_IntegerListResponse.Size(m)
}
func (m *IntegerListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegerListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IntegerListResponse proto.InternalMessageInfo

func (m *IntegerListResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IntegerListResponse) GetValues() []int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

//*
// @brief Range response, [begin, end)
type IntegerRangeResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Begin                int64            `protobuf:"varint,2,opt,name=begin,proto3" json:"begin,omitempty"`
	End                  int64            `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *IntegerRangeResponse) Reset()         { *m = IntegerRangeResponse{} }
func (m *IntegerRangeResponse) String() string { return proto.CompactTextString(m) }
func (*IntegerRangeResponse) ProtoMessage()    {}
func (*IntegerRangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{10}
}

func (m *IntegerRangeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntegerRangeResponse.Unmarshal(m, b)
}
func (m *IntegerRangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntegerRangeResponse.Marshal(b, m, deterministic)
}
func (m *IntegerRangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntegerRangeResponse.Merge(m, src)
}
func (m *IntegerRangeResponse) XXX_Size() int {
	return xxx_messageInfo_IntegerRangeResponse.Size(m)
}
func (m *IntegerRangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IntegerRangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IntegerRangeResponse proto.InternalMessageInfo

func (m *IntegerRangeResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IntegerRangeResponse) GetBegin() int64 {
	if m != nil {
		return m.Begin
	}
	return 0
}

func (m *IntegerRangeResponse) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

//*
// @brief Response of DescribeCollection
type CollectionDescription struct {
	Status               *commonpb.Status           `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Schema               *schemapb.CollectionSchema `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	Statistics           []*commonpb.KeyValuePair   `protobuf:"bytes,3,rep,name=statistics,proto3" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CollectionDescription) Reset()         { *m = CollectionDescription{} }
func (m *CollectionDescription) String() string { return proto.CompactTextString(m) }
func (*CollectionDescription) ProtoMessage()    {}
func (*CollectionDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{11}
}

func (m *CollectionDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionDescription.Unmarshal(m, b)
}
func (m *CollectionDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionDescription.Marshal(b, m, deterministic)
}
func (m *CollectionDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionDescription.Merge(m, src)
}
func (m *CollectionDescription) XXX_Size() int {
	return xxx_messageInfo_CollectionDescription.Size(m)
}
func (m *CollectionDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionDescription.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionDescription proto.InternalMessageInfo

func (m *CollectionDescription) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CollectionDescription) GetSchema() *schemapb.CollectionSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *CollectionDescription) GetStatistics() []*commonpb.KeyValuePair {
	if m != nil {
		return m.Statistics
	}
	return nil
}

//*
// @brief Response of DescribePartition
type PartitionDescription struct {
	Status               *commonpb.Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Name                 *PartitionName           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Statistics           []*commonpb.KeyValuePair `protobuf:"bytes,3,rep,name=statistics,proto3" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PartitionDescription) Reset()         { *m = PartitionDescription{} }
func (m *PartitionDescription) String() string { return proto.CompactTextString(m) }
func (*PartitionDescription) ProtoMessage()    {}
func (*PartitionDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{12}
}

func (m *PartitionDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionDescription.Unmarshal(m, b)
}
func (m *PartitionDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionDescription.Marshal(b, m, deterministic)
}
func (m *PartitionDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionDescription.Merge(m, src)
}
func (m *PartitionDescription) XXX_Size() int {
	return xxx_messageInfo_PartitionDescription.Size(m)
}
func (m *PartitionDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionDescription.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionDescription proto.InternalMessageInfo

func (m *PartitionDescription) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PartitionDescription) GetName() *PartitionName {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PartitionDescription) GetStatistics() []*commonpb.KeyValuePair {
	if m != nil {
		return m.Statistics
	}
	return nil
}

//*
// @brief Response of GetSysConfig
type SysConfigResponse struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Keys                 []string         `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	Values               []string         `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SysConfigResponse) Reset()         { *m = SysConfigResponse{} }
func (m *SysConfigResponse) String() string { return proto.CompactTextString(m) }
func (*SysConfigResponse) ProtoMessage()    {}
func (*SysConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{13}
}

func (m *SysConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SysConfigResponse.Unmarshal(m, b)
}
func (m *SysConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SysConfigResponse.Marshal(b, m, deterministic)
}
func (m *SysConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SysConfigResponse.Merge(m, src)
}
func (m *SysConfigResponse) XXX_Size() int {
	return xxx_messageInfo_SysConfigResponse.Size(m)
}
func (m *SysConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SysConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SysConfigResponse proto.InternalMessageInfo

func (m *SysConfigResponse) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SysConfigResponse) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *SysConfigResponse) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

//*
// @brief Scores of a query.
//        The default value of tag is "root".
//        It corresponds to the final score of each hit.
type Score struct {
	Tag                  string    `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Values               []float32 `protobuf:"fixed32,2,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Score) Reset()         { *m = Score{} }
func (m *Score) String() string { return proto.CompactTextString(m) }
func (*Score) ProtoMessage()    {}
func (*Score) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{14}
}

func (m *Score) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Score.Unmarshal(m, b)
}
func (m *Score) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Score.Marshal(b, m, deterministic)
}
func (m *Score) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Score.Merge(m, src)
}
func (m *Score) XXX_Size() int {
	return xxx_messageInfo_Score.Size(m)
}
func (m *Score) XXX_DiscardUnknown() {
	xxx_messageInfo_Score.DiscardUnknown(m)
}

var xxx_messageInfo_Score proto.InternalMessageInfo

func (m *Score) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Score) GetValues() []float32 {
	if m != nil {
		return m.Values
	}
	return nil
}

//*
// @brief Entities hit by query
type Hits struct {
	IDs                  []int64          `protobuf:"varint,1,rep,packed,name=IDs,proto3" json:"IDs,omitempty"`
	RowData              []*commonpb.Blob `protobuf:"bytes,2,rep,name=row_data,json=rowData,proto3" json:"row_data,omitempty"`
	Scores               []*Score         `protobuf:"bytes,3,rep,name=scores,proto3" json:"scores,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Hits) Reset()         { *m = Hits{} }
func (m *Hits) String() string { return proto.CompactTextString(m) }
func (*Hits) ProtoMessage()    {}
func (*Hits) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{15}
}

func (m *Hits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hits.Unmarshal(m, b)
}
func (m *Hits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hits.Marshal(b, m, deterministic)
}
func (m *Hits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hits.Merge(m, src)
}
func (m *Hits) XXX_Size() int {
	return xxx_messageInfo_Hits.Size(m)
}
func (m *Hits) XXX_DiscardUnknown() {
	xxx_messageInfo_Hits.DiscardUnknown(m)
}

var xxx_messageInfo_Hits proto.InternalMessageInfo

func (m *Hits) GetIDs() []int64 {
	if m != nil {
		return m.IDs
	}
	return nil
}

func (m *Hits) GetRowData() []*commonpb.Blob {
	if m != nil {
		return m.RowData
	}
	return nil
}

func (m *Hits) GetScores() []*Score {
	if m != nil {
		return m.Scores
	}
	return nil
}

//*
// @brief Query result
type QueryResult struct {
	Status               *commonpb.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Hits                 []*Hits          `protobuf:"bytes,2,rep,name=hits,proto3" json:"hits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *QueryResult) Reset()         { *m = QueryResult{} }
func (m *QueryResult) String() string { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()    {}
func (*QueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b40b84dd2f74cb, []int{16}
}

func (m *QueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResult.Unmarshal(m, b)
}
func (m *QueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResult.Marshal(b, m, deterministic)
}
func (m *QueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResult.Merge(m, src)
}
func (m *QueryResult) XXX_Size() int {
	return xxx_messageInfo_QueryResult.Size(m)
}
func (m *QueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResult proto.InternalMessageInfo

func (m *QueryResult) GetStatus() *commonpb.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *QueryResult) GetHits() []*Hits {
	if m != nil {
		return m.Hits
	}
	return nil
}

func init() {
	proto.RegisterEnum("milvus.proto.service.PlaceholderType", PlaceholderType_name, PlaceholderType_value)
	proto.RegisterType((*CollectionName)(nil), "milvus.proto.service.CollectionName")
	proto.RegisterType((*PartitionName)(nil), "milvus.proto.service.PartitionName")
	proto.RegisterType((*RowBatch)(nil), "milvus.proto.service.RowBatch")
	proto.RegisterType((*PlaceholderValue)(nil), "milvus.proto.service.PlaceholderValue")
	proto.RegisterType((*PlaceholderGroup)(nil), "milvus.proto.service.PlaceholderGroup")
	proto.RegisterType((*Query)(nil), "milvus.proto.service.Query")
	proto.RegisterType((*StringResponse)(nil), "milvus.proto.service.StringResponse")
	proto.RegisterType((*BoolResponse)(nil), "milvus.proto.service.BoolResponse")
	proto.RegisterType((*StringListResponse)(nil), "milvus.proto.service.StringListResponse")
	proto.RegisterType((*IntegerListResponse)(nil), "milvus.proto.service.IntegerListResponse")
	proto.RegisterType((*IntegerRangeResponse)(nil), "milvus.proto.service.IntegerRangeResponse")
	proto.RegisterType((*CollectionDescription)(nil), "milvus.proto.service.CollectionDescription")
	proto.RegisterType((*PartitionDescription)(nil), "milvus.proto.service.PartitionDescription")
	proto.RegisterType((*SysConfigResponse)(nil), "milvus.proto.service.SysConfigResponse")
	proto.RegisterType((*Score)(nil), "milvus.proto.service.Score")
	proto.RegisterType((*Hits)(nil), "milvus.proto.service.Hits")
	proto.RegisterType((*QueryResult)(nil), "milvus.proto.service.QueryResult")
}

func init() { proto.RegisterFile("service_msg.proto", fileDescriptor_b4b40b84dd2f74cb) }

var fileDescriptor_b4b40b84dd2f74cb = []byte{
	// 790 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x8e, 0xe3, 0x34,
	0x14, 0x26, 0x4d, 0x5a, 0x3a, 0xa7, 0x69, 0xb7, 0x35, 0x05, 0x85, 0x99, 0x9b, 0x92, 0xd5, 0x42,
	0x05, 0xa2, 0x15, 0xb3, 0x48, 0x68, 0x2f, 0x90, 0x68, 0x3b, 0x0b, 0xcc, 0xee, 0xaa, 0x33, 0xb8,
	0xd5, 0x4a, 0x0b, 0x12, 0x95, 0x9b, 0x98, 0xc4, 0x22, 0x8d, 0x23, 0xdb, 0x99, 0xaa, 0x7b, 0xcb,
	0x3b, 0xf0, 0x12, 0x3c, 0x08, 0x77, 0x3c, 0x13, 0x8a, 0x93, 0xfe, 0x0d, 0x45, 0x0c, 0xd3, 0xb9,
	0xb3, 0x4f, 0x7c, 0xce, 0xf7, 0x9d, 0xcf, 0xc7, 0x5f, 0xa0, 0x25, 0xa9, 0xb8, 0x61, 0x1e, 0x9d,
	0x2d, 0x64, 0xd0, 0x4b, 0x04, 0x57, 0x1c, 0xb5, 0x17, 0x2c, 0xba, 0x49, 0x65, 0xbe, 0xeb, 0x15,
	0xdf, 0x4f, 0x6d, 0x8f, 0x2f, 0x16, 0x3c, 0xce, 0xa3, 0xa7, 0xb6, 0xf4, 0x42, 0xba, 0x20, 0xf9,
	0xce, 0x7d, 0x06, 0x8d, 0x11, 0x8f, 0x22, 0xea, 0x29, 0xc6, 0xe3, 0x31, 0x59, 0x50, 0xf4, 0x09,
	0x3c, 0xf2, 0x36, 0x91, 0x59, 0x4c, 0x16, 0xd4, 0x31, 0x3a, 0x46, 0xf7, 0x04, 0x37, 0xbc, 0xbd,
	0x83, 0xee, 0x0b, 0xa8, 0x5f, 0x13, 0xa1, 0xd8, 0xff, 0xce, 0x44, 0x4d, 0x30, 0x15, 0x09, 0x9c,
	0x92, 0xfe, 0x98, 0x2d, 0xdd, 0x3f, 0x0c, 0xa8, 0x62, 0xbe, 0x1c, 0x12, 0xe5, 0x85, 0x77, 0xaf,
	0xf3, 0x18, 0xea, 0xc9, 0x9a, 0xc1, 0x6c, 0x5b, 0xd1, 0xde, 0x04, 0xa7, 0x24, 0x40, 0x5f, 0x42,
	0x55, 0xf0, 0xe5, 0xcc, 0x27, 0x8a, 0x38, 0x66, 0xc7, 0xec, 0xd6, 0xce, 0x3f, 0xec, 0xed, 0xc9,
	0x54, 0xa8, 0x33, 0x8c, 0xf8, 0x1c, 0xbf, 0x2b, 0xf8, 0xf2, 0x82, 0x28, 0x82, 0xce, 0xe0, 0x24,
	0x24, 0x32, 0x9c, 0xfd, 0x4a, 0x57, 0xd2, 0xb1, 0x3a, 0x66, 0xb7, 0x8c, 0xab, 0x59, 0xe0, 0x25,
	0x5d, 0x49, 0x77, 0x09, 0xcd, 0xeb, 0x88, 0x78, 0x34, 0xe4, 0x91, 0x4f, 0xc5, 0x6b, 0x12, 0xa5,
	0x9b, 0x9e, 0x8c, 0x4d, 0x4f, 0xe8, 0x19, 0x58, 0x6a, 0x95, 0x50, 0x4d, 0xaa, 0x71, 0xfe, 0xa4,
	0x77, 0xe8, 0x6e, 0x7a, 0x3b, 0x75, 0xa6, 0xab, 0x84, 0x62, 0x9d, 0x82, 0x3e, 0x80, 0xca, 0x4d,
	0x56, 0x55, 0x6a, 0xc6, 0x36, 0x2e, 0x76, 0xee, 0xcf, 0x7b, 0xc0, 0xdf, 0x09, 0x9e, 0x26, 0xe8,
	0x05, 0xd8, 0xc9, 0x36, 0x26, 0x1d, 0x43, 0xf7, 0xf8, 0xf1, 0x7f, 0xc2, 0x69, 0xda, 0x78, 0x2f,
	0xd7, 0xfd, 0xdd, 0x80, 0xf2, 0x0f, 0x29, 0x15, 0xab, 0xbb, 0xdf, 0xc1, 0x13, 0x68, 0xec, 0xdd,
	0x81, 0x74, 0x4a, 0x1d, 0xb3, 0x7b, 0x82, 0xeb, 0xbb, 0x97, 0x20, 0x33, 0x79, 0x7c, 0x19, 0x39,
	0x66, 0x2e, 0x8f, 0x2f, 0x23, 0xf4, 0x19, 0xb4, 0x76, 0xb0, 0x67, 0x41, 0xd6, 0x8c, 0x63, 0x75,
	0x8c, 0xae, 0x8d, 0x9b, 0xc9, 0xad, 0x26, 0xdd, 0x9f, 0xa0, 0x31, 0x51, 0x82, 0xc5, 0x01, 0xa6,
	0x32, 0xe1, 0xb1, 0xa4, 0xe8, 0x29, 0x54, 0xa4, 0x22, 0x2a, 0x95, 0x9a, 0x57, 0xed, 0xfc, 0xec,
	0xe0, 0xa5, 0x4e, 0xf4, 0x11, 0x5c, 0x1c, 0x45, 0x6d, 0x28, 0x6b, 0x25, 0x8b, 0x41, 0xc9, 0x37,
	0xee, 0x1b, 0xb0, 0x87, 0x9c, 0x47, 0x0f, 0x58, 0xba, 0xba, 0x2e, 0x4d, 0x00, 0xe5, 0xbc, 0x5f,
	0x31, 0xa9, 0x8e, 0x03, 0xd8, 0xce, 0x44, 0x2e, 0xf0, 0x7a, 0x26, 0xe6, 0xf0, 0xde, 0x65, 0xac,
	0x68, 0x40, 0xc5, 0x43, 0x63, 0x98, 0x1b, 0x0c, 0x09, 0xed, 0x02, 0x03, 0x93, 0x38, 0xa0, 0x47,
	0x2b, 0x35, 0xa7, 0x01, 0x8b, 0xb5, 0x52, 0x26, 0xce, 0x37, 0xd9, 0x80, 0xd0, 0xd8, 0xd7, 0x03,
	0x62, 0xe2, 0x6c, 0xe9, 0xfe, 0x65, 0xc0, 0xfb, 0x5b, 0x6f, 0xba, 0xa0, 0xd2, 0x13, 0x2c, 0xc9,
	0x96, 0xf7, 0x83, 0xfd, 0x1a, 0x2a, 0xb9, 0xf3, 0x69, 0xdc, 0xda, 0x3f, 0x1e, 0x64, 0xee, 0x8a,
	0x5b, 0xc0, 0x89, 0x0e, 0xe0, 0x22, 0x09, 0x0d, 0x00, 0xb2, 0x42, 0x4c, 0x2a, 0xe6, 0xc9, 0xc2,
	0x48, 0x3e, 0x3a, 0x88, 0xfb, 0x92, 0xae, 0xf4, 0xdb, 0xba, 0x26, 0x4c, 0xe0, 0x9d, 0x24, 0xf7,
	0x4f, 0x03, 0xda, 0x1b, 0xc7, 0x3c, 0xba, 0x9f, 0xaf, 0xc0, 0xd2, 0xcf, 0x32, 0xef, 0xe6, 0xf1,
	0xbf, 0xbc, 0xf7, 0x5d, 0x83, 0xc6, 0x3a, 0xe1, 0x21, 0x3a, 0x51, 0xd0, 0x9a, 0xac, 0xe4, 0x88,
	0xc7, 0xbf, 0xb0, 0x23, 0x5f, 0x24, 0x02, 0x4b, 0x5b, 0x6c, 0x3e, 0xd3, 0x7a, 0x7d, 0xcb, 0xfd,
	0xb6, 0x93, 0xfe, 0x05, 0x94, 0x27, 0x1e, 0x17, 0x87, 0xbc, 0x76, 0x7f, 0x70, 0x4b, 0x9b, 0x94,
	0xdf, 0x0c, 0xb0, 0xbe, 0x67, 0x4a, 0xfb, 0xcf, 0xe5, 0x45, 0x6e, 0x8e, 0x26, 0xce, 0x96, 0x7b,
	0xff, 0x85, 0xd2, 0x9d, 0xff, 0x0b, 0x59, 0x93, 0x19, 0x87, 0xb5, 0x70, 0x67, 0x87, 0x75, 0xd7,
	0x3c, 0x71, 0x71, 0xd4, 0x15, 0x50, 0xd3, 0xae, 0x8a, 0xa9, 0x4c, 0x23, 0x75, 0x3f, 0xa1, 0x7a,
	0x60, 0x85, 0x4c, 0xc9, 0x82, 0xea, 0xe9, 0x61, 0xd8, 0xac, 0x55, 0xac, 0xcf, 0x7d, 0xfa, 0x0d,
	0x3c, 0xba, 0xf5, 0x6f, 0x41, 0x55, 0xb0, 0xc6, 0x57, 0xe3, 0xe7, 0xcd, 0x77, 0x50, 0x0b, 0xea,
	0xaf, 0x9f, 0x8f, 0xa6, 0x57, 0x78, 0x36, 0xbc, 0x1c, 0x0f, 0xf0, 0x9b, 0xa6, 0x8f, 0x9a, 0x60,
	0x17, 0xa1, 0x6f, 0x5f, 0x5d, 0x0d, 0xa6, 0x4d, 0x3a, 0x1c, 0xfd, 0x38, 0x08, 0x98, 0x0a, 0xd3,
	0x79, 0xc6, 0xa8, 0xff, 0x96, 0x45, 0x11, 0x7b, 0xab, 0xa8, 0x17, 0xf6, 0x73, 0xe8, 0xcf, 0x7d,
	0x26, 0x95, 0x60, 0xf3, 0x54, 0x51, 0xbf, 0xcf, 0x62, 0x45, 0x45, 0x4c, 0xa2, 0xbe, 0xe6, 0xd3,
	0x2f, 0xf8, 0x24, 0xf3, 0x79, 0x45, 0x07, 0x9e, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x8a,
	0x8b, 0xe7, 0xad, 0x08, 0x00, 0x00,
}