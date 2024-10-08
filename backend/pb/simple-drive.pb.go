// proto/simple-drive.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: simple-drive.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ItemType int32

const (
	ItemType_FILE      ItemType = 0
	ItemType_DIRECTORY ItemType = 1
)

// Enum value maps for ItemType.
var (
	ItemType_name = map[int32]string{
		0: "FILE",
		1: "DIRECTORY",
	}
	ItemType_value = map[string]int32{
		"FILE":      0,
		"DIRECTORY": 1,
	}
)

func (x ItemType) Enum() *ItemType {
	p := new(ItemType)
	*p = x
	return p
}

func (x ItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_simple_drive_proto_enumTypes[0].Descriptor()
}

func (ItemType) Type() protoreflect.EnumType {
	return &file_simple_drive_proto_enumTypes[0]
}

func (x ItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemType.Descriptor instead.
func (ItemType) EnumDescriptor() ([]byte, []int) {
	return file_simple_drive_proto_rawDescGZIP(), []int{0}
}

// メッセージ定義
type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId       string                 `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name           string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type           ItemType               `protobuf:"varint,4,opt,name=type,proto3,enum=simpledrive.ItemType" json:"type,omitempty"`
	Size           *wrapperspb.Int64Value `protobuf:"bytes,5,opt,name=size,proto3" json:"size,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastModifiedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_modified_at,json=lastModifiedAt,proto3" json:"last_modified_at,omitempty"`
	ThumbnailId    string                 `protobuf:"bytes,9,opt,name=thumbnail_id,json=thumbnailId,proto3" json:"thumbnail_id,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_drive_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_simple_drive_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_simple_drive_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetType() ItemType {
	if x != nil {
		return x.Type
	}
	return ItemType_FILE
}

func (x *Item) GetSize() *wrapperspb.Int64Value {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *Item) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Item) GetLastModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastModifiedAt
	}
	return nil
}

func (x *Item) GetThumbnailId() string {
	if x != nil {
		return x.ThumbnailId
	}
	return ""
}

type ItemCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type     ItemType `protobuf:"varint,2,opt,name=type,proto3,enum=simpledrive.ItemType" json:"type,omitempty"`
	ParentId string   `protobuf:"bytes,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	File     []byte   `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"` // バイナリデータはbytes型で表現
}

func (x *ItemCreateRequest) Reset() {
	*x = ItemCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_drive_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemCreateRequest) ProtoMessage() {}

func (x *ItemCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simple_drive_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemCreateRequest.ProtoReflect.Descriptor instead.
func (*ItemCreateRequest) Descriptor() ([]byte, []int) {
	return file_simple_drive_proto_rawDescGZIP(), []int{1}
}

func (x *ItemCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ItemCreateRequest) GetType() ItemType {
	if x != nil {
		return x.Type
	}
	return ItemType_FILE
}

func (x *ItemCreateRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ItemCreateRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type ItemUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParentId string                  `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
}

func (x *ItemUpdateRequest) Reset() {
	*x = ItemUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_drive_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemUpdateRequest) ProtoMessage() {}

func (x *ItemUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simple_drive_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemUpdateRequest.ProtoReflect.Descriptor instead.
func (*ItemUpdateRequest) Descriptor() ([]byte, []int) {
	return file_simple_drive_proto_rawDescGZIP(), []int{2}
}

func (x *ItemUpdateRequest) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ItemUpdateRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

type ListItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId string `protobuf:"bytes,1,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size     int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListItemsRequest) Reset() {
	*x = ListItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_drive_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsRequest) ProtoMessage() {}

func (x *ListItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simple_drive_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsRequest.ProtoReflect.Descriptor instead.
func (*ListItemsRequest) Descriptor() ([]byte, []int) {
	return file_simple_drive_proto_rawDescGZIP(), []int{3}
}

func (x *ListItemsRequest) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *ListItemsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListItemsRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items  []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Paging *Paging `protobuf:"bytes,2,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *ListItemsResponse) Reset() {
	*x = ListItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_drive_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsResponse) ProtoMessage() {}

func (x *ListItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simple_drive_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsResponse.ProtoReflect.Descriptor instead.
func (*ListItemsResponse) Descriptor() ([]byte, []int) {
	return file_simple_drive_proto_rawDescGZIP(), []int{4}
}

func (x *ListItemsResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListItemsResponse) GetPaging() *Paging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_drive_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simple_drive_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_simple_drive_proto_rawDescGZIP(), []int{5}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_drive_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simple_drive_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_simple_drive_proto_rawDescGZIP(), []int{6}
}

func (x *HealthCheckResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Paging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPage int32 `protobuf:"varint,1,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	PageSize    int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	TotalItems  int64 `protobuf:"varint,3,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	TotalPages  int32 `protobuf:"varint,4,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
}

func (x *Paging) Reset() {
	*x = Paging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simple_drive_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paging) ProtoMessage() {}

func (x *Paging) ProtoReflect() protoreflect.Message {
	mi := &file_simple_drive_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paging.ProtoReflect.Descriptor instead.
func (*Paging) Descriptor() ([]byte, []int) {
	return file_simple_drive_proto_rawDescGZIP(), []int{7}
}

func (x *Paging) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *Paging) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Paging) GetTotalItems() int64 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *Paging) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

var File_simple_drive_proto protoreflect.FileDescriptor

var file_simple_drive_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xdf, 0x02, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x44, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x09, 0x52,
	0x10, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65,
	0x79, 0x22, 0x83, 0x01, 0x0a, 0x11, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x62, 0x0a, 0x11, 0x49, 0x74, 0x65, 0x6d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x22, 0x69, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22,
	0x14, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x06, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65,
	0x73, 0x2a, 0x23, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x49, 0x52, 0x45, 0x43,
	0x54, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x32, 0xad, 0x03, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x50, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1f, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1d, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1e, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x11,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x3f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x1e, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x42, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simple_drive_proto_rawDescOnce sync.Once
	file_simple_drive_proto_rawDescData = file_simple_drive_proto_rawDesc
)

func file_simple_drive_proto_rawDescGZIP() []byte {
	file_simple_drive_proto_rawDescOnce.Do(func() {
		file_simple_drive_proto_rawDescData = protoimpl.X.CompressGZIP(file_simple_drive_proto_rawDescData)
	})
	return file_simple_drive_proto_rawDescData
}

var file_simple_drive_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_simple_drive_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_simple_drive_proto_goTypes = []interface{}{
	(ItemType)(0),                  // 0: simpledrive.ItemType
	(*Item)(nil),                   // 1: simpledrive.Item
	(*ItemCreateRequest)(nil),      // 2: simpledrive.ItemCreateRequest
	(*ItemUpdateRequest)(nil),      // 3: simpledrive.ItemUpdateRequest
	(*ListItemsRequest)(nil),       // 4: simpledrive.ListItemsRequest
	(*ListItemsResponse)(nil),      // 5: simpledrive.ListItemsResponse
	(*HealthCheckRequest)(nil),     // 6: simpledrive.HealthCheckRequest
	(*HealthCheckResponse)(nil),    // 7: simpledrive.HealthCheckResponse
	(*Paging)(nil),                 // 8: simpledrive.Paging
	(*wrapperspb.Int64Value)(nil),  // 9: google.protobuf.Int64Value
	(*timestamppb.Timestamp)(nil),  // 10: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 11: google.protobuf.StringValue
	(*emptypb.Empty)(nil),          // 12: google.protobuf.Empty
}
var file_simple_drive_proto_depIdxs = []int32{
	0,  // 0: simpledrive.Item.type:type_name -> simpledrive.ItemType
	9,  // 1: simpledrive.Item.size:type_name -> google.protobuf.Int64Value
	10, // 2: simpledrive.Item.created_at:type_name -> google.protobuf.Timestamp
	10, // 3: simpledrive.Item.last_modified_at:type_name -> google.protobuf.Timestamp
	0,  // 4: simpledrive.ItemCreateRequest.type:type_name -> simpledrive.ItemType
	11, // 5: simpledrive.ItemUpdateRequest.name:type_name -> google.protobuf.StringValue
	1,  // 6: simpledrive.ListItemsResponse.items:type_name -> simpledrive.Item
	8,  // 7: simpledrive.ListItemsResponse.paging:type_name -> simpledrive.Paging
	6,  // 8: simpledrive.FileManager.HealthCheck:input_type -> simpledrive.HealthCheckRequest
	4,  // 9: simpledrive.FileManager.ListItems:input_type -> simpledrive.ListItemsRequest
	2,  // 10: simpledrive.FileManager.CreateItem:input_type -> simpledrive.ItemCreateRequest
	11, // 11: simpledrive.FileManager.GetItem:input_type -> google.protobuf.StringValue
	3,  // 12: simpledrive.FileManager.UpdateItem:input_type -> simpledrive.ItemUpdateRequest
	11, // 13: simpledrive.FileManager.DeleteItem:input_type -> google.protobuf.StringValue
	7,  // 14: simpledrive.FileManager.HealthCheck:output_type -> simpledrive.HealthCheckResponse
	5,  // 15: simpledrive.FileManager.ListItems:output_type -> simpledrive.ListItemsResponse
	1,  // 16: simpledrive.FileManager.CreateItem:output_type -> simpledrive.Item
	1,  // 17: simpledrive.FileManager.GetItem:output_type -> simpledrive.Item
	1,  // 18: simpledrive.FileManager.UpdateItem:output_type -> simpledrive.Item
	12, // 19: simpledrive.FileManager.DeleteItem:output_type -> google.protobuf.Empty
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_simple_drive_proto_init() }
func file_simple_drive_proto_init() {
	if File_simple_drive_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_simple_drive_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_simple_drive_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemCreateRequest); i {
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
		file_simple_drive_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemUpdateRequest); i {
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
		file_simple_drive_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsRequest); i {
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
		file_simple_drive_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsResponse); i {
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
		file_simple_drive_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_simple_drive_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
		file_simple_drive_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paging); i {
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
			RawDescriptor: file_simple_drive_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_simple_drive_proto_goTypes,
		DependencyIndexes: file_simple_drive_proto_depIdxs,
		EnumInfos:         file_simple_drive_proto_enumTypes,
		MessageInfos:      file_simple_drive_proto_msgTypes,
	}.Build()
	File_simple_drive_proto = out.File
	file_simple_drive_proto_rawDesc = nil
	file_simple_drive_proto_goTypes = nil
	file_simple_drive_proto_depIdxs = nil
}
