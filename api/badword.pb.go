// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: api/badword.proto

package api

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

type BadwordReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nama          string                 `protobuf:"bytes,1,opt,name=nama,proto3" json:"nama,omitempty"`
	JmlKata       int32                  `protobuf:"varint,2,opt,name=jmlKata,proto3" json:"jmlKata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BadwordReq) Reset() {
	*x = BadwordReq{}
	mi := &file_api_badword_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BadwordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadwordReq) ProtoMessage() {}

func (x *BadwordReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_badword_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadwordReq.ProtoReflect.Descriptor instead.
func (*BadwordReq) Descriptor() ([]byte, []int) {
	return file_api_badword_proto_rawDescGZIP(), []int{0}
}

func (x *BadwordReq) GetNama() string {
	if x != nil {
		return x.Nama
	}
	return ""
}

func (x *BadwordReq) GetJmlKata() int32 {
	if x != nil {
		return x.JmlKata
	}
	return 0
}

type BadwordRes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kalimat       string                 `protobuf:"bytes,1,opt,name=kalimat,proto3" json:"kalimat,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BadwordRes) Reset() {
	*x = BadwordRes{}
	mi := &file_api_badword_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BadwordRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadwordRes) ProtoMessage() {}

func (x *BadwordRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_badword_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadwordRes.ProtoReflect.Descriptor instead.
func (*BadwordRes) Descriptor() ([]byte, []int) {
	return file_api_badword_proto_rawDescGZIP(), []int{1}
}

func (x *BadwordRes) GetKalimat() string {
	if x != nil {
		return x.Kalimat
	}
	return ""
}

var File_api_badword_proto protoreflect.FileDescriptor

const file_api_badword_proto_rawDesc = "" +
	"\n" +
	"\x11api/badword.proto\x12\abadword\":\n" +
	"\n" +
	"BadwordReq\x12\x12\n" +
	"\x04nama\x18\x01 \x01(\tR\x04nama\x12\x18\n" +
	"\ajmlKata\x18\x02 \x01(\x05R\ajmlKata\"&\n" +
	"\n" +
	"BadwordRes\x12\x18\n" +
	"\akalimat\x18\x01 \x01(\tR\akalimat2H\n" +
	"\x0eBadwordService\x126\n" +
	"\n" +
	"SayBadword\x12\x13.badword.BadwordReq\x1a\x13.badword.BadwordResB\aZ\x05./apib\x06proto3"

var (
	file_api_badword_proto_rawDescOnce sync.Once
	file_api_badword_proto_rawDescData []byte
)

func file_api_badword_proto_rawDescGZIP() []byte {
	file_api_badword_proto_rawDescOnce.Do(func() {
		file_api_badword_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_badword_proto_rawDesc), len(file_api_badword_proto_rawDesc)))
	})
	return file_api_badword_proto_rawDescData
}

var file_api_badword_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_badword_proto_goTypes = []any{
	(*BadwordReq)(nil), // 0: badword.BadwordReq
	(*BadwordRes)(nil), // 1: badword.BadwordRes
}
var file_api_badword_proto_depIdxs = []int32{
	0, // 0: badword.BadwordService.SayBadword:input_type -> badword.BadwordReq
	1, // 1: badword.BadwordService.SayBadword:output_type -> badword.BadwordRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_badword_proto_init() }
func file_api_badword_proto_init() {
	if File_api_badword_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_badword_proto_rawDesc), len(file_api_badword_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_badword_proto_goTypes,
		DependencyIndexes: file_api_badword_proto_depIdxs,
		MessageInfos:      file_api_badword_proto_msgTypes,
	}.Build()
	File_api_badword_proto = out.File
	file_api_badword_proto_goTypes = nil
	file_api_badword_proto_depIdxs = nil
}
