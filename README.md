# tls-cert-test
Simple HTTP server to test TLS protocol versions and certificates.

## Getting started
1. Download a pre-compiled binary from the release [page](https://github.com/dassump/tls-cert-test/releases).
2. Run `tls-cert-test --help`

```shell
$ tls-cert-test --help
tls-cert-test (v1.0.0)

Simple HTTP server to test TLS protocol versions and certificates
https://github.com/dassump/tls-cert-test

Usage of tls-cert-test:
  -key string
        Certificate KEY file (default "certificate.key")
  -pem string
        Certificate PEM file (default "certificate.pem")
  -port uint
        first TCP port to listen, 4 sequential ports required (default 8000)
```

## Examples

### Default
```shell
$ tls-cert-test
2022/10/28 16:28:18 (:8001/TLS1.1) Starting with certificate certificate.pem and key certificate.key
2022/10/28 16:28:18 (:8003/TLS1.3) Starting with certificate certificate.pem and key certificate.key
2022/10/28 16:28:18 (:8000/TLS1.0) Starting with certificate certificate.pem and key certificate.key
2022/10/28 16:28:18 (:8002/TLS1.2) Starting with certificate certificate.pem and key certificate.key
```

### Custom TCP port
```shell
$ tls-cert-test -port 10000
2022/10/28 16:29:15 (:10000/TLS1.0) Starting with certificate certificate.pem and key certificate.key
2022/10/28 16:29:15 (:10002/TLS1.2) Starting with certificate certificate.pem and key certificate.key
2022/10/28 16:29:15 (:10001/TLS1.1) Starting with certificate certificate.pem and key certificate.key
2022/10/28 16:29:15 (:10003/TLS1.3) Starting with certificate certificate.pem and key certificate.key
```

### Custom certificate/key files
```shell
$ tls-cert-test -pem mycert.pem -key mykey.key 
2022/10/28 16:31:52 (:8001/TLS1.1) Starting with certificate mycert.pem and key mykey.key
2022/10/28 16:31:52 (:8000/TLS1.0) Starting with certificate mycert.pem and key mykey.key
2022/10/28 16:31:52 (:8002/TLS1.2) Starting with certificate mycert.pem and key mykey.key
2022/10/28 16:31:52 (:8003/TLS1.3) Starting with certificate mycert.pem and key mykey.key
```

## Tests

### Server
```shell
$ make cert
Generating a 2048 bit RSA private key
...................................................................................+++++
........+++++
writing new private key to 'certificate.key'
-----

$ tls-cert-test
2022/10/28 16:28:18 (:8001/TLS1.1) Starting with certificate certificate.pem and key certificate.key
2022/10/28 16:28:18 (:8003/TLS1.3) Starting with certificate certificate.pem and key certificate.key
2022/10/28 16:28:18 (:8000/TLS1.0) Starting with certificate certificate.pem and key certificate.key
2022/10/28 16:28:18 (:8002/TLS1.2) Starting with certificate certificate.pem and key certificate.key
```

### Client Openssl
```shell
$ openssl version
LibreSSL 3.3.6
```

### Client TLS1.0
```shell
$ echo -n "GET / HTTP/1.0\nHost: tls-cert-test.local\n\n" | openssl s_client -connect localhost:8000 -quiet -tls1
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, CN = tls-cert-test.local
verify error:num=18:self signed certificate
verify return:1
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, CN = tls-cert-test.local
verify return:1
HTTP/1.0 200 OK
Date: Fri, 28 Oct 2022 19:54:04 GMT
Content-Length: 93
Content-Type: text/plain; charset=utf-8

Protocol-Version: TLS1.0 | Cerificate-PEM: certificate.pem | Certificate-KEY: certificate.key
```

### Client TLS1.1
```shell
$ echo -n "GET / HTTP/1.0\nHost: tls-cert-test.local\n\n" | openssl s_client -connect localhost:8001 -quiet -tls1_1
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, CN = tls-cert-test.local
verify error:num=18:self signed certificate
verify return:1
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, CN = tls-cert-test.local
verify return:1
HTTP/1.0 200 OK
Date: Fri, 28 Oct 2022 19:57:40 GMT
Content-Length: 93
Content-Type: text/plain; charset=utf-8

Protocol-Version: TLS1.1 | Cerificate-PEM: certificate.pem | Certificate-KEY: certificate.key
```

### Client TLS1.2
```shell
$ echo -n "GET / HTTP/1.0\nHost: tls-cert-test.local\n\n" | openssl s_client -connect localhost:8002 -quiet -tls1_2
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, CN = tls-cert-test.local
verify error:num=18:self signed certificate
verify return:1
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, CN = tls-cert-test.local
verify return:1
HTTP/1.0 200 OK
Date: Fri, 28 Oct 2022 19:59:14 GMT
Content-Length: 93
Content-Type: text/plain; charset=utf-8

Protocol-Version: TLS1.2 | Cerificate-PEM: certificate.pem | Certificate-KEY: certificate.key
```


### Client TLS1.3
```shell
$ echo -n "GET / HTTP/1.0\nHost: tls-cert-test.local\n\n" | openssl s_client -connect localhost:8003 -quiet -tls1_3
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, CN = tls-cert-test.local
verify error:num=18:self signed certificate
verify return:1
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, CN = tls-cert-test.local
verify return:1
HTTP/1.0 200 OK
Date: Fri, 28 Oct 2022 19:59:24 GMT
Content-Length: 93
Content-Type: text/plain; charset=utf-8

Protocol-Version: TLS1.3 | Cerificate-PEM: certificate.pem | Certificate-KEY: certificate.key
```

### Errors
Sending request to port 8000/TLSv1 using a client that only supports TLSv1.2.

```shell
$ echo -n "GET / HTTP/1.0\nHost: tls-cert-test.local\n\n" | openssl s_client -connect localhost:8000 -quiet -tls1_2
8819942656:error:1404B42E:SSL routines:ST_CONNECT:tlsv1 alert protocol version:/AppleInternal/Library/BuildRoots/810eba08-405a-11ed-86e9-6af958a02716/Library/Caches/com.apple.xbs/Sources/libressl/libressl-3.3/ssl/tls13_lib.c:151:
```

The event is logged to the server's standard output.

```shell
2022/10/28 17:05:50 (:8000/TLS1.0) http: TLS handshake error from [::1]:56744: tls: received record with version 303 when expecting version 301
```

Request on port 8003/TLSv1.3 with TLSv1.1 client.

```shell
$ echo -n "GET / HTTP/1.0\nHost: tls-cert-test.local\n\n" | openssl s_client -connect localhost:8003 -quiet -tls1_1
8819942656:error:1404B42E:SSL routines:ST_CONNECT:tlsv1 alert protocol version:/AppleInternal/Library/BuildRoots/810eba08-405a-11ed-86e9-6af958a02716/Library/Caches/com.apple.xbs/Sources/libressl/libressl-3.3/ssl/tls13_lib.c:129:SSL alert number 70
```

```
2022/10/28 17:13:56 (:8003/TLS1.3) http: TLS handshake error from [::1]:56889: tls: client offered only unsupported versions: [302 301]
```

## Contributing
Bug reports and pull requests are welcome on GitHub at https://github.com/dassump/tls-cert-test.