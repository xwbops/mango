// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: table.proto

//import "types.proto";

package table

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

type CMDTable int32

const (
	CMDTable_IDApplyReq            CMDTable = 1  //
	CMDTable_IDApplyRsp            CMDTable = 2  //
	CMDTable_IDReleaseReq          CMDTable = 3  //
	CMDTable_IDReleaseRsp          CMDTable = 4  //
	CMDTable_IDSetPlayerToTableReq CMDTable = 5  //
	CMDTable_IDSetPlayerToTableRsp CMDTable = 6  //
	CMDTable_IDMatchTableReq       CMDTable = 7  //
	CMDTable_IDMatchTableRsp       CMDTable = 8  //
	CMDTable_IDGameMessage         CMDTable = 9  //
	CMDTable_IDWriteGameScore      CMDTable = 10 //
	CMDTable_IDGameOver            CMDTable = 11 //
)

// Enum value maps for CMDTable.
var (
	CMDTable_name = map[int32]string{
		1:  "IDApplyReq",
		2:  "IDApplyRsp",
		3:  "IDReleaseReq",
		4:  "IDReleaseRsp",
		5:  "IDSetPlayerToTableReq",
		6:  "IDSetPlayerToTableRsp",
		7:  "IDMatchTableReq",
		8:  "IDMatchTableRsp",
		9:  "IDGameMessage",
		10: "IDWriteGameScore",
		11: "IDGameOver",
	}
	CMDTable_value = map[string]int32{
		"IDApplyReq":            1,
		"IDApplyRsp":            2,
		"IDReleaseReq":          3,
		"IDReleaseRsp":          4,
		"IDSetPlayerToTableReq": 5,
		"IDSetPlayerToTableRsp": 6,
		"IDMatchTableReq":       7,
		"IDMatchTableRsp":       8,
		"IDGameMessage":         9,
		"IDWriteGameScore":      10,
		"IDGameOver":            11,
	}
)

func (x CMDTable) Enum() *CMDTable {
	p := new(CMDTable)
	*p = x
	return p
}

func (x CMDTable) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMDTable) Descriptor() protoreflect.EnumDescriptor {
	return file_table_proto_enumTypes[0].Descriptor()
}

func (CMDTable) Type() protoreflect.EnumType {
	return &file_table_proto_enumTypes[0]
}

func (x CMDTable) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CMDTable) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CMDTable(num)
	return nil
}

// Deprecated: Use CMDTable.Descriptor instead.
func (CMDTable) EnumDescriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{0}
}

type ApplyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyCount *uint32 `protobuf:"varint,1,opt,name=apply_count,json=applyCount" json:"apply_count,omitempty"`
}

func (x *ApplyReq) Reset() {
	*x = ApplyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyReq) ProtoMessage() {}

func (x *ApplyReq) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyReq.ProtoReflect.Descriptor instead.
func (*ApplyReq) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{0}
}

func (x *ApplyReq) GetApplyCount() uint32 {
	if x != nil && x.ApplyCount != nil {
		return *x.ApplyCount
	}
	return 0
}

type ApplyRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyCount *uint32  `protobuf:"varint,1,opt,name=apply_count,json=applyCount" json:"apply_count,omitempty"`
	TableIds   []uint64 `protobuf:"varint,2,rep,name=table_ids,json=tableIds" json:"table_ids,omitempty"`
}

func (x *ApplyRsp) Reset() {
	*x = ApplyRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyRsp) ProtoMessage() {}

func (x *ApplyRsp) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyRsp.ProtoReflect.Descriptor instead.
func (*ApplyRsp) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{1}
}

func (x *ApplyRsp) GetApplyCount() uint32 {
	if x != nil && x.ApplyCount != nil {
		return *x.ApplyCount
	}
	return 0
}

func (x *ApplyRsp) GetTableIds() []uint64 {
	if x != nil {
		return x.TableIds
	}
	return nil
}

type ReleaseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReleaseCount *uint32  `protobuf:"varint,1,opt,name=release_count,json=releaseCount" json:"release_count,omitempty"`
	TableIds     []uint64 `protobuf:"varint,2,rep,name=table_ids,json=tableIds" json:"table_ids,omitempty"`
}

func (x *ReleaseReq) Reset() {
	*x = ReleaseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseReq) ProtoMessage() {}

func (x *ReleaseReq) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseReq.ProtoReflect.Descriptor instead.
func (*ReleaseReq) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{2}
}

func (x *ReleaseReq) GetReleaseCount() uint32 {
	if x != nil && x.ReleaseCount != nil {
		return *x.ReleaseCount
	}
	return 0
}

func (x *ReleaseReq) GetTableIds() []uint64 {
	if x != nil {
		return x.TableIds
	}
	return nil
}

type ReleaseRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReleaseRsp) Reset() {
	*x = ReleaseRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleaseRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleaseRsp) ProtoMessage() {}

func (x *ReleaseRsp) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleaseRsp.ProtoReflect.Descriptor instead.
func (*ReleaseRsp) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{3}
}

type SetPlayerToTableReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableId    *uint64 `protobuf:"varint,1,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	UserId     *uint64 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	SeatId     *uint32 `protobuf:"varint,3,opt,name=seat_id,json=seatId" json:"seat_id,omitempty"`
	Gateconnid *uint64 `protobuf:"varint,4,opt,name=gateconnid" json:"gateconnid,omitempty"`
}

func (x *SetPlayerToTableReq) Reset() {
	*x = SetPlayerToTableReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPlayerToTableReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPlayerToTableReq) ProtoMessage() {}

func (x *SetPlayerToTableReq) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPlayerToTableReq.ProtoReflect.Descriptor instead.
func (*SetPlayerToTableReq) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{4}
}

func (x *SetPlayerToTableReq) GetTableId() uint64 {
	if x != nil && x.TableId != nil {
		return *x.TableId
	}
	return 0
}

func (x *SetPlayerToTableReq) GetUserId() uint64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *SetPlayerToTableReq) GetSeatId() uint32 {
	if x != nil && x.SeatId != nil {
		return *x.SeatId
	}
	return 0
}

func (x *SetPlayerToTableReq) GetGateconnid() uint64 {
	if x != nil && x.Gateconnid != nil {
		return *x.Gateconnid
	}
	return 0
}

type SetPlayerToTableRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetPlayerToTableRsp) Reset() {
	*x = SetPlayerToTableRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPlayerToTableRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPlayerToTableRsp) ProtoMessage() {}

func (x *SetPlayerToTableRsp) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPlayerToTableRsp.ProtoReflect.Descriptor instead.
func (*SetPlayerToTableRsp) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{5}
}

type MatchTableReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableId *uint64  `protobuf:"varint,1,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	Players []uint64 `protobuf:"varint,2,rep,name=players" json:"players,omitempty"`
}

func (x *MatchTableReq) Reset() {
	*x = MatchTableReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchTableReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchTableReq) ProtoMessage() {}

func (x *MatchTableReq) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchTableReq.ProtoReflect.Descriptor instead.
func (*MatchTableReq) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{6}
}

func (x *MatchTableReq) GetTableId() uint64 {
	if x != nil && x.TableId != nil {
		return *x.TableId
	}
	return 0
}

func (x *MatchTableReq) GetPlayers() []uint64 {
	if x != nil {
		return x.Players
	}
	return nil
}

type MatchTableRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MatchTableRsp) Reset() {
	*x = MatchTableRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchTableRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchTableRsp) ProtoMessage() {}

func (x *MatchTableRsp) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchTableRsp.ProtoReflect.Descriptor instead.
func (*MatchTableRsp) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{7}
}

type GameMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubCmdid *uint32 `protobuf:"varint,1,opt,name=sub_cmdid,json=subCmdid" json:"sub_cmdid,omitempty"`
	Data     []byte  `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (x *GameMessage) Reset() {
	*x = GameMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameMessage) ProtoMessage() {}

func (x *GameMessage) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameMessage.ProtoReflect.Descriptor instead.
func (*GameMessage) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{8}
}

func (x *GameMessage) GetSubCmdid() uint32 {
	if x != nil && x.SubCmdid != nil {
		return *x.SubCmdid
	}
	return 0
}

func (x *GameMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type WriteGameScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableId *uint64 `protobuf:"varint,1,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
}

func (x *WriteGameScore) Reset() {
	*x = WriteGameScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteGameScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteGameScore) ProtoMessage() {}

func (x *WriteGameScore) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteGameScore.ProtoReflect.Descriptor instead.
func (*WriteGameScore) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{9}
}

func (x *WriteGameScore) GetTableId() uint64 {
	if x != nil && x.TableId != nil {
		return *x.TableId
	}
	return 0
}

type GameOver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableId *uint64 `protobuf:"varint,1,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
}

func (x *GameOver) Reset() {
	*x = GameOver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameOver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameOver) ProtoMessage() {}

func (x *GameOver) ProtoReflect() protoreflect.Message {
	mi := &file_table_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameOver.ProtoReflect.Descriptor instead.
func (*GameOver) Descriptor() ([]byte, []int) {
	return file_table_proto_rawDescGZIP(), []int{10}
}

func (x *GameOver) GetTableId() uint64 {
	if x != nil && x.TableId != nil {
		return *x.TableId
	}
	return 0
}

var File_table_proto protoreflect.FileDescriptor

var file_table_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62,
	0x73, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x2b, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x48, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x73, 0x70,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x4e,
	0x0a, 0x0a, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x0c,
	0x0a, 0x0a, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x73, 0x70, 0x22, 0x82, 0x01, 0x0a,
	0x13, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x61, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x65, 0x61, 0x74, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x69,
	0x64, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x73, 0x70, 0x22, 0x44, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x0f,
	0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x73, 0x70, 0x22,
	0x3e, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6d, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x73, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x2b, 0x0a, 0x0e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x08,
	0x47, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x49, 0x64, 0x2a, 0xe7, 0x01, 0x0a, 0x08, 0x43, 0x4d, 0x44, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x44, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x10, 0x01,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x44, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x73, 0x70, 0x10, 0x02,
	0x12, 0x10, 0x0a, 0x0c, 0x49, 0x44, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x44, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52,
	0x73, 0x70, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x44, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x54, 0x6f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x10, 0x05, 0x12,
	0x19, 0x0a, 0x15, 0x49, 0x44, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x73, 0x70, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x44,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x10, 0x07, 0x12,
	0x13, 0x0a, 0x0f, 0x49, 0x44, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x73, 0x70, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x44, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x09, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x44, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x10, 0x0a, 0x12, 0x0e, 0x0a,
	0x0a, 0x49, 0x44, 0x47, 0x61, 0x6d, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x10, 0x0b, 0x42, 0x08, 0x5a,
	0x06, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65,
}

var (
	file_table_proto_rawDescOnce sync.Once
	file_table_proto_rawDescData = file_table_proto_rawDesc
)

func file_table_proto_rawDescGZIP() []byte {
	file_table_proto_rawDescOnce.Do(func() {
		file_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_table_proto_rawDescData)
	})
	return file_table_proto_rawDescData
}

var file_table_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_table_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_table_proto_goTypes = []interface{}{
	(CMDTable)(0),               // 0: bs.table.CMDTable
	(*ApplyReq)(nil),            // 1: bs.table.ApplyReq
	(*ApplyRsp)(nil),            // 2: bs.table.ApplyRsp
	(*ReleaseReq)(nil),          // 3: bs.table.ReleaseReq
	(*ReleaseRsp)(nil),          // 4: bs.table.ReleaseRsp
	(*SetPlayerToTableReq)(nil), // 5: bs.table.SetPlayerToTableReq
	(*SetPlayerToTableRsp)(nil), // 6: bs.table.SetPlayerToTableRsp
	(*MatchTableReq)(nil),       // 7: bs.table.MatchTableReq
	(*MatchTableRsp)(nil),       // 8: bs.table.MatchTableRsp
	(*GameMessage)(nil),         // 9: bs.table.GameMessage
	(*WriteGameScore)(nil),      // 10: bs.table.WriteGameScore
	(*GameOver)(nil),            // 11: bs.table.GameOver
}
var file_table_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_table_proto_init() }
func file_table_proto_init() {
	if File_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyReq); i {
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
		file_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyRsp); i {
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
		file_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseReq); i {
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
		file_table_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleaseRsp); i {
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
		file_table_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPlayerToTableReq); i {
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
		file_table_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPlayerToTableRsp); i {
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
		file_table_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchTableReq); i {
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
		file_table_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchTableRsp); i {
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
		file_table_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameMessage); i {
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
		file_table_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteGameScore); i {
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
		file_table_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameOver); i {
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
			RawDescriptor: file_table_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_table_proto_goTypes,
		DependencyIndexes: file_table_proto_depIdxs,
		EnumInfos:         file_table_proto_enumTypes,
		MessageInfos:      file_table_proto_msgTypes,
	}.Build()
	File_table_proto = out.File
	file_table_proto_rawDesc = nil
	file_table_proto_goTypes = nil
	file_table_proto_depIdxs = nil
}
