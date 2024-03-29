// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: game/v1/game.proto

package gamev1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Move int32

const (
	Move_MOVE_UNSPECIFIED Move = 0
	Move_MOVE_ROCK        Move = 1
	Move_MOVE_PAPER       Move = 2
	Move_MOVE_SCISSORS    Move = 3
)

// Enum value maps for Move.
var (
	Move_name = map[int32]string{
		0: "MOVE_UNSPECIFIED",
		1: "MOVE_ROCK",
		2: "MOVE_PAPER",
		3: "MOVE_SCISSORS",
	}
	Move_value = map[string]int32{
		"MOVE_UNSPECIFIED": 0,
		"MOVE_ROCK":        1,
		"MOVE_PAPER":       2,
		"MOVE_SCISSORS":    3,
	}
)

func (x Move) Enum() *Move {
	p := new(Move)
	*p = x
	return p
}

func (x Move) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Move) Descriptor() protoreflect.EnumDescriptor {
	return file_game_v1_game_proto_enumTypes[0].Descriptor()
}

func (Move) Type() protoreflect.EnumType {
	return &file_game_v1_game_proto_enumTypes[0]
}

func (x Move) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Move.Descriptor instead.
func (Move) EnumDescriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{0}
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Wins     int32  `protobuf:"varint,3,opt,name=wins,proto3" json:"wins,omitempty"`
	Losses   int32  `protobuf:"varint,4,opt,name=losses,proto3" json:"losses,omitempty"`
	Draws    int32  `protobuf:"varint,5,opt,name=draws,proto3" json:"draws,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Session) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *Session) GetWins() int32 {
	if x != nil {
		return x.Wins
	}
	return 0
}

func (x *Session) GetLosses() int32 {
	if x != nil {
		return x.Losses
	}
	return 0
}

func (x *Session) GetDraws() int32 {
	if x != nil {
		return x.Draws
	}
	return 0
}

type PlayGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId  string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	PlayerMove Move   `protobuf:"varint,2,opt,name=player_move,json=playerMove,proto3,enum=game.v1.Move" json:"player_move,omitempty"`
}

func (x *PlayGameRequest) Reset() {
	*x = PlayGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayGameRequest) ProtoMessage() {}

func (x *PlayGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayGameRequest.ProtoReflect.Descriptor instead.
func (*PlayGameRequest) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{1}
}

func (x *PlayGameRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *PlayGameRequest) GetPlayerMove() Move {
	if x != nil {
		return x.PlayerMove
	}
	return Move_MOVE_UNSPECIFIED
}

type PlayGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComputerMove Move     `protobuf:"varint,1,opt,name=computer_move,json=computerMove,proto3,enum=game.v1.Move" json:"computer_move,omitempty"`
	Session      *Session `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *PlayGameResponse) Reset() {
	*x = PlayGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayGameResponse) ProtoMessage() {}

func (x *PlayGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayGameResponse.ProtoReflect.Descriptor instead.
func (*PlayGameResponse) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{2}
}

func (x *PlayGameResponse) GetComputerMove() Move {
	if x != nil {
		return x.ComputerMove
	}
	return Move_MOVE_UNSPECIFIED
}

func (x *PlayGameResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type StartSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId *string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3,oneof" json:"player_id,omitempty"`
}

func (x *StartSessionRequest) Reset() {
	*x = StartSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartSessionRequest) ProtoMessage() {}

func (x *StartSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartSessionRequest.ProtoReflect.Descriptor instead.
func (*StartSessionRequest) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{3}
}

func (x *StartSessionRequest) GetPlayerId() string {
	if x != nil && x.PlayerId != nil {
		return *x.PlayerId
	}
	return ""
}

type StartSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *StartSessionResponse) Reset() {
	*x = StartSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartSessionResponse) ProtoMessage() {}

func (x *StartSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartSessionResponse.ProtoReflect.Descriptor instead.
func (*StartSessionResponse) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{4}
}

func (x *StartSessionResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type EndSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *EndSessionRequest) Reset() {
	*x = EndSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndSessionRequest) ProtoMessage() {}

func (x *EndSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndSessionRequest.ProtoReflect.Descriptor instead.
func (*EndSessionRequest) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{5}
}

func (x *EndSessionRequest) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

type EndSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *EndSessionResponse) Reset() {
	*x = EndSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_v1_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndSessionResponse) ProtoMessage() {}

func (x *EndSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndSessionResponse.ProtoReflect.Descriptor instead.
func (*EndSessionResponse) Descriptor() ([]byte, []int) {
	return file_game_v1_game_proto_rawDescGZIP(), []int{6}
}

func (x *EndSessionResponse) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

var File_game_v1_game_proto protoreflect.FileDescriptor

var file_game_v1_game_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x78, 0x0a,
	0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x77, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x64, 0x72, 0x61, 0x77, 0x73, 0x22, 0x60, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x0a, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x22, 0x72, 0x0a, 0x10, 0x50, 0x6c, 0x61,
	0x79, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x6f, 0x76, 0x65, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x76,
	0x65, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a,
	0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x11, 0x45, 0x6e, 0x64, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x12, 0x45, 0x6e, 0x64,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2a, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x4e, 0x0a, 0x04, 0x4d,
	0x6f, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x4f, 0x56,
	0x45, 0x5f, 0x52, 0x4f, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x4f, 0x56, 0x45,
	0x5f, 0x50, 0x41, 0x50, 0x45, 0x52, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x4f, 0x56, 0x45,
	0x5f, 0x53, 0x43, 0x49, 0x53, 0x53, 0x4f, 0x52, 0x53, 0x10, 0x03, 0x42, 0x9a, 0x01, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x47, 0x61, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x73, 0x61, 0x72, 0x69, 0x74, 0x65, 0x6b, 0x2f, 0x70, 0x61,
	0x70, 0x65, 0x72, 0x2d, 0x72, 0x6f, 0x63, 0x6b, 0x2d, 0x73, 0x63, 0x69, 0x73, 0x73, 0x6f, 0x72,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67,
	0x61, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x61, 0x6d, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x47, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07,
	0x47, 0x61, 0x6d, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x47, 0x61, 0x6d, 0x65, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08,
	0x47, 0x61, 0x6d, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_v1_game_proto_rawDescOnce sync.Once
	file_game_v1_game_proto_rawDescData = file_game_v1_game_proto_rawDesc
)

func file_game_v1_game_proto_rawDescGZIP() []byte {
	file_game_v1_game_proto_rawDescOnce.Do(func() {
		file_game_v1_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_v1_game_proto_rawDescData)
	})
	return file_game_v1_game_proto_rawDescData
}

var file_game_v1_game_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_v1_game_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_game_v1_game_proto_goTypes = []interface{}{
	(Move)(0),                    // 0: game.v1.Move
	(*Session)(nil),              // 1: game.v1.Session
	(*PlayGameRequest)(nil),      // 2: game.v1.PlayGameRequest
	(*PlayGameResponse)(nil),     // 3: game.v1.PlayGameResponse
	(*StartSessionRequest)(nil),  // 4: game.v1.StartSessionRequest
	(*StartSessionResponse)(nil), // 5: game.v1.StartSessionResponse
	(*EndSessionRequest)(nil),    // 6: game.v1.EndSessionRequest
	(*EndSessionResponse)(nil),   // 7: game.v1.EndSessionResponse
}
var file_game_v1_game_proto_depIdxs = []int32{
	0, // 0: game.v1.PlayGameRequest.player_move:type_name -> game.v1.Move
	0, // 1: game.v1.PlayGameResponse.computer_move:type_name -> game.v1.Move
	1, // 2: game.v1.PlayGameResponse.session:type_name -> game.v1.Session
	1, // 3: game.v1.StartSessionResponse.session:type_name -> game.v1.Session
	1, // 4: game.v1.EndSessionRequest.session:type_name -> game.v1.Session
	1, // 5: game.v1.EndSessionResponse.session:type_name -> game.v1.Session
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_game_v1_game_proto_init() }
func file_game_v1_game_proto_init() {
	if File_game_v1_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_v1_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_game_v1_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayGameRequest); i {
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
		file_game_v1_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayGameResponse); i {
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
		file_game_v1_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartSessionRequest); i {
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
		file_game_v1_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartSessionResponse); i {
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
		file_game_v1_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndSessionRequest); i {
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
		file_game_v1_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndSessionResponse); i {
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
	file_game_v1_game_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_v1_game_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_v1_game_proto_goTypes,
		DependencyIndexes: file_game_v1_game_proto_depIdxs,
		EnumInfos:         file_game_v1_game_proto_enumTypes,
		MessageInfos:      file_game_v1_game_proto_msgTypes,
	}.Build()
	File_game_v1_game_proto = out.File
	file_game_v1_game_proto_rawDesc = nil
	file_game_v1_game_proto_goTypes = nil
	file_game_v1_game_proto_depIdxs = nil
}
