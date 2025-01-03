// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: isuxportal/resources/team.proto

package resources

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

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LeaderId           int64               `protobuf:"varint,3,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	MemberIds          []int64             `protobuf:"varint,4,rep,packed,name=member_ids,json=memberIds,proto3" json:"member_ids,omitempty"`
	FinalParticipation bool                `protobuf:"varint,5,opt,name=final_participation,json=finalParticipation,proto3" json:"final_participation,omitempty"`
	Hidden             bool                `protobuf:"varint,6,opt,name=hidden,proto3" json:"hidden,omitempty"`
	Withdrawn          bool                `protobuf:"varint,7,opt,name=withdrawn,proto3" json:"withdrawn,omitempty"`
	Disqualified       bool                `protobuf:"varint,9,opt,name=disqualified,proto3" json:"disqualified,omitempty"`
	Student            *Team_StudentStatus `protobuf:"bytes,10,opt,name=student,proto3" json:"student,omitempty"`
	Detail             *Team_TeamDetail    `protobuf:"bytes,8,opt,name=detail,proto3" json:"detail,omitempty"`
	Leader             *Contestant         `protobuf:"bytes,16,opt,name=leader,proto3" json:"leader,omitempty"`
	Members            []*Contestant       `protobuf:"bytes,17,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	mi := &file_isuxportal_resources_team_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_resources_team_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_isuxportal_resources_team_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetLeaderId() int64 {
	if x != nil {
		return x.LeaderId
	}
	return 0
}

func (x *Team) GetMemberIds() []int64 {
	if x != nil {
		return x.MemberIds
	}
	return nil
}

func (x *Team) GetFinalParticipation() bool {
	if x != nil {
		return x.FinalParticipation
	}
	return false
}

func (x *Team) GetHidden() bool {
	if x != nil {
		return x.Hidden
	}
	return false
}

func (x *Team) GetWithdrawn() bool {
	if x != nil {
		return x.Withdrawn
	}
	return false
}

func (x *Team) GetDisqualified() bool {
	if x != nil {
		return x.Disqualified
	}
	return false
}

func (x *Team) GetStudent() *Team_StudentStatus {
	if x != nil {
		return x.Student
	}
	return nil
}

func (x *Team) GetDetail() *Team_TeamDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *Team) GetLeader() *Contestant {
	if x != nil {
		return x.Leader
	}
	return nil
}

func (x *Team) GetMembers() []*Contestant {
	if x != nil {
		return x.Members
	}
	return nil
}

type Team_StudentStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Team_StudentStatus) Reset() {
	*x = Team_StudentStatus{}
	mi := &file_isuxportal_resources_team_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Team_StudentStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team_StudentStatus) ProtoMessage() {}

func (x *Team_StudentStatus) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_resources_team_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team_StudentStatus.ProtoReflect.Descriptor instead.
func (*Team_StudentStatus) Descriptor() ([]byte, []int) {
	return file_isuxportal_resources_team_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Team_StudentStatus) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type Team_TeamDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailAddress string `protobuf:"bytes,1,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	InviteToken  string `protobuf:"bytes,16,opt,name=invite_token,json=inviteToken,proto3" json:"invite_token,omitempty"`
}

func (x *Team_TeamDetail) Reset() {
	*x = Team_TeamDetail{}
	mi := &file_isuxportal_resources_team_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Team_TeamDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team_TeamDetail) ProtoMessage() {}

func (x *Team_TeamDetail) ProtoReflect() protoreflect.Message {
	mi := &file_isuxportal_resources_team_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team_TeamDetail.ProtoReflect.Descriptor instead.
func (*Team_TeamDetail) Descriptor() ([]byte, []int) {
	return file_isuxportal_resources_team_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Team_TeamDetail) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *Team_TeamDetail) GetInviteToken() string {
	if x != nil {
		return x.InviteToken
	}
	return ""
}

var File_isuxportal_resources_team_proto protoreflect.FileDescriptor

var file_isuxportal_resources_team_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x25, 0x69,
	0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x05, 0x0a, 0x04, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x2f, 0x0a,
	0x13, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x66, 0x69, 0x6e, 0x61,
	0x6c, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x48, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x73, 0x75, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x2e, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x12, 0x43, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x54, 0x65, 0x61, 0x6d, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x3e, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x52,
	0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x73, 0x75, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x27, 0x0a, 0x0d, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x1a, 0x54, 0x0a, 0x0a, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0xf9, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x09, 0x54, 0x65, 0x61,
	0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x73, 0x75, 0x63, 0x6f, 0x6e, 0x2f, 0x69, 0x73, 0x75, 0x63,
	0x6f, 0x6e, 0x31, 0x34, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x72, 0x75, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x49,
	0x50, 0x52, 0xaa, 0x02, 0x1a, 0x49, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca,
	0x02, 0x1a, 0x49, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xe2, 0x02, 0x26, 0x49,
	0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x49, 0x73, 0x75, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_isuxportal_resources_team_proto_rawDescOnce sync.Once
	file_isuxportal_resources_team_proto_rawDescData = file_isuxportal_resources_team_proto_rawDesc
)

func file_isuxportal_resources_team_proto_rawDescGZIP() []byte {
	file_isuxportal_resources_team_proto_rawDescOnce.Do(func() {
		file_isuxportal_resources_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_isuxportal_resources_team_proto_rawDescData)
	})
	return file_isuxportal_resources_team_proto_rawDescData
}

var file_isuxportal_resources_team_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_isuxportal_resources_team_proto_goTypes = []any{
	(*Team)(nil),               // 0: isuxportal.proto.resources.Team
	(*Team_StudentStatus)(nil), // 1: isuxportal.proto.resources.Team.StudentStatus
	(*Team_TeamDetail)(nil),    // 2: isuxportal.proto.resources.Team.TeamDetail
	(*Contestant)(nil),         // 3: isuxportal.proto.resources.Contestant
}
var file_isuxportal_resources_team_proto_depIdxs = []int32{
	1, // 0: isuxportal.proto.resources.Team.student:type_name -> isuxportal.proto.resources.Team.StudentStatus
	2, // 1: isuxportal.proto.resources.Team.detail:type_name -> isuxportal.proto.resources.Team.TeamDetail
	3, // 2: isuxportal.proto.resources.Team.leader:type_name -> isuxportal.proto.resources.Contestant
	3, // 3: isuxportal.proto.resources.Team.members:type_name -> isuxportal.proto.resources.Contestant
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_isuxportal_resources_team_proto_init() }
func file_isuxportal_resources_team_proto_init() {
	if File_isuxportal_resources_team_proto != nil {
		return
	}
	file_isuxportal_resources_contestant_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_isuxportal_resources_team_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_isuxportal_resources_team_proto_goTypes,
		DependencyIndexes: file_isuxportal_resources_team_proto_depIdxs,
		MessageInfos:      file_isuxportal_resources_team_proto_msgTypes,
	}.Build()
	File_isuxportal_resources_team_proto = out.File
	file_isuxportal_resources_team_proto_rawDesc = nil
	file_isuxportal_resources_team_proto_goTypes = nil
	file_isuxportal_resources_team_proto_depIdxs = nil
}
