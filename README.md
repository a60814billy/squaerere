# Simple DNS server written by go language

[![hackmd-github-sync-badge](https://stage.hackmd.io/SA0Sg_MmT7qBCczdk4jBsg/badge)](https://stage.hackmd.io/SA0Sg_MmT7qBCczdk4jBsg)

The project is a POC of DNS server. Currency only support few dns server feature.

- Listen on UDP port 53
- Support one Server Query
- No Query Type check

## How to use

### requirements

- go 1.14 
- make

### Build server
```
$ make
go build -o ./dist/squaerere ./cmd/main.go
$ make run
```

### Test

```
$ nslookup
> server 127.0.0.1
Default server: 127.0.0.1
Address: 127.0.0.1#53
> test.com
Server:		127.0.0.1
Address:	127.0.0.1#53

Non-authoritative answer:
*** Can't find test.com: No answer

> abc.raccoon.me
Server:		127.0.0.1
Address:	127.0.0.1#53

Non-authoritative answer:
Name:	abc.raccoon.me
Address: 127.0.0.1
```

## References

- [DNS related RFCs list](https://www.statdns.com/rfc/)
- [RFC 1035 - DOMAIN NAMES - IMPLEMENTATION AND SPECIFICATION](https://tools.ietf.org/html/rfc1035)