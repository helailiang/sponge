// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.2
// source: api/types/types.proto

package types

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

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32     `protobuf:"varint,1,opt,name=page,proto3" json:"page"`      // page number, starting from 0
	Limit   int32     `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`    // lines per page
	Sort    string    `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort"`       // sorted fields, multi-column sorting separated by commas
	Columns []*Column `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns"` // query conditions
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_types_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_api_types_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_api_types_types_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Params) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Params) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *Params) GetColumns() []*Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`   // column name
	Exp   string `protobuf:"bytes,2,opt,name=exp,proto3" json:"exp"`     // expressions, default value is "=", support =, !=, >, >=, <, <=, like, in, notin
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value"` // column value
	Logic string `protobuf:"bytes,4,opt,name=logic,proto3" json:"logic"` // logical type, default value is "and", support &, and, ||, or
}

func (x *Column) Reset() {
	*x = Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_types_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_api_types_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_api_types_types_proto_rawDescGZIP(), []int{1}
}

func (x *Column) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Column) GetExp() string {
	if x != nil {
		return x.Exp
	}
	return ""
}

func (x *Column) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Column) GetLogic() string {
	if x != nil {
		return x.Logic
	}
	return ""
}

type Conditions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Columns []*Column `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns"` // query conditions
}

func (x *Conditions) Reset() {
	*x = Conditions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_types_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conditions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conditions) ProtoMessage() {}

func (x *Conditions) ProtoReflect() protoreflect.Message {
	mi := &file_api_types_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conditions.ProtoReflect.Descriptor instead.
func (*Conditions) Descriptor() ([]byte, []int) {
	return file_api_types_types_proto_rawDescGZIP(), []int{2}
}

func (x *Conditions) GetColumns() []*Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

var File_api_types_types_proto protoreflect.FileDescriptor

var file_api_types_types_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x6f,
	0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22,
	0x5a, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x22, 0x35, 0x0a, 0x0a, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x7a, 0x68, 0x75, 0x66, 0x75, 0x79, 0x69, 0x2f, 0x73, 0x70, 0x6f, 0x6e, 0x67, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_types_types_proto_rawDescOnce sync.Once
	file_api_types_types_proto_rawDescData = file_api_types_types_proto_rawDesc
)

func file_api_types_types_proto_rawDescGZIP() []byte {
	file_api_types_types_proto_rawDescOnce.Do(func() {
		file_api_types_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_types_types_proto_rawDescData)
	})
	return file_api_types_types_proto_rawDescData
}

var file_api_types_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_types_types_proto_goTypes = []interface{}{
	(*Params)(nil),     // 0: types.Params
	(*Column)(nil),     // 1: types.Column
	(*Conditions)(nil), // 2: types.Conditions
}
var file_api_types_types_proto_depIdxs = []int32{
	1, // 0: types.Params.columns:type_name -> types.Column
	1, // 1: types.Conditions.columns:type_name -> types.Column
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_types_types_proto_init() }
func file_api_types_types_proto_init() {
	if File_api_types_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_types_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
		file_api_types_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column); i {
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
		file_api_types_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conditions); i {
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
			RawDescriptor: file_api_types_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_types_types_proto_goTypes,
		DependencyIndexes: file_api_types_types_proto_depIdxs,
		MessageInfos:      file_api_types_types_proto_msgTypes,
	}.Build()
	File_api_types_types_proto = out.File
	file_api_types_types_proto_rawDesc = nil
	file_api_types_types_proto_goTypes = nil
	file_api_types_types_proto_depIdxs = nil
}
