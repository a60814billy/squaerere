package main

import (
	"fmt"
	"github.com/a60814billy/squaerere/internal/dnsPacketParser"
	"log"
	"net"
	"os"
	"strings"
)

var (
	DEBUG          = false
	BIND_ADDRESS   = "127.0.0.1:53"
	ZONE_DOMAIN    = "raccoon.me"
	STATIC_DNS_MAP = map[string]string{
		"abc.com": "127.0.0.1",
	}
)

func main() {
	fmt.Printf("Starting the Simple DNS Server at %s\n", BIND_ADDRESS)
	udpAddr, err := net.ResolveUDPAddr("udp", BIND_ADDRESS)
	if err != nil {
		log.Fatalln(err)
	}

	udpServer, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer udpServer.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			log.Printf("%s\n", err.Error())
			continue
		}
		go server(udpServer, addr, buf[:n])
	}
}

func server(srv *net.UDPConn, address net.Addr, buf []byte) {
	var header [12]byte
	copy(header[:], buf[0:12])
	queryHeader := dnsPacketParser.ParseDnsQueryHeader(header)
	queryRecord := dnsPacketParser.ParseQuerySection(buf[12:])

	if DEBUG {
		printPacket(queryHeader, buf, queryRecord)
	}

	respIP := ""
	if STATIC_DNS_MAP[queryRecord.Domain] != "" {
		respIP = STATIC_DNS_MAP[queryRecord.Domain]
	} else if strings.HasSuffix(queryRecord.Domain, ZONE_DOMAIN) {
		respIP = "127.0.0.1"
	}

	// response to client
	var respPacket []byte
	respHeader := dnsPacketParser.DnsQueryHeader{
		ID:     queryHeader.ID,
		QR:     1,
		OpCode: queryHeader.OpCode,
		RCode:  0,
	}

	if len(respIP) != 0 {
		respHeader.ANCount = 1
	}
	respPacket = append(respPacket, respHeader.ToBytes()...)

	if len(respIP) != 0 {
		ans := dnsPacketParser.AnswerRecord{
			Domain:       queryRecord.Domain,
			Type:         dnsPacketParser.TYPE_A,
			Class:        dnsPacketParser.CLASS_IN,
			TTL:          30,
			ResourceData: dnsPacketParser.ParseIP(respIP),
		}
		respPacket = append(respPacket, ans.ToBytes()...)
	}

	srv.WriteTo(respPacket, address)
}

func printPacket(queryHeader dnsPacketParser.DnsQueryHeader, buf []byte, queryRecord dnsPacketParser.QueryRecord) {
	fmt.Fprintf(os.Stdout, "ID: %X (%d)\n", queryHeader.ID, queryHeader.ID)
	fmt.Fprintf(os.Stdout, "QR: %d\n", queryHeader.QR)
	fmt.Fprintf(os.Stdout, "Opcode: %d\n", queryHeader.OpCode)
	fmt.Fprintf(os.Stdout, "AA (Authoritative Answer): %d,\nTC (TrunCation): %d,\nRD (Recursion Desired): %d,\nRA (Recursion Available): %d\n",
		queryHeader.AA, queryHeader.TC, queryHeader.RD, queryHeader.RA)
	fmt.Fprintf(os.Stdout, "RCode: %X\n", queryHeader.RCode)
	fmt.Fprintf(os.Stdout, "QDCOUNT: %d\n", queryHeader.QDCount)
	fmt.Fprintf(os.Stdout, "ANCount: %d\n", queryHeader.ANCount)
	fmt.Fprintf(os.Stdout, "NSCount: %d\n", queryHeader.NSCount)
	fmt.Fprintf(os.Stdout, "ARCount: %d\n", queryHeader.ARCount)

	fmt.Fprintf(os.Stdout, "%d, %X\n", len(buf), buf)
	fmt.Fprintf(os.Stdout, "Query Domain: %s %s\n",
		dnsPacketParser.TypeString(queryRecord.Type),
		queryRecord.Domain,
	)
}
