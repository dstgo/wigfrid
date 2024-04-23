// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: setting.proto

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

type RoomSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Slots       int64  `protobuf:"varint,4,opt,name=slots,proto3" json:"slots,omitempty"`
	Pvp         bool   `protobuf:"varint,5,opt,name=pvp,proto3" json:"pvp,omitempty"`
	Mode        string `protobuf:"bytes,6,opt,name=mode,proto3" json:"mode,omitempty"`
	Paused      bool   `protobuf:"varint,7,opt,name=paused,proto3" json:"paused,omitempty"`
	Vote        bool   `protobuf:"varint,8,opt,name=vote,proto3" json:"vote,omitempty"`
	Kick        bool   `protobuf:"varint,9,opt,name=kick,proto3" json:"kick,omitempty"`
}

func (x *RoomSetting) Reset() {
	*x = RoomSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomSetting) ProtoMessage() {}

func (x *RoomSetting) ProtoReflect() protoreflect.Message {
	mi := &file_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomSetting.ProtoReflect.Descriptor instead.
func (*RoomSetting) Descriptor() ([]byte, []int) {
	return file_setting_proto_rawDescGZIP(), []int{0}
}

func (x *RoomSetting) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoomSetting) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RoomSetting) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RoomSetting) GetSlots() int64 {
	if x != nil {
		return x.Slots
	}
	return 0
}

func (x *RoomSetting) GetPvp() bool {
	if x != nil {
		return x.Pvp
	}
	return false
}

func (x *RoomSetting) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *RoomSetting) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

func (x *RoomSetting) GetVote() bool {
	if x != nil {
		return x.Vote
	}
	return false
}

func (x *RoomSetting) GetKick() bool {
	if x != nil {
		return x.Kick
	}
	return false
}

type RawWorldSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawLuaString string `protobuf:"bytes,1,opt,name=rawLuaString,proto3" json:"rawLuaString,omitempty"`
}

func (x *RawWorldSetting) Reset() {
	*x = RawWorldSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawWorldSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawWorldSetting) ProtoMessage() {}

func (x *RawWorldSetting) ProtoReflect() protoreflect.Message {
	mi := &file_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawWorldSetting.ProtoReflect.Descriptor instead.
func (*RawWorldSetting) Descriptor() ([]byte, []int) {
	return file_setting_proto_rawDescGZIP(), []int{1}
}

func (x *RawWorldSetting) GetRawLuaString() string {
	if x != nil {
		return x.RawLuaString
	}
	return ""
}

type SettingPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SettingPair) Reset() {
	*x = SettingPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingPair) ProtoMessage() {}

func (x *SettingPair) ProtoReflect() protoreflect.Message {
	mi := &file_setting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingPair.ProtoReflect.Descriptor instead.
func (*SettingPair) Descriptor() ([]byte, []int) {
	return file_setting_proto_rawDescGZIP(), []int{2}
}

func (x *SettingPair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SettingPair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type WorldSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Master []*SettingPair `protobuf:"bytes,1,rep,name=master,proto3" json:"master,omitempty"`
	Cave   []*SettingPair `protobuf:"bytes,2,rep,name=cave,proto3" json:"cave,omitempty"`
}

func (x *WorldSetting) Reset() {
	*x = WorldSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldSetting) ProtoMessage() {}

func (x *WorldSetting) ProtoReflect() protoreflect.Message {
	mi := &file_setting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldSetting.ProtoReflect.Descriptor instead.
func (*WorldSetting) Descriptor() ([]byte, []int) {
	return file_setting_proto_rawDescGZIP(), []int{3}
}

func (x *WorldSetting) GetMaster() []*SettingPair {
	if x != nil {
		return x.Master
	}
	return nil
}

func (x *WorldSetting) GetCave() []*SettingPair {
	if x != nil {
		return x.Cave
	}
	return nil
}

var File_setting_proto protoreflect.FileDescriptor

var file_setting_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x76, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x70, 0x76, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x75, 0x73, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f,
	0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x63, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6b, 0x69,
	0x63, 0x6b, 0x22, 0x35, 0x0a, 0x0f, 0x52, 0x61, 0x77, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x61, 0x77, 0x4c, 0x75, 0x61, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x61, 0x77,
	0x4c, 0x75, 0x61, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x35, 0x0a, 0x0b, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x66, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x2c, 0x0a, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x52, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x04, 0x63, 0x61, 0x76, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x04, 0x63, 0x61, 0x76, 0x65, 0x32, 0x86, 0x03, 0x0a, 0x0e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x2e,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x37, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65,
	0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x1a, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x39, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x0e, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x43, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x77, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x13, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x61, 0x77, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x3f, 0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x52, 0x61, 0x77, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x52, 0x61, 0x77, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x1a, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x73, 0x74, 0x67, 0x6f, 0x2f, 0x77, 0x69, 0x67, 0x66, 0x72, 0x69, 0x64, 0x2f, 0x70, 0x62,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_setting_proto_rawDescOnce sync.Once
	file_setting_proto_rawDescData = file_setting_proto_rawDesc
)

func file_setting_proto_rawDescGZIP() []byte {
	file_setting_proto_rawDescOnce.Do(func() {
		file_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_setting_proto_rawDescData)
	})
	return file_setting_proto_rawDescData
}

var file_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_setting_proto_goTypes = []interface{}{
	(*RoomSetting)(nil),     // 0: setting.RoomSetting
	(*RawWorldSetting)(nil), // 1: setting.RawWorldSetting
	(*SettingPair)(nil),     // 2: setting.SettingPair
	(*WorldSetting)(nil),    // 3: setting.WorldSetting
	(*ContainerId)(nil),     // 4: daemon.ContainerId
	(*Notify)(nil),          // 5: result.Notify
}
var file_setting_proto_depIdxs = []int32{
	2, // 0: setting.WorldSetting.master:type_name -> setting.SettingPair
	2, // 1: setting.WorldSetting.cave:type_name -> setting.SettingPair
	4, // 2: setting.SettingService.GetRoomSetting:input_type -> daemon.ContainerId
	0, // 3: setting.SettingService.SaveRoomSetting:input_type -> setting.RoomSetting
	4, // 4: setting.SettingService.GetWorldSetting:input_type -> daemon.ContainerId
	3, // 5: setting.SettingService.SaveWorldSetting:input_type -> setting.WorldSetting
	4, // 6: setting.SettingService.GetRawWorldSetting:input_type -> daemon.ContainerId
	1, // 7: setting.SettingService.SaveRawWorldSetting:input_type -> setting.RawWorldSetting
	0, // 8: setting.SettingService.GetRoomSetting:output_type -> setting.RoomSetting
	5, // 9: setting.SettingService.SaveRoomSetting:output_type -> result.Notify
	3, // 10: setting.SettingService.GetWorldSetting:output_type -> setting.WorldSetting
	5, // 11: setting.SettingService.SaveWorldSetting:output_type -> result.Notify
	1, // 12: setting.SettingService.GetRawWorldSetting:output_type -> setting.RawWorldSetting
	5, // 13: setting.SettingService.SaveRawWorldSetting:output_type -> result.Notify
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_setting_proto_init() }
func file_setting_proto_init() {
	if File_setting_proto != nil {
		return
	}
	file_result_proto_init()
	file_daemon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomSetting); i {
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
		file_setting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawWorldSetting); i {
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
		file_setting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingPair); i {
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
		file_setting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldSetting); i {
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
			RawDescriptor: file_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_setting_proto_goTypes,
		DependencyIndexes: file_setting_proto_depIdxs,
		MessageInfos:      file_setting_proto_msgTypes,
	}.Build()
	File_setting_proto = out.File
	file_setting_proto_rawDesc = nil
	file_setting_proto_goTypes = nil
	file_setting_proto_depIdxs = nil
}
