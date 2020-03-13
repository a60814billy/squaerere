package dnsPacketParser

import (
	"encoding/binary"
	"strconv"
	"strings"
)

type QueryRecord struct {
	Domain string
	Type   uint16
	Class  uint16
}

type AnswerRecord struct {
	Domain             string
	Type               uint16
	Class              uint16
	TTL                uint32
	resourceDataLength uint16
	ResourceData       []byte
}

func (A *AnswerRecord) ToBytes() []byte {

	A.resourceDataLength = uint16(len(A.ResourceData))

	var res []byte
	res = append(res, domainToBytes(A.Domain)...)
	res = append(res, uint8(A.Type>>4), uint8(A.Type))
	res = append(res, uint8(A.Class>>4), uint8(A.Class))
	res = append(res, uint8(A.TTL>>24), uint8(A.TTL>>16), uint8(A.TTL>>8), uint8(A.TTL))
	res = append(res, uint8(A.resourceDataLength>>4), uint8(A.resourceDataLength))
	res = append(res, A.ResourceData...)
	return res
}

func domainToBytes(domain string) []byte {
	var r []byte
	labels := strings.Split(domain, ".")

	for _, label := range labels {
		r = append(r, uint8(len(label)))
		r = append(r, []byte(label)...)
	}
	r = append(r, 0)
	return r
}

func ParseQuerySection(body []byte) QueryRecord {
	var domain []string
	var idx int = 0

	for {
		labelLength := int(body[idx : idx+1][0])
		idx += 1
		if labelLength == 0 {
			break
		}
		labelName := body[idx : idx+labelLength]
		domain = append(domain, string(labelName))
		idx += labelLength
	}

	domainName := strings.Join(domain, ".")

	res := QueryRecord{
		Domain: domainName,
		Type:   binary.BigEndian.Uint16(body[idx : idx+2]),
		Class:  binary.BigEndian.Uint16(body[idx+2 : idx+4]),
	}

	return res
}

func ParseIP(s string) []byte {
	var r []byte
	ipSeg := strings.Split(s, ".")

	for _, seg := range ipSeg {
		i, err := strconv.ParseUint(seg, 10, 8)
		if err != nil {
			return nil
		}
		r = append(r, byte(i))
	}
	return r
}
