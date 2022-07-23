// 表示正在使用proto2命令

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: a.proto

package example

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

type Per_PhoneType int32

const (
	Per_MOBILE Per_PhoneType = 0 //手机号
	Per_HOME   Per_PhoneType = 1 //家庭联系电话
	Per_WORK   Per_PhoneType = 2 //工作联系电话
)

// Enum value maps for Per_PhoneType.
var (
	Per_PhoneType_name = map[int32]string{
		0: "MOBILE",
		1: "HOME",
		2: "WORK",
	}
	Per_PhoneType_value = map[string]int32{
		"MOBILE": 0,
		"HOME":   1,
		"WORK":   2,
	}
)

func (x Per_PhoneType) Enum() *Per_PhoneType {
	p := new(Per_PhoneType)
	*p = x
	return p
}

func (x Per_PhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Per_PhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_a_proto_enumTypes[0].Descriptor()
}

func (Per_PhoneType) Type() protoreflect.EnumType {
	return &file_a_proto_enumTypes[0]
}

func (x Per_PhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Per_PhoneType.Descriptor instead.
func (Per_PhoneType) EnumDescriptor() ([]byte, []int) {
	return file_a_proto_rawDescGZIP(), []int{0, 0}
}

//编译器将生成一个名为person的类
//类的字段信息包括姓名name,编号id,邮箱email，
//以及电话号码phones
type Per struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // (位置1)
	Id     int32              `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email  string             `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`   // (位置2)
	Phones []*Per_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"` // (位置4)
}

func (x *Per) Reset() {
	*x = Per{}
	if protoimpl.UnsafeEnabled {
		mi := &file_a_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Per) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Per) ProtoMessage() {}

func (x *Per) ProtoReflect() protoreflect.Message {
	mi := &file_a_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Per.ProtoReflect.Descriptor instead.
func (*Per) Descriptor() ([]byte, []int) {
	return file_a_proto_rawDescGZIP(), []int{0}
}

func (x *Per) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Per) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Per) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Per) GetPhones() []*Per_PhoneNumber {
	if x != nil {
		return x.Phones
	}
	return nil
}

// 通讯录消息体，包括一个Person类的people
type AddressBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	People []*Per `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
}

func (x *AddressBook) Reset() {
	*x = AddressBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_a_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressBook) ProtoMessage() {}

func (x *AddressBook) ProtoReflect() protoreflect.Message {
	mi := &file_a_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressBook.ProtoReflect.Descriptor instead.
func (*AddressBook) Descriptor() ([]byte, []int) {
	return file_a_proto_rawDescGZIP(), []int{1}
}

func (x *AddressBook) GetPeople() []*Per {
	if x != nil {
		return x.People
	}
	return nil
}

//电话号码phone消息体
//组成包括号码number、电话类型 type
type Per_PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string         `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type   *Per_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=Per_PhoneType,oneof" json:"type,omitempty"` // (位置3)
}

func (x *Per_PhoneNumber) Reset() {
	*x = Per_PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_a_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Per_PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Per_PhoneNumber) ProtoMessage() {}

func (x *Per_PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_a_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Per_PhoneNumber.ProtoReflect.Descriptor instead.
func (*Per_PhoneNumber) Descriptor() ([]byte, []int) {
	return file_a_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Per_PhoneNumber) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Per_PhoneNumber) GetType() Per_PhoneType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Per_MOBILE
}

var File_a_proto protoreflect.FileDescriptor

var file_a_proto_rawDesc = []byte{
	0x0a, 0x07, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x03, 0x50, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x28, 0x0a, 0x06, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x50, 0x65,
	0x72, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x73, 0x1a, 0x57, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x50, 0x65, 0x72,
	0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2b,
	0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d, 0x45, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x02, 0x22, 0x2b, 0x0a, 0x0b, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x0a, 0x06, 0x70, 0x65,
	0x6f, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x50, 0x65, 0x72,
	0x52, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_a_proto_rawDescOnce sync.Once
	file_a_proto_rawDescData = file_a_proto_rawDesc
)

func file_a_proto_rawDescGZIP() []byte {
	file_a_proto_rawDescOnce.Do(func() {
		file_a_proto_rawDescData = protoimpl.X.CompressGZIP(file_a_proto_rawDescData)
	})
	return file_a_proto_rawDescData
}

var file_a_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_a_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_a_proto_goTypes = []interface{}{
	(Per_PhoneType)(0),      // 0: Per.PhoneType
	(*Per)(nil),             // 1: Per
	(*AddressBook)(nil),     // 2: AddressBook
	(*Per_PhoneNumber)(nil), // 3: Per.PhoneNumber
}
var file_a_proto_depIdxs = []int32{
	3, // 0: Per.phones:type_name -> Per.PhoneNumber
	1, // 1: AddressBook.people:type_name -> Per
	0, // 2: Per.PhoneNumber.type:type_name -> Per.PhoneType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_a_proto_init() }
func file_a_proto_init() {
	if File_a_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_a_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Per); i {
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
		file_a_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressBook); i {
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
		file_a_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Per_PhoneNumber); i {
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
	file_a_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_a_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_a_proto_goTypes,
		DependencyIndexes: file_a_proto_depIdxs,
		EnumInfos:         file_a_proto_enumTypes,
		MessageInfos:      file_a_proto_msgTypes,
	}.Build()
	File_a_proto = out.File
	file_a_proto_rawDesc = nil
	file_a_proto_goTypes = nil
	file_a_proto_depIdxs = nil
}
