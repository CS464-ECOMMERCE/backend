// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: ecommerce.proto

package ecommerce_app

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CartItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	mi := &file_ecommerce_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{0}
}

func (x *CartItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Item          *CartItem              `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddItemRequest) Reset() {
	*x = AddItemRequest{}
	mi := &file_ecommerce_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemRequest) ProtoMessage() {}

func (x *AddItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemRequest.ProtoReflect.Descriptor instead.
func (*AddItemRequest) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{1}
}

func (x *AddItemRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddItemRequest) GetItem() *CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type EmptyCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyCartRequest) Reset() {
	*x = EmptyCartRequest{}
	mi := &file_ecommerce_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyCartRequest) ProtoMessage() {}

func (x *EmptyCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyCartRequest.ProtoReflect.Descriptor instead.
func (*EmptyCartRequest) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{2}
}

func (x *EmptyCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCartRequest) Reset() {
	*x = GetCartRequest{}
	mi := &file_ecommerce_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartRequest) ProtoMessage() {}

func (x *GetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartRequest.ProtoReflect.Descriptor instead.
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{3}
}

func (x *GetCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Cart struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*CartItem            `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cart) Reset() {
	*x = Cart{}
	mi := &file_ecommerce_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{4}
}

func (x *Cart) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Cart) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_ecommerce_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{5}
}

type Product struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price           float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Inventory       int32                  `protobuf:"varint,4,opt,name=inventory,proto3" json:"inventory,omitempty"`
	Description     string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Images          []string               `protobuf:"bytes,6,rep,name=images,proto3" json:"images,omitempty" gorm:"type:text[]"` // [bucketname+filename]
	StripePriceId   string                 `protobuf:"bytes,7,opt,name=stripe_price_id,json=stripePriceId,proto3" json:"stripe_price_id,omitempty"`
	StripeProductId string                 `protobuf:"bytes,8,opt,name=stripe_product_id,json=stripeProductId,proto3" json:"stripe_product_id,omitempty"`
	MerchantId      string                 `protobuf:"bytes,9,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_ecommerce_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{6}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetInventory() int32 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *Product) GetStripePriceId() string {
	if x != nil {
		return x.StripePriceId
	}
	return ""
}

func (x *Product) GetStripeProductId() string {
	if x != nil {
		return x.StripeProductId
	}
	return ""
}

func (x *Product) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

type UpdateProductImagesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProductImagesRequest) Reset() {
	*x = UpdateProductImagesRequest{}
	mi := &file_ecommerce_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductImagesRequest) ProtoMessage() {}

func (x *UpdateProductImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductImagesRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductImagesRequest) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateProductImagesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateProductImagesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PresignedUrl  string                 `protobuf:"bytes,1,opt,name=presigned_url,json=presignedUrl,proto3" json:"presigned_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProductImagesResponse) Reset() {
	*x = UpdateProductImagesResponse{}
	mi := &file_ecommerce_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductImagesResponse) ProtoMessage() {}

func (x *UpdateProductImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductImagesResponse.ProtoReflect.Descriptor instead.
func (*UpdateProductImagesResponse) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateProductImagesResponse) GetPresignedUrl() string {
	if x != nil {
		return x.PresignedUrl
	}
	return ""
}

type CreateProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price         float32                `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Inventory     int32                  `protobuf:"varint,3,opt,name=inventory,proto3" json:"inventory,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	MerchantId    string                 `protobuf:"bytes,5,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"` // repeated string categories = 6;
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	mi := &file_ecommerce_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{9}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetInventory() int32 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

type UpdateProductRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price           float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Inventory       int32                  `protobuf:"varint,4,opt,name=inventory,proto3" json:"inventory,omitempty"`
	Description     string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Images          []string               `protobuf:"bytes,6,rep,name=images,proto3" json:"images,omitempty"`
	StripePriceId   string                 `protobuf:"bytes,7,opt,name=stripe_price_id,json=stripePriceId,proto3" json:"stripe_price_id,omitempty"`
	StripeProductId string                 `protobuf:"bytes,8,opt,name=stripe_product_id,json=stripeProductId,proto3" json:"stripe_product_id,omitempty"`
	MerchantId      string                 `protobuf:"bytes,9,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"` // repeated string categories = 7;
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *UpdateProductRequest) Reset() {
	*x = UpdateProductRequest{}
	mi := &file_ecommerce_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRequest) ProtoMessage() {}

func (x *UpdateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductRequest) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProductRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateProductRequest) GetInventory() int32 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

func (x *UpdateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateProductRequest) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *UpdateProductRequest) GetStripePriceId() string {
	if x != nil {
		return x.StripePriceId
	}
	return ""
}

func (x *UpdateProductRequest) GetStripeProductId() string {
	if x != nil {
		return x.StripeProductId
	}
	return ""
}

func (x *UpdateProductRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

type DeleteProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MerchantId    string                 `protobuf:"bytes,2,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteProductRequest) Reset() {
	*x = DeleteProductRequest{}
	mi := &file_ecommerce_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductRequest) ProtoMessage() {}

func (x *DeleteProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteProductRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

type ListProductsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*Product             `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListProductsResponse) Reset() {
	*x = ListProductsResponse{}
	mi := &file_ecommerce_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResponse) ProtoMessage() {}

func (x *ListProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResponse.ProtoReflect.Descriptor instead.
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{12}
}

func (x *ListProductsResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type GetProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	mi := &file_ecommerce_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecommerce_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_ecommerce_proto_rawDescGZIP(), []int{13}
}

func (x *GetProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_ecommerce_proto protoreflect.FileDescriptor

var file_ecommerce_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x22, 0x45, 0x0a, 0x08,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x52, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x2b, 0x0a, 0x10, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x4a, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x90, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x22, 0xa1, 0x01, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x9d, 0x02,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74, 0x72, 0x69,
	0x70, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x47, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x23,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xbe, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19,
	0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x19, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43,
	0x61, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x32, 0xd5, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1f, 0x2e,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x43, 0x53, 0x34, 0x36, 0x34, 0x2d, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x52,
	0x43, 0x45, 0x2f, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x70, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_ecommerce_proto_rawDescOnce sync.Once
	file_ecommerce_proto_rawDescData []byte
)

func file_ecommerce_proto_rawDescGZIP() []byte {
	file_ecommerce_proto_rawDescOnce.Do(func() {
		file_ecommerce_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ecommerce_proto_rawDesc), len(file_ecommerce_proto_rawDesc)))
	})
	return file_ecommerce_proto_rawDescData
}

var file_ecommerce_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_ecommerce_proto_goTypes = []any{
	(*CartItem)(nil),                    // 0: ecommerce.CartItem
	(*AddItemRequest)(nil),              // 1: ecommerce.AddItemRequest
	(*EmptyCartRequest)(nil),            // 2: ecommerce.EmptyCartRequest
	(*GetCartRequest)(nil),              // 3: ecommerce.GetCartRequest
	(*Cart)(nil),                        // 4: ecommerce.Cart
	(*Empty)(nil),                       // 5: ecommerce.Empty
	(*Product)(nil),                     // 6: ecommerce.Product
	(*UpdateProductImagesRequest)(nil),  // 7: ecommerce.UpdateProductImagesRequest
	(*UpdateProductImagesResponse)(nil), // 8: ecommerce.UpdateProductImagesResponse
	(*CreateProductRequest)(nil),        // 9: ecommerce.CreateProductRequest
	(*UpdateProductRequest)(nil),        // 10: ecommerce.UpdateProductRequest
	(*DeleteProductRequest)(nil),        // 11: ecommerce.DeleteProductRequest
	(*ListProductsResponse)(nil),        // 12: ecommerce.ListProductsResponse
	(*GetProductRequest)(nil),           // 13: ecommerce.GetProductRequest
}
var file_ecommerce_proto_depIdxs = []int32{
	0,  // 0: ecommerce.AddItemRequest.item:type_name -> ecommerce.CartItem
	0,  // 1: ecommerce.Cart.items:type_name -> ecommerce.CartItem
	6,  // 2: ecommerce.ListProductsResponse.products:type_name -> ecommerce.Product
	1,  // 3: ecommerce.CartService.AddItem:input_type -> ecommerce.AddItemRequest
	3,  // 4: ecommerce.CartService.GetCart:input_type -> ecommerce.GetCartRequest
	2,  // 5: ecommerce.CartService.EmptyCart:input_type -> ecommerce.EmptyCartRequest
	5,  // 6: ecommerce.ProductService.ListProducts:input_type -> ecommerce.Empty
	13, // 7: ecommerce.ProductService.GetProduct:input_type -> ecommerce.GetProductRequest
	9,  // 8: ecommerce.ProductService.CreateProduct:input_type -> ecommerce.CreateProductRequest
	11, // 9: ecommerce.ProductService.DeleteProduct:input_type -> ecommerce.DeleteProductRequest
	10, // 10: ecommerce.ProductService.UpdateProduct:input_type -> ecommerce.UpdateProductRequest
	7,  // 11: ecommerce.ProductService.UpdateProductImages:input_type -> ecommerce.UpdateProductImagesRequest
	5,  // 12: ecommerce.CartService.AddItem:output_type -> ecommerce.Empty
	4,  // 13: ecommerce.CartService.GetCart:output_type -> ecommerce.Cart
	5,  // 14: ecommerce.CartService.EmptyCart:output_type -> ecommerce.Empty
	12, // 15: ecommerce.ProductService.ListProducts:output_type -> ecommerce.ListProductsResponse
	6,  // 16: ecommerce.ProductService.GetProduct:output_type -> ecommerce.Product
	6,  // 17: ecommerce.ProductService.CreateProduct:output_type -> ecommerce.Product
	5,  // 18: ecommerce.ProductService.DeleteProduct:output_type -> ecommerce.Empty
	6,  // 19: ecommerce.ProductService.UpdateProduct:output_type -> ecommerce.Product
	8,  // 20: ecommerce.ProductService.UpdateProductImages:output_type -> ecommerce.UpdateProductImagesResponse
	12, // [12:21] is the sub-list for method output_type
	3,  // [3:12] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_ecommerce_proto_init() }
func file_ecommerce_proto_init() {
	if File_ecommerce_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ecommerce_proto_rawDesc), len(file_ecommerce_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_ecommerce_proto_goTypes,
		DependencyIndexes: file_ecommerce_proto_depIdxs,
		MessageInfos:      file_ecommerce_proto_msgTypes,
	}.Build()
	File_ecommerce_proto = out.File
	file_ecommerce_proto_goTypes = nil
	file_ecommerce_proto_depIdxs = nil
}
