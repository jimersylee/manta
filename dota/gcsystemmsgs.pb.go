// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gcsystemmsgs.proto

package dota

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EGCSystemMsg int32

const (
	EGCSystemMsg_k_EGCMsgInvalid                           EGCSystemMsg = 0
	EGCSystemMsg_k_EGCMsgMulti                             EGCSystemMsg = 1
	EGCSystemMsg_k_EGCMsgGenericReply                      EGCSystemMsg = 10
	EGCSystemMsg_k_EGCMsgSystemBase                        EGCSystemMsg = 50
	EGCSystemMsg_k_EGCMsgAchievementAwarded                EGCSystemMsg = 51
	EGCSystemMsg_k_EGCMsgConCommand                        EGCSystemMsg = 52
	EGCSystemMsg_k_EGCMsgStartPlaying                      EGCSystemMsg = 53
	EGCSystemMsg_k_EGCMsgStopPlaying                       EGCSystemMsg = 54
	EGCSystemMsg_k_EGCMsgStartGameserver                   EGCSystemMsg = 55
	EGCSystemMsg_k_EGCMsgStopGameserver                    EGCSystemMsg = 56
	EGCSystemMsg_k_EGCMsgWGRequest                         EGCSystemMsg = 57
	EGCSystemMsg_k_EGCMsgWGResponse                        EGCSystemMsg = 58
	EGCSystemMsg_k_EGCMsgGetUserGameStatsSchema            EGCSystemMsg = 59
	EGCSystemMsg_k_EGCMsgGetUserGameStatsSchemaResponse    EGCSystemMsg = 60
	EGCSystemMsg_k_EGCMsgGetUserStatsDEPRECATED            EGCSystemMsg = 61
	EGCSystemMsg_k_EGCMsgGetUserStatsResponse              EGCSystemMsg = 62
	EGCSystemMsg_k_EGCMsgAppInfoUpdated                    EGCSystemMsg = 63
	EGCSystemMsg_k_EGCMsgValidateSession                   EGCSystemMsg = 64
	EGCSystemMsg_k_EGCMsgValidateSessionResponse           EGCSystemMsg = 65
	EGCSystemMsg_k_EGCMsgLookupAccountFromInput            EGCSystemMsg = 66
	EGCSystemMsg_k_EGCMsgSendHTTPRequest                   EGCSystemMsg = 67
	EGCSystemMsg_k_EGCMsgSendHTTPRequestResponse           EGCSystemMsg = 68
	EGCSystemMsg_k_EGCMsgPreTestSetup                      EGCSystemMsg = 69
	EGCSystemMsg_k_EGCMsgRecordSupportAction               EGCSystemMsg = 70
	EGCSystemMsg_k_EGCMsgGetAccountDetails_DEPRECATED      EGCSystemMsg = 71
	EGCSystemMsg_k_EGCMsgReceiveInterAppMessage            EGCSystemMsg = 73
	EGCSystemMsg_k_EGCMsgFindAccounts                      EGCSystemMsg = 74
	EGCSystemMsg_k_EGCMsgPostAlert                         EGCSystemMsg = 75
	EGCSystemMsg_k_EGCMsgGetLicenses                       EGCSystemMsg = 76
	EGCSystemMsg_k_EGCMsgGetUserStats                      EGCSystemMsg = 77
	EGCSystemMsg_k_EGCMsgGetCommands                       EGCSystemMsg = 78
	EGCSystemMsg_k_EGCMsgGetCommandsResponse               EGCSystemMsg = 79
	EGCSystemMsg_k_EGCMsgAddFreeLicense                    EGCSystemMsg = 80
	EGCSystemMsg_k_EGCMsgAddFreeLicenseResponse            EGCSystemMsg = 81
	EGCSystemMsg_k_EGCMsgGetIPLocation                     EGCSystemMsg = 82
	EGCSystemMsg_k_EGCMsgGetIPLocationResponse             EGCSystemMsg = 83
	EGCSystemMsg_k_EGCMsgSystemStatsSchema                 EGCSystemMsg = 84
	EGCSystemMsg_k_EGCMsgGetSystemStats                    EGCSystemMsg = 85
	EGCSystemMsg_k_EGCMsgGetSystemStatsResponse            EGCSystemMsg = 86
	EGCSystemMsg_k_EGCMsgSendEmail                         EGCSystemMsg = 87
	EGCSystemMsg_k_EGCMsgSendEmailResponse                 EGCSystemMsg = 88
	EGCSystemMsg_k_EGCMsgGetEmailTemplate                  EGCSystemMsg = 89
	EGCSystemMsg_k_EGCMsgGetEmailTemplateResponse          EGCSystemMsg = 90
	EGCSystemMsg_k_EGCMsgGrantGuestPass                    EGCSystemMsg = 91
	EGCSystemMsg_k_EGCMsgGrantGuestPassResponse            EGCSystemMsg = 92
	EGCSystemMsg_k_EGCMsgGetAccountDetails                 EGCSystemMsg = 93
	EGCSystemMsg_k_EGCMsgGetAccountDetailsResponse         EGCSystemMsg = 94
	EGCSystemMsg_k_EGCMsgGetPersonaNames                   EGCSystemMsg = 95
	EGCSystemMsg_k_EGCMsgGetPersonaNamesResponse           EGCSystemMsg = 96
	EGCSystemMsg_k_EGCMsgMultiplexMsg                      EGCSystemMsg = 97
	EGCSystemMsg_k_EGCMsgWebAPIRegisterInterfaces          EGCSystemMsg = 101
	EGCSystemMsg_k_EGCMsgWebAPIJobRequest                  EGCSystemMsg = 102
	EGCSystemMsg_k_EGCMsgWebAPIJobRequestHttpResponse      EGCSystemMsg = 104
	EGCSystemMsg_k_EGCMsgWebAPIJobRequestForwardResponse   EGCSystemMsg = 105
	EGCSystemMsg_k_EGCMsgMemCachedGet                      EGCSystemMsg = 200
	EGCSystemMsg_k_EGCMsgMemCachedGetResponse              EGCSystemMsg = 201
	EGCSystemMsg_k_EGCMsgMemCachedSet                      EGCSystemMsg = 202
	EGCSystemMsg_k_EGCMsgMemCachedDelete                   EGCSystemMsg = 203
	EGCSystemMsg_k_EGCMsgMemCachedStats                    EGCSystemMsg = 204
	EGCSystemMsg_k_EGCMsgMemCachedStatsResponse            EGCSystemMsg = 205
	EGCSystemMsg_k_EGCMsgSQLStats                          EGCSystemMsg = 210
	EGCSystemMsg_k_EGCMsgSQLStatsResponse                  EGCSystemMsg = 211
	EGCSystemMsg_k_EGCMsgMasterSetDirectory                EGCSystemMsg = 220
	EGCSystemMsg_k_EGCMsgMasterSetDirectoryResponse        EGCSystemMsg = 221
	EGCSystemMsg_k_EGCMsgMasterSetWebAPIRouting            EGCSystemMsg = 222
	EGCSystemMsg_k_EGCMsgMasterSetWebAPIRoutingResponse    EGCSystemMsg = 223
	EGCSystemMsg_k_EGCMsgMasterSetClientMsgRouting         EGCSystemMsg = 224
	EGCSystemMsg_k_EGCMsgMasterSetClientMsgRoutingResponse EGCSystemMsg = 225
	EGCSystemMsg_k_EGCMsgSetOptions                        EGCSystemMsg = 226
	EGCSystemMsg_k_EGCMsgSetOptionsResponse                EGCSystemMsg = 227
	EGCSystemMsg_k_EGCMsgSystemBase2                       EGCSystemMsg = 500
	EGCSystemMsg_k_EGCMsgGetPurchaseTrustStatus            EGCSystemMsg = 501
	EGCSystemMsg_k_EGCMsgGetPurchaseTrustStatusResponse    EGCSystemMsg = 502
	EGCSystemMsg_k_EGCMsgUpdateSession                     EGCSystemMsg = 503
	EGCSystemMsg_k_EGCMsgGCAccountVacStatusChange          EGCSystemMsg = 504
	EGCSystemMsg_k_EGCMsgCheckFriendship                   EGCSystemMsg = 505
	EGCSystemMsg_k_EGCMsgCheckFriendshipResponse           EGCSystemMsg = 506
	EGCSystemMsg_k_EGCMsgGetPartnerAccountLink             EGCSystemMsg = 507
	EGCSystemMsg_k_EGCMsgGetPartnerAccountLinkResponse     EGCSystemMsg = 508
	EGCSystemMsg_k_EGCMsgVSReportedSuspiciousActivity      EGCSystemMsg = 509
	EGCSystemMsg_k_EGCMsgDPPartnerMicroTxns                EGCSystemMsg = 512
	EGCSystemMsg_k_EGCMsgDPPartnerMicroTxnsResponse        EGCSystemMsg = 513
	EGCSystemMsg_k_EGCMsgGetIPASN                          EGCSystemMsg = 514
	EGCSystemMsg_k_EGCMsgGetIPASNResponse                  EGCSystemMsg = 515
	EGCSystemMsg_k_EGCMsgGetAppFriendsList                 EGCSystemMsg = 516
	EGCSystemMsg_k_EGCMsgGetAppFriendsListResponse         EGCSystemMsg = 517
	EGCSystemMsg_k_EGCMsgVacVerificationChange             EGCSystemMsg = 518
	EGCSystemMsg_k_EGCMsgAccountPhoneNumberChange          EGCSystemMsg = 519
	EGCSystemMsg_k_EGCMsgAccountTwoFactorChange            EGCSystemMsg = 520
	EGCSystemMsg_k_EGCMsgCheckClanMembership               EGCSystemMsg = 521
	EGCSystemMsg_k_EGCMsgCheckClanMembershipResponse       EGCSystemMsg = 522
)

var EGCSystemMsg_name = map[int32]string{
	0:   "k_EGCMsgInvalid",
	1:   "k_EGCMsgMulti",
	10:  "k_EGCMsgGenericReply",
	50:  "k_EGCMsgSystemBase",
	51:  "k_EGCMsgAchievementAwarded",
	52:  "k_EGCMsgConCommand",
	53:  "k_EGCMsgStartPlaying",
	54:  "k_EGCMsgStopPlaying",
	55:  "k_EGCMsgStartGameserver",
	56:  "k_EGCMsgStopGameserver",
	57:  "k_EGCMsgWGRequest",
	58:  "k_EGCMsgWGResponse",
	59:  "k_EGCMsgGetUserGameStatsSchema",
	60:  "k_EGCMsgGetUserGameStatsSchemaResponse",
	61:  "k_EGCMsgGetUserStatsDEPRECATED",
	62:  "k_EGCMsgGetUserStatsResponse",
	63:  "k_EGCMsgAppInfoUpdated",
	64:  "k_EGCMsgValidateSession",
	65:  "k_EGCMsgValidateSessionResponse",
	66:  "k_EGCMsgLookupAccountFromInput",
	67:  "k_EGCMsgSendHTTPRequest",
	68:  "k_EGCMsgSendHTTPRequestResponse",
	69:  "k_EGCMsgPreTestSetup",
	70:  "k_EGCMsgRecordSupportAction",
	71:  "k_EGCMsgGetAccountDetails_DEPRECATED",
	73:  "k_EGCMsgReceiveInterAppMessage",
	74:  "k_EGCMsgFindAccounts",
	75:  "k_EGCMsgPostAlert",
	76:  "k_EGCMsgGetLicenses",
	77:  "k_EGCMsgGetUserStats",
	78:  "k_EGCMsgGetCommands",
	79:  "k_EGCMsgGetCommandsResponse",
	80:  "k_EGCMsgAddFreeLicense",
	81:  "k_EGCMsgAddFreeLicenseResponse",
	82:  "k_EGCMsgGetIPLocation",
	83:  "k_EGCMsgGetIPLocationResponse",
	84:  "k_EGCMsgSystemStatsSchema",
	85:  "k_EGCMsgGetSystemStats",
	86:  "k_EGCMsgGetSystemStatsResponse",
	87:  "k_EGCMsgSendEmail",
	88:  "k_EGCMsgSendEmailResponse",
	89:  "k_EGCMsgGetEmailTemplate",
	90:  "k_EGCMsgGetEmailTemplateResponse",
	91:  "k_EGCMsgGrantGuestPass",
	92:  "k_EGCMsgGrantGuestPassResponse",
	93:  "k_EGCMsgGetAccountDetails",
	94:  "k_EGCMsgGetAccountDetailsResponse",
	95:  "k_EGCMsgGetPersonaNames",
	96:  "k_EGCMsgGetPersonaNamesResponse",
	97:  "k_EGCMsgMultiplexMsg",
	101: "k_EGCMsgWebAPIRegisterInterfaces",
	102: "k_EGCMsgWebAPIJobRequest",
	104: "k_EGCMsgWebAPIJobRequestHttpResponse",
	105: "k_EGCMsgWebAPIJobRequestForwardResponse",
	200: "k_EGCMsgMemCachedGet",
	201: "k_EGCMsgMemCachedGetResponse",
	202: "k_EGCMsgMemCachedSet",
	203: "k_EGCMsgMemCachedDelete",
	204: "k_EGCMsgMemCachedStats",
	205: "k_EGCMsgMemCachedStatsResponse",
	210: "k_EGCMsgSQLStats",
	211: "k_EGCMsgSQLStatsResponse",
	220: "k_EGCMsgMasterSetDirectory",
	221: "k_EGCMsgMasterSetDirectoryResponse",
	222: "k_EGCMsgMasterSetWebAPIRouting",
	223: "k_EGCMsgMasterSetWebAPIRoutingResponse",
	224: "k_EGCMsgMasterSetClientMsgRouting",
	225: "k_EGCMsgMasterSetClientMsgRoutingResponse",
	226: "k_EGCMsgSetOptions",
	227: "k_EGCMsgSetOptionsResponse",
	500: "k_EGCMsgSystemBase2",
	501: "k_EGCMsgGetPurchaseTrustStatus",
	502: "k_EGCMsgGetPurchaseTrustStatusResponse",
	503: "k_EGCMsgUpdateSession",
	504: "k_EGCMsgGCAccountVacStatusChange",
	505: "k_EGCMsgCheckFriendship",
	506: "k_EGCMsgCheckFriendshipResponse",
	507: "k_EGCMsgGetPartnerAccountLink",
	508: "k_EGCMsgGetPartnerAccountLinkResponse",
	509: "k_EGCMsgVSReportedSuspiciousActivity",
	512: "k_EGCMsgDPPartnerMicroTxns",
	513: "k_EGCMsgDPPartnerMicroTxnsResponse",
	514: "k_EGCMsgGetIPASN",
	515: "k_EGCMsgGetIPASNResponse",
	516: "k_EGCMsgGetAppFriendsList",
	517: "k_EGCMsgGetAppFriendsListResponse",
	518: "k_EGCMsgVacVerificationChange",
	519: "k_EGCMsgAccountPhoneNumberChange",
	520: "k_EGCMsgAccountTwoFactorChange",
	521: "k_EGCMsgCheckClanMembership",
	522: "k_EGCMsgCheckClanMembershipResponse",
}
var EGCSystemMsg_value = map[string]int32{
	"k_EGCMsgInvalid":                           0,
	"k_EGCMsgMulti":                             1,
	"k_EGCMsgGenericReply":                      10,
	"k_EGCMsgSystemBase":                        50,
	"k_EGCMsgAchievementAwarded":                51,
	"k_EGCMsgConCommand":                        52,
	"k_EGCMsgStartPlaying":                      53,
	"k_EGCMsgStopPlaying":                       54,
	"k_EGCMsgStartGameserver":                   55,
	"k_EGCMsgStopGameserver":                    56,
	"k_EGCMsgWGRequest":                         57,
	"k_EGCMsgWGResponse":                        58,
	"k_EGCMsgGetUserGameStatsSchema":            59,
	"k_EGCMsgGetUserGameStatsSchemaResponse":    60,
	"k_EGCMsgGetUserStatsDEPRECATED":            61,
	"k_EGCMsgGetUserStatsResponse":              62,
	"k_EGCMsgAppInfoUpdated":                    63,
	"k_EGCMsgValidateSession":                   64,
	"k_EGCMsgValidateSessionResponse":           65,
	"k_EGCMsgLookupAccountFromInput":            66,
	"k_EGCMsgSendHTTPRequest":                   67,
	"k_EGCMsgSendHTTPRequestResponse":           68,
	"k_EGCMsgPreTestSetup":                      69,
	"k_EGCMsgRecordSupportAction":               70,
	"k_EGCMsgGetAccountDetails_DEPRECATED":      71,
	"k_EGCMsgReceiveInterAppMessage":            73,
	"k_EGCMsgFindAccounts":                      74,
	"k_EGCMsgPostAlert":                         75,
	"k_EGCMsgGetLicenses":                       76,
	"k_EGCMsgGetUserStats":                      77,
	"k_EGCMsgGetCommands":                       78,
	"k_EGCMsgGetCommandsResponse":               79,
	"k_EGCMsgAddFreeLicense":                    80,
	"k_EGCMsgAddFreeLicenseResponse":            81,
	"k_EGCMsgGetIPLocation":                     82,
	"k_EGCMsgGetIPLocationResponse":             83,
	"k_EGCMsgSystemStatsSchema":                 84,
	"k_EGCMsgGetSystemStats":                    85,
	"k_EGCMsgGetSystemStatsResponse":            86,
	"k_EGCMsgSendEmail":                         87,
	"k_EGCMsgSendEmailResponse":                 88,
	"k_EGCMsgGetEmailTemplate":                  89,
	"k_EGCMsgGetEmailTemplateResponse":          90,
	"k_EGCMsgGrantGuestPass":                    91,
	"k_EGCMsgGrantGuestPassResponse":            92,
	"k_EGCMsgGetAccountDetails":                 93,
	"k_EGCMsgGetAccountDetailsResponse":         94,
	"k_EGCMsgGetPersonaNames":                   95,
	"k_EGCMsgGetPersonaNamesResponse":           96,
	"k_EGCMsgMultiplexMsg":                      97,
	"k_EGCMsgWebAPIRegisterInterfaces":          101,
	"k_EGCMsgWebAPIJobRequest":                  102,
	"k_EGCMsgWebAPIJobRequestHttpResponse":      104,
	"k_EGCMsgWebAPIJobRequestForwardResponse":   105,
	"k_EGCMsgMemCachedGet":                      200,
	"k_EGCMsgMemCachedGetResponse":              201,
	"k_EGCMsgMemCachedSet":                      202,
	"k_EGCMsgMemCachedDelete":                   203,
	"k_EGCMsgMemCachedStats":                    204,
	"k_EGCMsgMemCachedStatsResponse":            205,
	"k_EGCMsgSQLStats":                          210,
	"k_EGCMsgSQLStatsResponse":                  211,
	"k_EGCMsgMasterSetDirectory":                220,
	"k_EGCMsgMasterSetDirectoryResponse":        221,
	"k_EGCMsgMasterSetWebAPIRouting":            222,
	"k_EGCMsgMasterSetWebAPIRoutingResponse":    223,
	"k_EGCMsgMasterSetClientMsgRouting":         224,
	"k_EGCMsgMasterSetClientMsgRoutingResponse": 225,
	"k_EGCMsgSetOptions":                        226,
	"k_EGCMsgSetOptionsResponse":                227,
	"k_EGCMsgSystemBase2":                       500,
	"k_EGCMsgGetPurchaseTrustStatus":            501,
	"k_EGCMsgGetPurchaseTrustStatusResponse":    502,
	"k_EGCMsgUpdateSession":                     503,
	"k_EGCMsgGCAccountVacStatusChange":          504,
	"k_EGCMsgCheckFriendship":                   505,
	"k_EGCMsgCheckFriendshipResponse":           506,
	"k_EGCMsgGetPartnerAccountLink":             507,
	"k_EGCMsgGetPartnerAccountLinkResponse":     508,
	"k_EGCMsgVSReportedSuspiciousActivity":      509,
	"k_EGCMsgDPPartnerMicroTxns":                512,
	"k_EGCMsgDPPartnerMicroTxnsResponse":        513,
	"k_EGCMsgGetIPASN":                          514,
	"k_EGCMsgGetIPASNResponse":                  515,
	"k_EGCMsgGetAppFriendsList":                 516,
	"k_EGCMsgGetAppFriendsListResponse":         517,
	"k_EGCMsgVacVerificationChange":             518,
	"k_EGCMsgAccountPhoneNumberChange":          519,
	"k_EGCMsgAccountTwoFactorChange":            520,
	"k_EGCMsgCheckClanMembership":               521,
	"k_EGCMsgCheckClanMembershipResponse":       522,
}

func (x EGCSystemMsg) Enum() *EGCSystemMsg {
	p := new(EGCSystemMsg)
	*p = x
	return p
}
func (x EGCSystemMsg) String() string {
	return proto.EnumName(EGCSystemMsg_name, int32(x))
}
func (x *EGCSystemMsg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCSystemMsg_value, data, "EGCSystemMsg")
	if err != nil {
		return err
	}
	*x = EGCSystemMsg(value)
	return nil
}
func (EGCSystemMsg) EnumDescriptor() ([]byte, []int) { return fileDescriptor31, []int{0} }

type ESOMsg int32

const (
	ESOMsg_k_ESOMsg_Create                   ESOMsg = 21
	ESOMsg_k_ESOMsg_Update                   ESOMsg = 22
	ESOMsg_k_ESOMsg_Destroy                  ESOMsg = 23
	ESOMsg_k_ESOMsg_CacheSubscribed          ESOMsg = 24
	ESOMsg_k_ESOMsg_CacheUnsubscribed        ESOMsg = 25
	ESOMsg_k_ESOMsg_UpdateMultiple           ESOMsg = 26
	ESOMsg_k_ESOMsg_CacheSubscriptionRefresh ESOMsg = 28
	ESOMsg_k_ESOMsg_CacheSubscribedUpToDate  ESOMsg = 29
)

var ESOMsg_name = map[int32]string{
	21: "k_ESOMsg_Create",
	22: "k_ESOMsg_Update",
	23: "k_ESOMsg_Destroy",
	24: "k_ESOMsg_CacheSubscribed",
	25: "k_ESOMsg_CacheUnsubscribed",
	26: "k_ESOMsg_UpdateMultiple",
	28: "k_ESOMsg_CacheSubscriptionRefresh",
	29: "k_ESOMsg_CacheSubscribedUpToDate",
}
var ESOMsg_value = map[string]int32{
	"k_ESOMsg_Create":                   21,
	"k_ESOMsg_Update":                   22,
	"k_ESOMsg_Destroy":                  23,
	"k_ESOMsg_CacheSubscribed":          24,
	"k_ESOMsg_CacheUnsubscribed":        25,
	"k_ESOMsg_UpdateMultiple":           26,
	"k_ESOMsg_CacheSubscriptionRefresh": 28,
	"k_ESOMsg_CacheSubscribedUpToDate":  29,
}

func (x ESOMsg) Enum() *ESOMsg {
	p := new(ESOMsg)
	*p = x
	return p
}
func (x ESOMsg) String() string {
	return proto.EnumName(ESOMsg_name, int32(x))
}
func (x *ESOMsg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ESOMsg_value, data, "ESOMsg")
	if err != nil {
		return err
	}
	*x = ESOMsg(value)
	return nil
}
func (ESOMsg) EnumDescriptor() ([]byte, []int) { return fileDescriptor31, []int{1} }

type EGCBaseClientMsg int32

const (
	EGCBaseClientMsg_k_EMsgGCPingRequest                EGCBaseClientMsg = 3001
	EGCBaseClientMsg_k_EMsgGCPingResponse               EGCBaseClientMsg = 3002
	EGCBaseClientMsg_k_EMsgGCToClientPollConvarRequest  EGCBaseClientMsg = 3003
	EGCBaseClientMsg_k_EMsgGCToClientPollConvarResponse EGCBaseClientMsg = 3004
	EGCBaseClientMsg_k_EMsgGCClientWelcome              EGCBaseClientMsg = 4004
	EGCBaseClientMsg_k_EMsgGCServerWelcome              EGCBaseClientMsg = 4005
	EGCBaseClientMsg_k_EMsgGCClientHello                EGCBaseClientMsg = 4006
	EGCBaseClientMsg_k_EMsgGCServerHello                EGCBaseClientMsg = 4007
	EGCBaseClientMsg_k_EMsgGCClientConnectionStatus     EGCBaseClientMsg = 4009
	EGCBaseClientMsg_k_EMsgGCServerConnectionStatus     EGCBaseClientMsg = 4010
)

var EGCBaseClientMsg_name = map[int32]string{
	3001: "k_EMsgGCPingRequest",
	3002: "k_EMsgGCPingResponse",
	3003: "k_EMsgGCToClientPollConvarRequest",
	3004: "k_EMsgGCToClientPollConvarResponse",
	4004: "k_EMsgGCClientWelcome",
	4005: "k_EMsgGCServerWelcome",
	4006: "k_EMsgGCClientHello",
	4007: "k_EMsgGCServerHello",
	4009: "k_EMsgGCClientConnectionStatus",
	4010: "k_EMsgGCServerConnectionStatus",
}
var EGCBaseClientMsg_value = map[string]int32{
	"k_EMsgGCPingRequest":                3001,
	"k_EMsgGCPingResponse":               3002,
	"k_EMsgGCToClientPollConvarRequest":  3003,
	"k_EMsgGCToClientPollConvarResponse": 3004,
	"k_EMsgGCClientWelcome":              4004,
	"k_EMsgGCServerWelcome":              4005,
	"k_EMsgGCClientHello":                4006,
	"k_EMsgGCServerHello":                4007,
	"k_EMsgGCClientConnectionStatus":     4009,
	"k_EMsgGCServerConnectionStatus":     4010,
}

func (x EGCBaseClientMsg) Enum() *EGCBaseClientMsg {
	p := new(EGCBaseClientMsg)
	*p = x
	return p
}
func (x EGCBaseClientMsg) String() string {
	return proto.EnumName(EGCBaseClientMsg_name, int32(x))
}
func (x *EGCBaseClientMsg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCBaseClientMsg_value, data, "EGCBaseClientMsg")
	if err != nil {
		return err
	}
	*x = EGCBaseClientMsg(value)
	return nil
}
func (EGCBaseClientMsg) EnumDescriptor() ([]byte, []int) { return fileDescriptor31, []int{2} }

type EGCToGCMsg int32

const (
	EGCToGCMsg_k_EGCToGCMsgMasterAck                   EGCToGCMsg = 150
	EGCToGCMsg_k_EGCToGCMsgMasterAckResponse           EGCToGCMsg = 151
	EGCToGCMsg_k_EGCToGCMsgRouted                      EGCToGCMsg = 152
	EGCToGCMsg_k_EGCToGCMsgRoutedReply                 EGCToGCMsg = 153
	EGCToGCMsg_k_EMsgGCUpdateSubGCSessionInfo          EGCToGCMsg = 154
	EGCToGCMsg_k_EMsgGCRequestSubGCSessionInfo         EGCToGCMsg = 155
	EGCToGCMsg_k_EMsgGCRequestSubGCSessionInfoResponse EGCToGCMsg = 156
	EGCToGCMsg_k_EGCToGCMsgMasterStartupComplete       EGCToGCMsg = 157
	EGCToGCMsg_k_EMsgGCToGCSOCacheSubscribe            EGCToGCMsg = 158
	EGCToGCMsg_k_EMsgGCToGCSOCacheUnsubscribe          EGCToGCMsg = 159
	EGCToGCMsg_k_EMsgGCToGCLoadSessionSOCache          EGCToGCMsg = 160
	EGCToGCMsg_k_EMsgGCToGCLoadSessionSOCacheResponse  EGCToGCMsg = 161
	EGCToGCMsg_k_EMsgGCToGCUpdateSessionStats          EGCToGCMsg = 162
	EGCToGCMsg_k_EMsgGCToGCUniverseStartup             EGCToGCMsg = 163
	EGCToGCMsg_k_EMsgGCToGCUniverseStartupResponse     EGCToGCMsg = 164
	EGCToGCMsg_k_EMsgGCToGCForwardAccountDetails       EGCToGCMsg = 165
)

var EGCToGCMsg_name = map[int32]string{
	150: "k_EGCToGCMsgMasterAck",
	151: "k_EGCToGCMsgMasterAckResponse",
	152: "k_EGCToGCMsgRouted",
	153: "k_EGCToGCMsgRoutedReply",
	154: "k_EMsgGCUpdateSubGCSessionInfo",
	155: "k_EMsgGCRequestSubGCSessionInfo",
	156: "k_EMsgGCRequestSubGCSessionInfoResponse",
	157: "k_EGCToGCMsgMasterStartupComplete",
	158: "k_EMsgGCToGCSOCacheSubscribe",
	159: "k_EMsgGCToGCSOCacheUnsubscribe",
	160: "k_EMsgGCToGCLoadSessionSOCache",
	161: "k_EMsgGCToGCLoadSessionSOCacheResponse",
	162: "k_EMsgGCToGCUpdateSessionStats",
	163: "k_EMsgGCToGCUniverseStartup",
	164: "k_EMsgGCToGCUniverseStartupResponse",
	165: "k_EMsgGCToGCForwardAccountDetails",
}
var EGCToGCMsg_value = map[string]int32{
	"k_EGCToGCMsgMasterAck":                   150,
	"k_EGCToGCMsgMasterAckResponse":           151,
	"k_EGCToGCMsgRouted":                      152,
	"k_EGCToGCMsgRoutedReply":                 153,
	"k_EMsgGCUpdateSubGCSessionInfo":          154,
	"k_EMsgGCRequestSubGCSessionInfo":         155,
	"k_EMsgGCRequestSubGCSessionInfoResponse": 156,
	"k_EGCToGCMsgMasterStartupComplete":       157,
	"k_EMsgGCToGCSOCacheSubscribe":            158,
	"k_EMsgGCToGCSOCacheUnsubscribe":          159,
	"k_EMsgGCToGCLoadSessionSOCache":          160,
	"k_EMsgGCToGCLoadSessionSOCacheResponse":  161,
	"k_EMsgGCToGCUpdateSessionStats":          162,
	"k_EMsgGCToGCUniverseStartup":             163,
	"k_EMsgGCToGCUniverseStartupResponse":     164,
	"k_EMsgGCToGCForwardAccountDetails":       165,
}

func (x EGCToGCMsg) Enum() *EGCToGCMsg {
	p := new(EGCToGCMsg)
	*p = x
	return p
}
func (x EGCToGCMsg) String() string {
	return proto.EnumName(EGCToGCMsg_name, int32(x))
}
func (x *EGCToGCMsg) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EGCToGCMsg_value, data, "EGCToGCMsg")
	if err != nil {
		return err
	}
	*x = EGCToGCMsg(value)
	return nil
}
func (EGCToGCMsg) EnumDescriptor() ([]byte, []int) { return fileDescriptor31, []int{3} }

func init() {
	proto.RegisterEnum("dota.EGCSystemMsg", EGCSystemMsg_name, EGCSystemMsg_value)
	proto.RegisterEnum("dota.ESOMsg", ESOMsg_name, ESOMsg_value)
	proto.RegisterEnum("dota.EGCBaseClientMsg", EGCBaseClientMsg_name, EGCBaseClientMsg_value)
	proto.RegisterEnum("dota.EGCToGCMsg", EGCToGCMsg_name, EGCToGCMsg_value)
}

func init() { proto.RegisterFile("gcsystemmsgs.proto", fileDescriptor31) }

var fileDescriptor31 = []byte{
	// 1614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x57, 0x49, 0x73, 0xdb, 0xc8,
	0x15, 0x36, 0x42, 0x26, 0x87, 0xae, 0xa4, 0xd2, 0x6e, 0x5b, 0xa6, 0x25, 0x4b, 0x96, 0xf7, 0x45,
	0x4e, 0xf9, 0xe0, 0xec, 0x7b, 0x68, 0x90, 0xa2, 0xe4, 0x90, 0x12, 0x4d, 0x50, 0x52, 0x76, 0xa5,
	0x05, 0x3c, 0x91, 0x28, 0x81, 0x68, 0xa4, 0xbb, 0x21, 0x5b, 0x37, 0x67, 0x5f, 0x7e, 0x40, 0xf6,
	0xc4, 0x59, 0xec, 0xa4, 0x92, 0x7f, 0x90, 0xe5, 0x07, 0x64, 0x99, 0x99, 0xc3, 0xcc, 0x61, 0x2e,
	0xb3, 0xcf, 0x1f, 0x98, 0xc3, 0xec, 0x4b, 0xd5, 0x14, 0x80, 0x46, 0xa3, 0x41, 0x4a, 0x9a, 0x1b,
	0xd9, 0xdf, 0xf7, 0x5e, 0xbf, 0x7e, 0x3b, 0x10, 0x19, 0xb8, 0x62, 0x4f, 0x48, 0x18, 0x8d, 0xc4,
	0x40, 0x5c, 0x8f, 0x38, 0x93, 0x8c, 0x54, 0x3d, 0x26, 0xe9, 0xc2, 0xfd, 0x1a, 0xfa, 0x60, 0xb3,
	0x65, 0x3b, 0x29, 0xda, 0x11, 0x03, 0x72, 0x0c, 0x7d, 0x78, 0x67, 0xb3, 0xd9, 0xb2, 0x3b, 0x62,
	0xb0, 0x1c, 0xee, 0xd2, 0xc0, 0xf7, 0xf0, 0x11, 0x72, 0x14, 0x7d, 0x28, 0x3f, 0xec, 0xc4, 0x81,
	0xf4, 0xb1, 0x45, 0x4e, 0xa2, 0xe3, 0xf9, 0x51, 0x0b, 0x42, 0xe0, 0xbe, 0xdb, 0x83, 0x28, 0xd8,
	0xc3, 0x88, 0x9c, 0x40, 0x24, 0x47, 0x32, 0xb5, 0x37, 0xa9, 0x00, 0x7c, 0x83, 0x9c, 0x46, 0x33,
	0xf9, 0x79, 0xdd, 0x1d, 0xfa, 0xb0, 0x0b, 0x23, 0x08, 0x65, 0xfd, 0x0e, 0xe5, 0x1e, 0x78, 0xf8,
	0xa3, 0xa6, 0x9c, 0xcd, 0x42, 0x9b, 0x8d, 0x46, 0x34, 0xf4, 0xf0, 0xc7, 0xcc, 0x9b, 0x1c, 0x49,
	0xb9, 0xec, 0x06, 0x74, 0xcf, 0x0f, 0x07, 0xf8, 0xe3, 0xa4, 0x86, 0x8e, 0x15, 0x08, 0x8b, 0x72,
	0xe0, 0x13, 0xe4, 0x14, 0xaa, 0x95, 0x44, 0x5a, 0x74, 0x04, 0x02, 0xf8, 0x2e, 0x70, 0xfc, 0x49,
	0x32, 0x83, 0x4e, 0x98, 0x52, 0x06, 0xf6, 0x29, 0x32, 0x85, 0x8e, 0xe6, 0xd8, 0x46, 0xab, 0x07,
	0xdf, 0x89, 0x41, 0x48, 0xfc, 0x69, 0xd3, 0xb4, 0xe4, 0x58, 0x44, 0x2c, 0x14, 0x80, 0x3f, 0x43,
	0xce, 0xa1, 0xd3, 0x85, 0x13, 0xe4, 0x9a, 0x00, 0x9e, 0x68, 0x73, 0x24, 0x95, 0xc2, 0x71, 0x87,
	0x30, 0xa2, 0xf8, 0xb3, 0x64, 0x01, 0x5d, 0x3a, 0x9c, 0xa3, 0xf5, 0x7d, 0x6e, 0x1f, 0x7d, 0x29,
	0xaf, 0xd1, 0xec, 0xf6, 0x9a, 0x76, 0xbd, 0xdf, 0x6c, 0xe0, 0xcf, 0x93, 0x33, 0x68, 0x76, 0x3f,
	0x8e, 0xd6, 0xf2, 0x05, 0xf3, 0x81, 0xf5, 0x28, 0x5a, 0x0e, 0xb7, 0xd9, 0x5a, 0xe4, 0x51, 0x09,
	0x1e, 0xfe, 0xa2, 0xe9, 0x99, 0xf5, 0x24, 0xb8, 0x54, 0x82, 0x03, 0x42, 0xf8, 0x2c, 0xc4, 0x5f,
	0x22, 0xe7, 0xd1, 0xfc, 0x01, 0xa0, 0xd6, 0x5e, 0x37, 0x6d, 0x6c, 0x33, 0xb6, 0x13, 0x47, 0x75,
	0xd7, 0x65, 0x71, 0x28, 0x17, 0x39, 0x1b, 0x2d, 0x87, 0x51, 0x2c, 0xf1, 0xcd, 0x92, 0xff, 0x21,
	0xf4, 0x96, 0xfa, 0xfd, 0x6e, 0xee, 0x4c, 0xdb, 0xbc, 0x65, 0x0c, 0xd4, 0xb7, 0x34, 0xcc, 0xa0,
	0x77, 0x39, 0xf4, 0x41, 0x48, 0x07, 0x64, 0x1c, 0xe1, 0x26, 0x99, 0x47, 0xa7, 0x72, 0xa4, 0x07,
	0x2e, 0xe3, 0x9e, 0x13, 0x47, 0x11, 0xe3, 0xb2, 0xee, 0xca, 0xe4, 0x15, 0x8b, 0xe4, 0x0a, 0xba,
	0x60, 0x38, 0x48, 0x59, 0xd7, 0x00, 0x49, 0xfd, 0x40, 0x6c, 0x1a, 0xae, 0x6c, 0x99, 0x4f, 0xe9,
	0x81, 0x0b, 0xfe, 0x2e, 0x2c, 0x87, 0x12, 0x78, 0x3d, 0x8a, 0x3a, 0x20, 0x04, 0x1d, 0x00, 0x5e,
	0x36, 0x0d, 0x59, 0xf4, 0x43, 0x4f, 0xa9, 0x13, 0xf8, 0x96, 0x99, 0x2b, 0x5d, 0x26, 0x64, 0x3d,
	0x00, 0x2e, 0xf1, 0x97, 0xcd, 0xa4, 0x6c, 0x81, 0x6c, 0xfb, 0x2e, 0x84, 0x02, 0x04, 0x6e, 0x97,
	0x2b, 0xa6, 0x08, 0x1c, 0xee, 0x8c, 0x89, 0xa8, 0xcc, 0x17, 0x78, 0xc5, 0x7c, 0xab, 0x01, 0x68,
	0x37, 0xad, 0x96, 0x42, 0xed, 0x79, 0x8b, 0x1c, 0x40, 0x5d, 0x88, 0xbb, 0xe6, 0xeb, 0xca, 0x98,
	0x96, 0xbf, 0x4d, 0xa6, 0xd1, 0x94, 0x71, 0xc1, 0x72, 0xb7, 0xcd, 0x5c, 0x9a, 0xba, 0xb1, 0x47,
	0xce, 0xa2, 0xb9, 0x7d, 0x21, 0x2d, 0xed, 0x90, 0x39, 0x34, 0x5d, 0xae, 0x74, 0x33, 0xf3, 0xfb,
	0xa6, 0x71, 0x2d, 0x90, 0x06, 0x03, 0xaf, 0x8d, 0x65, 0xba, 0x81, 0x69, 0xf5, 0xeb, 0xa6, 0x83,
	0x93, 0x44, 0x69, 0x8e, 0xa8, 0x1f, 0xe0, 0x8d, 0xd2, 0xad, 0xf9, 0xb1, 0x96, 0xfa, 0x0a, 0x99,
	0x45, 0x27, 0x0d, 0xcd, 0x29, 0xda, 0x87, 0x51, 0x14, 0x50, 0x09, 0xf8, 0xab, 0xe4, 0x02, 0x3a,
	0x73, 0x10, 0xaa, 0x75, 0x7c, 0xad, 0x64, 0x39, 0xa7, 0xa1, 0x6c, 0x25, 0xd9, 0xd9, 0xa5, 0x42,
	0xe0, 0xaf, 0x97, 0x2c, 0x2f, 0x61, 0x5a, 0xfe, 0x1b, 0xa6, 0x89, 0x13, 0x29, 0x88, 0xbf, 0x49,
	0x2e, 0xa2, 0xb3, 0x07, 0xc2, 0x5a, 0xcb, 0xb7, 0xcc, 0x2a, 0x6a, 0x81, 0xec, 0x02, 0x17, 0x2c,
	0xa4, 0x2b, 0x49, 0xbb, 0xc2, 0x9b, 0x66, 0x15, 0x8d, 0x81, 0x5a, 0xc3, 0xb7, 0xcd, 0x94, 0x4b,
	0xfb, 0x76, 0x14, 0xc0, 0xdd, 0x8e, 0x18, 0x60, 0x6a, 0xfa, 0x61, 0x03, 0xb6, 0xea, 0xdd, 0xe5,
	0x1e, 0x0c, 0x7c, 0x21, 0x81, 0xa7, 0x15, 0xb0, 0x4d, 0x5d, 0x10, 0x18, 0x4c, 0x5f, 0x66, 0xac,
	0x5b, 0x6c, 0x2b, 0x2f, 0xe4, 0x6d, 0xb3, 0xd0, 0xc6, 0xd1, 0x25, 0x29, 0x23, 0x6d, 0xc7, 0x90,
	0x5c, 0x43, 0x97, 0x0f, 0x62, 0x2e, 0x32, 0x9e, 0x4c, 0x00, 0x4d, 0xf6, 0xc9, 0xb4, 0x61, 0x34,
	0x8c, 0x6c, 0xea, 0x0e, 0xc1, 0x6b, 0x81, 0xc4, 0xff, 0xb1, 0xc8, 0xd9, 0xa2, 0xf7, 0x99, 0x90,
	0x16, 0xfe, 0xaf, 0xb5, 0xaf, 0xb4, 0x03, 0x12, 0xff, 0xcf, 0x22, 0xb3, 0x85, 0x3f, 0x35, 0xd4,
	0x80, 0x00, 0x24, 0xe0, 0xff, 0x5b, 0xe4, 0x54, 0x11, 0xf3, 0x42, 0x30, 0xcd, 0xd6, 0x47, 0x2c,
	0x72, 0xbe, 0x08, 0x7a, 0x19, 0xd4, 0x57, 0x3f, 0x6a, 0x91, 0x29, 0x84, 0x75, 0x62, 0xde, 0x6e,
	0x67, 0xb2, 0x8f, 0x5b, 0x64, 0xae, 0x70, 0x62, 0x7e, 0xac, 0xa5, 0x9e, 0xb0, 0xc8, 0x7c, 0x31,
	0x16, 0x3b, 0x34, 0x89, 0x80, 0x03, 0xb2, 0xe1, 0x73, 0x70, 0x25, 0xe3, 0x7b, 0xf8, 0x29, 0x8b,
	0x5c, 0x46, 0xe7, 0x0e, 0x26, 0x68, 0x4d, 0x4f, 0x97, 0x8d, 0xcc, 0x89, 0x2a, 0xb8, 0x2c, 0x96,
	0xc9, 0x64, 0x7c, 0xc6, 0x22, 0xd7, 0x8a, 0x71, 0xb4, 0x3f, 0x49, 0x6b, 0x7c, 0xd6, 0x22, 0x97,
	0x8a, 0x44, 0xd5, 0x64, 0x3b, 0xf0, 0x21, 0x94, 0x49, 0xcb, 0x54, 0x4a, 0x9f, 0xb3, 0xc8, 0x75,
	0x74, 0xf5, 0x3d, 0x79, 0x5a, 0xef, 0xf3, 0x16, 0xa9, 0x19, 0x2b, 0x02, 0xc8, 0xd5, 0x28, 0xe9,
	0x2b, 0x02, 0xbf, 0x50, 0x72, 0x46, 0x01, 0x68, 0xc9, 0x17, 0x93, 0xb5, 0xe3, 0xd8, 0xe4, 0x72,
	0x71, 0x03, 0xbf, 0x5c, 0x31, 0x5f, 0x9f, 0x14, 0x44, 0xcc, 0xdd, 0x21, 0x15, 0xd0, 0xe7, 0xb1,
	0x90, 0x89, 0xcf, 0x63, 0x81, 0x5f, 0xa9, 0x98, 0xaf, 0xdf, 0x9f, 0xa4, 0xef, 0x7a, 0xb5, 0x42,
	0x66, 0x8a, 0xe6, 0x98, 0x0d, 0xd0, 0x7c, 0x52, 0xbe, 0x56, 0x21, 0x17, 0x8d, 0x3e, 0x62, 0xab,
	0x0a, 0x5e, 0xa7, 0x6e, 0xa6, 0xc4, 0x1e, 0xd2, 0x70, 0x00, 0xf8, 0xf5, 0x8a, 0x99, 0x72, 0xf6,
	0x10, 0xdc, 0x9d, 0x45, 0xee, 0x43, 0xe8, 0x89, 0xa1, 0x1f, 0xe1, 0x37, 0x2a, 0xe4, 0x42, 0x51,
	0xc3, 0x63, 0xa8, 0x36, 0xe3, 0xcd, 0x0a, 0x39, 0x57, 0x6a, 0xc4, 0x5d, 0xca, 0x65, 0x08, 0x5c,
	0x5d, 0xd9, 0xf6, 0xc3, 0x1d, 0xfc, 0x56, 0x85, 0x2c, 0xa0, 0x8b, 0x87, 0x72, 0xb4, 0xbe, 0xb7,
	0x2b, 0xe4, 0x6a, 0x51, 0xb6, 0xeb, 0x4e, 0x0f, 0x92, 0xd9, 0x09, 0x9e, 0x13, 0x8b, 0xc8, 0x77,
	0x7d, 0x16, 0x8b, 0x64, 0x8e, 0xee, 0xfa, 0x72, 0x0f, 0xbf, 0x53, 0x31, 0xc3, 0xd1, 0xe8, 0x2a,
	0xad, 0x1d, 0xdf, 0xe5, 0xac, 0x7f, 0x37, 0x14, 0xf8, 0x5e, 0xd5, 0xcc, 0xcd, 0x49, 0x82, 0xbe,
	0xf4, 0xbb, 0x55, 0xb3, 0x36, 0xd2, 0x69, 0x52, 0x77, 0x56, 0xf0, 0xf7, 0xaa, 0x66, 0x6d, 0xe4,
	0xc7, 0x5a, 0xea, 0xfb, 0x55, 0x72, 0xba, 0xdc, 0x47, 0xa3, 0x48, 0x79, 0xa8, 0xed, 0x0b, 0x89,
	0x7f, 0x50, 0x35, 0xf3, 0x73, 0x02, 0xd7, 0x7a, 0x7e, 0x58, 0x35, 0x5d, 0xb8, 0x4e, 0xdd, 0x75,
	0xe0, 0xfe, 0xb6, 0x9f, 0x4d, 0x33, 0x15, 0xaa, 0x1f, 0x55, 0xcd, 0x88, 0x2a, 0xc7, 0x75, 0x87,
	0x2c, 0x84, 0x95, 0x78, 0xb4, 0x05, 0x5c, 0xd1, 0x7e, 0x5c, 0x35, 0xd3, 0x4c, 0xd1, 0xfa, 0x77,
	0xd8, 0x22, 0x4d, 0x8a, 0x51, 0x91, 0x7e, 0x52, 0x25, 0x67, 0x8a, 0xb9, 0x9d, 0x06, 0xd6, 0x0e,
	0x68, 0xd8, 0x81, 0x44, 0x51, 0x1a, 0xfa, 0x9f, 0x56, 0xc9, 0x15, 0x74, 0xfe, 0x10, 0x86, 0xb6,
	0xfd, 0x67, 0xd5, 0x85, 0x97, 0x2c, 0xf4, 0x81, 0xa6, 0xb3, 0x5a, 0xec, 0xe6, 0xe9, 0xef, 0x4d,
	0x9b, 0x43, 0x32, 0xd1, 0xa6, 0x4a, 0x87, 0x59, 0x9a, 0xe2, 0x13, 0xe4, 0x78, 0xea, 0xee, 0xec,
	0xb0, 0x01, 0x42, 0x72, 0xb6, 0x87, 0x6b, 0xaa, 0x9d, 0x2b, 0xf9, 0xa4, 0x87, 0x39, 0xf1, 0x96,
	0x70, 0xb9, 0xbf, 0x05, 0x1e, 0x3e, 0xa9, 0xf6, 0x73, 0x03, 0x5d, 0x0b, 0x45, 0x81, 0x4f, 0xab,
	0x71, 0x64, 0x5e, 0x94, 0xcf, 0x14, 0x3c, 0xa3, 0x46, 0xda, 0xa4, 0xea, 0x28, 0x5b, 0x19, 0xb6,
	0x39, 0x88, 0x21, 0x9e, 0x55, 0x63, 0x67, 0x5f, 0x0b, 0xd6, 0xa2, 0x3e, 0x6b, 0x24, 0xd6, 0xcf,
	0x2d, 0x3c, 0xf6, 0x3e, 0x84, 0x9b, 0x2d, 0x3b, 0x29, 0x6d, 0xdd, 0x45, 0x54, 0xe5, 0xa7, 0xf5,
	0xd6, 0x4d, 0xdb, 0x49, 0x36, 0x86, 0xfe, 0x51, 0x53, 0x2d, 0xdf, 0x40, 0x94, 0xf3, 0xfe, 0x59,
	0x53, 0x09, 0x92, 0x42, 0x7d, 0x96, 0xe9, 0xea, 0xb2, 0x20, 0xb0, 0x59, 0xb8, 0x4b, 0x79, 0xae,
	0xe2, 0x5f, 0x35, 0x95, 0xc7, 0x07, 0xf2, 0x94, 0xc2, 0x7f, 0xd7, 0x54, 0x4f, 0x48, 0x89, 0x19,
	0x6d, 0x03, 0x02, 0x97, 0x8d, 0x00, 0x3f, 0x98, 0x37, 0x31, 0x27, 0xfd, 0xa0, 0xc8, 0xb1, 0x87,
	0xf3, 0xa6, 0xf5, 0x99, 0xdc, 0x12, 0x04, 0x01, 0xc3, 0x7f, 0x29, 0x21, 0x99, 0x54, 0x86, 0xfc,
	0x75, 0x5e, 0xa5, 0x9a, 0x21, 0x63, 0xb3, 0x30, 0x84, 0x74, 0xcd, 0x55, 0x1d, 0xed, 0x6f, 0x25,
	0x52, 0x26, 0x3e, 0x41, 0xfa, 0xfb, 0xfc, 0xc2, 0x93, 0x55, 0x84, 0x9a, 0xc9, 0xd3, 0xd2, 0x84,
	0xd3, 0x8d, 0x4d, 0xfd, 0xcf, 0x5a, 0x76, 0xdd, 0xdd, 0xc1, 0x3f, 0xb7, 0x74, 0xa9, 0x8c, 0x63,
	0xda, 0x09, 0xbf, 0x28, 0xda, 0xb7, 0xe2, 0x24, 0x0d, 0x1e, 0x3c, 0xfc, 0xcb, 0x62, 0xc2, 0x96,
	0x80, 0xec, 0xbb, 0xf0, 0x57, 0x96, 0x69, 0xaa, 0xea, 0xa7, 0xf1, 0x56, 0x62, 0x75, 0xda, 0x54,
	0x93, 0xcf, 0x14, 0xfc, 0x6b, 0x4b, 0xf5, 0xc4, 0x94, 0xa4, 0xe2, 0x33, 0xc1, 0xfa, 0x8d, 0x45,
	0x3e, 0x92, 0x2e, 0x14, 0x87, 0xb1, 0xb4, 0xbd, 0xbf, 0x2d, 0xc6, 0x58, 0xe9, 0x4d, 0xe9, 0x87,
	0x61, 0x1c, 0xd9, 0x6c, 0x14, 0xa5, 0x2b, 0xc0, 0xef, 0xf2, 0xf5, 0x42, 0x65, 0x41, 0xcb, 0x76,
	0x56, 0xcb, 0x29, 0x8a, 0x7f, 0x5f, 0x7a, 0x83, 0x41, 0x31, 0x2a, 0x05, 0xdf, 0x9f, 0x20, 0xb5,
	0x19, 0xf5, 0x94, 0x65, 0x8a, 0x8f, 0xff, 0x90, 0x0f, 0xe2, 0x43, 0x48, 0xfa, 0x05, 0x7f, 0x9c,
	0xd0, 0x58, 0x1a, 0x47, 0xd9, 0xa2, 0xf1, 0x27, 0x4b, 0x75, 0x9d, 0x82, 0x14, 0xfa, 0xbb, 0xc0,
	0x05, 0xa8, 0x87, 0xe2, 0x3f, 0x5b, 0xaa, 0xeb, 0x1c, 0xc4, 0xd0, 0x17, 0x3e, 0xb0, 0xca, 0x85,
	0xd3, 0xb2, 0xd5, 0x96, 0x36, 0xb6, 0xc9, 0x3e, 0xb4, 0x6e, 0xbe, 0x7f, 0xc9, 0xba, 0x67, 0x1d,
	0x79, 0x37, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xc5, 0x24, 0x72, 0x61, 0x10, 0x00, 0x00,
}
