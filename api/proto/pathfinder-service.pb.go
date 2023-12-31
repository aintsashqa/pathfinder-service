// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.0
// source: api/proto/pathfinder-service.proto

package proto

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

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_pathfinder_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_pathfinder_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_api_proto_pathfinder_service_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position    *Point  `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	Weight      float64 `protobuf:"fixed64,2,opt,name=weight,proto3" json:"weight,omitempty"`
	Unavailable bool    `protobuf:"varint,3,opt,name=unavailable,proto3" json:"unavailable,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_pathfinder_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_pathfinder_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_api_proto_pathfinder_service_proto_rawDescGZIP(), []int{1}
}

func (x *Object) GetPosition() *Point {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Object) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Object) GetUnavailable() bool {
	if x != nil {
		return x.Unavailable
	}
	return false
}

type GetPathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry          *Point    `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
	Final          *Point    `protobuf:"bytes,2,opt,name=final,proto3" json:"final,omitempty"`
	Objects        []*Object `protobuf:"bytes,3,rep,name=objects,proto3" json:"objects,omitempty"`
	Step           float64   `protobuf:"fixed64,4,opt,name=step,proto3" json:"step,omitempty"`
	UseExtraFields bool      `protobuf:"varint,5,opt,name=use_extra_fields,json=useExtraFields,proto3" json:"use_extra_fields,omitempty"`
}

func (x *GetPathRequest) Reset() {
	*x = GetPathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_pathfinder_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPathRequest) ProtoMessage() {}

func (x *GetPathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_pathfinder_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPathRequest.ProtoReflect.Descriptor instead.
func (*GetPathRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_pathfinder_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetPathRequest) GetEntry() *Point {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *GetPathRequest) GetFinal() *Point {
	if x != nil {
		return x.Final
	}
	return nil
}

func (x *GetPathRequest) GetObjects() []*Object {
	if x != nil {
		return x.Objects
	}
	return nil
}

func (x *GetPathRequest) GetStep() float64 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *GetPathRequest) GetUseExtraFields() bool {
	if x != nil {
		return x.UseExtraFields
	}
	return false
}

type GetPathResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path []*Point `protobuf:"bytes,1,rep,name=path,proto3" json:"path,omitempty"`
}

func (x *GetPathResponse) Reset() {
	*x = GetPathResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_pathfinder_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPathResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPathResponse) ProtoMessage() {}

func (x *GetPathResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_pathfinder_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPathResponse.ProtoReflect.Descriptor instead.
func (*GetPathResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_pathfinder_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetPathResponse) GetPath() []*Point {
	if x != nil {
		return x.Path
	}
	return nil
}

var File_api_proto_pathfinder_service_proto protoreflect.FileDescriptor

var file_api_proto_pathfinder_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x74, 0x68,
	0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x22, 0x23, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x01, 0x79, 0x22, 0x71, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x2d, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x6e, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x74,
	0x68, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x2c, 0x0a,
	0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x61, 0x74, 0x68, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12,
	0x28, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x38, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x74,
	0x68, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x32, 0x59, 0x0a, 0x11, 0x50, 0x61, 0x74, 0x68, 0x66, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x66, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13,
	0x5a, 0x11, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_pathfinder_service_proto_rawDescOnce sync.Once
	file_api_proto_pathfinder_service_proto_rawDescData = file_api_proto_pathfinder_service_proto_rawDesc
)

func file_api_proto_pathfinder_service_proto_rawDescGZIP() []byte {
	file_api_proto_pathfinder_service_proto_rawDescOnce.Do(func() {
		file_api_proto_pathfinder_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_pathfinder_service_proto_rawDescData)
	})
	return file_api_proto_pathfinder_service_proto_rawDescData
}

var file_api_proto_pathfinder_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_pathfinder_service_proto_goTypes = []interface{}{
	(*Point)(nil),           // 0: pathfinder.Point
	(*Object)(nil),          // 1: pathfinder.Object
	(*GetPathRequest)(nil),  // 2: pathfinder.GetPathRequest
	(*GetPathResponse)(nil), // 3: pathfinder.GetPathResponse
}
var file_api_proto_pathfinder_service_proto_depIdxs = []int32{
	0, // 0: pathfinder.Object.position:type_name -> pathfinder.Point
	0, // 1: pathfinder.GetPathRequest.entry:type_name -> pathfinder.Point
	0, // 2: pathfinder.GetPathRequest.final:type_name -> pathfinder.Point
	1, // 3: pathfinder.GetPathRequest.objects:type_name -> pathfinder.Object
	0, // 4: pathfinder.GetPathResponse.path:type_name -> pathfinder.Point
	2, // 5: pathfinder.PathfinderService.GetPath:input_type -> pathfinder.GetPathRequest
	3, // 6: pathfinder.PathfinderService.GetPath:output_type -> pathfinder.GetPathResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_pathfinder_service_proto_init() }
func file_api_proto_pathfinder_service_proto_init() {
	if File_api_proto_pathfinder_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_pathfinder_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_api_proto_pathfinder_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_api_proto_pathfinder_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPathRequest); i {
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
		file_api_proto_pathfinder_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPathResponse); i {
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
			RawDescriptor: file_api_proto_pathfinder_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_pathfinder_service_proto_goTypes,
		DependencyIndexes: file_api_proto_pathfinder_service_proto_depIdxs,
		MessageInfos:      file_api_proto_pathfinder_service_proto_msgTypes,
	}.Build()
	File_api_proto_pathfinder_service_proto = out.File
	file_api_proto_pathfinder_service_proto_rawDesc = nil
	file_api_proto_pathfinder_service_proto_goTypes = nil
	file_api_proto_pathfinder_service_proto_depIdxs = nil
}
