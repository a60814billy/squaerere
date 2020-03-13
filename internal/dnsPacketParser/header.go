package dnsPacketParser

import (
	"encoding/binary"
)

const (
	TYPE_A     uint16 = 1   //  1 a host address
	TYPE_NS    uint16 = 2   // 2 an authoritative name server
	TYPE_MD    uint16 = 3   // 3 a mail destination (Obsolete - use MX)
	TYPE_MF    uint16 = 4   // 4 a mail forwarder (Obsolete - use MX)
	TYPE_CNAME uint16 = 5   // 5 the canonical name for an alias
	TYPE_SOA   uint16 = 6   // 6 marks the start of a zone of authority
	TYPE_MB    uint16 = 7   // 7 a mailbox domain name (EXPERIMENTAL)
	TYPE_MG    uint16 = 8   // 8 a mail group member (EXPERIMENTAL)
	TYPE_MR    uint16 = 9   // 9 a mail rename domain name (EXPERIMENTAL)
	TYPE_NULL  uint16 = 10  // 10 a null RR (EXPERIMENTAL)
	TYPE_WKS   uint16 = 11  // 11 a well known service description
	TYPE_PTR   uint16 = 12  // 12 a domain name pointer
	TYPE_HINFO uint16 = 13  // 13 host information
	TYPE_MINFO uint16 = 14  // 14 mailbox or mail list information
	TYPE_MX    uint16 = 15  // 15 mail exchange
	TYPE_TXT   uint16 = 16  // 16 text strings
	TYPE_AXFR  uint16 = 252 // 252 A request for a transfer of an entire zone
	TYPE_MAILB uint16 = 253 // 253 A request for mailbox-related records (MB, MG or MR)
	TYPE_MAILA uint16 = 254 // 254 A request for mail agent RRs (Obsolete - see MX)
	TYPE_ALL   uint16 = 255 // 255 A request for all records

	CLASS_IN  uint16 = 1   // 1 the Internet
	CLASS_CS  uint16 = 2   // 2 the CSNET class (Obsolete - used only for examples in some obsolete RFCs)
	CLASS_CH  uint16 = 3   // 3 the CHAOS class
	CLASS_HS  uint16 = 4   // 4 Hesiod [Dyer 87]
	CLASS_ALL uint16 = 255 // 255 any class
)

func TypeString(t uint16) string {
	typeMap := map[uint16]string{
		1:  "A",
		2:  "NS",
		5:  "CNAME",
		6:  "SOA",
		12: "PTR",
		15: "MX",
		16: "TXT",
	}
	return typeMap[t]
}

type DnsQueryHeader struct {
	ID      uint16
	QR      uint8
	OpCode  uint8
	AA      uint8
	TC      uint8
	RD      uint8
	RA      uint8
	Z       uint8
	RCode   uint8
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

func (d *DnsQueryHeader) ToBytes() []byte {
	var r []byte
	r = append(r, uint8(d.ID>>8))
	r = append(r, uint8(d.ID))
	r = append(r, d.QR<<7|d.OpCode<<3|d.AA<<2|d.TC<<1|d.RD)
	r = append(r, d.RA<<7|d.Z<<3|d.RCode)
	r = append(r, uint8(d.QDCount>>8))
	r = append(r, uint8(d.QDCount))
	r = append(r, uint8(d.ANCount>>8))
	r = append(r, uint8(d.ANCount))
	r = append(r, uint8(d.NSCount>>8))
	r = append(r, uint8(d.NSCount))
	r = append(r, uint8(d.ARCount>>8))
	r = append(r, uint8(d.ARCount))

	return r
}

func ParseDnsQueryHeader(header [12]byte) DnsQueryHeader {
	res := DnsQueryHeader{}
	res.ID = binary.BigEndian.Uint16(header[0:2])
	res.QR = header[2:3][0] & 128 >> 7     // & 0x10 0=> query, 1=> response
	res.OpCode = header[2:3][0] & 120 >> 3 // & 0x78
	// aa, tc, rd, ra
	res.AA = header[2:3][0] & 4 >> 2
	res.TC = header[2:3][0] & 2 >> 1
	res.RD = header[2:3][0] & 1
	res.RA = header[3:4][0] & 128 >> 7
	res.Z = header[3:4][0] & 112 >> 4 // & 0x70
	res.RCode = header[3:4][0] & 15

	res.QDCount = binary.BigEndian.Uint16(header[4:6])
	res.ANCount = binary.BigEndian.Uint16(header[6:8])
	res.NSCount = binary.BigEndian.Uint16(header[8:10])
	res.ARCount = binary.BigEndian.Uint16(header[10:12])

	return res
}
