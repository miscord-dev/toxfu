// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: proto/api.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IPPrefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Bits    int32  `protobuf:"varint,2,opt,name=bits,proto3" json:"bits,omitempty"`
}

func (x *IPPrefix) Reset() {
	*x = IPPrefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPPrefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPPrefix) ProtoMessage() {}

func (x *IPPrefix) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPPrefix.ProtoReflect.Descriptor instead.
func (*IPPrefix) Descriptor() ([]byte, []int) {
	return file_proto_api_proto_rawDescGZIP(), []int{0}
}

func (x *IPPrefix) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *IPPrefix) GetBits() int32 {
	if x != nil {
		return x.Bits
	}
	return 0
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PublicKey          string         `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PublicDiscoKey     string         `protobuf:"bytes,3,opt,name=public_disco_key,json=publicDiscoKey,proto3" json:"public_disco_key,omitempty"`
	Endpoints          []string       `protobuf:"bytes,4,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Addresses          []*IPPrefix    `protobuf:"bytes,5,rep,name=addresses,proto3" json:"addresses,omitempty"`
	AdvertisedPrefixes []*IPPrefix    `protobuf:"bytes,6,rep,name=advertised_prefixes,json=advertisedPrefixes,proto3" json:"advertised_prefixes,omitempty"`
	Attribute          *NodeAttribute `protobuf:"bytes,7,opt,name=attribute,proto3" json:"attribute,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_proto_api_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Node) GetPublicDiscoKey() string {
	if x != nil {
		return x.PublicDiscoKey
	}
	return ""
}

func (x *Node) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *Node) GetAddresses() []*IPPrefix {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *Node) GetAdvertisedPrefixes() []*IPPrefix {
	if x != nil {
		return x.AdvertisedPrefixes
	}
	return nil
}

func (x *Node) GetAttribute() *NodeAttribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

type NodeAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostName string `protobuf:"bytes,1,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	Os       string `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	Goos     string `protobuf:"bytes,3,opt,name=goos,proto3" json:"goos,omitempty"`
	Goarch   string `protobuf:"bytes,4,opt,name=goarch,proto3" json:"goarch,omitempty"`
}

func (x *NodeAttribute) Reset() {
	*x = NodeAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeAttribute) ProtoMessage() {}

func (x *NodeAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeAttribute.ProtoReflect.Descriptor instead.
func (*NodeAttribute) Descriptor() ([]byte, []int) {
	return file_proto_api_proto_rawDescGZIP(), []int{2}
}

func (x *NodeAttribute) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *NodeAttribute) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *NodeAttribute) GetGoos() string {
	if x != nil {
		return x.Goos
	}
	return ""
}

func (x *NodeAttribute) GetGoarch() string {
	if x != nil {
		return x.Goarch
	}
	return ""
}

type NodeRefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey      string         `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PublicDiscoKey string         `protobuf:"bytes,2,opt,name=public_disco_key,json=publicDiscoKey,proto3" json:"public_disco_key,omitempty"`
	Endpoints      []string       `protobuf:"bytes,3,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Attribute      *NodeAttribute `protobuf:"bytes,4,opt,name=attribute,proto3" json:"attribute,omitempty"`
}

func (x *NodeRefreshRequest) Reset() {
	*x = NodeRefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeRefreshRequest) ProtoMessage() {}

func (x *NodeRefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeRefreshRequest.ProtoReflect.Descriptor instead.
func (*NodeRefreshRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_proto_rawDescGZIP(), []int{3}
}

func (x *NodeRefreshRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *NodeRefreshRequest) GetPublicDiscoKey() string {
	if x != nil {
		return x.PublicDiscoKey
	}
	return ""
}

func (x *NodeRefreshRequest) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *NodeRefreshRequest) GetAttribute() *NodeAttribute {
	if x != nil {
		return x.Attribute
	}
	return nil
}

type NodeRefreshResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SelfNode   *Node   `protobuf:"bytes,1,opt,name=self_node,json=selfNode,proto3" json:"self_node,omitempty"`
	StunServer string  `protobuf:"bytes,2,opt,name=stun_server,json=stunServer,proto3" json:"stun_server,omitempty"`
	Peers      []*Node `protobuf:"bytes,3,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *NodeRefreshResponse) Reset() {
	*x = NodeRefreshResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeRefreshResponse) ProtoMessage() {}

func (x *NodeRefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeRefreshResponse.ProtoReflect.Descriptor instead.
func (*NodeRefreshResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_proto_rawDescGZIP(), []int{4}
}

func (x *NodeRefreshResponse) GetSelfNode() *Node {
	if x != nil {
		return x.SelfNode
	}
	return nil
}

func (x *NodeRefreshResponse) GetStunServer() string {
	if x != nil {
		return x.StunServer
	}
	return ""
}

func (x *NodeRefreshResponse) GetPeers() []*Node {
	if x != nil {
		return x.Peers
	}
	return nil
}

var File_proto_api_proto protoreflect.FileDescriptor

var file_proto_api_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x08, 0x49, 0x50, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x69,
	0x74, 0x73, 0x22, 0xa2, 0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x50,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x40, 0x0a, 0x13, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x50, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52,
	0x12, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x09, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x68, 0x0a, 0x0d, 0x4e, 0x6f, 0x64, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x6f, 0x6f, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x6f, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x6f, 0x61,
	0x72, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x6f, 0x61, 0x72, 0x63,
	0x68, 0x22, 0xaf, 0x01, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x4b, 0x65,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x32, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x13, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x73,
	0x65, 0x6c, 0x66, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x08, 0x73, 0x65, 0x6c,
	0x66, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x75, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x75, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x32, 0x4f, 0x0a, 0x07, 0x4e, 0x6f, 0x64,
	0x65, 0x41, 0x50, 0x49, 0x12, 0x44, 0x0a, 0x07, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x74, 0x6f, 0x78, 0x66, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_proto_rawDescOnce sync.Once
	file_proto_api_proto_rawDescData = file_proto_api_proto_rawDesc
)

func file_proto_api_proto_rawDescGZIP() []byte {
	file_proto_api_proto_rawDescOnce.Do(func() {
		file_proto_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_proto_rawDescData)
	})
	return file_proto_api_proto_rawDescData
}

var file_proto_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_api_proto_goTypes = []interface{}{
	(*IPPrefix)(nil),            // 0: proto.IPPrefix
	(*Node)(nil),                // 1: proto.Node
	(*NodeAttribute)(nil),       // 2: proto.NodeAttribute
	(*NodeRefreshRequest)(nil),  // 3: proto.NodeRefreshRequest
	(*NodeRefreshResponse)(nil), // 4: proto.NodeRefreshResponse
}
var file_proto_api_proto_depIdxs = []int32{
	0, // 0: proto.Node.addresses:type_name -> proto.IPPrefix
	0, // 1: proto.Node.advertised_prefixes:type_name -> proto.IPPrefix
	2, // 2: proto.Node.attribute:type_name -> proto.NodeAttribute
	2, // 3: proto.NodeRefreshRequest.attribute:type_name -> proto.NodeAttribute
	1, // 4: proto.NodeRefreshResponse.self_node:type_name -> proto.Node
	1, // 5: proto.NodeRefreshResponse.peers:type_name -> proto.Node
	3, // 6: proto.NodeAPI.Refresh:input_type -> proto.NodeRefreshRequest
	4, // 7: proto.NodeAPI.Refresh:output_type -> proto.NodeRefreshResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_api_proto_init() }
func file_proto_api_proto_init() {
	if File_proto_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPPrefix); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeAttribute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeRefreshRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeRefreshResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_proto_goTypes,
		DependencyIndexes: file_proto_api_proto_depIdxs,
		MessageInfos:      file_proto_api_proto_msgTypes,
	}.Build()
	File_proto_api_proto = out.File
	file_proto_api_proto_rawDesc = nil
	file_proto_api_proto_goTypes = nil
	file_proto_api_proto_depIdxs = nil
}
