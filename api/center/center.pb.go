// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: center.proto

package center

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

type CMDCenter int32

const (
	CMDCenter_IDAppRegReq     CMDCenter = 1 //服务注册
	CMDCenter_IDAppRegRsp     CMDCenter = 2 //服务注册
	CMDCenter_IDAppState      CMDCenter = 3 //服务状态
	CMDCenter_IDAppOfflineReq CMDCenter = 4 //服务离线
	CMDCenter_IDPulseNotify   CMDCenter = 5 //服务心跳
	CMDCenter_IDAppUpdateReq  CMDCenter = 6 //更新请求
)

// Enum value maps for CMDCenter.
var (
	CMDCenter_name = map[int32]string{
		1: "IDAppRegReq",
		2: "IDAppRegRsp",
		3: "IDAppState",
		4: "IDAppOfflineReq",
		5: "IDPulseNotify",
		6: "IDAppUpdateReq",
	}
	CMDCenter_value = map[string]int32{
		"IDAppRegReq":     1,
		"IDAppRegRsp":     2,
		"IDAppState":      3,
		"IDAppOfflineReq": 4,
		"IDPulseNotify":   5,
		"IDAppUpdateReq":  6,
	}
)

func (x CMDCenter) Enum() *CMDCenter {
	p := new(CMDCenter)
	*p = x
	return p
}

func (x CMDCenter) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMDCenter) Descriptor() protoreflect.EnumDescriptor {
	return file_center_proto_enumTypes[0].Descriptor()
}

func (CMDCenter) Type() protoreflect.EnumType {
	return &file_center_proto_enumTypes[0]
}

func (x CMDCenter) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CMDCenter) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CMDCenter(num)
	return nil
}

// Deprecated: Use CMDCenter.Descriptor instead.
func (CMDCenter) EnumDescriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{0}
}

type AppStateNotifyState int32

const (
	AppStateNotify_Online   AppStateNotifyState = 1
	AppStateNotify_ReOnline AppStateNotifyState = 2
	AppStateNotify_OffLine  AppStateNotifyState = 3
)

// Enum value maps for AppStateNotifyState.
var (
	AppStateNotifyState_name = map[int32]string{
		1: "Online",
		2: "ReOnline",
		3: "OffLine",
	}
	AppStateNotifyState_value = map[string]int32{
		"Online":   1,
		"ReOnline": 2,
		"OffLine":  3,
	}
)

func (x AppStateNotifyState) Enum() *AppStateNotifyState {
	p := new(AppStateNotifyState)
	*p = x
	return p
}

func (x AppStateNotifyState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppStateNotifyState) Descriptor() protoreflect.EnumDescriptor {
	return file_center_proto_enumTypes[1].Descriptor()
}

func (AppStateNotifyState) Type() protoreflect.EnumType {
	return &file_center_proto_enumTypes[1]
}

func (x AppStateNotifyState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AppStateNotifyState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AppStateNotifyState(num)
	return nil
}

// Deprecated: Use AppStateNotifyState.Descriptor instead.
func (AppStateNotifyState) EnumDescriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{3, 0}
}

type AppPulseNotify_PulseAction int32

const (
	AppPulseNotify_HeartBeatReq AppPulseNotify_PulseAction = 0
	AppPulseNotify_HeartBeatRsp AppPulseNotify_PulseAction = 1
	AppPulseNotify_LogoutReq    AppPulseNotify_PulseAction = 2
	AppPulseNotify_LogoutRsp    AppPulseNotify_PulseAction = 3
	AppPulseNotify_AnotherLogin AppPulseNotify_PulseAction = 5 //顶掉了。
)

// Enum value maps for AppPulseNotify_PulseAction.
var (
	AppPulseNotify_PulseAction_name = map[int32]string{
		0: "HeartBeatReq",
		1: "HeartBeatRsp",
		2: "LogoutReq",
		3: "LogoutRsp",
		5: "AnotherLogin",
	}
	AppPulseNotify_PulseAction_value = map[string]int32{
		"HeartBeatReq": 0,
		"HeartBeatRsp": 1,
		"LogoutReq":    2,
		"LogoutRsp":    3,
		"AnotherLogin": 5,
	}
)

func (x AppPulseNotify_PulseAction) Enum() *AppPulseNotify_PulseAction {
	p := new(AppPulseNotify_PulseAction)
	*p = x
	return p
}

func (x AppPulseNotify_PulseAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppPulseNotify_PulseAction) Descriptor() protoreflect.EnumDescriptor {
	return file_center_proto_enumTypes[2].Descriptor()
}

func (AppPulseNotify_PulseAction) Type() protoreflect.EnumType {
	return &file_center_proto_enumTypes[2]
}

func (x AppPulseNotify_PulseAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AppPulseNotify_PulseAction) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AppPulseNotify_PulseAction(num)
	return nil
}

// Deprecated: Use AppPulseNotify_PulseAction.Descriptor instead.
func (AppPulseNotify_PulseAction) EnumDescriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{5, 0}
}

type AppUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogAddr *string `protobuf:"bytes,1,opt,name=log_addr,json=logAddr" json:"log_addr,omitempty"` // 日志地址信息(log_ip:port,log_level)
}

func (x *AppUpdateReq) Reset() {
	*x = AppUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_center_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppUpdateReq) ProtoMessage() {}

func (x *AppUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppUpdateReq.ProtoReflect.Descriptor instead.
func (*AppUpdateReq) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{0}
}

func (x *AppUpdateReq) GetLogAddr() string {
	if x != nil && x.LogAddr != nil {
		return *x.LogAddr
	}
	return ""
}

//服务注册
type RegisterAppReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthKey    *string `protobuf:"bytes,1,req,name=auth_key,json=authKey" json:"auth_key,omitempty"`
	AttData    *string `protobuf:"bytes,2,opt,name=att_data,json=attData" json:"att_data,omitempty"`
	MyAddress  *string `protobuf:"bytes,3,opt,name=my_address,json=myAddress" json:"my_address,omitempty"`
	AppType    *uint32 `protobuf:"varint,4,opt,name=app_type,json=appType" json:"app_type,omitempty"` //Router 或其他
	AppId      *uint32 `protobuf:"varint,5,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	ReregToken *string `protobuf:"bytes,6,opt,name=rereg_token,json=reregToken" json:"rereg_token,omitempty"` //如果中间网络断开了,可以使用rereg_token强行再次注册
	AppName    *string `protobuf:"bytes,7,opt,name=app_name,json=appName" json:"app_name,omitempty"`          //app的名称(一般为进程名)
}

func (x *RegisterAppReq) Reset() {
	*x = RegisterAppReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_center_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAppReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAppReq) ProtoMessage() {}

func (x *RegisterAppReq) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAppReq.ProtoReflect.Descriptor instead.
func (*RegisterAppReq) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterAppReq) GetAuthKey() string {
	if x != nil && x.AuthKey != nil {
		return *x.AuthKey
	}
	return ""
}

func (x *RegisterAppReq) GetAttData() string {
	if x != nil && x.AttData != nil {
		return *x.AttData
	}
	return ""
}

func (x *RegisterAppReq) GetMyAddress() string {
	if x != nil && x.MyAddress != nil {
		return *x.MyAddress
	}
	return ""
}

func (x *RegisterAppReq) GetAppType() uint32 {
	if x != nil && x.AppType != nil {
		return *x.AppType
	}
	return 0
}

func (x *RegisterAppReq) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *RegisterAppReq) GetReregToken() string {
	if x != nil && x.ReregToken != nil {
		return *x.ReregToken
	}
	return ""
}

func (x *RegisterAppReq) GetAppName() string {
	if x != nil && x.AppName != nil {
		return *x.AppName
	}
	return ""
}

//服务注册
type RegisterAppRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegResult  *uint32 `protobuf:"varint,1,req,name=reg_result,json=regResult" json:"reg_result,omitempty"`   //0表示成功，其它为错误码(rereg_token为出错内容)
	ReregToken *string `protobuf:"bytes,2,opt,name=rereg_token,json=reregToken" json:"rereg_token,omitempty"` //如果中间网络断开了,可以使用rereg_token强行再次注册
	CenterId   *uint32 `protobuf:"varint,3,opt,name=center_id,json=centerId" json:"center_id,omitempty"`
	AppType    *uint32 `protobuf:"varint,4,opt,name=app_type,json=appType" json:"app_type,omitempty"` //Router 或其他
	AppId      *uint32 `protobuf:"varint,5,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	AppName    *string `protobuf:"bytes,6,opt,name=app_name,json=appName" json:"app_name,omitempty"`          //app的名称(一般为进程名)
	AppAddress *string `protobuf:"bytes,7,opt,name=app_address,json=appAddress" json:"app_address,omitempty"` //监听地址
}

func (x *RegisterAppRsp) Reset() {
	*x = RegisterAppRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_center_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAppRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAppRsp) ProtoMessage() {}

func (x *RegisterAppRsp) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAppRsp.ProtoReflect.Descriptor instead.
func (*RegisterAppRsp) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterAppRsp) GetRegResult() uint32 {
	if x != nil && x.RegResult != nil {
		return *x.RegResult
	}
	return 0
}

func (x *RegisterAppRsp) GetReregToken() string {
	if x != nil && x.ReregToken != nil {
		return *x.ReregToken
	}
	return ""
}

func (x *RegisterAppRsp) GetCenterId() uint32 {
	if x != nil && x.CenterId != nil {
		return *x.CenterId
	}
	return 0
}

func (x *RegisterAppRsp) GetAppType() uint32 {
	if x != nil && x.AppType != nil {
		return *x.AppType
	}
	return 0
}

func (x *RegisterAppRsp) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *RegisterAppRsp) GetAppName() string {
	if x != nil && x.AppName != nil {
		return *x.AppName
	}
	return ""
}

func (x *RegisterAppRsp) GetAppAddress() string {
	if x != nil && x.AppAddress != nil {
		return *x.AppAddress
	}
	return ""
}

//服务状态
type AppStateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppState *uint32 `protobuf:"varint,1,req,name=app_state,json=appState" json:"app_state,omitempty"`
	CenterId *uint32 `protobuf:"varint,2,opt,name=center_id,json=centerId" json:"center_id,omitempty"`
	AppType  *uint32 `protobuf:"varint,4,opt,name=app_type,json=appType" json:"app_type,omitempty"`
	AppId    *uint32 `protobuf:"varint,5,opt,name=app_id,json=appId" json:"app_id,omitempty"`
}

func (x *AppStateNotify) Reset() {
	*x = AppStateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_center_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppStateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppStateNotify) ProtoMessage() {}

func (x *AppStateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppStateNotify.ProtoReflect.Descriptor instead.
func (*AppStateNotify) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{3}
}

func (x *AppStateNotify) GetAppState() uint32 {
	if x != nil && x.AppState != nil {
		return *x.AppState
	}
	return 0
}

func (x *AppStateNotify) GetCenterId() uint32 {
	if x != nil && x.CenterId != nil {
		return *x.CenterId
	}
	return 0
}

func (x *AppStateNotify) GetAppType() uint32 {
	if x != nil && x.AppType != nil {
		return *x.AppType
	}
	return 0
}

func (x *AppStateNotify) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

//服务离线
type AppOfflineReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOffline *bool   `protobuf:"varint,1,opt,name=is_offline,json=isOffline" json:"is_offline,omitempty"` //是否下线
	AppId     *uint32 `protobuf:"varint,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	AppType   *uint32 `protobuf:"varint,3,opt,name=app_type,json=appType" json:"app_type,omitempty"`
}

func (x *AppOfflineReq) Reset() {
	*x = AppOfflineReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_center_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppOfflineReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppOfflineReq) ProtoMessage() {}

func (x *AppOfflineReq) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppOfflineReq.ProtoReflect.Descriptor instead.
func (*AppOfflineReq) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{4}
}

func (x *AppOfflineReq) GetIsOffline() bool {
	if x != nil && x.IsOffline != nil {
		return *x.IsOffline
	}
	return false
}

func (x *AppOfflineReq) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *AppOfflineReq) GetAppType() uint32 {
	if x != nil && x.AppType != nil {
		return *x.AppType
	}
	return 0
}

//服务心跳
type AppPulseNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action    *AppPulseNotify_PulseAction `protobuf:"varint,1,req,name=action,enum=bs.center.AppPulseNotify_PulseAction" json:"action,omitempty"`
	PulseData *uint64                     `protobuf:"varint,2,opt,name=pulse_data,json=pulseData" json:"pulse_data,omitempty"`
	LoadData  *uint32                     `protobuf:"varint,3,opt,name=load_data,json=loadData" json:"load_data,omitempty"`
}

func (x *AppPulseNotify) Reset() {
	*x = AppPulseNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_center_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppPulseNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppPulseNotify) ProtoMessage() {}

func (x *AppPulseNotify) ProtoReflect() protoreflect.Message {
	mi := &file_center_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppPulseNotify.ProtoReflect.Descriptor instead.
func (*AppPulseNotify) Descriptor() ([]byte, []int) {
	return file_center_proto_rawDescGZIP(), []int{5}
}

func (x *AppPulseNotify) GetAction() AppPulseNotify_PulseAction {
	if x != nil && x.Action != nil {
		return *x.Action
	}
	return AppPulseNotify_HeartBeatReq
}

func (x *AppPulseNotify) GetPulseData() uint64 {
	if x != nil && x.PulseData != nil {
		return *x.PulseData
	}
	return 0
}

func (x *AppPulseNotify) GetLoadData() uint32 {
	if x != nil && x.LoadData != nil {
		return *x.LoadData
	}
	return 0
}

var File_center_proto protoreflect.FileDescriptor

var file_center_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x62, 0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0c, 0x41, 0x70, 0x70,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67,
	0x41, 0x64, 0x64, 0x72, 0x22, 0xd3, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x4b,
	0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x74, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x70, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x72, 0x65, 0x67, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x72, 0x65, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xdb, 0x01, 0x0a, 0x0e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x52, 0x73, 0x70, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x09, 0x72, 0x65, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x72, 0x65, 0x67, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x72, 0x65, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70,
	0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x70,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x70, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08,
	0x61, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x66,
	0x66, 0x4c, 0x69, 0x6e, 0x65, 0x10, 0x03, 0x22, 0x60, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x4f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x6f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0xee, 0x01, 0x0a, 0x0e, 0x41, 0x70,
	0x70, 0x50, 0x75, 0x6c, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x3d, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x62,
	0x73, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x50, 0x75, 0x6c, 0x73,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x50, 0x75, 0x6c, 0x73, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x75, 0x6c, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x70, 0x75, 0x6c, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c,
	0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0x61, 0x0a, 0x0b, 0x50, 0x75, 0x6c, 0x73, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x73, 0x70, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x52, 0x73, 0x70, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x6e, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x05, 0x2a, 0x79, 0x0a, 0x09, 0x43, 0x4d,
	0x44, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x44, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x67, 0x52, 0x65, 0x71, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x44, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x67, 0x52, 0x73, 0x70, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x44, 0x41,
	0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x44, 0x41,
	0x70, 0x70, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x10, 0x04, 0x12, 0x11,
	0x0a, 0x0d, 0x49, 0x44, 0x50, 0x75, 0x6c, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x10,
	0x05, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x44, 0x41, 0x70, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x10, 0x06, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
}

var (
	file_center_proto_rawDescOnce sync.Once
	file_center_proto_rawDescData = file_center_proto_rawDesc
)

func file_center_proto_rawDescGZIP() []byte {
	file_center_proto_rawDescOnce.Do(func() {
		file_center_proto_rawDescData = protoimpl.X.CompressGZIP(file_center_proto_rawDescData)
	})
	return file_center_proto_rawDescData
}

var file_center_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_center_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_center_proto_goTypes = []interface{}{
	(CMDCenter)(0),                  // 0: bs.center.CMDCenter
	(AppStateNotifyState)(0),        // 1: bs.center.AppStateNotify.state
	(AppPulseNotify_PulseAction)(0), // 2: bs.center.AppPulseNotify.PulseAction
	(*AppUpdateReq)(nil),            // 3: bs.center.AppUpdateReq
	(*RegisterAppReq)(nil),          // 4: bs.center.RegisterAppReq
	(*RegisterAppRsp)(nil),          // 5: bs.center.RegisterAppRsp
	(*AppStateNotify)(nil),          // 6: bs.center.AppStateNotify
	(*AppOfflineReq)(nil),           // 7: bs.center.AppOfflineReq
	(*AppPulseNotify)(nil),          // 8: bs.center.AppPulseNotify
}
var file_center_proto_depIdxs = []int32{
	2, // 0: bs.center.AppPulseNotify.action:type_name -> bs.center.AppPulseNotify.PulseAction
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_center_proto_init() }
func file_center_proto_init() {
	if File_center_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_center_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppUpdateReq); i {
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
		file_center_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAppReq); i {
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
		file_center_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAppRsp); i {
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
		file_center_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppStateNotify); i {
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
		file_center_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppOfflineReq); i {
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
		file_center_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppPulseNotify); i {
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
			RawDescriptor: file_center_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_center_proto_goTypes,
		DependencyIndexes: file_center_proto_depIdxs,
		EnumInfos:         file_center_proto_enumTypes,
		MessageInfos:      file_center_proto_msgTypes,
	}.Build()
	File_center_proto = out.File
	file_center_proto_rawDesc = nil
	file_center_proto_goTypes = nil
	file_center_proto_depIdxs = nil
}
