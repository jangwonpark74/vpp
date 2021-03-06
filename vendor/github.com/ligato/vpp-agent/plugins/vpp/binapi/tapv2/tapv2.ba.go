// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: /usr/share/vpp/api/tapv2.api.json

/*
 Package tapv2 is a generated from VPP binary API module 'tapv2'.

 It contains following objects:
	  6 messages
	  3 services

*/
package tapv2

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

/* Messages */

// TapCreateV2 represents the VPP binary API message 'tap_create_v2'.
//
//            "tap_create_v2",
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
//                "id"
//            ],
//            [
//                "u8",
//                "use_random_mac"
//            ],
//            [
//                "u8",
//                "mac_address",
//                6
//            ],
//            [
//                "u16",
//                "tx_ring_sz"
//            ],
//            [
//                "u16",
//                "rx_ring_sz"
//            ],
//            [
//                "u8",
//                "host_namespace_set"
//            ],
//            [
//                "u8",
//                "host_namespace",
//                64
//            ],
//            [
//                "u8",
//                "host_mac_addr_set"
//            ],
//            [
//                "u8",
//                "host_mac_addr",
//                6
//            ],
//            [
//                "u8",
//                "host_if_name_set"
//            ],
//            [
//                "u8",
//                "host_if_name",
//                64
//            ],
//            [
//                "u8",
//                "host_bridge_set"
//            ],
//            [
//                "u8",
//                "host_bridge",
//                64
//            ],
//            [
//                "u8",
//                "host_ip4_addr_set"
//            ],
//            [
//                "u8",
//                "host_ip4_addr",
//                4
//            ],
//            [
//                "u8",
//                "host_ip4_prefix_len"
//            ],
//            [
//                "u8",
//                "host_ip6_addr_set"
//            ],
//            [
//                "u8",
//                "host_ip6_addr",
//                16
//            ],
//            [
//                "u8",
//                "host_ip6_prefix_len"
//            ],
//            [
//                "u8",
//                "host_ip4_gw_set"
//            ],
//            [
//                "u8",
//                "host_ip4_gw",
//                4
//            ],
//            [
//                "u8",
//                "host_ip6_gw_set"
//            ],
//            [
//                "u8",
//                "host_ip6_gw",
//                16
//            ],
//            [
//                "u8",
//                "tag",
//                64
//            ],
//            {
//                "crc": "0x34ce8043"
//            }
//
type TapCreateV2 struct {
	ID               uint32
	UseRandomMac     uint8
	MacAddress       []byte `struc:"[6]byte"`
	TxRingSz         uint16
	RxRingSz         uint16
	HostNamespaceSet uint8
	HostNamespace    []byte `struc:"[64]byte"`
	HostMacAddrSet   uint8
	HostMacAddr      []byte `struc:"[6]byte"`
	HostIfNameSet    uint8
	HostIfName       []byte `struc:"[64]byte"`
	HostBridgeSet    uint8
	HostBridge       []byte `struc:"[64]byte"`
	HostIP4AddrSet   uint8
	HostIP4Addr      []byte `struc:"[4]byte"`
	HostIP4PrefixLen uint8
	HostIP6AddrSet   uint8
	HostIP6Addr      []byte `struc:"[16]byte"`
	HostIP6PrefixLen uint8
	HostIP4GwSet     uint8
	HostIP4Gw        []byte `struc:"[4]byte"`
	HostIP6GwSet     uint8
	HostIP6Gw        []byte `struc:"[16]byte"`
	Tag              []byte `struc:"[64]byte"`
}

func (*TapCreateV2) GetMessageName() string {
	return "tap_create_v2"
}
func (*TapCreateV2) GetCrcString() string {
	return "34ce8043"
}
func (*TapCreateV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapCreateV2Reply represents the VPP binary API message 'tap_create_v2_reply'.
//
//            "tap_create_v2_reply",
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
type TapCreateV2Reply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*TapCreateV2Reply) GetMessageName() string {
	return "tap_create_v2_reply"
}
func (*TapCreateV2Reply) GetCrcString() string {
	return "fda5941f"
}
func (*TapCreateV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// TapDeleteV2 represents the VPP binary API message 'tap_delete_v2'.
//
//            "tap_delete_v2",
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
type TapDeleteV2 struct {
	SwIfIndex uint32
}

func (*TapDeleteV2) GetMessageName() string {
	return "tap_delete_v2"
}
func (*TapDeleteV2) GetCrcString() string {
	return "529cb13f"
}
func (*TapDeleteV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// TapDeleteV2Reply represents the VPP binary API message 'tap_delete_v2_reply'.
//
//            "tap_delete_v2_reply",
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
type TapDeleteV2Reply struct {
	Retval int32
}

func (*TapDeleteV2Reply) GetMessageName() string {
	return "tap_delete_v2_reply"
}
func (*TapDeleteV2Reply) GetCrcString() string {
	return "e8d4e804"
}
func (*TapDeleteV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// SwInterfaceTapV2Dump represents the VPP binary API message 'sw_interface_tap_v2_dump'.
//
//            "sw_interface_tap_v2_dump",
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
type SwInterfaceTapV2Dump struct{}

func (*SwInterfaceTapV2Dump) GetMessageName() string {
	return "sw_interface_tap_v2_dump"
}
func (*SwInterfaceTapV2Dump) GetCrcString() string {
	return "51077d14"
}
func (*SwInterfaceTapV2Dump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// SwInterfaceTapV2Details represents the VPP binary API message 'sw_interface_tap_v2_details'.
//
//            "sw_interface_tap_v2_details",
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
//                "u32",
//                "id"
//            ],
//            [
//                "u8",
//                "dev_name",
//                64
//            ],
//            [
//                "u16",
//                "tx_ring_sz"
//            ],
//            [
//                "u16",
//                "rx_ring_sz"
//            ],
//            [
//                "u8",
//                "host_mac_addr",
//                6
//            ],
//            [
//                "u8",
//                "host_if_name",
//                64
//            ],
//            [
//                "u8",
//                "host_namespace",
//                64
//            ],
//            [
//                "u8",
//                "host_bridge",
//                64
//            ],
//            [
//                "u8",
//                "host_ip4_addr",
//                4
//            ],
//            [
//                "u8",
//                "host_ip4_prefix_len"
//            ],
//            [
//                "u8",
//                "host_ip6_addr",
//                16
//            ],
//            [
//                "u8",
//                "host_ip6_prefix_len"
//            ],
//            {
//                "crc": "0xb4c58229"
//            }
//
type SwInterfaceTapV2Details struct {
	SwIfIndex        uint32
	ID               uint32
	DevName          []byte `struc:"[64]byte"`
	TxRingSz         uint16
	RxRingSz         uint16
	HostMacAddr      []byte `struc:"[6]byte"`
	HostIfName       []byte `struc:"[64]byte"`
	HostNamespace    []byte `struc:"[64]byte"`
	HostBridge       []byte `struc:"[64]byte"`
	HostIP4Addr      []byte `struc:"[4]byte"`
	HostIP4PrefixLen uint8
	HostIP6Addr      []byte `struc:"[16]byte"`
	HostIP6PrefixLen uint8
}

func (*SwInterfaceTapV2Details) GetMessageName() string {
	return "sw_interface_tap_v2_details"
}
func (*SwInterfaceTapV2Details) GetCrcString() string {
	return "b4c58229"
}
func (*SwInterfaceTapV2Details) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

/* Services */

type Services interface {
	DumpSwInterfaceTapV2(*SwInterfaceTapV2Dump) (*SwInterfaceTapV2Details, error)
	TapCreateV2(*TapCreateV2) (*TapCreateV2Reply, error)
	TapDeleteV2(*TapDeleteV2) (*TapDeleteV2Reply, error)
}

func init() {
	api.RegisterMessage((*TapCreateV2)(nil), "tapv2.TapCreateV2")
	api.RegisterMessage((*TapCreateV2Reply)(nil), "tapv2.TapCreateV2Reply")
	api.RegisterMessage((*TapDeleteV2)(nil), "tapv2.TapDeleteV2")
	api.RegisterMessage((*TapDeleteV2Reply)(nil), "tapv2.TapDeleteV2Reply")
	api.RegisterMessage((*SwInterfaceTapV2Dump)(nil), "tapv2.SwInterfaceTapV2Dump")
	api.RegisterMessage((*SwInterfaceTapV2Details)(nil), "tapv2.SwInterfaceTapV2Details")
}
