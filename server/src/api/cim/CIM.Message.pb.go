// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CIM.Message.proto

package cim

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 发送消息
type CIMMsgData struct {
	// cmd id:		0x0301
	FromUserId           uint64         `protobuf:"varint,1,opt,name=from_user_id,json=fromUserId,proto3" json:"from_user_id,omitempty"`
	FromNickName         string         `protobuf:"bytes,2,opt,name=from_nick_name,json=fromNickName,proto3" json:"from_nick_name,omitempty"`
	ToSessionId          uint64         `protobuf:"varint,3,opt,name=to_session_id,json=toSessionId,proto3" json:"to_session_id,omitempty"`
	MsgId                string         `protobuf:"bytes,4,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	CreateTime           int32          `protobuf:"varint,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	MsgType              CIMMsgType     `protobuf:"varint,6,opt,name=msg_type,json=msgType,proto3,enum=CIM.Def.CIMMsgType" json:"msg_type,omitempty"`
	SessionType          CIMSessionType `protobuf:"varint,7,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	MsgData              []byte         `protobuf:"bytes,8,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CIMMsgData) Reset()         { *m = CIMMsgData{} }
func (m *CIMMsgData) String() string { return proto.CompactTextString(m) }
func (*CIMMsgData) ProtoMessage()    {}
func (*CIMMsgData) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e1c4ce9d6da003, []int{0}
}

func (m *CIMMsgData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMMsgData.Unmarshal(m, b)
}
func (m *CIMMsgData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMMsgData.Marshal(b, m, deterministic)
}
func (m *CIMMsgData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMMsgData.Merge(m, src)
}
func (m *CIMMsgData) XXX_Size() int {
	return xxx_messageInfo_CIMMsgData.Size(m)
}
func (m *CIMMsgData) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMMsgData.DiscardUnknown(m)
}

var xxx_messageInfo_CIMMsgData proto.InternalMessageInfo

func (m *CIMMsgData) GetFromUserId() uint64 {
	if m != nil {
		return m.FromUserId
	}
	return 0
}

func (m *CIMMsgData) GetFromNickName() string {
	if m != nil {
		return m.FromNickName
	}
	return ""
}

func (m *CIMMsgData) GetToSessionId() uint64 {
	if m != nil {
		return m.ToSessionId
	}
	return 0
}

func (m *CIMMsgData) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *CIMMsgData) GetCreateTime() int32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *CIMMsgData) GetMsgType() CIMMsgType {
	if m != nil {
		return m.MsgType
	}
	return CIMMsgType_kCIM_MSG_TYPE_UNKNOWN
}

func (m *CIMMsgData) GetSessionType() CIMSessionType {
	if m != nil {
		return m.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (m *CIMMsgData) GetMsgData() []byte {
	if m != nil {
		return m.MsgData
	}
	return nil
}

// 消息收到回复
type CIMMsgDataAck struct {
	// cmd id:		0x0302
	FromUserId  uint64         `protobuf:"varint,1,opt,name=from_user_id,json=fromUserId,proto3" json:"from_user_id,omitempty"`
	ToSessionId uint64         `protobuf:"varint,2,opt,name=to_session_id,json=toSessionId,proto3" json:"to_session_id,omitempty"`
	MsgId       string         `protobuf:"bytes,4,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	ServerMsgId uint64         `protobuf:"varint,3,opt,name=server_msg_id,json=serverMsgId,proto3" json:"server_msg_id,omitempty"`
	ResCode     CIMResCode     `protobuf:"varint,5,opt,name=res_code,json=resCode,proto3,enum=CIM.Def.CIMResCode" json:"res_code,omitempty"`
	SessionType CIMSessionType `protobuf:"varint,6,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	//optional
	CreateTime           int32    `protobuf:"varint,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIMMsgDataAck) Reset()         { *m = CIMMsgDataAck{} }
func (m *CIMMsgDataAck) String() string { return proto.CompactTextString(m) }
func (*CIMMsgDataAck) ProtoMessage()    {}
func (*CIMMsgDataAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e1c4ce9d6da003, []int{1}
}

func (m *CIMMsgDataAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMMsgDataAck.Unmarshal(m, b)
}
func (m *CIMMsgDataAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMMsgDataAck.Marshal(b, m, deterministic)
}
func (m *CIMMsgDataAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMMsgDataAck.Merge(m, src)
}
func (m *CIMMsgDataAck) XXX_Size() int {
	return xxx_messageInfo_CIMMsgDataAck.Size(m)
}
func (m *CIMMsgDataAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMMsgDataAck.DiscardUnknown(m)
}

var xxx_messageInfo_CIMMsgDataAck proto.InternalMessageInfo

func (m *CIMMsgDataAck) GetFromUserId() uint64 {
	if m != nil {
		return m.FromUserId
	}
	return 0
}

func (m *CIMMsgDataAck) GetToSessionId() uint64 {
	if m != nil {
		return m.ToSessionId
	}
	return 0
}

func (m *CIMMsgDataAck) GetMsgId() string {
	if m != nil {
		return m.MsgId
	}
	return ""
}

func (m *CIMMsgDataAck) GetServerMsgId() uint64 {
	if m != nil {
		return m.ServerMsgId
	}
	return 0
}

func (m *CIMMsgDataAck) GetResCode() CIMResCode {
	if m != nil {
		return m.ResCode
	}
	return CIMResCode_kCIM_RES_CODE_OK
}

func (m *CIMMsgDataAck) GetSessionType() CIMSessionType {
	if m != nil {
		return m.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (m *CIMMsgDataAck) GetCreateTime() int32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

// 消息已读回复（我方）
type CIMMsgDataReadAck struct {
	// cmd id:		0x0303
	UserId               uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionId            uint64         `protobuf:"varint,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	MsgId                uint64         `protobuf:"varint,3,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	SessionType          CIMSessionType `protobuf:"varint,4,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CIMMsgDataReadAck) Reset()         { *m = CIMMsgDataReadAck{} }
func (m *CIMMsgDataReadAck) String() string { return proto.CompactTextString(m) }
func (*CIMMsgDataReadAck) ProtoMessage()    {}
func (*CIMMsgDataReadAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e1c4ce9d6da003, []int{2}
}

func (m *CIMMsgDataReadAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMMsgDataReadAck.Unmarshal(m, b)
}
func (m *CIMMsgDataReadAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMMsgDataReadAck.Marshal(b, m, deterministic)
}
func (m *CIMMsgDataReadAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMMsgDataReadAck.Merge(m, src)
}
func (m *CIMMsgDataReadAck) XXX_Size() int {
	return xxx_messageInfo_CIMMsgDataReadAck.Size(m)
}
func (m *CIMMsgDataReadAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMMsgDataReadAck.DiscardUnknown(m)
}

var xxx_messageInfo_CIMMsgDataReadAck proto.InternalMessageInfo

func (m *CIMMsgDataReadAck) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMMsgDataReadAck) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *CIMMsgDataReadAck) GetMsgId() uint64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *CIMMsgDataReadAck) GetSessionType() CIMSessionType {
	if m != nil {
		return m.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

// 已读消息通知（对方）
type CIMMsgDataReadNotify struct {
	// cmd id:		0x0304
	UserId               uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionId            uint64         `protobuf:"varint,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	MsgId                uint64         `protobuf:"varint,3,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	SessionType          CIMSessionType `protobuf:"varint,4,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CIMMsgDataReadNotify) Reset()         { *m = CIMMsgDataReadNotify{} }
func (m *CIMMsgDataReadNotify) String() string { return proto.CompactTextString(m) }
func (*CIMMsgDataReadNotify) ProtoMessage()    {}
func (*CIMMsgDataReadNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e1c4ce9d6da003, []int{3}
}

func (m *CIMMsgDataReadNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMMsgDataReadNotify.Unmarshal(m, b)
}
func (m *CIMMsgDataReadNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMMsgDataReadNotify.Marshal(b, m, deterministic)
}
func (m *CIMMsgDataReadNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMMsgDataReadNotify.Merge(m, src)
}
func (m *CIMMsgDataReadNotify) XXX_Size() int {
	return xxx_messageInfo_CIMMsgDataReadNotify.Size(m)
}
func (m *CIMMsgDataReadNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMMsgDataReadNotify.DiscardUnknown(m)
}

var xxx_messageInfo_CIMMsgDataReadNotify proto.InternalMessageInfo

func (m *CIMMsgDataReadNotify) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMMsgDataReadNotify) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *CIMMsgDataReadNotify) GetMsgId() uint64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *CIMMsgDataReadNotify) GetSessionType() CIMSessionType {
	if m != nil {
		return m.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

// 某个会话最新的消息ID请求
type CIMGetLatestMsgIdReq struct {
	// cmd id: 		0x030b
	UserId               uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionType          CIMSessionType `protobuf:"varint,2,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	SessionId            uint64         `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CIMGetLatestMsgIdReq) Reset()         { *m = CIMGetLatestMsgIdReq{} }
func (m *CIMGetLatestMsgIdReq) String() string { return proto.CompactTextString(m) }
func (*CIMGetLatestMsgIdReq) ProtoMessage()    {}
func (*CIMGetLatestMsgIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e1c4ce9d6da003, []int{4}
}

func (m *CIMGetLatestMsgIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMGetLatestMsgIdReq.Unmarshal(m, b)
}
func (m *CIMGetLatestMsgIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMGetLatestMsgIdReq.Marshal(b, m, deterministic)
}
func (m *CIMGetLatestMsgIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMGetLatestMsgIdReq.Merge(m, src)
}
func (m *CIMGetLatestMsgIdReq) XXX_Size() int {
	return xxx_messageInfo_CIMGetLatestMsgIdReq.Size(m)
}
func (m *CIMGetLatestMsgIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMGetLatestMsgIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_CIMGetLatestMsgIdReq proto.InternalMessageInfo

func (m *CIMGetLatestMsgIdReq) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMGetLatestMsgIdReq) GetSessionType() CIMSessionType {
	if m != nil {
		return m.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (m *CIMGetLatestMsgIdReq) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

type CIMGetLatestMsgIdRsp struct {
	// cmd id:		0x030c
	UserId               uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionType          CIMSessionType `protobuf:"varint,2,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	SessionId            uint64         `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	LatestMsgId          string         `protobuf:"bytes,4,opt,name=latest_msg_id,json=latestMsgId,proto3" json:"latest_msg_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CIMGetLatestMsgIdRsp) Reset()         { *m = CIMGetLatestMsgIdRsp{} }
func (m *CIMGetLatestMsgIdRsp) String() string { return proto.CompactTextString(m) }
func (*CIMGetLatestMsgIdRsp) ProtoMessage()    {}
func (*CIMGetLatestMsgIdRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e1c4ce9d6da003, []int{5}
}

func (m *CIMGetLatestMsgIdRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMGetLatestMsgIdRsp.Unmarshal(m, b)
}
func (m *CIMGetLatestMsgIdRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMGetLatestMsgIdRsp.Marshal(b, m, deterministic)
}
func (m *CIMGetLatestMsgIdRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMGetLatestMsgIdRsp.Merge(m, src)
}
func (m *CIMGetLatestMsgIdRsp) XXX_Size() int {
	return xxx_messageInfo_CIMGetLatestMsgIdRsp.Size(m)
}
func (m *CIMGetLatestMsgIdRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMGetLatestMsgIdRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CIMGetLatestMsgIdRsp proto.InternalMessageInfo

func (m *CIMGetLatestMsgIdRsp) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMGetLatestMsgIdRsp) GetSessionType() CIMSessionType {
	if m != nil {
		return m.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (m *CIMGetLatestMsgIdRsp) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *CIMGetLatestMsgIdRsp) GetLatestMsgId() string {
	if m != nil {
		return m.LatestMsgId
	}
	return ""
}

// 批量查询消息详情请求(20条内)
type CIMGetMsgByIdReq struct {
	// cmd id: 		0x030d
	UserId               uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionType          CIMSessionType `protobuf:"varint,2,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	SessionId            uint64         `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	MsgIdList            []string       `protobuf:"bytes,4,rep,name=msg_id_list,json=msgIdList,proto3" json:"msg_id_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CIMGetMsgByIdReq) Reset()         { *m = CIMGetMsgByIdReq{} }
func (m *CIMGetMsgByIdReq) String() string { return proto.CompactTextString(m) }
func (*CIMGetMsgByIdReq) ProtoMessage()    {}
func (*CIMGetMsgByIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e1c4ce9d6da003, []int{6}
}

func (m *CIMGetMsgByIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMGetMsgByIdReq.Unmarshal(m, b)
}
func (m *CIMGetMsgByIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMGetMsgByIdReq.Marshal(b, m, deterministic)
}
func (m *CIMGetMsgByIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMGetMsgByIdReq.Merge(m, src)
}
func (m *CIMGetMsgByIdReq) XXX_Size() int {
	return xxx_messageInfo_CIMGetMsgByIdReq.Size(m)
}
func (m *CIMGetMsgByIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMGetMsgByIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_CIMGetMsgByIdReq proto.InternalMessageInfo

func (m *CIMGetMsgByIdReq) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMGetMsgByIdReq) GetSessionType() CIMSessionType {
	if m != nil {
		return m.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (m *CIMGetMsgByIdReq) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *CIMGetMsgByIdReq) GetMsgIdList() []string {
	if m != nil {
		return m.MsgIdList
	}
	return nil
}

type CIMGetMsgByIdRsp struct {
	// cmd id:		0x030e
	UserId               uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionType          CIMSessionType `protobuf:"varint,2,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	SessionId            uint64         `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	MsgList              []*CIMMsgInfo  `protobuf:"bytes,4,rep,name=msg_list,json=msgList,proto3" json:"msg_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CIMGetMsgByIdRsp) Reset()         { *m = CIMGetMsgByIdRsp{} }
func (m *CIMGetMsgByIdRsp) String() string { return proto.CompactTextString(m) }
func (*CIMGetMsgByIdRsp) ProtoMessage()    {}
func (*CIMGetMsgByIdRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e1c4ce9d6da003, []int{7}
}

func (m *CIMGetMsgByIdRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMGetMsgByIdRsp.Unmarshal(m, b)
}
func (m *CIMGetMsgByIdRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMGetMsgByIdRsp.Marshal(b, m, deterministic)
}
func (m *CIMGetMsgByIdRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMGetMsgByIdRsp.Merge(m, src)
}
func (m *CIMGetMsgByIdRsp) XXX_Size() int {
	return xxx_messageInfo_CIMGetMsgByIdRsp.Size(m)
}
func (m *CIMGetMsgByIdRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMGetMsgByIdRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CIMGetMsgByIdRsp proto.InternalMessageInfo

func (m *CIMGetMsgByIdRsp) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMGetMsgByIdRsp) GetSessionType() CIMSessionType {
	if m != nil {
		return m.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (m *CIMGetMsgByIdRsp) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *CIMGetMsgByIdRsp) GetMsgList() []*CIMMsgInfo {
	if m != nil {
		return m.MsgList
	}
	return nil
}

func init() {
	proto.RegisterType((*CIMMsgData)(nil), "CIM.Message.CIMMsgData")
	proto.RegisterType((*CIMMsgDataAck)(nil), "CIM.Message.CIMMsgDataAck")
	proto.RegisterType((*CIMMsgDataReadAck)(nil), "CIM.Message.CIMMsgDataReadAck")
	proto.RegisterType((*CIMMsgDataReadNotify)(nil), "CIM.Message.CIMMsgDataReadNotify")
	proto.RegisterType((*CIMGetLatestMsgIdReq)(nil), "CIM.Message.CIMGetLatestMsgIdReq")
	proto.RegisterType((*CIMGetLatestMsgIdRsp)(nil), "CIM.Message.CIMGetLatestMsgIdRsp")
	proto.RegisterType((*CIMGetMsgByIdReq)(nil), "CIM.Message.CIMGetMsgByIdReq")
	proto.RegisterType((*CIMGetMsgByIdRsp)(nil), "CIM.Message.CIMGetMsgByIdRsp")
}

func init() { proto.RegisterFile("CIM.Message.proto", fileDescriptor_51e1c4ce9d6da003) }

var fileDescriptor_51e1c4ce9d6da003 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0x4f, 0x8f, 0x93, 0x40,
	0x14, 0xc0, 0x43, 0xff, 0xf7, 0xd1, 0x6e, 0x14, 0x35, 0x45, 0x13, 0x95, 0x10, 0x0f, 0x9c, 0x38,
	0xe8, 0xcd, 0x9b, 0xdb, 0x4d, 0x94, 0x64, 0xe9, 0x01, 0xd7, 0x8b, 0x17, 0x32, 0x0b, 0x8f, 0x4a,
	0xda, 0xe9, 0x54, 0xde, 0xd4, 0xa4, 0x5f, 0xc1, 0xef, 0xa0, 0x89, 0x17, 0x8f, 0x1e, 0xf5, 0xeb,
	0x19, 0x66, 0x58, 0xbb, 0x85, 0x35, 0xe9, 0x9e, 0xba, 0x37, 0x78, 0xc3, 0x7b, 0xfc, 0xde, 0xef,
	0x31, 0x03, 0xdc, 0x9f, 0x06, 0xa1, 0x1f, 0x22, 0x11, 0x9b, 0xa3, 0xbf, 0x2e, 0x84, 0x14, 0x96,
	0x79, 0x2d, 0xf4, 0x64, 0x5c, 0xde, 0x9c, 0x61, 0xa6, 0xd7, 0xdc, 0x3f, 0x2d, 0x80, 0x69, 0x10,
	0x86, 0x34, 0x3f, 0x63, 0x92, 0x59, 0x0e, 0x8c, 0xb2, 0x42, 0xf0, 0x78, 0x43, 0x58, 0xc4, 0x79,
	0x6a, 0x1b, 0x8e, 0xe1, 0x75, 0x22, 0x28, 0x63, 0x1f, 0x08, 0x8b, 0x20, 0xb5, 0x5e, 0xc0, 0x89,
	0x7a, 0x62, 0x95, 0x27, 0x8b, 0x78, 0xc5, 0x38, 0xda, 0x2d, 0xc7, 0xf0, 0x86, 0x91, 0xca, 0x9b,
	0xe5, 0xc9, 0x62, 0xc6, 0x38, 0x5a, 0x2e, 0x8c, 0xa5, 0x88, 0x09, 0x89, 0x72, 0xb1, 0x2a, 0x0b,
	0xb5, 0x55, 0x21, 0x53, 0x8a, 0xf7, 0x3a, 0x16, 0xa4, 0xd6, 0x23, 0xe8, 0x71, 0x9a, 0x97, 0x8b,
	0x1d, 0x55, 0xa1, 0xcb, 0x69, 0x1e, 0xa4, 0xd6, 0x73, 0x30, 0x93, 0x02, 0x99, 0xc4, 0x58, 0xe6,
	0x1c, 0xed, 0xae, 0x63, 0x78, 0xdd, 0x08, 0x74, 0xe8, 0x22, 0xe7, 0x68, 0xf9, 0x30, 0x28, 0xf3,
	0xe4, 0x76, 0x8d, 0x76, 0xcf, 0x31, 0xbc, 0x93, 0x97, 0x0f, 0xfc, 0xab, 0xa6, 0x74, 0x2b, 0x17,
	0xdb, 0x35, 0x46, 0x7d, 0xae, 0x2f, 0xac, 0xd7, 0x30, 0xba, 0x02, 0x51, 0x39, 0x7d, 0x95, 0x33,
	0xb9, 0x9e, 0x53, 0x41, 0xa9, 0x3c, 0x93, 0x76, 0x37, 0xd6, 0x63, 0xfd, 0xae, 0x94, 0x49, 0x66,
	0x0f, 0x1c, 0xc3, 0x1b, 0xa9, 0xb2, 0xa5, 0x2a, 0xf7, 0x5b, 0x0b, 0xc6, 0x3b, 0x73, 0x6f, 0x92,
	0xc5, 0x01, 0xf2, 0x1a, 0x5a, 0x5a, 0x07, 0x6b, 0x71, 0x61, 0x4c, 0x58, 0x7c, 0xc1, 0x22, 0xae,
	0x56, 0x2b, 0xa3, 0x3a, 0x18, 0xaa, 0x67, 0x7c, 0x18, 0x14, 0x48, 0x71, 0x22, 0x52, 0xed, 0xad,
	0x66, 0x26, 0x42, 0x9a, 0x8a, 0x14, 0xa3, 0x7e, 0xa1, 0x2f, 0x1a, 0x66, 0x7a, 0xb7, 0x30, 0x53,
	0x1b, 0x53, 0xbf, 0x3e, 0x26, 0xf7, 0xbb, 0xa1, 0xbe, 0xc5, 0xca, 0x4f, 0x84, 0x2c, 0x2d, 0x1d,
	0x4d, 0xa0, 0xbf, 0xaf, 0xa7, 0xb7, 0xd1, 0x6a, 0x9e, 0x02, 0x34, 0xbc, 0x0c, 0xe9, 0x06, 0x2b,
	0xba, 0xef, 0xca, 0x4a, 0xbd, 0x83, 0xce, 0xe1, 0x1d, 0xb8, 0x3f, 0x0c, 0x78, 0xb8, 0x0f, 0x38,
	0x13, 0x32, 0xcf, 0xb6, 0x77, 0x89, 0xf1, 0xab, 0x66, 0x7c, 0x8b, 0xf2, 0x9c, 0x49, 0x24, 0xa9,
	0xe6, 0x1c, 0xe1, 0xe7, 0xff, 0x33, 0xd6, 0xdf, 0xd6, 0xba, 0xc5, 0x4c, 0xf7, 0xfb, 0x6b, 0xd7,
	0xfa, 0x73, 0x7f, 0xdd, 0x08, 0x43, 0xeb, 0x63, 0xc0, 0x94, 0xfb, 0x61, 0xa9, 0x28, 0xe2, 0xbd,
	0xdd, 0x62, 0x2e, 0x77, 0x68, 0xee, 0x4f, 0x03, 0xee, 0x69, 0xe0, 0x90, 0xe6, 0xa7, 0xdb, 0xa3,
	0x99, 0xb3, 0x9e, 0x81, 0xa9, 0x29, 0xe3, 0x65, 0x4e, 0xd2, 0xee, 0x38, 0x6d, 0x6f, 0x18, 0x0d,
	0xd5, 0xe7, 0x71, 0x9e, 0x93, 0x74, 0x7f, 0x37, 0x40, 0x8f, 0x64, 0xb5, 0x3a, 0x5b, 0xff, 0x51,
	0x9a, 0x8d, 0xb3, 0x35, 0x58, 0x65, 0x42, 0x1d, 0x82, 0x25, 0xf8, 0xa9, 0x03, 0x93, 0x44, 0x70,
	0x3f, 0x11, 0x59, 0x86, 0x98, 0x7c, 0x62, 0x52, 0xff, 0x56, 0x2e, 0x37, 0xd9, 0xbb, 0xf6, 0xc7,
	0x76, 0x92, 0xf3, 0xcb, 0x9e, 0x0a, 0xbc, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x98, 0xc5, 0x3e,
	0x76, 0x98, 0x06, 0x00, 0x00,
}
