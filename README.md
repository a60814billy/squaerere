# Simple DNS server written by go language

This project is a simple DNS server POC. It use map to define DNS record.
Currency not supports recursive query and multi-queries.

## How to use

### Build server
```
$ make
$ sudo ./dist/dns-server
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

