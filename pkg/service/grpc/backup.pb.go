// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backup.proto

package backup_rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Data struct {
	DataType             string   `protobuf:"bytes,1,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_65240d19de191688, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *Data) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Data) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Error struct {
	Info                 string   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_65240d19de191688, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type SaveRequest struct {
	Data                 *Data    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	TaskName             string   `protobuf:"bytes,2,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveRequest) Reset()         { *m = SaveRequest{} }
func (m *SaveRequest) String() string { return proto.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()    {}
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65240d19de191688, []int{2}
}

func (m *SaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveRequest.Unmarshal(m, b)
}
func (m *SaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveRequest.Marshal(b, m, deterministic)
}
func (m *SaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveRequest.Merge(m, src)
}
func (m *SaveRequest) XXX_Size() int {
	return xxx_messageInfo_SaveRequest.Size(m)
}
func (m *SaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveRequest proto.InternalMessageInfo

func (m *SaveRequest) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SaveRequest) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

type SaveResponse struct {
	Err                  *Error   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveResponse) Reset()         { *m = SaveResponse{} }
func (m *SaveResponse) String() string { return proto.CompactTextString(m) }
func (*SaveResponse) ProtoMessage()    {}
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65240d19de191688, []int{3}
}

func (m *SaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveResponse.Unmarshal(m, b)
}
func (m *SaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveResponse.Marshal(b, m, deterministic)
}
func (m *SaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveResponse.Merge(m, src)
}
func (m *SaveResponse) XXX_Size() int {
	return xxx_messageInfo_SaveResponse.Size(m)
}
func (m *SaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveResponse proto.InternalMessageInfo

func (m *SaveResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type CloseRequest struct {
	TaskName             string   `protobuf:"bytes,1,opt,name=TaskName,proto3" json:"TaskName,omitempty"`
	TypeName             string   `protobuf:"bytes,2,opt,name=TypeName,proto3" json:"TypeName,omitempty"`
	Time                 int64    `protobuf:"varint,3,opt,name=Time,proto3" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseRequest) Reset()         { *m = CloseRequest{} }
func (m *CloseRequest) String() string { return proto.CompactTextString(m) }
func (*CloseRequest) ProtoMessage()    {}
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65240d19de191688, []int{4}
}

func (m *CloseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseRequest.Unmarshal(m, b)
}
func (m *CloseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseRequest.Marshal(b, m, deterministic)
}
func (m *CloseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseRequest.Merge(m, src)
}
func (m *CloseRequest) XXX_Size() int {
	return xxx_messageInfo_CloseRequest.Size(m)
}
func (m *CloseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseRequest proto.InternalMessageInfo

func (m *CloseRequest) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *CloseRequest) GetTypeName() string {
	if m != nil {
		return m.TypeName
	}
	return ""
}

func (m *CloseRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type CloseResponse struct {
	Err                  *Error   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseResponse) Reset()         { *m = CloseResponse{} }
func (m *CloseResponse) String() string { return proto.CompactTextString(m) }
func (*CloseResponse) ProtoMessage()    {}
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65240d19de191688, []int{5}
}

func (m *CloseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseResponse.Unmarshal(m, b)
}
func (m *CloseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseResponse.Marshal(b, m, deterministic)
}
func (m *CloseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseResponse.Merge(m, src)
}
func (m *CloseResponse) XXX_Size() int {
	return xxx_messageInfo_CloseResponse.Size(m)
}
func (m *CloseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseResponse proto.InternalMessageInfo

func (m *CloseResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*Data)(nil), "backup_rpc.Data")
	proto.RegisterType((*Error)(nil), "backup_rpc.Error")
	proto.RegisterType((*SaveRequest)(nil), "backup_rpc.SaveRequest")
	proto.RegisterType((*SaveResponse)(nil), "backup_rpc.SaveResponse")
	proto.RegisterType((*CloseRequest)(nil), "backup_rpc.CloseRequest")
	proto.RegisterType((*CloseResponse)(nil), "backup_rpc.CloseResponse")
}

func init() { proto.RegisterFile("backup.proto", fileDescriptor_65240d19de191688) }

var fileDescriptor_65240d19de191688 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x31, 0x29, 0x28, 0xb9, 0xa6, 0x12, 0x78, 0x21, 0xb4, 0x4b, 0x65, 0x18, 0x32, 0x75,
	0x68, 0x79, 0x01, 0x0a, 0x4c, 0x48, 0x15, 0x72, 0x3b, 0xb1, 0x54, 0x6e, 0x39, 0xa4, 0xaa, 0x34,
	0x36, 0x8e, 0x5b, 0x29, 0x2f, 0xc2, 0xf3, 0xa2, 0x73, 0x1c, 0x08, 0x82, 0x85, 0xed, 0xee, 0x7e,
	0xfb, 0xbb, 0xff, 0xce, 0x86, 0x74, 0xa5, 0xd6, 0xdb, 0xbd, 0x19, 0x19, 0xab, 0x9d, 0xe6, 0x50,
	0x67, 0x4b, 0x6b, 0xd6, 0xe2, 0x11, 0x3a, 0xf7, 0xca, 0x29, 0x3e, 0x80, 0xe4, 0x45, 0x39, 0xb5,
	0x74, 0x95, 0xc1, 0x8c, 0x0d, 0x59, 0x9e, 0xc8, 0x98, 0x0a, 0x8b, 0xca, 0x20, 0xe7, 0xd0, 0x71,
	0x9b, 0x1d, 0x66, 0xc7, 0x43, 0x96, 0x47, 0xd2, 0xc7, 0x54, 0x23, 0x3d, 0x8b, 0x86, 0x2c, 0x4f,
	0xa5, 0x8f, 0xc5, 0x00, 0x4e, 0x1e, 0xac, 0xd5, 0x96, 0xc4, 0x4d, 0xf1, 0xaa, 0x03, 0xc8, 0xc7,
	0xe2, 0x09, 0xba, 0x73, 0x75, 0x40, 0x89, 0xef, 0x7b, 0x2c, 0x1d, 0xbf, 0x0e, 0xf7, 0xe9, 0x48,
	0x77, 0x7c, 0x36, 0xfa, 0xf6, 0x34, 0x22, 0x43, 0x35, 0x91, 0x6c, 0x39, 0x55, 0x6e, 0x97, 0x85,
	0x0a, 0xed, 0x13, 0x19, 0x53, 0x61, 0xa6, 0x76, 0x28, 0x26, 0x90, 0xd6, 0xc4, 0xd2, 0xe8, 0xa2,
	0x44, 0x7e, 0x05, 0x11, 0x5a, 0x1b, 0x88, 0xe7, 0x6d, 0xa2, 0x77, 0x25, 0x49, 0x15, 0xcf, 0x90,
	0xde, 0xbd, 0xe9, 0xf2, 0xcb, 0x47, 0x1f, 0xe2, 0x45, 0x00, 0x36, 0x73, 0x37, 0xb9, 0xd7, 0x2a,
	0x83, 0xb3, 0x56, 0xf3, 0x26, 0xa7, 0x11, 0x17, 0xb4, 0x93, 0xa8, 0xde, 0x09, 0xc5, 0xe2, 0x06,
	0x7a, 0x81, 0xfd, 0x0f, 0x47, 0xe3, 0x0f, 0x06, 0xbd, 0xa9, 0x57, 0xe6, 0x68, 0x0f, 0x9b, 0x35,
	0xf2, 0x5b, 0x88, 0x69, 0x30, 0xff, 0x30, 0x17, 0xed, 0x5b, 0xad, 0x05, 0xf6, 0xb3, 0xdf, 0x42,
	0xdd, 0x55, 0x1c, 0xe5, 0x8c, 0x4f, 0x21, 0xf1, 0x56, 0xfc, 0xfb, 0xfd, 0x38, 0xda, 0x9e, 0xbe,
	0x7f, 0xf9, 0x87, 0xd2, 0x50, 0x56, 0xa7, 0xfe, 0xbb, 0x4c, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x99, 0x7f, 0x53, 0x73, 0x3e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackupServiceClient is the client API for BackupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackupServiceClient interface {
	SaveData(ctx context.Context, opts ...grpc.CallOption) (BackupService_SaveDataClient, error)
	CloseType(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error)
}

type backupServiceClient struct {
	cc *grpc.ClientConn
}

func NewBackupServiceClient(cc *grpc.ClientConn) BackupServiceClient {
	return &backupServiceClient{cc}
}

func (c *backupServiceClient) SaveData(ctx context.Context, opts ...grpc.CallOption) (BackupService_SaveDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BackupService_serviceDesc.Streams[0], "/backup_rpc.BackupService/SaveData", opts...)
	if err != nil {
		return nil, err
	}
	x := &backupServiceSaveDataClient{stream}
	return x, nil
}

type BackupService_SaveDataClient interface {
	Send(*SaveRequest) error
	CloseAndRecv() (*SaveResponse, error)
	grpc.ClientStream
}

type backupServiceSaveDataClient struct {
	grpc.ClientStream
}

func (x *backupServiceSaveDataClient) Send(m *SaveRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *backupServiceSaveDataClient) CloseAndRecv() (*SaveResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SaveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backupServiceClient) CloseType(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error) {
	out := new(CloseResponse)
	err := c.cc.Invoke(ctx, "/backup_rpc.BackupService/CloseType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServiceServer is the server API for BackupService service.
type BackupServiceServer interface {
	SaveData(BackupService_SaveDataServer) error
	CloseType(context.Context, *CloseRequest) (*CloseResponse, error)
}

// UnimplementedBackupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBackupServiceServer struct {
}

func (*UnimplementedBackupServiceServer) SaveData(srv BackupService_SaveDataServer) error {
	return status.Errorf(codes.Unimplemented, "method SaveData not implemented")
}
func (*UnimplementedBackupServiceServer) CloseType(ctx context.Context, req *CloseRequest) (*CloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseType not implemented")
}

func RegisterBackupServiceServer(s *grpc.Server, srv BackupServiceServer) {
	s.RegisterService(&_BackupService_serviceDesc, srv)
}

func _BackupService_SaveData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BackupServiceServer).SaveData(&backupServiceSaveDataServer{stream})
}

type BackupService_SaveDataServer interface {
	SendAndClose(*SaveResponse) error
	Recv() (*SaveRequest, error)
	grpc.ServerStream
}

type backupServiceSaveDataServer struct {
	grpc.ServerStream
}

func (x *backupServiceSaveDataServer) SendAndClose(m *SaveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *backupServiceSaveDataServer) Recv() (*SaveRequest, error) {
	m := new(SaveRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BackupService_CloseType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).CloseType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup_rpc.BackupService/CloseType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).CloseType(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backup_rpc.BackupService",
	HandlerType: (*BackupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CloseType",
			Handler:    _BackupService_CloseType_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SaveData",
			Handler:       _BackupService_SaveData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "backup.proto",
}
