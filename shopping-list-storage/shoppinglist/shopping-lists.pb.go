// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: shopping-lists.proto

package shoppinglist

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetCustomerShoppingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetCustomerShoppingListRequest) Reset() {
	*x = GetCustomerShoppingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_lists_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerShoppingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerShoppingListRequest) ProtoMessage() {}

func (x *GetCustomerShoppingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_lists_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerShoppingListRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerShoppingListRequest) Descriptor() ([]byte, []int) {
	return file_shopping_lists_proto_rawDescGZIP(), []int{0}
}

func (x *GetCustomerShoppingListRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type RemoveCustomerShoppingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int32 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *RemoveCustomerShoppingListRequest) Reset() {
	*x = RemoveCustomerShoppingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_lists_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCustomerShoppingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCustomerShoppingListRequest) ProtoMessage() {}

func (x *RemoveCustomerShoppingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_lists_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCustomerShoppingListRequest.ProtoReflect.Descriptor instead.
func (*RemoveCustomerShoppingListRequest) Descriptor() ([]byte, []int) {
	return file_shopping_lists_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveCustomerShoppingListRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type ChangeCustomerShoppingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId   int32                 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ShoppingList *CustomerShoppingList `protobuf:"bytes,2,opt,name=shopping_list,json=shoppingList,proto3" json:"shopping_list,omitempty"`
}

func (x *ChangeCustomerShoppingListRequest) Reset() {
	*x = ChangeCustomerShoppingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_lists_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeCustomerShoppingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeCustomerShoppingListRequest) ProtoMessage() {}

func (x *ChangeCustomerShoppingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_lists_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeCustomerShoppingListRequest.ProtoReflect.Descriptor instead.
func (*ChangeCustomerShoppingListRequest) Descriptor() ([]byte, []int) {
	return file_shopping_lists_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeCustomerShoppingListRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *ChangeCustomerShoppingListRequest) GetShoppingList() *CustomerShoppingList {
	if x != nil {
		return x.ShoppingList
	}
	return nil
}

type ChangeCustomerShoppingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ChangeCustomerShoppingListResponse) Reset() {
	*x = ChangeCustomerShoppingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_lists_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeCustomerShoppingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeCustomerShoppingListResponse) ProtoMessage() {}

func (x *ChangeCustomerShoppingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_lists_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeCustomerShoppingListResponse.ProtoReflect.Descriptor instead.
func (*ChangeCustomerShoppingListResponse) Descriptor() ([]byte, []int) {
	return file_shopping_lists_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeCustomerShoppingListResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RemoveCustomerShoppingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveCustomerShoppingListResponse) Reset() {
	*x = RemoveCustomerShoppingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_lists_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCustomerShoppingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCustomerShoppingListResponse) ProtoMessage() {}

func (x *RemoveCustomerShoppingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_lists_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCustomerShoppingListResponse.ProtoReflect.Descriptor instead.
func (*RemoveCustomerShoppingListResponse) Descriptor() ([]byte, []int) {
	return file_shopping_lists_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveCustomerShoppingListResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CustomerShoppingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int32                        `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Items      []*CustomerShoppingList_Item `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CustomerShoppingList) Reset() {
	*x = CustomerShoppingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_lists_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerShoppingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerShoppingList) ProtoMessage() {}

func (x *CustomerShoppingList) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_lists_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerShoppingList.ProtoReflect.Descriptor instead.
func (*CustomerShoppingList) Descriptor() ([]byte, []int) {
	return file_shopping_lists_proto_rawDescGZIP(), []int{5}
}

func (x *CustomerShoppingList) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CustomerShoppingList) GetItems() []*CustomerShoppingList_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetCustomerShoppingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Empty                bool                  `protobuf:"varint,1,opt,name=Empty,proto3" json:"Empty,omitempty"`
	CustomerShoppingList *CustomerShoppingList `protobuf:"bytes,2,opt,name=customer_shopping_list,json=customerShoppingList,proto3" json:"customer_shopping_list,omitempty"`
}

func (x *GetCustomerShoppingListResponse) Reset() {
	*x = GetCustomerShoppingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_lists_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerShoppingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerShoppingListResponse) ProtoMessage() {}

func (x *GetCustomerShoppingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_lists_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerShoppingListResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerShoppingListResponse) Descriptor() ([]byte, []int) {
	return file_shopping_lists_proto_rawDescGZIP(), []int{6}
}

func (x *GetCustomerShoppingListResponse) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

func (x *GetCustomerShoppingListResponse) GetCustomerShoppingList() *CustomerShoppingList {
	if x != nil {
		return x.CustomerShoppingList
	}
	return nil
}

type CustomerShoppingList_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId       int32 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemQuantity int32 `protobuf:"varint,2,opt,name=item_quantity,json=itemQuantity,proto3" json:"item_quantity,omitempty"`
}

func (x *CustomerShoppingList_Item) Reset() {
	*x = CustomerShoppingList_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_lists_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerShoppingList_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerShoppingList_Item) ProtoMessage() {}

func (x *CustomerShoppingList_Item) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_lists_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerShoppingList_Item.ProtoReflect.Descriptor instead.
func (*CustomerShoppingList_Item) Descriptor() ([]byte, []int) {
	return file_shopping_lists_proto_rawDescGZIP(), []int{5, 0}
}

func (x *CustomerShoppingList_Item) GetItemId() int32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *CustomerShoppingList_Item) GetItemQuantity() int32 {
	if x != nil {
		return x.ItemQuantity
	}
	return 0
}

var File_shopping_lists_proto protoreflect.FileDescriptor

var file_shopping_lists_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x44, 0x0a, 0x21, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8d, 0x01,
	0x0a, 0x21, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x0c, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3e, 0x0a,
	0x22, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x3e, 0x0a,
	0x22, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xbc, 0x01,
	0x0a, 0x14, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x44, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x69, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x91, 0x01, 0x0a,
	0x1f, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x58, 0x0a, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x14, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x32, 0x98, 0x03, 0x0a, 0x14, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x78, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x81, 0x01, 0x0a, 0x1a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2f, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x6c, 0x69, 0x73,
	0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x81, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x0d, 0x2f,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x74, 0xaa, 0x02, 0x14, 0x53,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shopping_lists_proto_rawDescOnce sync.Once
	file_shopping_lists_proto_rawDescData = file_shopping_lists_proto_rawDesc
)

func file_shopping_lists_proto_rawDescGZIP() []byte {
	file_shopping_lists_proto_rawDescOnce.Do(func() {
		file_shopping_lists_proto_rawDescData = protoimpl.X.CompressGZIP(file_shopping_lists_proto_rawDescData)
	})
	return file_shopping_lists_proto_rawDescData
}

var file_shopping_lists_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_shopping_lists_proto_goTypes = []interface{}{
	(*GetCustomerShoppingListRequest)(nil),     // 0: shoppinglist.GetCustomerShoppingListRequest
	(*RemoveCustomerShoppingListRequest)(nil),  // 1: shoppinglist.RemoveCustomerShoppingListRequest
	(*ChangeCustomerShoppingListRequest)(nil),  // 2: shoppinglist.ChangeCustomerShoppingListRequest
	(*ChangeCustomerShoppingListResponse)(nil), // 3: shoppinglist.ChangeCustomerShoppingListResponse
	(*RemoveCustomerShoppingListResponse)(nil), // 4: shoppinglist.RemoveCustomerShoppingListResponse
	(*CustomerShoppingList)(nil),               // 5: shoppinglist.CustomerShoppingList
	(*GetCustomerShoppingListResponse)(nil),    // 6: shoppinglist.GetCustomerShoppingListResponse
	(*CustomerShoppingList_Item)(nil),          // 7: shoppinglist.CustomerShoppingList.Item
}
var file_shopping_lists_proto_depIdxs = []int32{
	5, // 0: shoppinglist.ChangeCustomerShoppingListRequest.shopping_list:type_name -> shoppinglist.CustomerShoppingList
	7, // 1: shoppinglist.CustomerShoppingList.items:type_name -> shoppinglist.CustomerShoppingList.Item
	5, // 2: shoppinglist.GetCustomerShoppingListResponse.customer_shopping_list:type_name -> shoppinglist.CustomerShoppingList
	0, // 3: shoppinglist.ShoppingListsStorage.GetCustomerShoppingList:input_type -> shoppinglist.GetCustomerShoppingListRequest
	2, // 4: shoppinglist.ShoppingListsStorage.ChangeCustomerShoppingList:input_type -> shoppinglist.ChangeCustomerShoppingListRequest
	1, // 5: shoppinglist.ShoppingListsStorage.RemoveCustomerShoppingList:input_type -> shoppinglist.RemoveCustomerShoppingListRequest
	6, // 6: shoppinglist.ShoppingListsStorage.GetCustomerShoppingList:output_type -> shoppinglist.GetCustomerShoppingListResponse
	3, // 7: shoppinglist.ShoppingListsStorage.ChangeCustomerShoppingList:output_type -> shoppinglist.ChangeCustomerShoppingListResponse
	4, // 8: shoppinglist.ShoppingListsStorage.RemoveCustomerShoppingList:output_type -> shoppinglist.RemoveCustomerShoppingListResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_shopping_lists_proto_init() }
func file_shopping_lists_proto_init() {
	if File_shopping_lists_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shopping_lists_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerShoppingListRequest); i {
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
		file_shopping_lists_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCustomerShoppingListRequest); i {
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
		file_shopping_lists_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeCustomerShoppingListRequest); i {
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
		file_shopping_lists_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeCustomerShoppingListResponse); i {
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
		file_shopping_lists_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCustomerShoppingListResponse); i {
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
		file_shopping_lists_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerShoppingList); i {
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
		file_shopping_lists_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerShoppingListResponse); i {
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
		file_shopping_lists_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerShoppingList_Item); i {
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
			RawDescriptor: file_shopping_lists_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shopping_lists_proto_goTypes,
		DependencyIndexes: file_shopping_lists_proto_depIdxs,
		MessageInfos:      file_shopping_lists_proto_msgTypes,
	}.Build()
	File_shopping_lists_proto = out.File
	file_shopping_lists_proto_rawDesc = nil
	file_shopping_lists_proto_goTypes = nil
	file_shopping_lists_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShoppingListsStorageClient is the client API for ShoppingListsStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShoppingListsStorageClient interface {
	GetCustomerShoppingList(ctx context.Context, in *GetCustomerShoppingListRequest, opts ...grpc.CallOption) (*GetCustomerShoppingListResponse, error)
	ChangeCustomerShoppingList(ctx context.Context, in *ChangeCustomerShoppingListRequest, opts ...grpc.CallOption) (*ChangeCustomerShoppingListResponse, error)
	RemoveCustomerShoppingList(ctx context.Context, in *RemoveCustomerShoppingListRequest, opts ...grpc.CallOption) (*RemoveCustomerShoppingListResponse, error)
}

type shoppingListsStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewShoppingListsStorageClient(cc grpc.ClientConnInterface) ShoppingListsStorageClient {
	return &shoppingListsStorageClient{cc}
}

func (c *shoppingListsStorageClient) GetCustomerShoppingList(ctx context.Context, in *GetCustomerShoppingListRequest, opts ...grpc.CallOption) (*GetCustomerShoppingListResponse, error) {
	out := new(GetCustomerShoppingListResponse)
	err := c.cc.Invoke(ctx, "/shoppinglist.ShoppingListsStorage/GetCustomerShoppingList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingListsStorageClient) ChangeCustomerShoppingList(ctx context.Context, in *ChangeCustomerShoppingListRequest, opts ...grpc.CallOption) (*ChangeCustomerShoppingListResponse, error) {
	out := new(ChangeCustomerShoppingListResponse)
	err := c.cc.Invoke(ctx, "/shoppinglist.ShoppingListsStorage/ChangeCustomerShoppingList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingListsStorageClient) RemoveCustomerShoppingList(ctx context.Context, in *RemoveCustomerShoppingListRequest, opts ...grpc.CallOption) (*RemoveCustomerShoppingListResponse, error) {
	out := new(RemoveCustomerShoppingListResponse)
	err := c.cc.Invoke(ctx, "/shoppinglist.ShoppingListsStorage/RemoveCustomerShoppingList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoppingListsStorageServer is the server API for ShoppingListsStorage service.
type ShoppingListsStorageServer interface {
	GetCustomerShoppingList(context.Context, *GetCustomerShoppingListRequest) (*GetCustomerShoppingListResponse, error)
	ChangeCustomerShoppingList(context.Context, *ChangeCustomerShoppingListRequest) (*ChangeCustomerShoppingListResponse, error)
	RemoveCustomerShoppingList(context.Context, *RemoveCustomerShoppingListRequest) (*RemoveCustomerShoppingListResponse, error)
}

// UnimplementedShoppingListsStorageServer can be embedded to have forward compatible implementations.
type UnimplementedShoppingListsStorageServer struct {
}

func (*UnimplementedShoppingListsStorageServer) GetCustomerShoppingList(context.Context, *GetCustomerShoppingListRequest) (*GetCustomerShoppingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerShoppingList not implemented")
}
func (*UnimplementedShoppingListsStorageServer) ChangeCustomerShoppingList(context.Context, *ChangeCustomerShoppingListRequest) (*ChangeCustomerShoppingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeCustomerShoppingList not implemented")
}
func (*UnimplementedShoppingListsStorageServer) RemoveCustomerShoppingList(context.Context, *RemoveCustomerShoppingListRequest) (*RemoveCustomerShoppingListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCustomerShoppingList not implemented")
}

func RegisterShoppingListsStorageServer(s *grpc.Server, srv ShoppingListsStorageServer) {
	s.RegisterService(&_ShoppingListsStorage_serviceDesc, srv)
}

func _ShoppingListsStorage_GetCustomerShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingListsStorageServer).GetCustomerShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppinglist.ShoppingListsStorage/GetCustomerShoppingList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingListsStorageServer).GetCustomerShoppingList(ctx, req.(*GetCustomerShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingListsStorage_ChangeCustomerShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeCustomerShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingListsStorageServer).ChangeCustomerShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppinglist.ShoppingListsStorage/ChangeCustomerShoppingList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingListsStorageServer).ChangeCustomerShoppingList(ctx, req.(*ChangeCustomerShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingListsStorage_RemoveCustomerShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCustomerShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingListsStorageServer).RemoveCustomerShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shoppinglist.ShoppingListsStorage/RemoveCustomerShoppingList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingListsStorageServer).RemoveCustomerShoppingList(ctx, req.(*RemoveCustomerShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShoppingListsStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shoppinglist.ShoppingListsStorage",
	HandlerType: (*ShoppingListsStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerShoppingList",
			Handler:    _ShoppingListsStorage_GetCustomerShoppingList_Handler,
		},
		{
			MethodName: "ChangeCustomerShoppingList",
			Handler:    _ShoppingListsStorage_ChangeCustomerShoppingList_Handler,
		},
		{
			MethodName: "RemoveCustomerShoppingList",
			Handler:    _ShoppingListsStorage_RemoveCustomerShoppingList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shopping-lists.proto",
}
