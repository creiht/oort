// Code generated by protoc-gen-go.
// source: ort_api.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	ort_api.proto

It has these top-level messages:
	FileRequest
	LookupRequest
	Attr
	File
	WriteResponse
	DirEntries
	DirEnt
	FileEnt
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

// FileRequest is the file inode
type FileRequest struct {
	Inode uint64 `protobuf:"varint,1,opt,name=inode" json:"inode,omitempty"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto1.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}

// LookupRequest
type LookupRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parent uint64 `protobuf:"varint,2,opt,name=parent" json:"parent,omitempty"`
}

func (m *LookupRequest) Reset()         { *m = LookupRequest{} }
func (m *LookupRequest) String() string { return proto1.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()    {}

// Attr. fields are optional by default in proto3, so
// clients don't have to send all fields when performing an
// attr update for example. These might not all be needed
// but i got tired of constantly forgetting fields.
type Attr struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Inode  uint64 `protobuf:"varint,2,opt,name=inode" json:"inode,omitempty"`
	Atime  int64  `protobuf:"varint,3,opt,name=atime" json:"atime,omitempty"`
	Mtime  int64  `protobuf:"varint,4,opt,name=mtime" json:"mtime,omitempty"`
	Ctime  int64  `protobuf:"varint,5,opt,name=ctime" json:"ctime,omitempty"`
	Crtime int64  `protobuf:"varint,6,opt,name=crtime" json:"crtime,omitempty"`
	Mode   uint32 `protobuf:"varint,7,opt,name=mode" json:"mode,omitempty"`
	Valid  int32  `protobuf:"varint,9,opt,name=valid" json:"valid,omitempty"`
	Parent string `protobuf:"bytes,10,opt,name=parent" json:"parent,omitempty"`
	Size   uint64 `protobuf:"varint,11,opt,name=size" json:"size,omitempty"`
}

func (m *Attr) Reset()         { *m = Attr{} }
func (m *Attr) String() string { return proto1.CompactTextString(m) }
func (*Attr) ProtoMessage()    {}

// File contains the files name and its contents
type File struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Inode   uint64 `protobuf:"varint,2,opt,name=inode" json:"inode,omitempty"`
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto1.CompactTextString(m) }
func (*File) ProtoMessage()    {}

// WriteRepsonse place holder. Maybe use an enum so
// we can map to fuse errors ?
type WriteResponse struct {
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto1.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}

// DirEntries just contains a list of directory entries
type DirEntries struct {
	DirEntries  []*DirEnt  `protobuf:"bytes,1,rep" json:"DirEntries,omitempty"`
	FileEntries []*FileEnt `protobuf:"bytes,2,rep" json:"FileEntries,omitempty"`
}

func (m *DirEntries) Reset()         { *m = DirEntries{} }
func (m *DirEntries) String() string { return proto1.CompactTextString(m) }
func (*DirEntries) ProtoMessage()    {}

func (m *DirEntries) GetDirEntries() []*DirEnt {
	if m != nil {
		return m.DirEntries
	}
	return nil
}

func (m *DirEntries) GetFileEntries() []*FileEnt {
	if m != nil {
		return m.FileEntries
	}
	return nil
}

// DirEnt is a directory entry
type DirEnt struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parent uint64 `protobuf:"varint,2,opt,name=parent" json:"parent,omitempty"`
	Attr   *Attr  `protobuf:"bytes,3,opt,name=attr" json:"attr,omitempty"`
}

func (m *DirEnt) Reset()         { *m = DirEnt{} }
func (m *DirEnt) String() string { return proto1.CompactTextString(m) }
func (*DirEnt) ProtoMessage()    {}

func (m *DirEnt) GetAttr() *Attr {
	if m != nil {
		return m.Attr
	}
	return nil
}

type FileEnt struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parent uint64 `protobuf:"varint,2,opt,name=parent" json:"parent,omitempty"`
	Attr   *Attr  `protobuf:"bytes,3,opt,name=attr" json:"attr,omitempty"`
}

func (m *FileEnt) Reset()         { *m = FileEnt{} }
func (m *FileEnt) String() string { return proto1.CompactTextString(m) }
func (*FileEnt) ProtoMessage()    {}

func (m *FileEnt) GetAttr() *Attr {
	if m != nil {
		return m.Attr
	}
	return nil
}

// Client API for Api service

type ApiClient interface {
	SetAttr(ctx context.Context, in *Attr, opts ...grpc.CallOption) (*Attr, error)
	GetAttr(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*Attr, error)
	Read(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*File, error)
	Write(ctx context.Context, in *File, opts ...grpc.CallOption) (*WriteResponse, error)
	MkDir(ctx context.Context, in *DirEnt, opts ...grpc.CallOption) (*DirEnt, error)
	Create(ctx context.Context, in *FileEnt, opts ...grpc.CallOption) (*FileEnt, error)
	Remove(ctx context.Context, in *DirEnt, opts ...grpc.CallOption) (*WriteResponse, error)
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*DirEnt, error)
	ReadDirAll(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*DirEntries, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) SetAttr(ctx context.Context, in *Attr, opts ...grpc.CallOption) (*Attr, error) {
	out := new(Attr)
	err := grpc.Invoke(ctx, "/proto.Api/SetAttr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetAttr(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*Attr, error) {
	out := new(Attr)
	err := grpc.Invoke(ctx, "/proto.Api/GetAttr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Read(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := grpc.Invoke(ctx, "/proto.Api/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Write(ctx context.Context, in *File, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/proto.Api/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) MkDir(ctx context.Context, in *DirEnt, opts ...grpc.CallOption) (*DirEnt, error) {
	out := new(DirEnt)
	err := grpc.Invoke(ctx, "/proto.Api/MkDir", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Create(ctx context.Context, in *FileEnt, opts ...grpc.CallOption) (*FileEnt, error) {
	out := new(FileEnt)
	err := grpc.Invoke(ctx, "/proto.Api/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Remove(ctx context.Context, in *DirEnt, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/proto.Api/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*DirEnt, error) {
	out := new(DirEnt)
	err := grpc.Invoke(ctx, "/proto.Api/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ReadDirAll(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*DirEntries, error) {
	out := new(DirEntries)
	err := grpc.Invoke(ctx, "/proto.Api/ReadDirAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiServer interface {
	SetAttr(context.Context, *Attr) (*Attr, error)
	GetAttr(context.Context, *FileRequest) (*Attr, error)
	Read(context.Context, *FileRequest) (*File, error)
	Write(context.Context, *File) (*WriteResponse, error)
	MkDir(context.Context, *DirEnt) (*DirEnt, error)
	Create(context.Context, *FileEnt) (*FileEnt, error)
	Remove(context.Context, *DirEnt) (*WriteResponse, error)
	Lookup(context.Context, *LookupRequest) (*DirEnt, error)
	ReadDirAll(context.Context, *FileRequest) (*DirEntries, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_SetAttr_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(Attr)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).SetAttr(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_GetAttr_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(FileRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).GetAttr(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_Read_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(FileRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).Read(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_Write_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(File)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).Write(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_MkDir_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(DirEnt)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).MkDir(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_Create_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(FileEnt)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_Remove_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(DirEnt)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).Remove(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_Lookup_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(LookupRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).Lookup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_ReadDirAll_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(FileRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).ReadDirAll(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetAttr",
			Handler:    _Api_SetAttr_Handler,
		},
		{
			MethodName: "GetAttr",
			Handler:    _Api_GetAttr_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Api_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _Api_Write_Handler,
		},
		{
			MethodName: "MkDir",
			Handler:    _Api_MkDir_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Api_Create_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Api_Remove_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _Api_Lookup_Handler,
		},
		{
			MethodName: "ReadDirAll",
			Handler:    _Api_ReadDirAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
