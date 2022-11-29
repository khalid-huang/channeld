// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: tps.proto

package tpspb

import (
	unrealpb "channeld.clewcat.com/channeld/pkg/unrealpb"
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

type TestRepChannelData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameState              *unrealpb.GameStateBase                    `protobuf:"bytes,1,opt,name=gameState,proto3" json:"gameState,omitempty"`
	ActorStates            map[uint32]*unrealpb.ActorState            `protobuf:"bytes,2,rep,name=actorStates,proto3" json:"actorStates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PawnStates             map[uint32]*unrealpb.PawnState             `protobuf:"bytes,3,rep,name=pawnStates,proto3" json:"pawnStates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CharacterStates        map[uint32]*unrealpb.CharacterState        `protobuf:"bytes,4,rep,name=characterStates,proto3" json:"characterStates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PlayerStates           map[uint32]*unrealpb.PlayerState           `protobuf:"bytes,5,rep,name=playerStates,proto3" json:"playerStates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ControllerStates       map[uint32]*unrealpb.ControllerState       `protobuf:"bytes,6,rep,name=controllerStates,proto3" json:"controllerStates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PlayerControllerStates map[uint32]*unrealpb.PlayerControllerState `protobuf:"bytes,7,rep,name=playerControllerStates,proto3" json:"playerControllerStates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ActorComponentStates   map[uint32]*unrealpb.ActorComponentState   `protobuf:"bytes,8,rep,name=actorComponentStates,proto3" json:"actorComponentStates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SceneComponentStates   map[uint32]*unrealpb.SceneComponentState   `protobuf:"bytes,9,rep,name=sceneComponentStates,proto3" json:"sceneComponentStates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestRepChannelData) Reset() {
	*x = TestRepChannelData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRepChannelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRepChannelData) ProtoMessage() {}

func (x *TestRepChannelData) ProtoReflect() protoreflect.Message {
	mi := &file_tps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRepChannelData.ProtoReflect.Descriptor instead.
func (*TestRepChannelData) Descriptor() ([]byte, []int) {
	return file_tps_proto_rawDescGZIP(), []int{0}
}

func (x *TestRepChannelData) GetGameState() *unrealpb.GameStateBase {
	if x != nil {
		return x.GameState
	}
	return nil
}

func (x *TestRepChannelData) GetActorStates() map[uint32]*unrealpb.ActorState {
	if x != nil {
		return x.ActorStates
	}
	return nil
}

func (x *TestRepChannelData) GetPawnStates() map[uint32]*unrealpb.PawnState {
	if x != nil {
		return x.PawnStates
	}
	return nil
}

func (x *TestRepChannelData) GetCharacterStates() map[uint32]*unrealpb.CharacterState {
	if x != nil {
		return x.CharacterStates
	}
	return nil
}

func (x *TestRepChannelData) GetPlayerStates() map[uint32]*unrealpb.PlayerState {
	if x != nil {
		return x.PlayerStates
	}
	return nil
}

func (x *TestRepChannelData) GetControllerStates() map[uint32]*unrealpb.ControllerState {
	if x != nil {
		return x.ControllerStates
	}
	return nil
}

func (x *TestRepChannelData) GetPlayerControllerStates() map[uint32]*unrealpb.PlayerControllerState {
	if x != nil {
		return x.PlayerControllerStates
	}
	return nil
}

func (x *TestRepChannelData) GetActorComponentStates() map[uint32]*unrealpb.ActorComponentState {
	if x != nil {
		return x.ActorComponentStates
	}
	return nil
}

func (x *TestRepChannelData) GetSceneComponentStates() map[uint32]*unrealpb.SceneComponentState {
	if x != nil {
		return x.SceneComponentStates
	}
	return nil
}

type GlobalChannelData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameState *unrealpb.GameStateBase `protobuf:"bytes,1,opt,name=gameState,proto3" json:"gameState,omitempty"`
}

func (x *GlobalChannelData) Reset() {
	*x = GlobalChannelData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalChannelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalChannelData) ProtoMessage() {}

func (x *GlobalChannelData) ProtoReflect() protoreflect.Message {
	mi := &file_tps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalChannelData.ProtoReflect.Descriptor instead.
func (*GlobalChannelData) Descriptor() ([]byte, []int) {
	return file_tps_proto_rawDescGZIP(), []int{1}
}

func (x *GlobalChannelData) GetGameState() *unrealpb.GameStateBase {
	if x != nil {
		return x.GameState
	}
	return nil
}

var File_tps_proto protoreflect.FileDescriptor

var file_tps_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x70, 0x73,
	0x70, 0x62, 0x1a, 0x13, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x0c, 0x0a, 0x12, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35,
	0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x52, 0x09, 0x67, 0x61, 0x6d, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x70, 0x73,
	0x70, 0x62, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x0a, 0x70, 0x61, 0x77, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x70, 0x73, 0x70, 0x62, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x50, 0x61, 0x77, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x77, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x58,
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x74, 0x70, 0x73, 0x70, 0x62, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x74, 0x70, 0x73, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x74, 0x70, 0x73, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x6d, 0x0a, 0x16, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x74, 0x70, 0x73, 0x70, 0x62, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x67, 0x0a, 0x14, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x70, 0x73, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x67,
	0x0a, 0x14, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74,
	0x70, 0x73, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x14, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x54, 0x0a, 0x10, 0x41, 0x63, 0x74, 0x6f, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75,
	0x6e, 0x72, 0x65, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x52, 0x0a,
	0x0f, 0x50, 0x61, 0x77, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x77,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x5c, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x6e, 0x72,
	0x65, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x56, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x6c, 0x70, 0x62,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5e, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x6a, 0x0a, 0x1b, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x6c,
	0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x66, 0x0a, 0x19, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x66, 0x0a, 0x19, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x75, 0x6e, 0x72, 0x65,
	0x61, 0x6c, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x4a, 0x0a, 0x11, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x6e,
	0x72, 0x65, 0x61, 0x6c, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42,
	0x35, 0x5a, 0x33, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x64, 0x2e, 0x63, 0x6c, 0x65, 0x77,
	0x63, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x64, 0x2d, 0x75, 0x65, 0x2d, 0x74, 0x70, 0x73,
	0x2f, 0x74, 0x70, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tps_proto_rawDescOnce sync.Once
	file_tps_proto_rawDescData = file_tps_proto_rawDesc
)

func file_tps_proto_rawDescGZIP() []byte {
	file_tps_proto_rawDescOnce.Do(func() {
		file_tps_proto_rawDescData = protoimpl.X.CompressGZIP(file_tps_proto_rawDescData)
	})
	return file_tps_proto_rawDescData
}

var file_tps_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_tps_proto_goTypes = []interface{}{
	(*TestRepChannelData)(nil),             // 0: tpspb.TestRepChannelData
	(*GlobalChannelData)(nil),              // 1: tpspb.GlobalChannelData
	nil,                                    // 2: tpspb.TestRepChannelData.ActorStatesEntry
	nil,                                    // 3: tpspb.TestRepChannelData.PawnStatesEntry
	nil,                                    // 4: tpspb.TestRepChannelData.CharacterStatesEntry
	nil,                                    // 5: tpspb.TestRepChannelData.PlayerStatesEntry
	nil,                                    // 6: tpspb.TestRepChannelData.ControllerStatesEntry
	nil,                                    // 7: tpspb.TestRepChannelData.PlayerControllerStatesEntry
	nil,                                    // 8: tpspb.TestRepChannelData.ActorComponentStatesEntry
	nil,                                    // 9: tpspb.TestRepChannelData.SceneComponentStatesEntry
	(*unrealpb.GameStateBase)(nil),         // 10: unrealpb.GameStateBase
	(*unrealpb.ActorState)(nil),            // 11: unrealpb.ActorState
	(*unrealpb.PawnState)(nil),             // 12: unrealpb.PawnState
	(*unrealpb.CharacterState)(nil),        // 13: unrealpb.CharacterState
	(*unrealpb.PlayerState)(nil),           // 14: unrealpb.PlayerState
	(*unrealpb.ControllerState)(nil),       // 15: unrealpb.ControllerState
	(*unrealpb.PlayerControllerState)(nil), // 16: unrealpb.PlayerControllerState
	(*unrealpb.ActorComponentState)(nil),   // 17: unrealpb.ActorComponentState
	(*unrealpb.SceneComponentState)(nil),   // 18: unrealpb.SceneComponentState
}
var file_tps_proto_depIdxs = []int32{
	10, // 0: tpspb.TestRepChannelData.gameState:type_name -> unrealpb.GameStateBase
	2,  // 1: tpspb.TestRepChannelData.actorStates:type_name -> tpspb.TestRepChannelData.ActorStatesEntry
	3,  // 2: tpspb.TestRepChannelData.pawnStates:type_name -> tpspb.TestRepChannelData.PawnStatesEntry
	4,  // 3: tpspb.TestRepChannelData.characterStates:type_name -> tpspb.TestRepChannelData.CharacterStatesEntry
	5,  // 4: tpspb.TestRepChannelData.playerStates:type_name -> tpspb.TestRepChannelData.PlayerStatesEntry
	6,  // 5: tpspb.TestRepChannelData.controllerStates:type_name -> tpspb.TestRepChannelData.ControllerStatesEntry
	7,  // 6: tpspb.TestRepChannelData.playerControllerStates:type_name -> tpspb.TestRepChannelData.PlayerControllerStatesEntry
	8,  // 7: tpspb.TestRepChannelData.actorComponentStates:type_name -> tpspb.TestRepChannelData.ActorComponentStatesEntry
	9,  // 8: tpspb.TestRepChannelData.sceneComponentStates:type_name -> tpspb.TestRepChannelData.SceneComponentStatesEntry
	10, // 9: tpspb.GlobalChannelData.gameState:type_name -> unrealpb.GameStateBase
	11, // 10: tpspb.TestRepChannelData.ActorStatesEntry.value:type_name -> unrealpb.ActorState
	12, // 11: tpspb.TestRepChannelData.PawnStatesEntry.value:type_name -> unrealpb.PawnState
	13, // 12: tpspb.TestRepChannelData.CharacterStatesEntry.value:type_name -> unrealpb.CharacterState
	14, // 13: tpspb.TestRepChannelData.PlayerStatesEntry.value:type_name -> unrealpb.PlayerState
	15, // 14: tpspb.TestRepChannelData.ControllerStatesEntry.value:type_name -> unrealpb.ControllerState
	16, // 15: tpspb.TestRepChannelData.PlayerControllerStatesEntry.value:type_name -> unrealpb.PlayerControllerState
	17, // 16: tpspb.TestRepChannelData.ActorComponentStatesEntry.value:type_name -> unrealpb.ActorComponentState
	18, // 17: tpspb.TestRepChannelData.SceneComponentStatesEntry.value:type_name -> unrealpb.SceneComponentState
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_tps_proto_init() }
func file_tps_proto_init() {
	if File_tps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRepChannelData); i {
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
		file_tps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalChannelData); i {
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
			RawDescriptor: file_tps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tps_proto_goTypes,
		DependencyIndexes: file_tps_proto_depIdxs,
		MessageInfos:      file_tps_proto_msgTypes,
	}.Build()
	File_tps_proto = out.File
	file_tps_proto_rawDesc = nil
	file_tps_proto_goTypes = nil
	file_tps_proto_depIdxs = nil
}
