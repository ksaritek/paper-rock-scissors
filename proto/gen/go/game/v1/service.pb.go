// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: game/v1/service.proto

package gamev1

import (
	reflect "reflect"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_game_v1_service_proto protoreflect.FileDescriptor

var file_game_v1_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xb9, 0x03, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x72, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x2d, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a,
	0x01, 0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x2f, 0x70, 0x6c, 0x61, 0x79, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x12, 0x6a, 0x0a, 0x0a, 0x45, 0x6e,
	0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x65, 0x6e, 0x64, 0x2d, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x66, 0x92, 0x41, 0x63, 0x12, 0x0c, 0x47, 0x61, 0x6d,
	0x65, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x53, 0x0a, 0x20, 0x46, 0x69, 0x6e,
	0x64, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74,
	0x20, 0x47, 0x61, 0x6d, 0x65, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x73, 0x61, 0x72, 0x69, 0x74, 0x65, 0x6b, 0x2f, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x2d, 0x72, 0x6f, 0x63, 0x6b, 0x2d, 0x73, 0x63, 0x69, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x42, 0x9d,
	0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x73, 0x61, 0x72, 0x69,
	0x74, 0x65, 0x6b, 0x2f, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2d, 0x72, 0x6f, 0x63, 0x6b, 0x2d, 0x73,
	0x63, 0x69, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x61, 0x6d,
	0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x47, 0x61, 0x6d, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13,
	0x47, 0x61, 0x6d, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_game_v1_service_proto_goTypes = []interface{}{
	(*StartSessionRequest)(nil),  // 0: game.v1.StartSessionRequest
	(*PlayGameRequest)(nil),      // 1: game.v1.PlayGameRequest
	(*EndSessionRequest)(nil),    // 2: game.v1.EndSessionRequest
	(*StartSessionResponse)(nil), // 3: game.v1.StartSessionResponse
	(*PlayGameResponse)(nil),     // 4: game.v1.PlayGameResponse
	(*EndSessionResponse)(nil),   // 5: game.v1.EndSessionResponse
}
var file_game_v1_service_proto_depIdxs = []int32{
	0, // 0: game.v1.GameService.StartSession:input_type -> game.v1.StartSessionRequest
	1, // 1: game.v1.GameService.PlayGame:input_type -> game.v1.PlayGameRequest
	2, // 2: game.v1.GameService.EndSession:input_type -> game.v1.EndSessionRequest
	3, // 3: game.v1.GameService.StartSession:output_type -> game.v1.StartSessionResponse
	4, // 4: game.v1.GameService.PlayGame:output_type -> game.v1.PlayGameResponse
	5, // 5: game.v1.GameService.EndSession:output_type -> game.v1.EndSessionResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_v1_service_proto_init() }
func file_game_v1_service_proto_init() {
	if File_game_v1_service_proto != nil {
		return
	}
	file_game_v1_game_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_v1_service_proto_goTypes,
		DependencyIndexes: file_game_v1_service_proto_depIdxs,
	}.Build()
	File_game_v1_service_proto = out.File
	file_game_v1_service_proto_rawDesc = nil
	file_game_v1_service_proto_goTypes = nil
	file_game_v1_service_proto_depIdxs = nil
}
