package dnsPacketParser

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

var typeMap map[uint16]string

func init() {
	typeMap = map[uint16]string{
		1:   "A",
		2:   "NS",
		3:   "MD",
		4:   "MF",
		5:   "CNAME",
		6:   "SOA",
		7:   "MB",
		8:   "MG",
		9:   "MR",
		10:  "NULL",
		11:  "WKS",
		12:  "PTR",
		15:  "MX",
		16:  "TXT",
		252: "AXFR",
		253: "MAILB",
		254: "MAILA",
		255: "ALL",
	}
}

func TypeString(t uint16) string {
	return typeMap[t]
}
