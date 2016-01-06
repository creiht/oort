// Code generated by protoc-gen-go.
// source: group_api.proto
// DO NOT EDIT!

/*
Package groupproto is a generated protocol buffer package.

It is generated from these files:
	group_api.proto

It has these top-level messages:
	EmptyMsg
	KeyGroupValue
	KeyGroup
	Key
	WriteResponse
	LookupResponse
	LookupGroupResponse
	LookupGroupItem
	ReadResponse
	DelResponse
*/
package groupproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EmptyMsg struct {
}

func (m *EmptyMsg) Reset()         { *m = EmptyMsg{} }
func (m *EmptyMsg) String() string { return proto.CompactTextString(m) }
func (*EmptyMsg) ProtoMessage()    {}

type KeyGroupValue struct {
	Keya     uint64 `protobuf:"varint,1,opt,name=keya" json:"keya,omitempty"`
	Keyb     uint64 `protobuf:"varint,2,opt,name=keyb" json:"keyb,omitempty"`
	Namekeya uint64 `protobuf:"varint,3,opt,name=namekeya" json:"namekeya,omitempty"`
	Namekeyb uint64 `protobuf:"varint,4,opt,name=namekeyb" json:"namekeyb,omitempty"`
	Value    []byte `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Ts       int64  `protobuf:"varint,6,opt,name=ts" json:"ts,omitempty"`
}

func (m *KeyGroupValue) Reset()         { *m = KeyGroupValue{} }
func (m *KeyGroupValue) String() string { return proto.CompactTextString(m) }
func (*KeyGroupValue) ProtoMessage()    {}

type KeyGroup struct {
	Keya     uint64 `protobuf:"varint,1,opt,name=keya" json:"keya,omitempty"`
	Keyb     uint64 `protobuf:"varint,2,opt,name=keyb" json:"keyb,omitempty"`
	Namekeya uint64 `protobuf:"varint,3,opt,name=namekeya" json:"namekeya,omitempty"`
	Namekeyb uint64 `protobuf:"varint,4,opt,name=namekeyb" json:"namekeyb,omitempty"`
	Ts       int64  `protobuf:"varint,5,opt,name=ts" json:"ts,omitempty"`
}

func (m *KeyGroup) Reset()         { *m = KeyGroup{} }
func (m *KeyGroup) String() string { return proto.CompactTextString(m) }
func (*KeyGroup) ProtoMessage()    {}

type Key struct {
	Keya uint64 `protobuf:"varint,1,opt,name=keya" json:"keya,omitempty"`
	Keyb uint64 `protobuf:"varint,2,opt,name=keyb" json:"keyb,omitempty"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}

type WriteResponse struct {
	Ts  int64  `protobuf:"varint,1,opt,name=ts" json:"ts,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}

type LookupResponse struct {
	Ts     int64  `protobuf:"varint,1,opt,name=ts" json:"ts,omitempty"`
	Length uint32 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	Err    string `protobuf:"bytes,3,opt,name=err" json:"err,omitempty"`
}

func (m *LookupResponse) Reset()         { *m = LookupResponse{} }
func (m *LookupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()    {}

type LookupGroupResponse struct {
	Items []*LookupGroupItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *LookupGroupResponse) Reset()         { *m = LookupGroupResponse{} }
func (m *LookupGroupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupGroupResponse) ProtoMessage()    {}

func (m *LookupGroupResponse) GetItems() []*LookupGroupItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type LookupGroupItem struct {
	Namekeya uint64 `protobuf:"varint,1,opt,name=namekeya" json:"namekeya,omitempty"`
	Namekeyb uint64 `protobuf:"varint,2,opt,name=namekeyb" json:"namekeyb,omitempty"`
	Ts       uint64 `protobuf:"varint,3,opt,name=ts" json:"ts,omitempty"`
	Length   uint32 `protobuf:"varint,4,opt,name=length" json:"length,omitempty"`
}

func (m *LookupGroupItem) Reset()         { *m = LookupGroupItem{} }
func (m *LookupGroupItem) String() string { return proto.CompactTextString(m) }
func (*LookupGroupItem) ProtoMessage()    {}

type ReadResponse struct {
	Ts    int64  `protobuf:"varint,1,opt,name=ts" json:"ts,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Err   string `protobuf:"bytes,3,opt,name=err" json:"err,omitempty"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}

type DelResponse struct {
	Ts  int64  `protobuf:"varint,1,opt,name=ts" json:"ts,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *DelResponse) Reset()         { *m = DelResponse{} }
func (m *DelResponse) String() string { return proto.CompactTextString(m) }
func (*DelResponse) ProtoMessage()    {}

func init() {
	proto.RegisterType((*EmptyMsg)(nil), "groupproto.EmptyMsg")
	proto.RegisterType((*KeyGroupValue)(nil), "groupproto.KeyGroupValue")
	proto.RegisterType((*KeyGroup)(nil), "groupproto.KeyGroup")
	proto.RegisterType((*Key)(nil), "groupproto.Key")
	proto.RegisterType((*WriteResponse)(nil), "groupproto.WriteResponse")
	proto.RegisterType((*LookupResponse)(nil), "groupproto.LookupResponse")
	proto.RegisterType((*LookupGroupResponse)(nil), "groupproto.LookupGroupResponse")
	proto.RegisterType((*LookupGroupItem)(nil), "groupproto.LookupGroupItem")
	proto.RegisterType((*ReadResponse)(nil), "groupproto.ReadResponse")
	proto.RegisterType((*DelResponse)(nil), "groupproto.DelResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for GroupStore service

type GroupStoreClient interface {
	Write(ctx context.Context, in *KeyGroupValue, opts ...grpc.CallOption) (*WriteResponse, error)
	Lookup(ctx context.Context, in *KeyGroup, opts ...grpc.CallOption) (*LookupResponse, error)
	LookupGroup(ctx context.Context, in *Key, opts ...grpc.CallOption) (*LookupGroupResponse, error)
	Read(ctx context.Context, in *KeyGroup, opts ...grpc.CallOption) (*ReadResponse, error)
	Delete(ctx context.Context, in *KeyGroup, opts ...grpc.CallOption) (*DelResponse, error)
}

type groupStoreClient struct {
	cc *grpc.ClientConn
}

func NewGroupStoreClient(cc *grpc.ClientConn) GroupStoreClient {
	return &groupStoreClient{cc}
}

func (c *groupStoreClient) Write(ctx context.Context, in *KeyGroupValue, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/groupproto.GroupStore/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupStoreClient) Lookup(ctx context.Context, in *KeyGroup, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := grpc.Invoke(ctx, "/groupproto.GroupStore/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupStoreClient) LookupGroup(ctx context.Context, in *Key, opts ...grpc.CallOption) (*LookupGroupResponse, error) {
	out := new(LookupGroupResponse)
	err := grpc.Invoke(ctx, "/groupproto.GroupStore/LookupGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupStoreClient) Read(ctx context.Context, in *KeyGroup, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := grpc.Invoke(ctx, "/groupproto.GroupStore/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupStoreClient) Delete(ctx context.Context, in *KeyGroup, opts ...grpc.CallOption) (*DelResponse, error) {
	out := new(DelResponse)
	err := grpc.Invoke(ctx, "/groupproto.GroupStore/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupStore service

type GroupStoreServer interface {
	Write(context.Context, *KeyGroupValue) (*WriteResponse, error)
	Lookup(context.Context, *KeyGroup) (*LookupResponse, error)
	LookupGroup(context.Context, *Key) (*LookupGroupResponse, error)
	Read(context.Context, *KeyGroup) (*ReadResponse, error)
	Delete(context.Context, *KeyGroup) (*DelResponse, error)
}

func RegisterGroupStoreServer(s *grpc.Server, srv GroupStoreServer) {
	s.RegisterService(&_GroupStore_serviceDesc, srv)
}

func _GroupStore_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(KeyGroupValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupStoreServer).Write(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GroupStore_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(KeyGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupStoreServer).Lookup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GroupStore_LookupGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupStoreServer).LookupGroup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GroupStore_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(KeyGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupStoreServer).Read(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GroupStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(KeyGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GroupStoreServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _GroupStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "groupproto.GroupStore",
	HandlerType: (*GroupStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _GroupStore_Write_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _GroupStore_Lookup_Handler,
		},
		{
			MethodName: "LookupGroup",
			Handler:    _GroupStore_LookupGroup_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _GroupStore_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GroupStore_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}