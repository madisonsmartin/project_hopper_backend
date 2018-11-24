// Code generated by protoc-gen-go. DO NOT EDIT.
// source: algorithm.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateAlgorithmCommand struct {
	Algorithm            *Algorithm `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	CreatedOn            int64      `protobuf:"varint,2,opt,name=createdOn,proto3" json:"createdOn,omitempty"`
	Id                   string     `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateAlgorithmCommand) Reset()         { *m = CreateAlgorithmCommand{} }
func (m *CreateAlgorithmCommand) String() string { return proto.CompactTextString(m) }
func (*CreateAlgorithmCommand) ProtoMessage()    {}
func (*CreateAlgorithmCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c381a4f1e580eed, []int{0}
}

func (m *CreateAlgorithmCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAlgorithmCommand.Unmarshal(m, b)
}
func (m *CreateAlgorithmCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAlgorithmCommand.Marshal(b, m, deterministic)
}
func (m *CreateAlgorithmCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAlgorithmCommand.Merge(m, src)
}
func (m *CreateAlgorithmCommand) XXX_Size() int {
	return xxx_messageInfo_CreateAlgorithmCommand.Size(m)
}
func (m *CreateAlgorithmCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAlgorithmCommand.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAlgorithmCommand proto.InternalMessageInfo

func (m *CreateAlgorithmCommand) GetAlgorithm() *Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return nil
}

func (m *CreateAlgorithmCommand) GetCreatedOn() int64 {
	if m != nil {
		return m.CreatedOn
	}
	return 0
}

func (m *CreateAlgorithmCommand) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AssociateFileCommand struct {
	Algorithm            *Algorithm     `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	AlgorithmFile        *AlgorithmFile `protobuf:"bytes,2,opt,name=algorithmFile,proto3" json:"algorithmFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AssociateFileCommand) Reset()         { *m = AssociateFileCommand{} }
func (m *AssociateFileCommand) String() string { return proto.CompactTextString(m) }
func (*AssociateFileCommand) ProtoMessage()    {}
func (*AssociateFileCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c381a4f1e580eed, []int{1}
}

func (m *AssociateFileCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssociateFileCommand.Unmarshal(m, b)
}
func (m *AssociateFileCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssociateFileCommand.Marshal(b, m, deterministic)
}
func (m *AssociateFileCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssociateFileCommand.Merge(m, src)
}
func (m *AssociateFileCommand) XXX_Size() int {
	return xxx_messageInfo_AssociateFileCommand.Size(m)
}
func (m *AssociateFileCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_AssociateFileCommand.DiscardUnknown(m)
}

var xxx_messageInfo_AssociateFileCommand proto.InternalMessageInfo

func (m *AssociateFileCommand) GetAlgorithm() *Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return nil
}

func (m *AssociateFileCommand) GetAlgorithmFile() *AlgorithmFile {
	if m != nil {
		return m.AlgorithmFile
	}
	return nil
}

type GetAlgorithmQuery struct {
	Algorithm            *Algorithm `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	CreatedOn            int64      `protobuf:"varint,2,opt,name=createdOn,proto3" json:"createdOn,omitempty"`
	Id                   string     `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAlgorithmQuery) Reset()         { *m = GetAlgorithmQuery{} }
func (m *GetAlgorithmQuery) String() string { return proto.CompactTextString(m) }
func (*GetAlgorithmQuery) ProtoMessage()    {}
func (*GetAlgorithmQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c381a4f1e580eed, []int{2}
}

func (m *GetAlgorithmQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAlgorithmQuery.Unmarshal(m, b)
}
func (m *GetAlgorithmQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAlgorithmQuery.Marshal(b, m, deterministic)
}
func (m *GetAlgorithmQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAlgorithmQuery.Merge(m, src)
}
func (m *GetAlgorithmQuery) XXX_Size() int {
	return xxx_messageInfo_GetAlgorithmQuery.Size(m)
}
func (m *GetAlgorithmQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAlgorithmQuery.DiscardUnknown(m)
}

var xxx_messageInfo_GetAlgorithmQuery proto.InternalMessageInfo

func (m *GetAlgorithmQuery) GetAlgorithm() *Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return nil
}

func (m *GetAlgorithmQuery) GetCreatedOn() int64 {
	if m != nil {
		return m.CreatedOn
	}
	return 0
}

func (m *GetAlgorithmQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Algorithm struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	FileIDs              []string `protobuf:"bytes,5,rep,name=fileIDs,proto3" json:"fileIDs,omitempty"`
	DatasetIDs           []string `protobuf:"bytes,6,rep,name=datasetIDs,proto3" json:"datasetIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Algorithm) Reset()         { *m = Algorithm{} }
func (m *Algorithm) String() string { return proto.CompactTextString(m) }
func (*Algorithm) ProtoMessage()    {}
func (*Algorithm) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c381a4f1e580eed, []int{3}
}

func (m *Algorithm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Algorithm.Unmarshal(m, b)
}
func (m *Algorithm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Algorithm.Marshal(b, m, deterministic)
}
func (m *Algorithm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Algorithm.Merge(m, src)
}
func (m *Algorithm) XXX_Size() int {
	return xxx_messageInfo_Algorithm.Size(m)
}
func (m *Algorithm) XXX_DiscardUnknown() {
	xxx_messageInfo_Algorithm.DiscardUnknown(m)
}

var xxx_messageInfo_Algorithm proto.InternalMessageInfo

func (m *Algorithm) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Algorithm) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Algorithm) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Algorithm) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Algorithm) GetFileIDs() []string {
	if m != nil {
		return m.FileIDs
	}
	return nil
}

func (m *Algorithm) GetDatasetIDs() []string {
	if m != nil {
		return m.DatasetIDs
	}
	return nil
}

type AlgorithmFile struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Filetype             string   `protobuf:"bytes,3,opt,name=filetype,proto3" json:"filetype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlgorithmFile) Reset()         { *m = AlgorithmFile{} }
func (m *AlgorithmFile) String() string { return proto.CompactTextString(m) }
func (*AlgorithmFile) ProtoMessage()    {}
func (*AlgorithmFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c381a4f1e580eed, []int{4}
}

func (m *AlgorithmFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlgorithmFile.Unmarshal(m, b)
}
func (m *AlgorithmFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlgorithmFile.Marshal(b, m, deterministic)
}
func (m *AlgorithmFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlgorithmFile.Merge(m, src)
}
func (m *AlgorithmFile) XXX_Size() int {
	return xxx_messageInfo_AlgorithmFile.Size(m)
}
func (m *AlgorithmFile) XXX_DiscardUnknown() {
	xxx_messageInfo_AlgorithmFile.DiscardUnknown(m)
}

var xxx_messageInfo_AlgorithmFile proto.InternalMessageInfo

func (m *AlgorithmFile) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *AlgorithmFile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AlgorithmFile) GetFiletype() string {
	if m != nil {
		return m.Filetype
	}
	return ""
}

// Events
type AlgorithmCreatedEvent struct {
	Algorithm            *Algorithm `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Id                   string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AlgorithmCreatedEvent) Reset()         { *m = AlgorithmCreatedEvent{} }
func (m *AlgorithmCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*AlgorithmCreatedEvent) ProtoMessage()    {}
func (*AlgorithmCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c381a4f1e580eed, []int{5}
}

func (m *AlgorithmCreatedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlgorithmCreatedEvent.Unmarshal(m, b)
}
func (m *AlgorithmCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlgorithmCreatedEvent.Marshal(b, m, deterministic)
}
func (m *AlgorithmCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlgorithmCreatedEvent.Merge(m, src)
}
func (m *AlgorithmCreatedEvent) XXX_Size() int {
	return xxx_messageInfo_AlgorithmCreatedEvent.Size(m)
}
func (m *AlgorithmCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AlgorithmCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AlgorithmCreatedEvent proto.InternalMessageInfo

func (m *AlgorithmCreatedEvent) GetAlgorithm() *Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return nil
}

func (m *AlgorithmCreatedEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MultipleAlgorithms struct {
	Algorithms           []*Algorithm `protobuf:"bytes,1,rep,name=algorithms,proto3" json:"algorithms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MultipleAlgorithms) Reset()         { *m = MultipleAlgorithms{} }
func (m *MultipleAlgorithms) String() string { return proto.CompactTextString(m) }
func (*MultipleAlgorithms) ProtoMessage()    {}
func (*MultipleAlgorithms) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c381a4f1e580eed, []int{6}
}

func (m *MultipleAlgorithms) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultipleAlgorithms.Unmarshal(m, b)
}
func (m *MultipleAlgorithms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultipleAlgorithms.Marshal(b, m, deterministic)
}
func (m *MultipleAlgorithms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultipleAlgorithms.Merge(m, src)
}
func (m *MultipleAlgorithms) XXX_Size() int {
	return xxx_messageInfo_MultipleAlgorithms.Size(m)
}
func (m *MultipleAlgorithms) XXX_DiscardUnknown() {
	xxx_messageInfo_MultipleAlgorithms.DiscardUnknown(m)
}

var xxx_messageInfo_MultipleAlgorithms proto.InternalMessageInfo

func (m *MultipleAlgorithms) GetAlgorithms() []*Algorithm {
	if m != nil {
		return m.Algorithms
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateAlgorithmCommand)(nil), "pb.CreateAlgorithmCommand")
	proto.RegisterType((*AssociateFileCommand)(nil), "pb.AssociateFileCommand")
	proto.RegisterType((*GetAlgorithmQuery)(nil), "pb.GetAlgorithmQuery")
	proto.RegisterType((*Algorithm)(nil), "pb.Algorithm")
	proto.RegisterType((*AlgorithmFile)(nil), "pb.AlgorithmFile")
	proto.RegisterType((*AlgorithmCreatedEvent)(nil), "pb.AlgorithmCreatedEvent")
	proto.RegisterType((*MultipleAlgorithms)(nil), "pb.MultipleAlgorithms")
}

func init() { proto.RegisterFile("algorithm.proto", fileDescriptor_8c381a4f1e580eed) }

var fileDescriptor_8c381a4f1e580eed = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xdd, 0x49, 0xd7, 0x6a, 0xee, 0x6e, 0x5c, 0xf6, 0xea, 0x96, 0xa1, 0x88, 0x84, 0x3c, 0x15,
	0xc4, 0x80, 0x11, 0x54, 0x7c, 0x10, 0x4a, 0xfd, 0xc0, 0x07, 0x11, 0xa3, 0x2f, 0x3e, 0x4e, 0x9b,
	0x31, 0x0e, 0x24, 0x99, 0x90, 0x99, 0x2e, 0x2c, 0xf8, 0x4f, 0xfc, 0x17, 0xfe, 0x0a, 0x7f, 0x96,
	0xcc, 0xb4, 0x9d, 0x7c, 0x2d, 0x88, 0xc2, 0xbe, 0xf5, 0xcc, 0x3d, 0xe7, 0xdc, 0xdb, 0xc3, 0xbd,
	0x81, 0x33, 0x56, 0xe4, 0xb2, 0x11, 0xfa, 0x7b, 0x19, 0xd7, 0x8d, 0xd4, 0x12, 0xbd, 0x7a, 0x1d,
	0x29, 0x98, 0xad, 0x1a, 0xce, 0x34, 0x5f, 0x1e, 0x8a, 0x2b, 0x59, 0x96, 0xac, 0xca, 0xf0, 0x11,
	0xf8, 0x4e, 0x40, 0x49, 0x48, 0x16, 0x27, 0x49, 0x10, 0xd7, 0xeb, 0xd8, 0x11, 0xd3, 0xb6, 0x8e,
	0x0f, 0xc0, 0xdf, 0x58, 0x9b, 0xec, 0x63, 0x45, 0xbd, 0x90, 0x2c, 0x26, 0x69, 0xfb, 0x80, 0x77,
	0xc1, 0x13, 0x19, 0x9d, 0x84, 0x64, 0xe1, 0xa7, 0x9e, 0xc8, 0xa2, 0x1f, 0x70, 0x7f, 0xa9, 0x94,
	0xdc, 0x08, 0xa6, 0xf9, 0x5b, 0x51, 0xf0, 0xff, 0x6a, 0xf9, 0x1c, 0x02, 0x07, 0x8c, 0x89, 0x6d,
	0x7b, 0x92, 0x9c, 0xf7, 0x04, 0xa6, 0x90, 0xf6, 0x79, 0x51, 0x05, 0xe7, 0xef, 0xb8, 0x76, 0x94,
	0x4f, 0x5b, 0xde, 0x5c, 0xdd, 0xe4, 0xbf, 0xfd, 0x49, 0xc0, 0x77, 0x36, 0x88, 0x70, 0x5c, 0xb1,
	0x92, 0xdb, 0x1e, 0x7e, 0x6a, 0x7f, 0x23, 0x85, 0xdb, 0x97, 0xbc, 0x51, 0x42, 0xee, 0xdc, 0xfc,
	0xf4, 0x00, 0x87, 0x5e, 0x38, 0x83, 0xa9, 0xd2, 0x4c, 0x6f, 0x15, 0x3d, 0xb6, 0x6f, 0x7b, 0x64,
	0x1c, 0xbe, 0x89, 0x82, 0xbf, 0x7f, 0xad, 0xe8, 0xad, 0x70, 0x62, 0x1c, 0xf6, 0x10, 0x1f, 0x02,
	0x64, 0x4c, 0x33, 0xc5, 0xb5, 0x29, 0x4e, 0x6d, 0xb1, 0xf3, 0x12, 0x7d, 0x85, 0xa0, 0x97, 0x96,
	0xb1, 0xda, 0xc8, 0x4a, 0xf3, 0x4a, 0xdb, 0x19, 0x4f, 0xd3, 0x03, 0x74, 0xa3, 0x7b, 0x9d, 0xd1,
	0xe7, 0x70, 0xc7, 0x74, 0xd2, 0x57, 0x35, 0xdf, 0x8f, 0xe9, 0x70, 0xf4, 0x05, 0x2e, 0xda, 0xad,
	0xda, 0xc5, 0xf3, 0xe6, 0xd2, 0x18, 0xfd, 0x53, 0xd8, 0xbb, 0x08, 0x3c, 0x17, 0xe7, 0x0a, 0xf0,
	0xc3, 0xb6, 0xd0, 0xa2, 0x2e, 0xda, 0x9d, 0x55, 0xf8, 0x18, 0xc0, 0x49, 0x14, 0x25, 0xe1, 0x64,
	0xec, 0xd9, 0x21, 0x24, 0xbf, 0x09, 0xa0, 0xab, 0x2c, 0xf3, 0xbc, 0xe1, 0x39, 0xd3, 0x1c, 0x5f,
	0xc1, 0xd9, 0xe0, 0x1a, 0x70, 0x6e, 0x4c, 0xae, 0x3f, 0x91, 0x79, 0xbf, 0x41, 0x74, 0x84, 0xcf,
	0xe0, 0xb4, 0xbb, 0x5a, 0x78, 0x61, 0x08, 0xa3, 0x65, 0x1b, 0xeb, 0x5e, 0x42, 0xd0, 0x3b, 0x08,
	0xa4, 0x96, 0x71, 0xcd, 0x8d, 0x8c, 0xb4, 0xc9, 0x2f, 0x02, 0xf7, 0xfa, 0xfe, 0x9f, 0xb5, 0x6c,
	0x38, 0xc6, 0x83, 0x59, 0xfa, 0xc2, 0xf1, 0x0c, 0x2f, 0x20, 0xe8, 0xf2, 0xd5, 0x50, 0x30, 0x33,
	0x70, 0x9c, 0x7c, 0x74, 0x84, 0x4f, 0xc6, 0xa9, 0xfd, 0xa5, 0xd9, 0x7a, 0x6a, 0xbf, 0x40, 0x4f,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x70, 0xf4, 0x86, 0x99, 0x94, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AlgorithmAggregateClient is the client API for AlgorithmAggregate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlgorithmAggregateClient interface {
	CreateAlgorithm(ctx context.Context, in *CreateAlgorithmCommand, opts ...grpc.CallOption) (*Algorithm, error)
	GetAlgorithm(ctx context.Context, in *GetAlgorithmQuery, opts ...grpc.CallOption) (*Algorithm, error)
	AssociateFile(ctx context.Context, in *AssociateFileCommand, opts ...grpc.CallOption) (*Algorithm, error)
}

type algorithmAggregateClient struct {
	cc *grpc.ClientConn
}

func NewAlgorithmAggregateClient(cc *grpc.ClientConn) AlgorithmAggregateClient {
	return &algorithmAggregateClient{cc}
}

func (c *algorithmAggregateClient) CreateAlgorithm(ctx context.Context, in *CreateAlgorithmCommand, opts ...grpc.CallOption) (*Algorithm, error) {
	out := new(Algorithm)
	err := c.cc.Invoke(ctx, "/pb.AlgorithmAggregate/CreateAlgorithm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algorithmAggregateClient) GetAlgorithm(ctx context.Context, in *GetAlgorithmQuery, opts ...grpc.CallOption) (*Algorithm, error) {
	out := new(Algorithm)
	err := c.cc.Invoke(ctx, "/pb.AlgorithmAggregate/GetAlgorithm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algorithmAggregateClient) AssociateFile(ctx context.Context, in *AssociateFileCommand, opts ...grpc.CallOption) (*Algorithm, error) {
	out := new(Algorithm)
	err := c.cc.Invoke(ctx, "/pb.AlgorithmAggregate/AssociateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlgorithmAggregateServer is the server API for AlgorithmAggregate service.
type AlgorithmAggregateServer interface {
	CreateAlgorithm(context.Context, *CreateAlgorithmCommand) (*Algorithm, error)
	GetAlgorithm(context.Context, *GetAlgorithmQuery) (*Algorithm, error)
	AssociateFile(context.Context, *AssociateFileCommand) (*Algorithm, error)
}

func RegisterAlgorithmAggregateServer(s *grpc.Server, srv AlgorithmAggregateServer) {
	s.RegisterService(&_AlgorithmAggregate_serviceDesc, srv)
}

func _AlgorithmAggregate_CreateAlgorithm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlgorithmCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmAggregateServer).CreateAlgorithm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AlgorithmAggregate/CreateAlgorithm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmAggregateServer).CreateAlgorithm(ctx, req.(*CreateAlgorithmCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlgorithmAggregate_GetAlgorithm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlgorithmQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmAggregateServer).GetAlgorithm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AlgorithmAggregate/GetAlgorithm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmAggregateServer).GetAlgorithm(ctx, req.(*GetAlgorithmQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlgorithmAggregate_AssociateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssociateFileCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmAggregateServer).AssociateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AlgorithmAggregate/AssociateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmAggregateServer).AssociateFile(ctx, req.(*AssociateFileCommand))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlgorithmAggregate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AlgorithmAggregate",
	HandlerType: (*AlgorithmAggregateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAlgorithm",
			Handler:    _AlgorithmAggregate_CreateAlgorithm_Handler,
		},
		{
			MethodName: "GetAlgorithm",
			Handler:    _AlgorithmAggregate_GetAlgorithm_Handler,
		},
		{
			MethodName: "AssociateFile",
			Handler:    _AlgorithmAggregate_AssociateFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "algorithm.proto",
}

// AlgorithmQueryStoreClient is the client API for AlgorithmQueryStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlgorithmQueryStoreClient interface {
	GetAlgorithm(ctx context.Context, in *Algorithm, opts ...grpc.CallOption) (*Algorithm, error)
	GetAlgorithms(ctx context.Context, in *Algorithm, opts ...grpc.CallOption) (*MultipleAlgorithms, error)
	CreateAlgorithm(ctx context.Context, in *Algorithm, opts ...grpc.CallOption) (*Algorithm, error)
}

type algorithmQueryStoreClient struct {
	cc *grpc.ClientConn
}

func NewAlgorithmQueryStoreClient(cc *grpc.ClientConn) AlgorithmQueryStoreClient {
	return &algorithmQueryStoreClient{cc}
}

func (c *algorithmQueryStoreClient) GetAlgorithm(ctx context.Context, in *Algorithm, opts ...grpc.CallOption) (*Algorithm, error) {
	out := new(Algorithm)
	err := c.cc.Invoke(ctx, "/pb.AlgorithmQueryStore/GetAlgorithm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algorithmQueryStoreClient) GetAlgorithms(ctx context.Context, in *Algorithm, opts ...grpc.CallOption) (*MultipleAlgorithms, error) {
	out := new(MultipleAlgorithms)
	err := c.cc.Invoke(ctx, "/pb.AlgorithmQueryStore/GetAlgorithms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *algorithmQueryStoreClient) CreateAlgorithm(ctx context.Context, in *Algorithm, opts ...grpc.CallOption) (*Algorithm, error) {
	out := new(Algorithm)
	err := c.cc.Invoke(ctx, "/pb.AlgorithmQueryStore/CreateAlgorithm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlgorithmQueryStoreServer is the server API for AlgorithmQueryStore service.
type AlgorithmQueryStoreServer interface {
	GetAlgorithm(context.Context, *Algorithm) (*Algorithm, error)
	GetAlgorithms(context.Context, *Algorithm) (*MultipleAlgorithms, error)
	CreateAlgorithm(context.Context, *Algorithm) (*Algorithm, error)
}

func RegisterAlgorithmQueryStoreServer(s *grpc.Server, srv AlgorithmQueryStoreServer) {
	s.RegisterService(&_AlgorithmQueryStore_serviceDesc, srv)
}

func _AlgorithmQueryStore_GetAlgorithm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Algorithm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmQueryStoreServer).GetAlgorithm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AlgorithmQueryStore/GetAlgorithm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmQueryStoreServer).GetAlgorithm(ctx, req.(*Algorithm))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlgorithmQueryStore_GetAlgorithms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Algorithm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmQueryStoreServer).GetAlgorithms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AlgorithmQueryStore/GetAlgorithms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmQueryStoreServer).GetAlgorithms(ctx, req.(*Algorithm))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlgorithmQueryStore_CreateAlgorithm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Algorithm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlgorithmQueryStoreServer).CreateAlgorithm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AlgorithmQueryStore/CreateAlgorithm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlgorithmQueryStoreServer).CreateAlgorithm(ctx, req.(*Algorithm))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlgorithmQueryStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AlgorithmQueryStore",
	HandlerType: (*AlgorithmQueryStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAlgorithm",
			Handler:    _AlgorithmQueryStore_GetAlgorithm_Handler,
		},
		{
			MethodName: "GetAlgorithms",
			Handler:    _AlgorithmQueryStore_GetAlgorithms_Handler,
		},
		{
			MethodName: "CreateAlgorithm",
			Handler:    _AlgorithmQueryStore_CreateAlgorithm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "algorithm.proto",
}
