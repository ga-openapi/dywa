// Code generated by protoc-gen-go. DO NOT EDIT.
// source: douyin-webapp/video.proto

package douyin_webapp // import "github.com/dev-openapi/douyin-webapp"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateImageTextReq struct {
	OpenId               string                   `protobuf:"bytes,1,opt,name=open_id,json=openId" json:"open_id,omitempty"`
	Body                 *CreateImageTextReq_Body `protobuf:"bytes,10,opt,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CreateImageTextReq) Reset()         { *m = CreateImageTextReq{} }
func (m *CreateImageTextReq) String() string { return proto.CompactTextString(m) }
func (*CreateImageTextReq) ProtoMessage()    {}
func (*CreateImageTextReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_666eec00951b5219, []int{0}
}
func (m *CreateImageTextReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImageTextReq.Unmarshal(m, b)
}
func (m *CreateImageTextReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImageTextReq.Marshal(b, m, deterministic)
}
func (dst *CreateImageTextReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImageTextReq.Merge(dst, src)
}
func (m *CreateImageTextReq) XXX_Size() int {
	return xxx_messageInfo_CreateImageTextReq.Size(m)
}
func (m *CreateImageTextReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImageTextReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImageTextReq proto.InternalMessageInfo

func (m *CreateImageTextReq) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *CreateImageTextReq) GetBody() *CreateImageTextReq_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

type CreateImageTextReq_Body struct {
	ImageList            []string `protobuf:"bytes,1,rep,name=image_list,json=imageList" json:"image_list,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	AtUsers              []string `protobuf:"bytes,3,rep,name=at_users,json=atUsers" json:"at_users,omitempty"`
	MicroAppTitle        string   `protobuf:"bytes,4,opt,name=micro_app_title,json=microAppTitle" json:"micro_app_title,omitempty"`
	MicroAppUrl          string   `protobuf:"bytes,5,opt,name=micro_app_url,json=microAppUrl" json:"micro_app_url,omitempty"`
	MicroAppiId          string   `protobuf:"bytes,6,opt,name=micro_appi_id,json=microAppiId" json:"micro_appi_id,omitempty"`
	PoiId                string   `protobuf:"bytes,7,opt,name=poi_id,json=poiId" json:"poi_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImageTextReq_Body) Reset()         { *m = CreateImageTextReq_Body{} }
func (m *CreateImageTextReq_Body) String() string { return proto.CompactTextString(m) }
func (*CreateImageTextReq_Body) ProtoMessage()    {}
func (*CreateImageTextReq_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_666eec00951b5219, []int{0, 0}
}
func (m *CreateImageTextReq_Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImageTextReq_Body.Unmarshal(m, b)
}
func (m *CreateImageTextReq_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImageTextReq_Body.Marshal(b, m, deterministic)
}
func (dst *CreateImageTextReq_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImageTextReq_Body.Merge(dst, src)
}
func (m *CreateImageTextReq_Body) XXX_Size() int {
	return xxx_messageInfo_CreateImageTextReq_Body.Size(m)
}
func (m *CreateImageTextReq_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImageTextReq_Body.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImageTextReq_Body proto.InternalMessageInfo

func (m *CreateImageTextReq_Body) GetImageList() []string {
	if m != nil {
		return m.ImageList
	}
	return nil
}

func (m *CreateImageTextReq_Body) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *CreateImageTextReq_Body) GetAtUsers() []string {
	if m != nil {
		return m.AtUsers
	}
	return nil
}

func (m *CreateImageTextReq_Body) GetMicroAppTitle() string {
	if m != nil {
		return m.MicroAppTitle
	}
	return ""
}

func (m *CreateImageTextReq_Body) GetMicroAppUrl() string {
	if m != nil {
		return m.MicroAppUrl
	}
	return ""
}

func (m *CreateImageTextReq_Body) GetMicroAppiId() string {
	if m != nil {
		return m.MicroAppiId
	}
	return ""
}

func (m *CreateImageTextReq_Body) GetPoiId() string {
	if m != nil {
		return m.PoiId
	}
	return ""
}

type CreateImageTextRes struct {
	Extra                *Extra                   `protobuf:"bytes,1,opt,name=extra" json:"extra,omitempty"`
	Data                 *CreateImageTextRes_Data `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CreateImageTextRes) Reset()         { *m = CreateImageTextRes{} }
func (m *CreateImageTextRes) String() string { return proto.CompactTextString(m) }
func (*CreateImageTextRes) ProtoMessage()    {}
func (*CreateImageTextRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_666eec00951b5219, []int{1}
}
func (m *CreateImageTextRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImageTextRes.Unmarshal(m, b)
}
func (m *CreateImageTextRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImageTextRes.Marshal(b, m, deterministic)
}
func (dst *CreateImageTextRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImageTextRes.Merge(dst, src)
}
func (m *CreateImageTextRes) XXX_Size() int {
	return xxx_messageInfo_CreateImageTextRes.Size(m)
}
func (m *CreateImageTextRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImageTextRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImageTextRes proto.InternalMessageInfo

func (m *CreateImageTextRes) GetExtra() *Extra {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *CreateImageTextRes) GetData() *CreateImageTextRes_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateImageTextRes_Data struct {
	ItemId               string   `protobuf:"bytes,1,opt,name=item_id,json=itemId" json:"item_id,omitempty"`
	VideoId              string   `protobuf:"bytes,2,opt,name=video_id,json=videoId" json:"video_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImageTextRes_Data) Reset()         { *m = CreateImageTextRes_Data{} }
func (m *CreateImageTextRes_Data) String() string { return proto.CompactTextString(m) }
func (*CreateImageTextRes_Data) ProtoMessage()    {}
func (*CreateImageTextRes_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_666eec00951b5219, []int{1, 0}
}
func (m *CreateImageTextRes_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImageTextRes_Data.Unmarshal(m, b)
}
func (m *CreateImageTextRes_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImageTextRes_Data.Marshal(b, m, deterministic)
}
func (dst *CreateImageTextRes_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImageTextRes_Data.Merge(dst, src)
}
func (m *CreateImageTextRes_Data) XXX_Size() int {
	return xxx_messageInfo_CreateImageTextRes_Data.Size(m)
}
func (m *CreateImageTextRes_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImageTextRes_Data.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImageTextRes_Data proto.InternalMessageInfo

func (m *CreateImageTextRes_Data) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *CreateImageTextRes_Data) GetVideoId() string {
	if m != nil {
		return m.VideoId
	}
	return ""
}

type UploadImageReq struct {
	OpenId               string               `protobuf:"bytes,1,opt,name=open_id,json=openId" json:"open_id,omitempty"`
	Body                 *UploadImageReq_Body `protobuf:"bytes,10,opt,name=body" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UploadImageReq) Reset()         { *m = UploadImageReq{} }
func (m *UploadImageReq) String() string { return proto.CompactTextString(m) }
func (*UploadImageReq) ProtoMessage()    {}
func (*UploadImageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_666eec00951b5219, []int{2}
}
func (m *UploadImageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageReq.Unmarshal(m, b)
}
func (m *UploadImageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageReq.Marshal(b, m, deterministic)
}
func (dst *UploadImageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageReq.Merge(dst, src)
}
func (m *UploadImageReq) XXX_Size() int {
	return xxx_messageInfo_UploadImageReq.Size(m)
}
func (m *UploadImageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageReq.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageReq proto.InternalMessageInfo

func (m *UploadImageReq) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *UploadImageReq) GetBody() *UploadImageReq_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

type UploadImageReq_Body struct {
	Image                []byte   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadImageReq_Body) Reset()         { *m = UploadImageReq_Body{} }
func (m *UploadImageReq_Body) String() string { return proto.CompactTextString(m) }
func (*UploadImageReq_Body) ProtoMessage()    {}
func (*UploadImageReq_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_666eec00951b5219, []int{2, 0}
}
func (m *UploadImageReq_Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageReq_Body.Unmarshal(m, b)
}
func (m *UploadImageReq_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageReq_Body.Marshal(b, m, deterministic)
}
func (dst *UploadImageReq_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageReq_Body.Merge(dst, src)
}
func (m *UploadImageReq_Body) XXX_Size() int {
	return xxx_messageInfo_UploadImageReq_Body.Size(m)
}
func (m *UploadImageReq_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageReq_Body.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageReq_Body proto.InternalMessageInfo

func (m *UploadImageReq_Body) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

type Image struct {
	ImageId              string   `protobuf:"bytes,1,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
	Height               int32    `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Width                int32    `protobuf:"varint,3,opt,name=width" json:"width,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_666eec00951b5219, []int{3}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *Image) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Image) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

type UploadImageRes struct {
	Extra                *Extra               `protobuf:"bytes,1,opt,name=extra" json:"extra,omitempty"`
	Data                 *UploadImageRes_Data `protobuf:"bytes,10,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UploadImageRes) Reset()         { *m = UploadImageRes{} }
func (m *UploadImageRes) String() string { return proto.CompactTextString(m) }
func (*UploadImageRes) ProtoMessage()    {}
func (*UploadImageRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_666eec00951b5219, []int{4}
}
func (m *UploadImageRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageRes.Unmarshal(m, b)
}
func (m *UploadImageRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageRes.Marshal(b, m, deterministic)
}
func (dst *UploadImageRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageRes.Merge(dst, src)
}
func (m *UploadImageRes) XXX_Size() int {
	return xxx_messageInfo_UploadImageRes.Size(m)
}
func (m *UploadImageRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageRes.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageRes proto.InternalMessageInfo

func (m *UploadImageRes) GetExtra() *Extra {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *UploadImageRes) GetData() *UploadImageRes_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type UploadImageRes_Data struct {
	Image                *Image   `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadImageRes_Data) Reset()         { *m = UploadImageRes_Data{} }
func (m *UploadImageRes_Data) String() string { return proto.CompactTextString(m) }
func (*UploadImageRes_Data) ProtoMessage()    {}
func (*UploadImageRes_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_666eec00951b5219, []int{4, 0}
}
func (m *UploadImageRes_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageRes_Data.Unmarshal(m, b)
}
func (m *UploadImageRes_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageRes_Data.Marshal(b, m, deterministic)
}
func (dst *UploadImageRes_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageRes_Data.Merge(dst, src)
}
func (m *UploadImageRes_Data) XXX_Size() int {
	return xxx_messageInfo_UploadImageRes_Data.Size(m)
}
func (m *UploadImageRes_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageRes_Data.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageRes_Data proto.InternalMessageInfo

func (m *UploadImageRes_Data) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateImageTextReq)(nil), "open.douyin.com.CreateImageTextReq")
	proto.RegisterType((*CreateImageTextReq_Body)(nil), "open.douyin.com.CreateImageTextReq.Body")
	proto.RegisterType((*CreateImageTextRes)(nil), "open.douyin.com.CreateImageTextRes")
	proto.RegisterType((*CreateImageTextRes_Data)(nil), "open.douyin.com.CreateImageTextRes.Data")
	proto.RegisterType((*UploadImageReq)(nil), "open.douyin.com.UploadImageReq")
	proto.RegisterType((*UploadImageReq_Body)(nil), "open.douyin.com.UploadImageReq.Body")
	proto.RegisterType((*Image)(nil), "open.douyin.com.Image")
	proto.RegisterType((*UploadImageRes)(nil), "open.douyin.com.UploadImageRes")
	proto.RegisterType((*UploadImageRes_Data)(nil), "open.douyin.com.UploadImageRes.Data")
}

func init() { proto.RegisterFile("douyin-webapp/video.proto", fileDescriptor_video_666eec00951b5219) }

var fileDescriptor_video_666eec00951b5219 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x93, 0x38, 0x21, 0x9b, 0x96, 0x48, 0x2b, 0x28, 0xae, 0x55, 0x44, 0x30, 0x55, 0x94,
	0x43, 0x6a, 0x8b, 0x50, 0xa4, 0x2a, 0xe2, 0x42, 0x81, 0x83, 0x25, 0x0e, 0xc8, 0x34, 0x1c, 0xb8,
	0x58, 0x9b, 0xec, 0x2a, 0x59, 0xc9, 0xc9, 0x2e, 0xf6, 0x3a, 0x4d, 0x8e, 0x70, 0xe0, 0x8e, 0x78,
	0x0c, 0x24, 0x9e, 0x80, 0x17, 0xe0, 0xcc, 0x91, 0x2b, 0x0f, 0x82, 0x76, 0x36, 0xd0, 0xfc, 0x89,
	0x86, 0x9b, 0xbf, 0x9d, 0x6f, 0x7e, 0xbf, 0xf1, 0xa0, 0x43, 0x2a, 0xf2, 0x39, 0x9f, 0x9c, 0x5c,
	0xb2, 0x3e, 0x91, 0x32, 0x98, 0x72, 0xca, 0x84, 0x2f, 0x53, 0xa1, 0x04, 0xae, 0x0b, 0xc9, 0x26,
	0xbe, 0xb1, 0xfb, 0x03, 0x31, 0x76, 0x8f, 0x86, 0x42, 0x0c, 0x13, 0x16, 0x10, 0xc9, 0x03, 0x32,
	0x99, 0x08, 0x45, 0x14, 0x17, 0x93, 0xcc, 0xd0, 0xdd, 0xb5, 0x48, 0x6c, 0xa6, 0x52, 0x62, 0x4c,
	0xde, 0xf7, 0x02, 0xc2, 0xcf, 0x52, 0x46, 0x14, 0x0b, 0xc7, 0x64, 0xc8, 0x2e, 0xd8, 0x4c, 0x45,
	0xec, 0x1d, 0xbe, 0x83, 0x2a, 0x3a, 0x45, 0xcc, 0xa9, 0x63, 0x35, 0xac, 0x56, 0x35, 0x2a, 0x6b,
	0x18, 0x52, 0xfc, 0x04, 0x95, 0xfa, 0x82, 0xce, 0x1d, 0xd4, 0xb0, 0x5a, 0xb5, 0x4e, 0xcb, 0x5f,
	0x2b, 0xc4, 0xdf, 0x8c, 0xe5, 0x9f, 0x0b, 0x3a, 0x8f, 0xc0, 0xcb, 0xfd, 0x69, 0xa1, 0x92, 0x86,
	0xf8, 0x2e, 0x42, 0x5c, 0x73, 0xe2, 0x84, 0x67, 0xca, 0xb1, 0x1a, 0xc5, 0x56, 0x35, 0xaa, 0xc2,
	0xcb, 0x4b, 0x9e, 0x29, 0x8c, 0x51, 0x49, 0xb1, 0x99, 0x72, 0x0a, 0x90, 0x1b, 0xbe, 0xf1, 0x21,
	0xba, 0x41, 0x54, 0x9c, 0x67, 0x2c, 0xcd, 0x9c, 0x22, 0x38, 0x54, 0x88, 0xea, 0x69, 0x88, 0x9b,
	0xa8, 0x3e, 0xe6, 0x83, 0x54, 0xc4, 0x44, 0xca, 0x58, 0x71, 0x95, 0x30, 0xa7, 0x04, 0x9e, 0xfb,
	0xf0, 0xfc, 0x54, 0xca, 0x0b, 0xfd, 0x88, 0x3d, 0xb4, 0x7f, 0xc5, 0xcb, 0xd3, 0xc4, 0xb1, 0x81,
	0x55, 0xfb, 0xc3, 0xea, 0xa5, 0xc9, 0x0a, 0x87, 0xeb, 0xfe, 0xcb, 0xab, 0x1c, 0x1e, 0x52, 0x7c,
	0x1b, 0x95, 0xa5, 0x00, 0x63, 0x05, 0x8c, 0xb6, 0x14, 0x3c, 0xa4, 0xde, 0x37, 0x6b, 0xcb, 0x2c,
	0x33, 0xdc, 0x46, 0x36, 0x4c, 0x1c, 0x26, 0x59, 0xeb, 0x1c, 0x6c, 0xcc, 0xec, 0x85, 0xb6, 0x46,
	0x86, 0xa4, 0x07, 0x4c, 0x89, 0x22, 0xd0, 0xfa, 0x4e, 0x03, 0xce, 0xfc, 0xe7, 0x44, 0x91, 0x08,
	0xbc, 0xdc, 0x2e, 0x2a, 0x69, 0xa4, 0xf5, 0xe3, 0x8a, 0x8d, 0x97, 0xf4, 0xd3, 0x30, 0xa4, 0x7a,
	0x8a, 0xb0, 0x48, 0xda, 0x62, 0xa6, 0x5b, 0x01, 0x1c, 0x52, 0xef, 0xbd, 0x85, 0x6e, 0xf6, 0x64,
	0x22, 0x08, 0x85, 0xe8, 0xff, 0x5c, 0x83, 0xb3, 0x95, 0x35, 0x38, 0xde, 0xa8, 0x72, 0x35, 0xce,
	0xf2, 0x0a, 0x1c, 0x2d, 0x36, 0xe0, 0x16, 0xb2, 0x41, 0x6f, 0x08, 0xbc, 0x17, 0x19, 0xe0, 0xbd,
	0x42, 0x36, 0x38, 0xe9, 0x3a, 0xcd, 0x82, 0xfc, 0x4d, 0x5d, 0x01, 0x1c, 0x52, 0x7c, 0x80, 0xca,
	0x23, 0xc6, 0x87, 0x23, 0xb3, 0x1e, 0x76, 0xb4, 0x40, 0x3a, 0xe2, 0x25, 0xa7, 0x6a, 0xe4, 0x14,
	0xe1, 0xd9, 0x00, 0xef, 0xeb, 0x7a, 0x57, 0xff, 0x2b, 0xc8, 0xd9, 0x42, 0x90, 0x9d, 0x5a, 0x5d,
	0x11, 0xe3, 0x74, 0x21, 0x46, 0x7b, 0xb9, 0xd5, 0x6d, 0xf9, 0x8c, 0xb3, 0x21, 0x75, 0xbe, 0x14,
	0xd0, 0xde, 0x1b, 0x2d, 0xc9, 0x6b, 0x96, 0x4e, 0xf9, 0x80, 0xe1, 0x4f, 0x16, 0xaa, 0xaf, 0xa9,
	0x8e, 0x1f, 0xec, 0xf0, 0xe3, 0xb9, 0x3b, 0x90, 0x32, 0xef, 0xf4, 0xc3, 0x8f, 0x5f, 0x9f, 0x0b,
	0xbe, 0xd7, 0x84, 0xdb, 0x61, 0xb8, 0xc1, 0xf4, 0xa1, 0x39, 0x37, 0xc1, 0x00, 0x1c, 0x62, 0x23,
	0x85, 0xfe, 0x0d, 0xbb, 0xa0, 0x22, 0xfe, 0x68, 0xa1, 0xda, 0x52, 0xe3, 0xf8, 0xde, 0x35, 0x1b,
	0xe0, 0x5e, 0x43, 0xc8, 0xbc, 0xc7, 0x50, 0x47, 0xe0, 0xdd, 0xdf, 0x5a, 0x47, 0x0e, 0x64, 0x53,
	0x47, 0x17, 0xe9, 0x12, 0xda, 0xe3, 0x3c, 0x51, 0xfc, 0xbc, 0xf9, 0xf6, 0x78, 0xc8, 0xd5, 0x28,
	0xef, 0xeb, 0x98, 0x01, 0x65, 0xd3, 0x13, 0x9d, 0xe7, 0x2a, 0xc4, 0xe2, 0xe6, 0xf5, 0xcb, 0x70,
	0xee, 0x1e, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x48, 0xb0, 0x0e, 0xa5, 0x55, 0x05, 0x00, 0x00,
}