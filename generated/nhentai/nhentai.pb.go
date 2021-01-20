// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: nhentai.proto

package nhentai

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

type GalleryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GalleryRequest) Reset() {
	*x = GalleryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nhentai_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GalleryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GalleryRequest) ProtoMessage() {}

func (x *GalleryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nhentai_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GalleryRequest.ProtoReflect.Descriptor instead.
func (*GalleryRequest) Descriptor() ([]byte, []int) {
	return file_nhentai_proto_rawDescGZIP(), []int{0}
}

func (x *GalleryRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GalleryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MediaId      int32                   `protobuf:"varint,2,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	Title        *GalleryResponse_Title  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Images       *GalleryResponse_Images `protobuf:"bytes,4,opt,name=images,proto3" json:"images,omitempty"`
	Scanlator    string                  `protobuf:"bytes,5,opt,name=scanlator,proto3" json:"scanlator,omitempty"`
	UploadDate   string                  `protobuf:"bytes,6,opt,name=upload_date,json=uploadDate,proto3" json:"upload_date,omitempty"`
	Tags         []*GalleryResponse_Tag  `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	NumPages     int32                   `protobuf:"varint,8,opt,name=num_pages,json=numPages,proto3" json:"num_pages,omitempty"`
	NumFavorites int32                   `protobuf:"varint,9,opt,name=num_favorites,json=numFavorites,proto3" json:"num_favorites,omitempty"`
}

func (x *GalleryResponse) Reset() {
	*x = GalleryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nhentai_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GalleryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GalleryResponse) ProtoMessage() {}

func (x *GalleryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nhentai_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GalleryResponse.ProtoReflect.Descriptor instead.
func (*GalleryResponse) Descriptor() ([]byte, []int) {
	return file_nhentai_proto_rawDescGZIP(), []int{1}
}

func (x *GalleryResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GalleryResponse) GetMediaId() int32 {
	if x != nil {
		return x.MediaId
	}
	return 0
}

func (x *GalleryResponse) GetTitle() *GalleryResponse_Title {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *GalleryResponse) GetImages() *GalleryResponse_Images {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *GalleryResponse) GetScanlator() string {
	if x != nil {
		return x.Scanlator
	}
	return ""
}

func (x *GalleryResponse) GetUploadDate() string {
	if x != nil {
		return x.UploadDate
	}
	return ""
}

func (x *GalleryResponse) GetTags() []*GalleryResponse_Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GalleryResponse) GetNumPages() int32 {
	if x != nil {
		return x.NumPages
	}
	return 0
}

func (x *GalleryResponse) GetNumFavorites() int32 {
	if x != nil {
		return x.NumFavorites
	}
	return 0
}

type GalleryResponse_Title struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Japanese string `protobuf:"bytes,1,opt,name=japanese,proto3" json:"japanese,omitempty"`
	English  string `protobuf:"bytes,2,opt,name=english,proto3" json:"english,omitempty"`
	Pretty   string `protobuf:"bytes,3,opt,name=pretty,proto3" json:"pretty,omitempty"`
}

func (x *GalleryResponse_Title) Reset() {
	*x = GalleryResponse_Title{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nhentai_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GalleryResponse_Title) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GalleryResponse_Title) ProtoMessage() {}

func (x *GalleryResponse_Title) ProtoReflect() protoreflect.Message {
	mi := &file_nhentai_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GalleryResponse_Title.ProtoReflect.Descriptor instead.
func (*GalleryResponse_Title) Descriptor() ([]byte, []int) {
	return file_nhentai_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GalleryResponse_Title) GetJapanese() string {
	if x != nil {
		return x.Japanese
	}
	return ""
}

func (x *GalleryResponse_Title) GetEnglish() string {
	if x != nil {
		return x.English
	}
	return ""
}

func (x *GalleryResponse_Title) GetPretty() string {
	if x != nil {
		return x.Pretty
	}
	return ""
}

type GalleryResponse_Images struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pages     []*GalleryResponse_Images_Image `protobuf:"bytes,1,rep,name=pages,proto3" json:"pages,omitempty"`
	Cover     *GalleryResponse_Images_Image   `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`
	Thumbnail *GalleryResponse_Images_Image   `protobuf:"bytes,3,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
}

func (x *GalleryResponse_Images) Reset() {
	*x = GalleryResponse_Images{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nhentai_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GalleryResponse_Images) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GalleryResponse_Images) ProtoMessage() {}

func (x *GalleryResponse_Images) ProtoReflect() protoreflect.Message {
	mi := &file_nhentai_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GalleryResponse_Images.ProtoReflect.Descriptor instead.
func (*GalleryResponse_Images) Descriptor() ([]byte, []int) {
	return file_nhentai_proto_rawDescGZIP(), []int{1, 1}
}

func (x *GalleryResponse_Images) GetPages() []*GalleryResponse_Images_Image {
	if x != nil {
		return x.Pages
	}
	return nil
}

func (x *GalleryResponse_Images) GetCover() *GalleryResponse_Images_Image {
	if x != nil {
		return x.Cover
	}
	return nil
}

func (x *GalleryResponse_Images) GetThumbnail() *GalleryResponse_Images_Image {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

type GalleryResponse_Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Url   string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Count int32  `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GalleryResponse_Tag) Reset() {
	*x = GalleryResponse_Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nhentai_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GalleryResponse_Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GalleryResponse_Tag) ProtoMessage() {}

func (x *GalleryResponse_Tag) ProtoReflect() protoreflect.Message {
	mi := &file_nhentai_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GalleryResponse_Tag.ProtoReflect.Descriptor instead.
func (*GalleryResponse_Tag) Descriptor() ([]byte, []int) {
	return file_nhentai_proto_rawDescGZIP(), []int{1, 2}
}

func (x *GalleryResponse_Tag) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GalleryResponse_Tag) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GalleryResponse_Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GalleryResponse_Tag) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GalleryResponse_Tag) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GalleryResponse_Images_Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T string `protobuf:"bytes,1,opt,name=t,proto3" json:"t,omitempty"`
	W int32  `protobuf:"varint,2,opt,name=w,proto3" json:"w,omitempty"`
	H int32  `protobuf:"varint,3,opt,name=h,proto3" json:"h,omitempty"`
}

func (x *GalleryResponse_Images_Image) Reset() {
	*x = GalleryResponse_Images_Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nhentai_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GalleryResponse_Images_Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GalleryResponse_Images_Image) ProtoMessage() {}

func (x *GalleryResponse_Images_Image) ProtoReflect() protoreflect.Message {
	mi := &file_nhentai_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GalleryResponse_Images_Image.ProtoReflect.Descriptor instead.
func (*GalleryResponse_Images_Image) Descriptor() ([]byte, []int) {
	return file_nhentai_proto_rawDescGZIP(), []int{1, 1, 0}
}

func (x *GalleryResponse_Images_Image) GetT() string {
	if x != nil {
		return x.T
	}
	return ""
}

func (x *GalleryResponse_Images_Image) GetW() int32 {
	if x != nil {
		return x.W
	}
	return 0
}

func (x *GalleryResponse_Images_Image) GetH() int32 {
	if x != nil {
		return x.H
	}
	return 0
}

var File_nhentai_proto protoreflect.FileDescriptor

var file_nhentai_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x68, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6e, 0x68, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x99, 0x06, 0x0a, 0x0f, 0x47,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x68, 0x65, 0x6e, 0x74,
	0x61, 0x69, 0x2e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x37, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x6e, 0x68, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63, 0x61, 0x6e,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x61,
	0x6e, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x68, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e,
	0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x75,
	0x6d, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e,
	0x75, 0x6d, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x1a, 0x55, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x61, 0x70, 0x61, 0x6e, 0x65, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x61, 0x70, 0x61, 0x6e, 0x65, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x74, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x74,
	0x74, 0x79, 0x1a, 0xfa, 0x01, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x3b, 0x0a,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e,
	0x68, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x05, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x68, 0x65, 0x6e,
	0x74, 0x61, 0x69, 0x2e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x68, 0x65,
	0x6e, 0x74, 0x61, 0x69, 0x2e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x1a, 0x31, 0x0a, 0x05,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x77, 0x12, 0x0c, 0x0a, 0x01, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x68, 0x1a,
	0x65, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x45, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x12, 0x3e, 0x0a,
	0x07, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x12, 0x17, 0x2e, 0x6e, 0x68, 0x65, 0x6e, 0x74,
	0x61, 0x69, 0x2e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x68, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x2e, 0x47, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a,
	0x09, 0x2e, 0x3b, 0x6e, 0x68, 0x65, 0x6e, 0x74, 0x61, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_nhentai_proto_rawDescOnce sync.Once
	file_nhentai_proto_rawDescData = file_nhentai_proto_rawDesc
)

func file_nhentai_proto_rawDescGZIP() []byte {
	file_nhentai_proto_rawDescOnce.Do(func() {
		file_nhentai_proto_rawDescData = protoimpl.X.CompressGZIP(file_nhentai_proto_rawDescData)
	})
	return file_nhentai_proto_rawDescData
}

var file_nhentai_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_nhentai_proto_goTypes = []interface{}{
	(*GalleryRequest)(nil),               // 0: nhentai.GalleryRequest
	(*GalleryResponse)(nil),              // 1: nhentai.GalleryResponse
	(*GalleryResponse_Title)(nil),        // 2: nhentai.GalleryResponse.Title
	(*GalleryResponse_Images)(nil),       // 3: nhentai.GalleryResponse.Images
	(*GalleryResponse_Tag)(nil),          // 4: nhentai.GalleryResponse.Tag
	(*GalleryResponse_Images_Image)(nil), // 5: nhentai.GalleryResponse.Images.Image
}
var file_nhentai_proto_depIdxs = []int32{
	2, // 0: nhentai.GalleryResponse.title:type_name -> nhentai.GalleryResponse.Title
	3, // 1: nhentai.GalleryResponse.images:type_name -> nhentai.GalleryResponse.Images
	4, // 2: nhentai.GalleryResponse.tags:type_name -> nhentai.GalleryResponse.Tag
	5, // 3: nhentai.GalleryResponse.Images.pages:type_name -> nhentai.GalleryResponse.Images.Image
	5, // 4: nhentai.GalleryResponse.Images.cover:type_name -> nhentai.GalleryResponse.Images.Image
	5, // 5: nhentai.GalleryResponse.Images.thumbnail:type_name -> nhentai.GalleryResponse.Images.Image
	0, // 6: nhentai.API.Gallery:input_type -> nhentai.GalleryRequest
	1, // 7: nhentai.API.Gallery:output_type -> nhentai.GalleryResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_nhentai_proto_init() }
func file_nhentai_proto_init() {
	if File_nhentai_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nhentai_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GalleryRequest); i {
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
		file_nhentai_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GalleryResponse); i {
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
		file_nhentai_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GalleryResponse_Title); i {
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
		file_nhentai_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GalleryResponse_Images); i {
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
		file_nhentai_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GalleryResponse_Tag); i {
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
		file_nhentai_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GalleryResponse_Images_Image); i {
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
			RawDescriptor: file_nhentai_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nhentai_proto_goTypes,
		DependencyIndexes: file_nhentai_proto_depIdxs,
		MessageInfos:      file_nhentai_proto_msgTypes,
	}.Build()
	File_nhentai_proto = out.File
	file_nhentai_proto_rawDesc = nil
	file_nhentai_proto_goTypes = nil
	file_nhentai_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	Gallery(ctx context.Context, in *GalleryRequest, opts ...grpc.CallOption) (*GalleryResponse, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Gallery(ctx context.Context, in *GalleryRequest, opts ...grpc.CallOption) (*GalleryResponse, error) {
	out := new(GalleryResponse)
	err := c.cc.Invoke(ctx, "/nhentai.API/Gallery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	Gallery(context.Context, *GalleryRequest) (*GalleryResponse, error)
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) Gallery(context.Context, *GalleryRequest) (*GalleryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gallery not implemented")
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Gallery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GalleryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Gallery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nhentai.API/Gallery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Gallery(ctx, req.(*GalleryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nhentai.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Gallery",
			Handler:    _API_Gallery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nhentai.proto",
}