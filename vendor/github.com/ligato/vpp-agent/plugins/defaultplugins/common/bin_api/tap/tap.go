// Code generated by govpp binapi-generator DO NOT EDIT.
// Package tap represents the VPP binary API of the 'tap' VPP module.
// Generated from '/usr/share/vpp/api/tap.api.json'
package tap

import "git.fd.io/govpp.git/api"

// VlApiVersion contains version of the API.
const VlAPIVersion = 0x3a0725de

// TapConnect represents the VPP binary API message 'tap_connect'.
// Generated from '/usr/share/vpp/api/tap.api.json', line 29:
//
//            "tap_connect",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "use_random_mac"
//            ],
//            [
//                "u8",
//                "tap_name",
//                64
//            ],
//            [
//                "u8",
//                "mac_address",
//                6
//            ],
//            [
//                "u8",
//                "renumber"
//            ],
//            [
//                "u32",
//                "custom_dev_instance"
//            ],
//            [
//                "u8",
//                "ip4_address_set"
//            ],
//            [
//                "u8",
//                "ip4_address",
//                4
//            ],
//            [
//                "u8",
//                "ip4_mask_width"
//            ],
//            [
//                "u8",
//                "ip6_address_set"
//            ],
//            [
//                "u8",
//                "ip6_address",
//                16
//            ],
//            [
//                "u8",
//                "ip6_mask_width"
//            ],
//            [
//                "u8",
//                "tag",
//                64
//            ],
//            {
//                "crc": "0x9b9c396f"
//            }
//
type TapConnect struct {
	UseRandomMac      uint8
	TapName           []byte `struc:"[64]byte"`
	MacAddress        []byte `struc:"[6]byte"`
	Renumber          uint8
	CustomDevInstance uint32
	IP4AddressSet     uint8
	IP4Address        []byte `struc:"[4]byte"`
	IP4MaskWidth      uint8
	IP6AddressSet     uint8
	IP6Address        []byte `struc:"[16]byte"`
	IP6MaskWidth      uint8
	Tag               []byte `struc:"[64]byte"`
}

func (*TapConnect) GetMessageName() string {
	return "tap_connect"
}
func (*TapConnect) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*TapConnect) GetCrcString() string {
	return "9b9c396f"
}
func NewTapConnect() api.Message {
	return &TapConnect{}
}

// TapConnectReply represents the VPP binary API message 'tap_connect_reply'.
// Generated from '/usr/share/vpp/api/tap.api.json', line 100:
//
//            "tap_connect_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            {
//                "crc": "0xfda5941f"
//            }
//
type TapConnectReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*TapConnectReply) GetMessageName() string {
	return "tap_connect_reply"
}
func (*TapConnectReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*TapConnectReply) GetCrcString() string {
	return "fda5941f"
}
func NewTapConnectReply() api.Message {
	return &TapConnectReply{}
}

// TapModify represents the VPP binary API message 'tap_modify'.
// Generated from '/usr/share/vpp/api/tap.api.json', line 122:
//
//            "tap_modify",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u8",
//                "use_random_mac"
//            ],
//            [
//                "u8",
//                "tap_name",
//                64
//            ],
//            [
//                "u8",
//                "mac_address",
//                6
//            ],
//            [
//                "u8",
//                "renumber"
//            ],
//            [
//                "u32",
//                "custom_dev_instance"
//            ],
//            {
//                "crc": "0x8047ae5c"
//            }
//
type TapModify struct {
	SwIfIndex         uint32
	UseRandomMac      uint8
	TapName           []byte `struc:"[64]byte"`
	MacAddress        []byte `struc:"[6]byte"`
	Renumber          uint8
	CustomDevInstance uint32
}

func (*TapModify) GetMessageName() string {
	return "tap_modify"
}
func (*TapModify) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*TapModify) GetCrcString() string {
	return "8047ae5c"
}
func NewTapModify() api.Message {
	return &TapModify{}
}

// TapModifyReply represents the VPP binary API message 'tap_modify_reply'.
// Generated from '/usr/share/vpp/api/tap.api.json', line 166:
//
//            "tap_modify_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            {
//                "crc": "0xfda5941f"
//            }
//
type TapModifyReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*TapModifyReply) GetMessageName() string {
	return "tap_modify_reply"
}
func (*TapModifyReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*TapModifyReply) GetCrcString() string {
	return "fda5941f"
}
func NewTapModifyReply() api.Message {
	return &TapModifyReply{}
}

// TapDelete represents the VPP binary API message 'tap_delete'.
// Generated from '/usr/share/vpp/api/tap.api.json', line 188:
//
//            "tap_delete",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            {
//                "crc": "0x529cb13f"
//            }
//
type TapDelete struct {
	SwIfIndex uint32
}

func (*TapDelete) GetMessageName() string {
	return "tap_delete"
}
func (*TapDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*TapDelete) GetCrcString() string {
	return "529cb13f"
}
func NewTapDelete() api.Message {
	return &TapDelete{}
}

// TapDeleteReply represents the VPP binary API message 'tap_delete_reply'.
// Generated from '/usr/share/vpp/api/tap.api.json', line 210:
//
//            "tap_delete_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type TapDeleteReply struct {
	Retval int32
}

func (*TapDeleteReply) GetMessageName() string {
	return "tap_delete_reply"
}
func (*TapDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*TapDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func NewTapDeleteReply() api.Message {
	return &TapDeleteReply{}
}

// SwInterfaceTapDump represents the VPP binary API message 'sw_interface_tap_dump'.
// Generated from '/usr/share/vpp/api/tap.api.json', line 228:
//
//            "sw_interface_tap_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type SwInterfaceTapDump struct {
}

func (*SwInterfaceTapDump) GetMessageName() string {
	return "sw_interface_tap_dump"
}
func (*SwInterfaceTapDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func (*SwInterfaceTapDump) GetCrcString() string {
	return "51077d14"
}
func NewSwInterfaceTapDump() api.Message {
	return &SwInterfaceTapDump{}
}

// SwInterfaceTapDetails represents the VPP binary API message 'sw_interface_tap_details'.
// Generated from '/usr/share/vpp/api/tap.api.json', line 246:
//
//            "sw_interface_tap_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u8",
//                "dev_name",
//                64
//            ],
//            {
//                "crc": "0x76229a57"
//            }
//
type SwInterfaceTapDetails struct {
	SwIfIndex uint32
	DevName   []byte `struc:"[64]byte"`
}

func (*SwInterfaceTapDetails) GetMessageName() string {
	return "sw_interface_tap_details"
}
func (*SwInterfaceTapDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func (*SwInterfaceTapDetails) GetCrcString() string {
	return "76229a57"
}
func NewSwInterfaceTapDetails() api.Message {
	return &SwInterfaceTapDetails{}
}
