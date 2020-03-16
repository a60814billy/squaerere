package dnsPacketParser

import (
	"encoding/binary"
	"errors"
)

/**
4.1.1. Header section format
The header contains the following fields:
                                    1  1  1  1  1  1
      0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                      ID                       |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                    QDCOUNT                    |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                    ANCOUNT                    |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                    NSCOUNT                    |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                    ARCOUNT                    |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
*/
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

func ParseDnsQueryHeader(header [12]byte) (*DnsQueryHeader, error) {
	res := &DnsQueryHeader{}
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

	if res.Z != 0 {
		return nil, errors.New("invalid DNS header, the Z field must be zero")
	}

	return res, nil
}
