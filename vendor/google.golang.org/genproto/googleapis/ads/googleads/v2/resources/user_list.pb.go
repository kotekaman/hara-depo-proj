// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/user_list.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A user list. This is a list of users a customer may target.
type UserList struct {
	// Immutable. The resource name of the user list.
	// User list resource names have the form:
	//
	// `customers/{customer_id}/userLists/{user_list_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Id of the user list.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. A flag that indicates if a user may edit a list. Depends on the list
	// ownership and list type. For example, external remarketing user lists are
	// not editable.
	//
	// This field is read-only.
	ReadOnly *wrappers.BoolValue `protobuf:"bytes,3,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	// Name of this user list. Depending on its access_reason, the user list name
	// may not be unique (e.g. if access_reason=SHARED)
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of this user list.
	Description *wrappers.StringValue `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Membership status of this user list. Indicates whether a user list is open
	// or active. Only open user lists can accumulate more users and can be
	// targeted to.
	MembershipStatus enums.UserListMembershipStatusEnum_UserListMembershipStatus `protobuf:"varint,6,opt,name=membership_status,json=membershipStatus,proto3,enum=google.ads.googleads.v2.enums.UserListMembershipStatusEnum_UserListMembershipStatus" json:"membership_status,omitempty"`
	// An ID from external system. It is used by user list sellers to correlate
	// IDs on their systems.
	IntegrationCode *wrappers.StringValue `protobuf:"bytes,7,opt,name=integration_code,json=integrationCode,proto3" json:"integration_code,omitempty"`
	// Number of days a user's cookie stays on your list since its most recent
	// addition to the list. This field must be between 0 and 540 inclusive.
	// However, for CRM based userlists, this field can be set to 10000 which
	// means no expiration.
	//
	// It'll be ignored for logical_user_list.
	MembershipLifeSpan *wrappers.Int64Value `protobuf:"bytes,8,opt,name=membership_life_span,json=membershipLifeSpan,proto3" json:"membership_life_span,omitempty"`
	// Output only. Estimated number of users in this user list, on the Google Display Network.
	// This value is null if the number of users has not yet been determined.
	//
	// This field is read-only.
	SizeForDisplay *wrappers.Int64Value `protobuf:"bytes,9,opt,name=size_for_display,json=sizeForDisplay,proto3" json:"size_for_display,omitempty"`
	// Output only. Size range in terms of number of users of the UserList, on the Google
	// Display Network.
	//
	// This field is read-only.
	SizeRangeForDisplay enums.UserListSizeRangeEnum_UserListSizeRange `protobuf:"varint,10,opt,name=size_range_for_display,json=sizeRangeForDisplay,proto3,enum=google.ads.googleads.v2.enums.UserListSizeRangeEnum_UserListSizeRange" json:"size_range_for_display,omitempty"`
	// Output only. Estimated number of users in this user list in the google.com domain.
	// These are the users available for targeting in Search campaigns.
	// This value is null if the number of users has not yet been determined.
	//
	// This field is read-only.
	SizeForSearch *wrappers.Int64Value `protobuf:"bytes,11,opt,name=size_for_search,json=sizeForSearch,proto3" json:"size_for_search,omitempty"`
	// Output only. Size range in terms of number of users of the UserList, for Search ads.
	//
	// This field is read-only.
	SizeRangeForSearch enums.UserListSizeRangeEnum_UserListSizeRange `protobuf:"varint,12,opt,name=size_range_for_search,json=sizeRangeForSearch,proto3,enum=google.ads.googleads.v2.enums.UserListSizeRangeEnum_UserListSizeRange" json:"size_range_for_search,omitempty"`
	// Output only. Type of this list.
	//
	// This field is read-only.
	Type enums.UserListTypeEnum_UserListType `protobuf:"varint,13,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.UserListTypeEnum_UserListType" json:"type,omitempty"`
	// Indicating the reason why this user list membership status is closed. It is
	// only populated on lists that were automatically closed due to inactivity,
	// and will be cleared once the list membership status becomes open.
	ClosingReason enums.UserListClosingReasonEnum_UserListClosingReason `protobuf:"varint,14,opt,name=closing_reason,json=closingReason,proto3,enum=google.ads.googleads.v2.enums.UserListClosingReasonEnum_UserListClosingReason" json:"closing_reason,omitempty"`
	// Output only. Indicates the reason this account has been granted access to the list.
	// The reason can be SHARED, OWNED, LICENSED or SUBSCRIBED.
	//
	// This field is read-only.
	AccessReason enums.AccessReasonEnum_AccessReason `protobuf:"varint,15,opt,name=access_reason,json=accessReason,proto3,enum=google.ads.googleads.v2.enums.AccessReasonEnum_AccessReason" json:"access_reason,omitempty"`
	// Indicates if this share is still enabled. When a UserList is shared with
	// the user this field is set to ENABLED. Later the userList owner can decide
	// to revoke the share and make it DISABLED.
	// The default value of this field is set to ENABLED.
	AccountUserListStatus enums.UserListAccessStatusEnum_UserListAccessStatus `protobuf:"varint,16,opt,name=account_user_list_status,json=accountUserListStatus,proto3,enum=google.ads.googleads.v2.enums.UserListAccessStatusEnum_UserListAccessStatus" json:"account_user_list_status,omitempty"`
	// Indicates if this user list is eligible for Google Search Network.
	EligibleForSearch *wrappers.BoolValue `protobuf:"bytes,17,opt,name=eligible_for_search,json=eligibleForSearch,proto3" json:"eligible_for_search,omitempty"`
	// Output only. Indicates this user list is eligible for Google Display Network.
	//
	// This field is read-only.
	EligibleForDisplay *wrappers.BoolValue `protobuf:"bytes,18,opt,name=eligible_for_display,json=eligibleForDisplay,proto3" json:"eligible_for_display,omitempty"`
	// The user list.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to UserList:
	//	*UserList_CrmBasedUserList
	//	*UserList_SimilarUserList
	//	*UserList_RuleBasedUserList
	//	*UserList_LogicalUserList
	//	*UserList_BasicUserList
	UserList             isUserList_UserList `protobuf_oneof:"user_list"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e9e9b976e83eeb7, []int{0}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *UserList) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UserList) GetReadOnly() *wrappers.BoolValue {
	if m != nil {
		return m.ReadOnly
	}
	return nil
}

func (m *UserList) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserList) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *UserList) GetMembershipStatus() enums.UserListMembershipStatusEnum_UserListMembershipStatus {
	if m != nil {
		return m.MembershipStatus
	}
	return enums.UserListMembershipStatusEnum_UNSPECIFIED
}

func (m *UserList) GetIntegrationCode() *wrappers.StringValue {
	if m != nil {
		return m.IntegrationCode
	}
	return nil
}

func (m *UserList) GetMembershipLifeSpan() *wrappers.Int64Value {
	if m != nil {
		return m.MembershipLifeSpan
	}
	return nil
}

func (m *UserList) GetSizeForDisplay() *wrappers.Int64Value {
	if m != nil {
		return m.SizeForDisplay
	}
	return nil
}

func (m *UserList) GetSizeRangeForDisplay() enums.UserListSizeRangeEnum_UserListSizeRange {
	if m != nil {
		return m.SizeRangeForDisplay
	}
	return enums.UserListSizeRangeEnum_UNSPECIFIED
}

func (m *UserList) GetSizeForSearch() *wrappers.Int64Value {
	if m != nil {
		return m.SizeForSearch
	}
	return nil
}

func (m *UserList) GetSizeRangeForSearch() enums.UserListSizeRangeEnum_UserListSizeRange {
	if m != nil {
		return m.SizeRangeForSearch
	}
	return enums.UserListSizeRangeEnum_UNSPECIFIED
}

func (m *UserList) GetType() enums.UserListTypeEnum_UserListType {
	if m != nil {
		return m.Type
	}
	return enums.UserListTypeEnum_UNSPECIFIED
}

func (m *UserList) GetClosingReason() enums.UserListClosingReasonEnum_UserListClosingReason {
	if m != nil {
		return m.ClosingReason
	}
	return enums.UserListClosingReasonEnum_UNSPECIFIED
}

func (m *UserList) GetAccessReason() enums.AccessReasonEnum_AccessReason {
	if m != nil {
		return m.AccessReason
	}
	return enums.AccessReasonEnum_UNSPECIFIED
}

func (m *UserList) GetAccountUserListStatus() enums.UserListAccessStatusEnum_UserListAccessStatus {
	if m != nil {
		return m.AccountUserListStatus
	}
	return enums.UserListAccessStatusEnum_UNSPECIFIED
}

func (m *UserList) GetEligibleForSearch() *wrappers.BoolValue {
	if m != nil {
		return m.EligibleForSearch
	}
	return nil
}

func (m *UserList) GetEligibleForDisplay() *wrappers.BoolValue {
	if m != nil {
		return m.EligibleForDisplay
	}
	return nil
}

type isUserList_UserList interface {
	isUserList_UserList()
}

type UserList_CrmBasedUserList struct {
	CrmBasedUserList *common.CrmBasedUserListInfo `protobuf:"bytes,19,opt,name=crm_based_user_list,json=crmBasedUserList,proto3,oneof"`
}

type UserList_SimilarUserList struct {
	SimilarUserList *common.SimilarUserListInfo `protobuf:"bytes,20,opt,name=similar_user_list,json=similarUserList,proto3,oneof"`
}

type UserList_RuleBasedUserList struct {
	RuleBasedUserList *common.RuleBasedUserListInfo `protobuf:"bytes,21,opt,name=rule_based_user_list,json=ruleBasedUserList,proto3,oneof"`
}

type UserList_LogicalUserList struct {
	LogicalUserList *common.LogicalUserListInfo `protobuf:"bytes,22,opt,name=logical_user_list,json=logicalUserList,proto3,oneof"`
}

type UserList_BasicUserList struct {
	BasicUserList *common.BasicUserListInfo `protobuf:"bytes,23,opt,name=basic_user_list,json=basicUserList,proto3,oneof"`
}

func (*UserList_CrmBasedUserList) isUserList_UserList() {}

func (*UserList_SimilarUserList) isUserList_UserList() {}

func (*UserList_RuleBasedUserList) isUserList_UserList() {}

func (*UserList_LogicalUserList) isUserList_UserList() {}

func (*UserList_BasicUserList) isUserList_UserList() {}

func (m *UserList) GetUserList() isUserList_UserList {
	if m != nil {
		return m.UserList
	}
	return nil
}

func (m *UserList) GetCrmBasedUserList() *common.CrmBasedUserListInfo {
	if x, ok := m.GetUserList().(*UserList_CrmBasedUserList); ok {
		return x.CrmBasedUserList
	}
	return nil
}

func (m *UserList) GetSimilarUserList() *common.SimilarUserListInfo {
	if x, ok := m.GetUserList().(*UserList_SimilarUserList); ok {
		return x.SimilarUserList
	}
	return nil
}

func (m *UserList) GetRuleBasedUserList() *common.RuleBasedUserListInfo {
	if x, ok := m.GetUserList().(*UserList_RuleBasedUserList); ok {
		return x.RuleBasedUserList
	}
	return nil
}

func (m *UserList) GetLogicalUserList() *common.LogicalUserListInfo {
	if x, ok := m.GetUserList().(*UserList_LogicalUserList); ok {
		return x.LogicalUserList
	}
	return nil
}

func (m *UserList) GetBasicUserList() *common.BasicUserListInfo {
	if x, ok := m.GetUserList().(*UserList_BasicUserList); ok {
		return x.BasicUserList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UserList) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UserList_CrmBasedUserList)(nil),
		(*UserList_SimilarUserList)(nil),
		(*UserList_RuleBasedUserList)(nil),
		(*UserList_LogicalUserList)(nil),
		(*UserList_BasicUserList)(nil),
	}
}

func init() {
	proto.RegisterType((*UserList)(nil), "google.ads.googleads.v2.resources.UserList")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/user_list.proto", fileDescriptor_4e9e9b976e83eeb7)
}

var fileDescriptor_4e9e9b976e83eeb7 = []byte{
	// 1020 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x4f, 0x6f, 0xdb, 0x36,
	0x18, 0xc6, 0x6b, 0x27, 0xed, 0x12, 0x26, 0x8e, 0x63, 0x26, 0xe9, 0xb4, 0xac, 0xd8, 0x92, 0x0d,
	0x05, 0xb2, 0x1d, 0xa4, 0xd6, 0xed, 0x86, 0xc1, 0x0d, 0x36, 0xc8, 0xd9, 0x9a, 0x34, 0x48, 0xd3,
	0x4e, 0x6e, 0x73, 0xd8, 0x02, 0x08, 0xb4, 0x44, 0x2b, 0x04, 0x24, 0x52, 0x20, 0xa5, 0x0c, 0x6e,
	0xd1, 0x61, 0x3b, 0xec, 0x8b, 0xec, 0x34, 0xec, 0xa3, 0xec, 0x53, 0xf4, 0xdc, 0x8f, 0xd0, 0xd3,
	0x20, 0x4a, 0x94, 0x29, 0xa7, 0xae, 0x75, 0xe8, 0x8d, 0xe2, 0xfb, 0xfe, 0xde, 0x87, 0x0f, 0xff,
	0x89, 0xe0, 0x6e, 0xc0, 0x58, 0x10, 0x62, 0x0b, 0xf9, 0xc2, 0xca, 0x9b, 0x59, 0xeb, 0xb2, 0x6b,
	0x71, 0x2c, 0x58, 0xca, 0x3d, 0x2c, 0xac, 0x54, 0x60, 0xee, 0x86, 0x44, 0x24, 0x66, 0xcc, 0x59,
	0xc2, 0xe0, 0x6e, 0x9e, 0x67, 0x22, 0x5f, 0x98, 0x25, 0x62, 0x5e, 0x76, 0xcd, 0x12, 0xd9, 0xb6,
	0x66, 0x55, 0xf5, 0x58, 0x14, 0x31, 0x3a, 0x29, 0x29, 0xf2, 0x9a, 0xdb, 0x33, 0x87, 0x81, 0x69,
	0x1a, 0x09, 0x0b, 0x79, 0x1e, 0x16, 0xc2, 0xe5, 0x18, 0x09, 0x46, 0x0b, 0xe4, 0xc1, 0xfb, 0x91,
	0x52, 0xc2, 0x2d, 0x60, 0x91, 0xa0, 0x24, 0x55, 0x7a, 0xfb, 0x75, 0x61, 0x2f, 0x64, 0x82, 0xd0,
	0xa0, 0x2a, 0xfd, 0x43, 0x5d, 0x3a, 0xc2, 0xd1, 0x10, 0x73, 0x71, 0x41, 0xe2, 0xaa, 0xfc, 0x77,
	0x75, 0x0b, 0x08, 0xf2, 0x02, 0xbb, 0x1c, 0xd1, 0x00, 0x17, 0x64, 0xb7, 0x2e, 0x99, 0x8c, 0x63,
	0xc5, 0x7c, 0xae, 0x98, 0x98, 0x58, 0x23, 0x82, 0x43, 0xdf, 0x1d, 0xe2, 0x0b, 0x74, 0x49, 0x18,
	0x2f, 0x12, 0x3e, 0xd1, 0x12, 0xd4, 0x22, 0x16, 0xa1, 0xcf, 0x8a, 0x90, 0xfc, 0x1a, 0xa6, 0x23,
	0xeb, 0x37, 0x8e, 0xe2, 0x18, 0x73, 0xe5, 0xe4, 0x96, 0x86, 0x22, 0x4a, 0x59, 0x82, 0x12, 0xc2,
	0x68, 0x11, 0xfd, 0xe2, 0x9f, 0x0e, 0x58, 0x7a, 0x2e, 0x30, 0x3f, 0x21, 0x22, 0x81, 0xa7, 0xa0,
	0xa5, 0x8a, 0xbb, 0x14, 0x45, 0xd8, 0x68, 0xec, 0x34, 0xf6, 0x96, 0xfb, 0x5f, 0xbd, 0xb6, 0xaf,
	0xbf, 0xb5, 0xbf, 0x04, 0xbb, 0x93, 0xbd, 0x54, 0xb4, 0x62, 0x22, 0x4c, 0x8f, 0x45, 0x96, 0xaa,
	0xe0, 0xac, 0x2a, 0xfe, 0x14, 0x45, 0x18, 0xde, 0x01, 0x4d, 0xe2, 0x1b, 0xcd, 0x9d, 0xc6, 0xde,
	0x4a, 0xf7, 0xd3, 0x82, 0x31, 0xd5, 0x38, 0xcd, 0x47, 0x34, 0xf9, 0xf6, 0xfe, 0x19, 0x0a, 0x53,
	0xdc, 0x5f, 0x78, 0x6d, 0x2f, 0x38, 0x4d, 0xe2, 0xc3, 0x7d, 0xb0, 0xcc, 0x31, 0xf2, 0x5d, 0x46,
	0xc3, 0xb1, 0xb1, 0x20, 0xc1, 0xed, 0x2b, 0x60, 0x9f, 0xb1, 0x50, 0xe3, 0x96, 0x32, 0xe2, 0x09,
	0x0d, 0xc7, 0xf0, 0x0e, 0x58, 0x94, 0xc3, 0x5e, 0x94, 0xe0, 0xad, 0x2b, 0xe0, 0x20, 0xe1, 0x84,
	0x06, 0x12, 0x75, 0x64, 0x26, 0xfc, 0x1e, 0xac, 0xf8, 0x58, 0x78, 0x9c, 0xc4, 0xd9, 0xa4, 0x18,
	0xd7, 0x6b, 0x80, 0x3a, 0x00, 0xff, 0x6c, 0x80, 0xce, 0x95, 0x2d, 0x64, 0xdc, 0xd8, 0x69, 0xec,
	0xad, 0x75, 0x9f, 0x99, 0xb3, 0x8e, 0xa1, 0xdc, 0x09, 0xa6, 0x9a, 0xb4, 0xc7, 0x25, 0x3f, 0x90,
	0xf8, 0x4f, 0x34, 0x8d, 0x66, 0x06, 0x9d, 0xf5, 0x68, 0xaa, 0x07, 0x1e, 0x82, 0x75, 0x42, 0x13,
	0x1c, 0x70, 0xb9, 0xb0, 0xae, 0xc7, 0x7c, 0x6c, 0x7c, 0x54, 0xc3, 0x48, 0x5b, 0xa3, 0x0e, 0x98,
	0x8f, 0xe1, 0x63, 0xb0, 0xa9, 0x79, 0x09, 0xc9, 0x08, 0xbb, 0x22, 0x46, 0xd4, 0x58, 0x9a, 0xbb,
	0x80, 0x0e, 0x9c, 0x80, 0x27, 0x64, 0x84, 0x07, 0x31, 0xa2, 0xf0, 0x18, 0xac, 0xcb, 0xc3, 0x31,
	0x62, 0xdc, 0xf5, 0x89, 0x88, 0x43, 0x34, 0x36, 0x96, 0x6b, 0xee, 0x85, 0xb5, 0x8c, 0x7c, 0xc8,
	0xf8, 0x8f, 0x39, 0x07, 0xff, 0x68, 0x80, 0x9b, 0x93, 0x93, 0x56, 0x29, 0x09, 0xe4, 0x64, 0x3f,
	0xac, 0x39, 0xd9, 0x03, 0xf2, 0x02, 0x3b, 0x59, 0x8d, 0xca, 0x2c, 0x97, 0xbd, 0xb9, 0xfa, 0x86,
	0x50, 0xdf, 0xda, 0x10, 0x8e, 0x40, 0xbb, 0xb4, 0x23, 0x30, 0xe2, 0xde, 0x85, 0xb1, 0x52, 0xd3,
	0x4d, 0xab, 0x70, 0x33, 0x90, 0x18, 0xfc, 0x1d, 0x6c, 0x4d, 0x79, 0x29, 0xea, 0xad, 0x7e, 0x78,
	0x2b, 0x50, 0xb7, 0x52, 0xe8, 0x3f, 0x07, 0x8b, 0xd9, 0xdd, 0x63, 0xb4, 0xa4, 0xdc, 0x7e, 0x4d,
	0xb9, 0x67, 0xe3, 0xb8, 0xaa, 0x94, 0x75, 0xe4, 0x22, 0xb2, 0x1c, 0x4c, 0xc1, 0x5a, 0xf5, 0x2e,
	0x36, 0xd6, 0xa4, 0xc0, 0x69, 0x4d, 0x81, 0x83, 0x1c, 0x76, 0x24, 0x5b, 0x51, 0xaa, 0x44, 0x9c,
	0x96, 0xa7, 0x7f, 0xc2, 0x11, 0x68, 0x55, 0x7e, 0x3e, 0x46, 0xbb, 0x96, 0x2d, 0x5b, 0x32, 0x9a,
	0x98, 0xde, 0x91, 0xdb, 0x5a, 0x45, 0x5a, 0x17, 0xfc, 0xab, 0x01, 0x0c, 0xe4, 0x79, 0x2c, 0xa5,
	0x89, 0xab, 0x5d, 0xff, 0xf9, 0x89, 0x5f, 0x97, 0x9a, 0x27, 0x35, 0x9d, 0xe6, 0x52, 0xef, 0x38,
	0xed, 0x7a, 0xc0, 0xd9, 0x2a, 0xd4, 0xca, 0x95, 0xcd, 0x8f, 0xfb, 0x31, 0xd8, 0xc0, 0x21, 0x09,
	0xc8, 0x30, 0xac, 0xec, 0x9d, 0xce, 0xbc, 0xcb, 0xd2, 0xe9, 0x28, 0x6c, 0xb2, 0x13, 0x7e, 0x06,
	0x9b, 0x95, 0x5a, 0xea, 0x4c, 0xc1, 0x7a, 0x37, 0x2f, 0xd4, 0x2a, 0xaa, 0x63, 0x82, 0xc1, 0x86,
	0xc7, 0x23, 0x77, 0x88, 0x04, 0xf6, 0x27, 0xf3, 0x64, 0x6c, 0xc8, 0x8a, 0xf7, 0x67, 0x4e, 0x50,
	0xfe, 0xec, 0x30, 0x0f, 0x78, 0xd4, 0xcf, 0x48, 0xe5, 0xf9, 0x11, 0x1d, 0xb1, 0xa3, 0x6b, 0xce,
	0xba, 0x37, 0xd5, 0x0f, 0x47, 0xa0, 0x23, 0x48, 0x44, 0x42, 0xc4, 0x35, 0x91, 0x4d, 0x29, 0x72,
	0x6f, 0x9e, 0xc8, 0x20, 0x07, 0x75, 0x0d, 0xe9, 0xe7, 0xe8, 0x9a, 0xd3, 0x16, 0xd5, 0x18, 0xbc,
	0x00, 0x9b, 0x3c, 0x0d, 0xf1, 0x15, 0x3f, 0x5b, 0x52, 0xea, 0x9b, 0x79, 0x52, 0x4e, 0x1a, 0xe2,
	0x77, 0x19, 0xea, 0xf0, 0xe9, 0x00, 0x44, 0xa0, 0x13, 0xb2, 0x80, 0x78, 0x28, 0xd4, 0x64, 0x6e,
	0xd6, 0x73, 0x74, 0x92, 0x83, 0x53, 0x22, 0xed, 0xb0, 0xda, 0x0d, 0x7f, 0x05, 0xed, 0x21, 0x12,
	0xc4, 0xd3, 0x04, 0x3e, 0x96, 0x02, 0x77, 0xe7, 0x09, 0xf4, 0x33, 0x6c, 0xaa, 0x7c, 0x6b, 0xa8,
	0x77, 0xf6, 0x9c, 0x37, 0xf6, 0x93, 0x1a, 0x4f, 0x04, 0xf8, 0xb5, 0x97, 0x8a, 0x84, 0x45, 0x98,
	0x0b, 0xeb, 0xa5, 0x6a, 0xbe, 0x92, 0xcf, 0xa2, 0x2c, 0x2c, 0xac, 0x97, 0xe5, 0xe0, 0x5e, 0xf5,
	0x57, 0xc0, 0x72, 0xf9, 0xd5, 0x7f, 0xdb, 0x00, 0xb7, 0x3d, 0x16, 0x99, 0x73, 0x1f, 0xb7, 0xfd,
	0x96, 0x12, 0x7b, 0x9a, 0x6d, 0xdc, 0xa7, 0x8d, 0x5f, 0x8e, 0x0b, 0x26, 0x60, 0x21, 0xa2, 0x81,
	0xc9, 0x78, 0x60, 0x05, 0x98, 0xca, 0x6d, 0x6d, 0x4d, 0xc6, 0xf9, 0x9e, 0x27, 0xf6, 0x83, 0xb2,
	0xf5, 0x77, 0x73, 0xe1, 0xd0, 0xb6, 0xff, 0x6d, 0xee, 0x1e, 0xe6, 0x25, 0x6d, 0x5f, 0x98, 0x79,
	0x33, 0x6b, 0x9d, 0x75, 0x4d, 0x47, 0x65, 0xfe, 0xa7, 0x72, 0xce, 0x6d, 0x5f, 0x9c, 0x97, 0x39,
	0xe7, 0x67, 0xdd, 0xf3, 0x32, 0xe7, 0x4d, 0xf3, 0x76, 0x1e, 0xe8, 0xf5, 0x6c, 0x5f, 0xf4, 0x7a,
	0x65, 0x56, 0xaf, 0x77, 0xd6, 0xed, 0xf5, 0xca, 0xbc, 0xe1, 0x0d, 0x39, 0xd8, 0x7b, 0xff, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x93, 0x6e, 0x15, 0xbf, 0x0e, 0x0c, 0x00, 0x00,
}