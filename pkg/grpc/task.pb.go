// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Interface exported by the server.
type Task struct {
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetCreatedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Task) GetUpdatedAt() *timestamppb.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type GetAllTasksReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllTasksReq) Reset()         { *m = GetAllTasksReq{} }
func (m *GetAllTasksReq) String() string { return proto.CompactTextString(m) }
func (*GetAllTasksReq) ProtoMessage()    {}
func (*GetAllTasksReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{1}
}

func (m *GetAllTasksReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllTasksReq.Unmarshal(m, b)
}
func (m *GetAllTasksReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllTasksReq.Marshal(b, m, deterministic)
}
func (m *GetAllTasksReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllTasksReq.Merge(m, src)
}
func (m *GetAllTasksReq) XXX_Size() int {
	return xxx_messageInfo_GetAllTasksReq.Size(m)
}
func (m *GetAllTasksReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllTasksReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllTasksReq proto.InternalMessageInfo

type GetAllTasksRes struct {
	Task                 []*Task  `protobuf:"bytes,1,rep,name=task,proto3" json:"task,omitempty"`
	TotalTask            int64    `protobuf:"varint,2,opt,name=total_task,json=totalTask,proto3" json:"total_task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllTasksRes) Reset()         { *m = GetAllTasksRes{} }
func (m *GetAllTasksRes) String() string { return proto.CompactTextString(m) }
func (*GetAllTasksRes) ProtoMessage()    {}
func (*GetAllTasksRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{2}
}

func (m *GetAllTasksRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllTasksRes.Unmarshal(m, b)
}
func (m *GetAllTasksRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllTasksRes.Marshal(b, m, deterministic)
}
func (m *GetAllTasksRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllTasksRes.Merge(m, src)
}
func (m *GetAllTasksRes) XXX_Size() int {
	return xxx_messageInfo_GetAllTasksRes.Size(m)
}
func (m *GetAllTasksRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllTasksRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllTasksRes proto.InternalMessageInfo

func (m *GetAllTasksRes) GetTask() []*Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *GetAllTasksRes) GetTotalTask() int64 {
	if m != nil {
		return m.TotalTask
	}
	return 0
}

type GetSingleTaskReq struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSingleTaskReq) Reset()         { *m = GetSingleTaskReq{} }
func (m *GetSingleTaskReq) String() string { return proto.CompactTextString(m) }
func (*GetSingleTaskReq) ProtoMessage()    {}
func (*GetSingleTaskReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{3}
}

func (m *GetSingleTaskReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSingleTaskReq.Unmarshal(m, b)
}
func (m *GetSingleTaskReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSingleTaskReq.Marshal(b, m, deterministic)
}
func (m *GetSingleTaskReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSingleTaskReq.Merge(m, src)
}
func (m *GetSingleTaskReq) XXX_Size() int {
	return xxx_messageInfo_GetSingleTaskReq.Size(m)
}
func (m *GetSingleTaskReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSingleTaskReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetSingleTaskReq proto.InternalMessageInfo

func (m *GetSingleTaskReq) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type GetSingleTaskRes struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSingleTaskRes) Reset()         { *m = GetSingleTaskRes{} }
func (m *GetSingleTaskRes) String() string { return proto.CompactTextString(m) }
func (*GetSingleTaskRes) ProtoMessage()    {}
func (*GetSingleTaskRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{4}
}

func (m *GetSingleTaskRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSingleTaskRes.Unmarshal(m, b)
}
func (m *GetSingleTaskRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSingleTaskRes.Marshal(b, m, deterministic)
}
func (m *GetSingleTaskRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSingleTaskRes.Merge(m, src)
}
func (m *GetSingleTaskRes) XXX_Size() int {
	return xxx_messageInfo_GetSingleTaskRes.Size(m)
}
func (m *GetSingleTaskRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSingleTaskRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetSingleTaskRes proto.InternalMessageInfo

func (m *GetSingleTaskRes) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type CreateTaskReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskReq) Reset()         { *m = CreateTaskReq{} }
func (m *CreateTaskReq) String() string { return proto.CompactTextString(m) }
func (*CreateTaskReq) ProtoMessage()    {}
func (*CreateTaskReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{5}
}

func (m *CreateTaskReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskReq.Unmarshal(m, b)
}
func (m *CreateTaskReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskReq.Marshal(b, m, deterministic)
}
func (m *CreateTaskReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskReq.Merge(m, src)
}
func (m *CreateTaskReq) XXX_Size() int {
	return xxx_messageInfo_CreateTaskReq.Size(m)
}
func (m *CreateTaskReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskReq proto.InternalMessageInfo

func (m *CreateTaskReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateTaskReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateTaskRes struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskRes) Reset()         { *m = CreateTaskRes{} }
func (m *CreateTaskRes) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRes) ProtoMessage()    {}
func (*CreateTaskRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{6}
}

func (m *CreateTaskRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRes.Unmarshal(m, b)
}
func (m *CreateTaskRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRes.Marshal(b, m, deterministic)
}
func (m *CreateTaskRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRes.Merge(m, src)
}
func (m *CreateTaskRes) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRes.Size(m)
}
func (m *CreateTaskRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRes proto.InternalMessageInfo

func (m *CreateTaskRes) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type DeleteTaskReq struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskReq) Reset()         { *m = DeleteTaskReq{} }
func (m *DeleteTaskReq) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskReq) ProtoMessage()    {}
func (*DeleteTaskReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{7}
}

func (m *DeleteTaskReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskReq.Unmarshal(m, b)
}
func (m *DeleteTaskReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskReq.Marshal(b, m, deterministic)
}
func (m *DeleteTaskReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskReq.Merge(m, src)
}
func (m *DeleteTaskReq) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskReq.Size(m)
}
func (m *DeleteTaskReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskReq proto.InternalMessageInfo

func (m *DeleteTaskReq) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *DeleteTaskReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type DeleteTaskRes struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskRes) Reset()         { *m = DeleteTaskRes{} }
func (m *DeleteTaskRes) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskRes) ProtoMessage()    {}
func (*DeleteTaskRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{8}
}

func (m *DeleteTaskRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskRes.Unmarshal(m, b)
}
func (m *DeleteTaskRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskRes.Marshal(b, m, deterministic)
}
func (m *DeleteTaskRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskRes.Merge(m, src)
}
func (m *DeleteTaskRes) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskRes.Size(m)
}
func (m *DeleteTaskRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskRes proto.InternalMessageInfo

func (m *DeleteTaskRes) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "todo.Task")
	proto.RegisterType((*GetAllTasksReq)(nil), "todo.GetAllTasksReq")
	proto.RegisterType((*GetAllTasksRes)(nil), "todo.GetAllTasksRes")
	proto.RegisterType((*GetSingleTaskReq)(nil), "todo.GetSingleTaskReq")
	proto.RegisterType((*GetSingleTaskRes)(nil), "todo.GetSingleTaskRes")
	proto.RegisterType((*CreateTaskReq)(nil), "todo.CreateTaskReq")
	proto.RegisterType((*CreateTaskRes)(nil), "todo.CreateTaskRes")
	proto.RegisterType((*DeleteTaskReq)(nil), "todo.DeleteTaskReq")
	proto.RegisterType((*DeleteTaskRes)(nil), "todo.DeleteTaskRes")
}

func init() {
	proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff)
}

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x86, 0x99, 0x34, 0x56, 0x7a, 0x4a, 0xd6, 0x3a, 0x8a, 0x86, 0x80, 0x5a, 0x72, 0x55, 0x10,
	0x52, 0x88, 0x20, 0x88, 0x82, 0x44, 0x85, 0x65, 0xaf, 0x84, 0xec, 0x82, 0xe0, 0x4d, 0x98, 0xee,
	0x1c, 0x43, 0xd8, 0xb4, 0x93, 0x66, 0x4e, 0x7d, 0x00, 0x5f, 0xc0, 0xe7, 0xf1, 0xed, 0x64, 0x26,
	0x4d, 0x9b, 0x68, 0x96, 0xdc, 0xcd, 0xfc, 0xf9, 0xff, 0xf9, 0x33, 0xdf, 0x49, 0x00, 0x48, 0xe8,
	0xbb, 0xa8, 0xaa, 0x15, 0x29, 0xee, 0x92, 0x92, 0x2a, 0x78, 0x95, 0x2b, 0x95, 0x97, 0xb8, 0xb6,
	0xda, 0xe6, 0xf0, 0x63, 0x4d, 0xc5, 0x16, 0x35, 0x89, 0x6d, 0xd5, 0xd8, 0xc2, 0x3f, 0x0c, 0xdc,
	0x1b, 0xa1, 0xef, 0xf8, 0x05, 0x38, 0x85, 0xf4, 0xd9, 0x92, 0xad, 0x66, 0xa9, 0x53, 0x48, 0xfe,
	0x1c, 0x1e, 0x1e, 0x34, 0xd6, 0x59, 0x21, 0x7d, 0xc7, 0x8a, 0x53, 0xb3, 0xbd, 0x92, 0x9c, 0x83,
	0xbb, 0x13, 0x5b, 0xf4, 0x27, 0x56, 0xb5, 0x6b, 0xfe, 0x0e, 0xe0, 0xb6, 0x46, 0x41, 0x28, 0x33,
	0x41, 0xbe, 0xbb, 0x64, 0xab, 0x79, 0x1c, 0x44, 0x4d, 0x77, 0xd4, 0x76, 0x47, 0x37, 0x6d, 0x77,
	0x3a, 0x3b, 0xba, 0x13, 0x32, 0xd1, 0x43, 0x25, 0xdb, 0xe8, 0x83, 0xf1, 0xe8, 0xd1, 0x9d, 0x50,
	0xb8, 0x80, 0x8b, 0x4b, 0xa4, 0xa4, 0x2c, 0xcd, 0x05, 0x74, 0x8a, 0xfb, 0xf0, 0xeb, 0x3f, 0x8a,
	0xe6, 0x2f, 0xc1, 0x35, 0x50, 0x7c, 0xb6, 0x9c, 0xac, 0xe6, 0x31, 0x44, 0x86, 0x4a, 0x64, 0x9e,
	0xa6, 0x56, 0xe7, 0x2f, 0x00, 0x48, 0x91, 0x28, 0x33, 0xeb, 0x32, 0x37, 0x9d, 0xa4, 0x33, 0xab,
	0x18, 0x53, 0xf8, 0x1a, 0x16, 0x97, 0x48, 0xd7, 0xc5, 0x2e, 0x2f, 0xd1, 0xa6, 0x70, 0x6f, 0xc8,
	0x18, 0x73, 0x76, 0xc2, 0x35, 0x35, 0xdb, 0x2b, 0x19, 0xc6, 0xff, 0x99, 0xbb, 0xfd, 0x6c, 0xa8,
	0x3f, 0xfc, 0x00, 0xde, 0x67, 0xcb, 0xa2, 0x73, 0x7a, 0xcb, 0x9d, 0x0d, 0x72, 0x77, 0xce, 0xdc,
	0xc3, 0x75, 0x3f, 0x3d, 0x5e, 0x97, 0x80, 0xf7, 0x05, 0x4b, 0xa4, 0xd1, 0xcb, 0xdc, 0x3b, 0x7f,
	0xd3, 0xd9, 0x3d, 0x62, 0xb4, 0x33, 0xfe, 0xcd, 0xe0, 0x71, 0xe3, 0x15, 0x12, 0xeb, 0x6b, 0xac,
	0x7f, 0x16, 0xb7, 0xe6, 0x93, 0x99, 0x77, 0x46, 0xc5, 0x9f, 0x36, 0xb1, 0xfe, 0x3c, 0x83, 0x21,
	0x55, 0xf3, 0x8f, 0xe0, 0xf5, 0x38, 0xf3, 0x67, 0x27, 0x5b, 0x6f, 0x52, 0xc1, 0xb0, 0xae, 0xe3,
	0x5f, 0x0c, 0x16, 0x66, 0xfd, 0xad, 0x2e, 0x08, 0xdb, 0x17, 0x7a, 0x0b, 0x70, 0x66, 0xc9, 0x9f,
	0x34, 0xd1, 0xde, 0x6c, 0x82, 0x01, 0x51, 0x9b, 0xdc, 0x99, 0x47, 0x9b, 0xeb, 0x41, 0x0e, 0x06,
	0x44, 0xfd, 0xe9, 0xd1, 0x77, 0x8f, 0x54, 0xae, 0x9a, 0x5f, 0xf3, 0x7d, 0xb5, 0xd9, 0x4c, 0xed,
	0xea, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x69, 0xdf, 0xc6, 0xc6, 0x03, 0x00, 0x00,
}
