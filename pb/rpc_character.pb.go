// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.0
// source: services/characters/rpc_character.proto

package pb

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

type CreateCharacterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ShortDescription string `protobuf:"bytes,2,opt,name=shortDescription,proto3" json:"shortDescription,omitempty"`
	WorldId          int32  `protobuf:"varint,3,opt,name=worldId,proto3" json:"worldId,omitempty"`
	SystemId         int32  `protobuf:"varint,4,opt,name=systemId,proto3" json:"systemId,omitempty"`
}

func (x *CreateCharacterRequest) Reset() {
	*x = CreateCharacterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_characters_rpc_character_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCharacterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCharacterRequest) ProtoMessage() {}

func (x *CreateCharacterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_characters_rpc_character_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCharacterRequest.ProtoReflect.Descriptor instead.
func (*CreateCharacterRequest) Descriptor() ([]byte, []int) {
	return file_services_characters_rpc_character_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCharacterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCharacterRequest) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

func (x *CreateCharacterRequest) GetWorldId() int32 {
	if x != nil {
		return x.WorldId
	}
	return 0
}

func (x *CreateCharacterRequest) GetSystemId() int32 {
	if x != nil {
		return x.SystemId
	}
	return 0
}

type CreateCharacterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Character *Character  `protobuf:"bytes,1,opt,name=character,proto3" json:"character,omitempty"`
	Module    *ViewModule `protobuf:"bytes,2,opt,name=module,proto3" json:"module,omitempty"`
}

func (x *CreateCharacterResponse) Reset() {
	*x = CreateCharacterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_characters_rpc_character_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCharacterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCharacterResponse) ProtoMessage() {}

func (x *CreateCharacterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_characters_rpc_character_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCharacterResponse.ProtoReflect.Descriptor instead.
func (*CreateCharacterResponse) Descriptor() ([]byte, []int) {
	return file_services_characters_rpc_character_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCharacterResponse) GetCharacter() *Character {
	if x != nil {
		return x.Character
	}
	return nil
}

func (x *CreateCharacterResponse) GetModule() *ViewModule {
	if x != nil {
		return x.Module
	}
	return nil
}

type UpdateCharacterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CharacterId      int32   `protobuf:"varint,1,opt,name=characterId,proto3" json:"characterId,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	ShortDescription *string `protobuf:"bytes,3,opt,name=shortDescription,proto3,oneof" json:"shortDescription,omitempty"`
	Public           *bool   `protobuf:"varint,4,opt,name=public,proto3,oneof" json:"public,omitempty"`
	WorldId          *int32  `protobuf:"varint,5,opt,name=worldId,proto3,oneof" json:"worldId,omitempty"`
	SystemId         *int32  `protobuf:"varint,6,opt,name=systemId,proto3,oneof" json:"systemId,omitempty"`
}

func (x *UpdateCharacterRequest) Reset() {
	*x = UpdateCharacterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_characters_rpc_character_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCharacterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCharacterRequest) ProtoMessage() {}

func (x *UpdateCharacterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_characters_rpc_character_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCharacterRequest.ProtoReflect.Descriptor instead.
func (*UpdateCharacterRequest) Descriptor() ([]byte, []int) {
	return file_services_characters_rpc_character_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCharacterRequest) GetCharacterId() int32 {
	if x != nil {
		return x.CharacterId
	}
	return 0
}

func (x *UpdateCharacterRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateCharacterRequest) GetShortDescription() string {
	if x != nil && x.ShortDescription != nil {
		return *x.ShortDescription
	}
	return ""
}

func (x *UpdateCharacterRequest) GetPublic() bool {
	if x != nil && x.Public != nil {
		return *x.Public
	}
	return false
}

func (x *UpdateCharacterRequest) GetWorldId() int32 {
	if x != nil && x.WorldId != nil {
		return *x.WorldId
	}
	return 0
}

func (x *UpdateCharacterRequest) GetSystemId() int32 {
	if x != nil && x.SystemId != nil {
		return *x.SystemId
	}
	return 0
}

type UploadCharacterImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CharacterId int32  `protobuf:"varint,1,opt,name=characterId,proto3" json:"characterId,omitempty"`
	Data        []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	ImageTypeId int32  `protobuf:"varint,3,opt,name=imageTypeId,proto3" json:"imageTypeId,omitempty"`
}

func (x *UploadCharacterImageRequest) Reset() {
	*x = UploadCharacterImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_characters_rpc_character_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadCharacterImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadCharacterImageRequest) ProtoMessage() {}

func (x *UploadCharacterImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_characters_rpc_character_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadCharacterImageRequest.ProtoReflect.Descriptor instead.
func (*UploadCharacterImageRequest) Descriptor() ([]byte, []int) {
	return file_services_characters_rpc_character_proto_rawDescGZIP(), []int{3}
}

func (x *UploadCharacterImageRequest) GetCharacterId() int32 {
	if x != nil {
		return x.CharacterId
	}
	return 0
}

func (x *UploadCharacterImageRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UploadCharacterImageRequest) GetImageTypeId() int32 {
	if x != nil {
		return x.ImageTypeId
	}
	return 0
}

type GetCharactersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Public   *bool   `protobuf:"varint,1,opt,name=public,proto3,oneof" json:"public,omitempty"`
	Tags     []int32 `protobuf:"varint,2,rep,packed,name=tags,proto3" json:"tags,omitempty"`
	WorldId  *int32  `protobuf:"varint,3,opt,name=worldId,proto3,oneof" json:"worldId,omitempty"`
	SystemId *int32  `protobuf:"varint,4,opt,name=systemId,proto3,oneof" json:"systemId,omitempty"`
	OrderBy  *string `protobuf:"bytes,5,opt,name=orderBy,proto3,oneof" json:"orderBy,omitempty"`
	Limit    *int32  `protobuf:"varint,6,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset   *int32  `protobuf:"varint,7,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *GetCharactersRequest) Reset() {
	*x = GetCharactersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_characters_rpc_character_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCharactersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCharactersRequest) ProtoMessage() {}

func (x *GetCharactersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_characters_rpc_character_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCharactersRequest.ProtoReflect.Descriptor instead.
func (*GetCharactersRequest) Descriptor() ([]byte, []int) {
	return file_services_characters_rpc_character_proto_rawDescGZIP(), []int{4}
}

func (x *GetCharactersRequest) GetPublic() bool {
	if x != nil && x.Public != nil {
		return *x.Public
	}
	return false
}

func (x *GetCharactersRequest) GetTags() []int32 {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetCharactersRequest) GetWorldId() int32 {
	if x != nil && x.WorldId != nil {
		return *x.WorldId
	}
	return 0
}

func (x *GetCharactersRequest) GetSystemId() int32 {
	if x != nil && x.SystemId != nil {
		return *x.SystemId
	}
	return 0
}

func (x *GetCharactersRequest) GetOrderBy() string {
	if x != nil && x.OrderBy != nil {
		return *x.OrderBy
	}
	return ""
}

func (x *GetCharactersRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetCharactersRequest) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

type GetCharactersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CharacterIds []int32 `protobuf:"varint,1,rep,packed,name=characterIds,proto3" json:"characterIds,omitempty"`
	TotalCount   int32   `protobuf:"varint,2,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
}

func (x *GetCharactersResponse) Reset() {
	*x = GetCharactersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_characters_rpc_character_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCharactersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCharactersResponse) ProtoMessage() {}

func (x *GetCharactersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_characters_rpc_character_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCharactersResponse.ProtoReflect.Descriptor instead.
func (*GetCharactersResponse) Descriptor() ([]byte, []int) {
	return file_services_characters_rpc_character_proto_rawDescGZIP(), []int{5}
}

func (x *GetCharactersResponse) GetCharacterIds() []int32 {
	if x != nil {
		return x.CharacterIds
	}
	return nil
}

func (x *GetCharactersResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type GetCharacterByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CharacterId int32 `protobuf:"varint,1,opt,name=characterId,proto3" json:"characterId,omitempty"`
}

func (x *GetCharacterByIdRequest) Reset() {
	*x = GetCharacterByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_characters_rpc_character_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCharacterByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCharacterByIdRequest) ProtoMessage() {}

func (x *GetCharacterByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_characters_rpc_character_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCharacterByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCharacterByIdRequest) Descriptor() ([]byte, []int) {
	return file_services_characters_rpc_character_proto_rawDescGZIP(), []int{6}
}

func (x *GetCharacterByIdRequest) GetCharacterId() int32 {
	if x != nil {
		return x.CharacterId
	}
	return 0
}

var File_services_characters_rpc_character_proto protoreflect.FileDescriptor

var file_services_characters_rpc_character_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x23, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x22, 0x6e, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x06, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x56, 0x69, 0x65, 0x77, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x22, 0xa3, 0x02, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x10, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x06, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x49,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x08, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x13, 0x0a, 0x11, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x75, 0x0a, 0x1b, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x22,
	0xa3, 0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x07, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x08, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x5b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68,
	0x65, 0x2d, 0x6d, 0x65, 0x64, 0x6f, 0x2f, 0x74, 0x61, 0x6c, 0x65, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_services_characters_rpc_character_proto_rawDescOnce sync.Once
	file_services_characters_rpc_character_proto_rawDescData = file_services_characters_rpc_character_proto_rawDesc
)

func file_services_characters_rpc_character_proto_rawDescGZIP() []byte {
	file_services_characters_rpc_character_proto_rawDescOnce.Do(func() {
		file_services_characters_rpc_character_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_characters_rpc_character_proto_rawDescData)
	})
	return file_services_characters_rpc_character_proto_rawDescData
}

var file_services_characters_rpc_character_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_services_characters_rpc_character_proto_goTypes = []interface{}{
	(*CreateCharacterRequest)(nil),      // 0: pb.CreateCharacterRequest
	(*CreateCharacterResponse)(nil),     // 1: pb.CreateCharacterResponse
	(*UpdateCharacterRequest)(nil),      // 2: pb.UpdateCharacterRequest
	(*UploadCharacterImageRequest)(nil), // 3: pb.UploadCharacterImageRequest
	(*GetCharactersRequest)(nil),        // 4: pb.GetCharactersRequest
	(*GetCharactersResponse)(nil),       // 5: pb.GetCharactersResponse
	(*GetCharacterByIdRequest)(nil),     // 6: pb.GetCharacterByIdRequest
	(*Character)(nil),                   // 7: pb.Character
	(*ViewModule)(nil),                  // 8: pb.ViewModule
}
var file_services_characters_rpc_character_proto_depIdxs = []int32{
	7, // 0: pb.CreateCharacterResponse.character:type_name -> pb.Character
	8, // 1: pb.CreateCharacterResponse.module:type_name -> pb.ViewModule
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_services_characters_rpc_character_proto_init() }
func file_services_characters_rpc_character_proto_init() {
	if File_services_characters_rpc_character_proto != nil {
		return
	}
	file_services_characters_character_proto_init()
	file_services_modules_module_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_services_characters_rpc_character_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCharacterRequest); i {
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
		file_services_characters_rpc_character_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCharacterResponse); i {
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
		file_services_characters_rpc_character_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCharacterRequest); i {
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
		file_services_characters_rpc_character_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadCharacterImageRequest); i {
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
		file_services_characters_rpc_character_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCharactersRequest); i {
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
		file_services_characters_rpc_character_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCharactersResponse); i {
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
		file_services_characters_rpc_character_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCharacterByIdRequest); i {
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
	file_services_characters_rpc_character_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_services_characters_rpc_character_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_characters_rpc_character_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_characters_rpc_character_proto_goTypes,
		DependencyIndexes: file_services_characters_rpc_character_proto_depIdxs,
		MessageInfos:      file_services_characters_rpc_character_proto_msgTypes,
	}.Build()
	File_services_characters_rpc_character_proto = out.File
	file_services_characters_rpc_character_proto_rawDesc = nil
	file_services_characters_rpc_character_proto_goTypes = nil
	file_services_characters_rpc_character_proto_depIdxs = nil
}
