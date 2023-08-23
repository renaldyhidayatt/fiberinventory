// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: productkeluar.proto

package pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ProductKeluar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Qty        string               `protobuf:"bytes,2,opt,name=qty,proto3" json:"qty,omitempty"`
	Product    *Product             `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
	ProductId  string               `protobuf:"bytes,4,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Category   *Category            `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	CategoryId string               `protobuf:"bytes,6,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	CreatedAt  *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ProductKeluar) Reset() {
	*x = ProductKeluar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productkeluar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductKeluar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductKeluar) ProtoMessage() {}

func (x *ProductKeluar) ProtoReflect() protoreflect.Message {
	mi := &file_productkeluar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductKeluar.ProtoReflect.Descriptor instead.
func (*ProductKeluar) Descriptor() ([]byte, []int) {
	return file_productkeluar_proto_rawDescGZIP(), []int{0}
}

func (x *ProductKeluar) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductKeluar) GetQty() string {
	if x != nil {
		return x.Qty
	}
	return ""
}

func (x *ProductKeluar) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *ProductKeluar) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ProductKeluar) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *ProductKeluar) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *ProductKeluar) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProductKeluar) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ProductKeluarsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProductKeluarsRequest) Reset() {
	*x = ProductKeluarsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productkeluar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductKeluarsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductKeluarsRequest) ProtoMessage() {}

func (x *ProductKeluarsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_productkeluar_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductKeluarsRequest.ProtoReflect.Descriptor instead.
func (*ProductKeluarsRequest) Descriptor() ([]byte, []int) {
	return file_productkeluar_proto_rawDescGZIP(), []int{1}
}

type ProductKeluarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductKeluar *ProductKeluar `protobuf:"bytes,1,opt,name=product_keluar,json=productKeluar,proto3" json:"product_keluar,omitempty"`
}

func (x *ProductKeluarResponse) Reset() {
	*x = ProductKeluarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productkeluar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductKeluarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductKeluarResponse) ProtoMessage() {}

func (x *ProductKeluarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_productkeluar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductKeluarResponse.ProtoReflect.Descriptor instead.
func (*ProductKeluarResponse) Descriptor() ([]byte, []int) {
	return file_productkeluar_proto_rawDescGZIP(), []int{2}
}

func (x *ProductKeluarResponse) GetProductKeluar() *ProductKeluar {
	if x != nil {
		return x.ProductKeluar
	}
	return nil
}

type ProductKeluarsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductKeluars []*ProductKeluar `protobuf:"bytes,1,rep,name=product_keluars,json=productKeluars,proto3" json:"product_keluars,omitempty"`
}

func (x *ProductKeluarsResponse) Reset() {
	*x = ProductKeluarsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productkeluar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductKeluarsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductKeluarsResponse) ProtoMessage() {}

func (x *ProductKeluarsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_productkeluar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductKeluarsResponse.ProtoReflect.Descriptor instead.
func (*ProductKeluarsResponse) Descriptor() ([]byte, []int) {
	return file_productkeluar_proto_rawDescGZIP(), []int{3}
}

func (x *ProductKeluarsResponse) GetProductKeluars() []*ProductKeluar {
	if x != nil {
		return x.ProductKeluars
	}
	return nil
}

type CreateProductKeluarInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qty        string `protobuf:"bytes,1,opt,name=qty,proto3" json:"qty,omitempty"`
	ProductId  string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	CategoryId string `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *CreateProductKeluarInput) Reset() {
	*x = CreateProductKeluarInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productkeluar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductKeluarInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductKeluarInput) ProtoMessage() {}

func (x *CreateProductKeluarInput) ProtoReflect() protoreflect.Message {
	mi := &file_productkeluar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductKeluarInput.ProtoReflect.Descriptor instead.
func (*CreateProductKeluarInput) Descriptor() ([]byte, []int) {
	return file_productkeluar_proto_rawDescGZIP(), []int{4}
}

func (x *CreateProductKeluarInput) GetQty() string {
	if x != nil {
		return x.Qty
	}
	return ""
}

func (x *CreateProductKeluarInput) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CreateProductKeluarInput) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type UpdateProductKeluarInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Qty        string `protobuf:"bytes,2,opt,name=qty,proto3" json:"qty,omitempty"`
	ProductId  string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	CategoryId string `protobuf:"bytes,4,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *UpdateProductKeluarInput) Reset() {
	*x = UpdateProductKeluarInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productkeluar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductKeluarInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductKeluarInput) ProtoMessage() {}

func (x *UpdateProductKeluarInput) ProtoReflect() protoreflect.Message {
	mi := &file_productkeluar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductKeluarInput.ProtoReflect.Descriptor instead.
func (*UpdateProductKeluarInput) Descriptor() ([]byte, []int) {
	return file_productkeluar_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateProductKeluarInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductKeluarInput) GetQty() string {
	if x != nil {
		return x.Qty
	}
	return ""
}

func (x *UpdateProductKeluarInput) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *UpdateProductKeluarInput) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type ProductKeluarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductKeluarRequest) Reset() {
	*x = ProductKeluarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productkeluar_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductKeluarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductKeluarRequest) ProtoMessage() {}

func (x *ProductKeluarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_productkeluar_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductKeluarRequest.ProtoReflect.Descriptor instead.
func (*ProductKeluarRequest) Descriptor() ([]byte, []int) {
	return file_productkeluar_proto_rawDescGZIP(), []int{6}
}

func (x *ProductKeluarRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteProductKeluarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteProductKeluarResponse) Reset() {
	*x = DeleteProductKeluarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_productkeluar_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductKeluarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductKeluarResponse) ProtoMessage() {}

func (x *DeleteProductKeluarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_productkeluar_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductKeluarResponse.ProtoReflect.Descriptor instead.
func (*DeleteProductKeluarResponse) Descriptor() ([]byte, []int) {
	return file_productkeluar_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteProductKeluarResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_productkeluar_proto protoreflect.FileDescriptor

var file_productkeluar_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x02, 0x0a, 0x0d, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x71,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x71, 0x74, 0x79, 0x12, 0x25, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b,
	0x65, 0x6c, 0x75, 0x61, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a,
	0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x6b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61,
	0x72, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72,
	0x22, 0x54, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b,
	0x65, 0x6c, 0x75, 0x61, 0x72, 0x73, 0x22, 0x6c, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x71, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x22, 0x7c, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x71,
	0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c,
	0x75, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x1b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0xa7, 0x03, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b,
	0x65, 0x6c, 0x75, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c,
	0x75, 0x61, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65,
	0x6c, 0x75, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75,
	0x61, 0x72, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72,
	0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c,
	0x75, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x12, 0x1c,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c, 0x75, 0x61, 0x72,
	0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c,
	0x75, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4b, 0x65, 0x6c,
	0x75, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1c, 0x5a,
	0x1a, 0x66, 0x69, 0x62, 0x65, 0x72, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_productkeluar_proto_rawDescOnce sync.Once
	file_productkeluar_proto_rawDescData = file_productkeluar_proto_rawDesc
)

func file_productkeluar_proto_rawDescGZIP() []byte {
	file_productkeluar_proto_rawDescOnce.Do(func() {
		file_productkeluar_proto_rawDescData = protoimpl.X.CompressGZIP(file_productkeluar_proto_rawDescData)
	})
	return file_productkeluar_proto_rawDescData
}

var file_productkeluar_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_productkeluar_proto_goTypes = []interface{}{
	(*ProductKeluar)(nil),               // 0: pb.ProductKeluar
	(*ProductKeluarsRequest)(nil),       // 1: pb.ProductKeluarsRequest
	(*ProductKeluarResponse)(nil),       // 2: pb.ProductKeluarResponse
	(*ProductKeluarsResponse)(nil),      // 3: pb.ProductKeluarsResponse
	(*CreateProductKeluarInput)(nil),    // 4: pb.CreateProductKeluarInput
	(*UpdateProductKeluarInput)(nil),    // 5: pb.UpdateProductKeluarInput
	(*ProductKeluarRequest)(nil),        // 6: pb.ProductKeluarRequest
	(*DeleteProductKeluarResponse)(nil), // 7: pb.DeleteProductKeluarResponse
	(*Product)(nil),                     // 8: pb.Product
	(*Category)(nil),                    // 9: pb.Category
	(*timestamp.Timestamp)(nil),         // 10: google.protobuf.Timestamp
}
var file_productkeluar_proto_depIdxs = []int32{
	8,  // 0: pb.ProductKeluar.product:type_name -> pb.Product
	9,  // 1: pb.ProductKeluar.category:type_name -> pb.Category
	10, // 2: pb.ProductKeluar.created_at:type_name -> google.protobuf.Timestamp
	10, // 3: pb.ProductKeluar.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 4: pb.ProductKeluarResponse.product_keluar:type_name -> pb.ProductKeluar
	0,  // 5: pb.ProductKeluarsResponse.product_keluars:type_name -> pb.ProductKeluar
	4,  // 6: pb.ProductKeluarService.CreateProductKeluar:input_type -> pb.CreateProductKeluarInput
	1,  // 7: pb.ProductKeluarService.GetProductKeluars:input_type -> pb.ProductKeluarsRequest
	6,  // 8: pb.ProductKeluarService.GetProductKeluar:input_type -> pb.ProductKeluarRequest
	5,  // 9: pb.ProductKeluarService.UpdateProductKeluar:input_type -> pb.UpdateProductKeluarInput
	6,  // 10: pb.ProductKeluarService.DeleteProductKeluar:input_type -> pb.ProductKeluarRequest
	2,  // 11: pb.ProductKeluarService.CreateProductKeluar:output_type -> pb.ProductKeluarResponse
	3,  // 12: pb.ProductKeluarService.GetProductKeluars:output_type -> pb.ProductKeluarsResponse
	2,  // 13: pb.ProductKeluarService.GetProductKeluar:output_type -> pb.ProductKeluarResponse
	2,  // 14: pb.ProductKeluarService.UpdateProductKeluar:output_type -> pb.ProductKeluarResponse
	7,  // 15: pb.ProductKeluarService.DeleteProductKeluar:output_type -> pb.DeleteProductKeluarResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_productkeluar_proto_init() }
func file_productkeluar_proto_init() {
	if File_productkeluar_proto != nil {
		return
	}
	file_category_proto_init()
	file_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_productkeluar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductKeluar); i {
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
		file_productkeluar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductKeluarsRequest); i {
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
		file_productkeluar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductKeluarResponse); i {
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
		file_productkeluar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductKeluarsResponse); i {
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
		file_productkeluar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductKeluarInput); i {
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
		file_productkeluar_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductKeluarInput); i {
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
		file_productkeluar_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductKeluarRequest); i {
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
		file_productkeluar_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductKeluarResponse); i {
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
			RawDescriptor: file_productkeluar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_productkeluar_proto_goTypes,
		DependencyIndexes: file_productkeluar_proto_depIdxs,
		MessageInfos:      file_productkeluar_proto_msgTypes,
	}.Build()
	File_productkeluar_proto = out.File
	file_productkeluar_proto_rawDesc = nil
	file_productkeluar_proto_goTypes = nil
	file_productkeluar_proto_depIdxs = nil
}
