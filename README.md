# tls-cert-test
Simple HTTP server to test TLS protocol versions and certificates.

## Getting started
1. Download a pre-compiled binary from the release [page](https://github.com/dassump/tls-cert-test/releases).
2. Run `tls-cert-test --help`

```shell
$ tls-cert-test --help
tls-cert-test (v1.1.0)

Simple HTTP server to test TLS protocol versions and certificates
https://github.com/dassump/tls-cert-test

Usage of tls-cert-test:
  -ip string
        IP address to listen (default "0.0.0.0")
  -key string
        Certificate KEY file (default "certificate.key")
  -pem string
        Certificate PEM file (default "certificate.pem")
  -port uint
        First TCP port to listen, 4 sequential ports required (default 8000)
```

## Examples

### Default
```shell
$ tls-cert-test
2022/12/12 18:22:07 (CERTIFICATE) loaded from certificate.pem and certificate.key with chain size 1
2022/12/12 18:22:07 (CERTIFICATE) parsed tls-cert-test.local, issued by tls-cert-test.local, valid until 2023-12-12 21:21:50 +0000 UTC
2022/12/12 18:22:07 (CERTIFICATE) x509: “tls-cert-test.local” certificate is not trusted
2022/12/12 18:22:07 (0.0.0.0:8000/TLS1.0) server started
2022/12/12 18:22:07 (0.0.0.0:8001/TLS1.1) server started
2022/12/12 18:22:07 (0.0.0.0:8002/TLS1.2) server started
2022/12/12 18:22:07 (0.0.0.0:8003/TLS1.3) server started
```

### Custom TCP port
```shell
$ tls-cert-test -port 10000
2022/12/12 18:22:07 (CERTIFICATE) loaded from certificate.pem and certificate.key with chain size 1
2022/12/12 18:22:07 (CERTIFICATE) parsed tls-cert-test.local, issued by tls-cert-test.local, valid until 2023-12-12 21:21:50 +0000 UTC
2022/12/12 18:22:07 (CERTIFICATE) x509: “tls-cert-test.local” certificate is not trusted
2022/12/12 18:22:07 (0.0.0.0:10000/TLS1.0) server started
2022/12/12 18:22:07 (0.0.0.0:10001/TLS1.1) server started
2022/12/12 18:22:07 (0.0.0.0:10002/TLS1.2) server started
2022/12/12 18:22:07 (0.0.0.0:10003/TLS1.3) server started
```

### Custom certificate/key files
```shell
$ tls-cert-test -pem mycert.pem -key mykey.key
2022/12/12 18:23:10 (CERTIFICATE) loaded from mycert.pem and mycert.key with chain size 3
2022/12/12 18:23:10 (CERTIFICATE) parsed *.mydomain.local, issued by My Root CA, valid until 2023-12-12 21:21:50 +0000 UTC
2022/12/12 18:23:10 (CERTIFICATE) everything seems to be ok
2022/12/12 18:23:10 (0.0.0.0:8000/TLS1.0) server started
2022/12/12 18:23:10 (0.0.0.0:8003/TLS1.3) server started
2022/12/12 18:23:10 (0.0.0.0:8002/TLS1.2) server started
2022/12/12 18:23:10 (0.0.0.0:8001/TLS1.1) server started
```

## Tests

### Server
```shell
$ make cert
Generating a 2048 bit RSA private key
................+++++
...................................+++++
writing new private key to 'certificate.key'
-----

$ tls-cert-test
2022/12/12 18:22:07 (CERTIFICATE) loaded from certificate.pem and certificate.key with chain size 1
2022/12/12 18:22:07 (CERTIFICATE) parsed tls-cert-test.local, issued by tls-cert-test.local, valid until 2023-12-12 21:21:50 +0000 UTC
2022/12/12 18:22:07 (CERTIFICATE) x509: “tls-cert-test.local” certificate is not trusted
2022/12/12 18:22:07 (0.0.0.0:8000/TLS1.0) server started
2022/12/12 18:22:07 (0.0.0.0:8001/TLS1.1) server started
2022/12/12 18:22:07 (0.0.0.0:8002/TLS1.2) server started
2022/12/12 18:22:07 (0.0.0.0:8003/TLS1.3) server started
```

### Client Openssl
```shell
$ openssl version
LibreSSL 3.3.6
```

### Client TLS1.0
```shell
$ echo -n "GET / HTTP/1.0\nHost: tls-cert-test.local\n\n" | openssl s_client -connect localhost:8000 -quiet -tls1
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, OU = XX, CN = tls-cert-test.local
verify error:num=18:self signed certificate
verify return:1
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, OU = XX, CN = tls-cert-test.local
verify return:1
HTTP/1.0 200 OK
Date: Mon, 12 Dec 2022 21:26:57 GMT
Content-Type: text/plain; charset=utf-8

protocol: TLS1.0
certificate:
  subject:
    commonname: tls-cert-test.local
    serialnumber: ""
    country: XX
    organization: tls-cert-test
    organizationalunit: XX
    locality: XX
    province: XX
    streetaddress: ""
    postalcode: ""
    names:
    - XX
    - XX
    - XX
    - tls-cert-test
    - XX
    - tls-cert-test.local
  issuer:
    commonname: tls-cert-test.local
    serialnumber: ""
    country: XX
    organization: tls-cert-test
    organizationalunit: XX
    locality: XX
    province: XX
    streetaddress: ""
    postalcode: ""
    names:
    - XX
    - XX
    - XX
    - tls-cert-test
    - XX
    - tls-cert-test.local
  serialnumber: "17890694716263038441"
  sans: []
  notbefore: 2022-12-12T21:21:50Z
  notafter: 2023-12-12T21:21:50Z
  signaturealgorithm: SHA256WithRSA
  aki: ""
  ski: ""
  rawpem: |
    -----BEGIN CERTIFICATE-----
    MIIDUDCCAjgCCQD4SIQOM9D96TANBgkqhkiG9w0BAQsFADBqMQswCQYDVQQGEwJY
    WDELMAkGA1UECAwCWFgxCzAJBgNVBAcMAlhYMRYwFAYDVQQKDA10bHMtY2VydC10
    ZXN0MQswCQYDVQQLDAJYWDEcMBoGA1UEAwwTdGxzLWNlcnQtdGVzdC5sb2NhbDAe
    Fw0yMjEyMTIyMTIxNTBaFw0yMzEyMTIyMTIxNTBaMGoxCzAJBgNVBAYTAlhYMQsw
    CQYDVQQIDAJYWDELMAkGA1UEBwwCWFgxFjAUBgNVBAoMDXRscy1jZXJ0LXRlc3Qx
    CzAJBgNVBAsMAlhYMRwwGgYDVQQDDBN0bHMtY2VydC10ZXN0LmxvY2FsMIIBIjAN
    BgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvI34kcstaD4P3y8xl9+O0AGVfSq1
    fYQQnI4mQSM1aY6Gf7TOO20deoBo2Cxggkf9pnEpTemGZ6YP9M+hXKoCW4nldXRG
    mzNDsGOL4AT1I5d9tcpy56asHGM+AObGRYNCYv5Vt020z62ZZYf5pnduXTIdjRn7
    5C0b5mkgkGCQEm2PGRZL3kV8vvcoROsjhemfBd+yWUYgxGW86XTQCeCT3d6rZD2k
    Oe7WEtf+tKY4uiT/EEFXI5udKfAw8O8Eb2g/aAfVs6Nv+JUR9uyykNQS1wCjOjZB
    B+gvVs/LwAAHQBiSnb2Cj27rTs2MdTALNH5t/GK5K3U0DLf3n2p6+M2psQIDAQAB
    MA0GCSqGSIb3DQEBCwUAA4IBAQAinzdxD9cVQ97URgpuH39Ke7T97+S3AK8uC2XB
    jAOcchnPpdMcY86Tff98HaqwIH6JP0a5abEdPQQ7iyLEMsc+CEN2AA4xsWI55ufC
    gwPGM4StYQvfgjTrGorl0qiiYb9qHJLOdTiGPFbiABvEQTT3PoCynJIoZsA1gbvc
    U2pEL291LD1rABHu5dtUGg3Agjbi1KdiVSv7vyRbNRHkH+C+olk/go2LeDktAsUk
    JpjC9FLXcmY/fXcJxQ+8a/iy2Tn3zbRtKgex4n5SJn6Eg3rvi3t7Zk/TF/EQnayy
    ibsii9gJwb0taf7m0ezpDNkciMvcRDjICRfZy0J4DrB/LkgO
    -----END CERTIFICATE-----
```

### Client TLS1.1
```shell
$ echo -n "GET / HTTP/1.0\nHost: tls-cert-test.local\n\n" | openssl s_client -connect localhost:8001 -quiet -tls1_1
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, OU = XX, CN = tls-cert-test.local
verify error:num=18:self signed certificate
verify return:1
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, OU = XX, CN = tls-cert-test.local
verify return:1
HTTP/1.0 200 OK
Date: Mon, 12 Dec 2022 21:28:09 GMT
Content-Type: text/plain; charset=utf-8

protocol: TLS1.1
certificate:
  subject:
    commonname: tls-cert-test.local
    serialnumber: ""
    country: XX
    organization: tls-cert-test
    organizationalunit: XX
    locality: XX
    province: XX
    streetaddress: ""
    postalcode: ""
    names:
    - XX
    - XX
    - XX
    - tls-cert-test
    - XX
    - tls-cert-test.local
  issuer:
    commonname: tls-cert-test.local
    serialnumber: ""
    country: XX
    organization: tls-cert-test
    organizationalunit: XX
    locality: XX
    province: XX
    streetaddress: ""
    postalcode: ""
    names:
    - XX
    - XX
    - XX
    - tls-cert-test
    - XX
    - tls-cert-test.local
  serialnumber: "17890694716263038441"
  sans: []
  notbefore: 2022-12-12T21:21:50Z
  notafter: 2023-12-12T21:21:50Z
  signaturealgorithm: SHA256WithRSA
  aki: ""
  ski: ""
  rawpem: |
    -----BEGIN CERTIFICATE-----
    MIIDUDCCAjgCCQD4SIQOM9D96TANBgkqhkiG9w0BAQsFADBqMQswCQYDVQQGEwJY
    WDELMAkGA1UECAwCWFgxCzAJBgNVBAcMAlhYMRYwFAYDVQQKDA10bHMtY2VydC10
    ZXN0MQswCQYDVQQLDAJYWDEcMBoGA1UEAwwTdGxzLWNlcnQtdGVzdC5sb2NhbDAe
    Fw0yMjEyMTIyMTIxNTBaFw0yMzEyMTIyMTIxNTBaMGoxCzAJBgNVBAYTAlhYMQsw
    CQYDVQQIDAJYWDELMAkGA1UEBwwCWFgxFjAUBgNVBAoMDXRscy1jZXJ0LXRlc3Qx
    CzAJBgNVBAsMAlhYMRwwGgYDVQQDDBN0bHMtY2VydC10ZXN0LmxvY2FsMIIBIjAN
    BgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvI34kcstaD4P3y8xl9+O0AGVfSq1
    fYQQnI4mQSM1aY6Gf7TOO20deoBo2Cxggkf9pnEpTemGZ6YP9M+hXKoCW4nldXRG
    mzNDsGOL4AT1I5d9tcpy56asHGM+AObGRYNCYv5Vt020z62ZZYf5pnduXTIdjRn7
    5C0b5mkgkGCQEm2PGRZL3kV8vvcoROsjhemfBd+yWUYgxGW86XTQCeCT3d6rZD2k
    Oe7WEtf+tKY4uiT/EEFXI5udKfAw8O8Eb2g/aAfVs6Nv+JUR9uyykNQS1wCjOjZB
    B+gvVs/LwAAHQBiSnb2Cj27rTs2MdTALNH5t/GK5K3U0DLf3n2p6+M2psQIDAQAB
    MA0GCSqGSIb3DQEBCwUAA4IBAQAinzdxD9cVQ97URgpuH39Ke7T97+S3AK8uC2XB
    jAOcchnPpdMcY86Tff98HaqwIH6JP0a5abEdPQQ7iyLEMsc+CEN2AA4xsWI55ufC
    gwPGM4StYQvfgjTrGorl0qiiYb9qHJLOdTiGPFbiABvEQTT3PoCynJIoZsA1gbvc
    U2pEL291LD1rABHu5dtUGg3Agjbi1KdiVSv7vyRbNRHkH+C+olk/go2LeDktAsUk
    JpjC9FLXcmY/fXcJxQ+8a/iy2Tn3zbRtKgex4n5SJn6Eg3rvi3t7Zk/TF/EQnayy
    ibsii9gJwb0taf7m0ezpDNkciMvcRDjICRfZy0J4DrB/LkgO
    -----END CERTIFICATE-----
```

### Client TLS1.2
```shell
$ echo -n "GET / HTTP/1.0\nHost: tls-cert-test.local\n\n" | openssl s_client -connect localhost:8002 -quiet -tls1_2
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, OU = XX, CN = tls-cert-test.local
verify error:num=18:self signed certificate
verify return:1
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, OU = XX, CN = tls-cert-test.local
verify return:1
HTTP/1.0 200 OK
Date: Mon, 12 Dec 2022 21:28:26 GMT
Content-Type: text/plain; charset=utf-8

protocol: TLS1.2
certificate:
  subject:
    commonname: tls-cert-test.local
    serialnumber: ""
    country: XX
    organization: tls-cert-test
    organizationalunit: XX
    locality: XX
    province: XX
    streetaddress: ""
    postalcode: ""
    names:
    - XX
    - XX
    - XX
    - tls-cert-test
    - XX
    - tls-cert-test.local
  issuer:
    commonname: tls-cert-test.local
    serialnumber: ""
    country: XX
    organization: tls-cert-test
    organizationalunit: XX
    locality: XX
    province: XX
    streetaddress: ""
    postalcode: ""
    names:
    - XX
    - XX
    - XX
    - tls-cert-test
    - XX
    - tls-cert-test.local
  serialnumber: "17890694716263038441"
  sans: []
  notbefore: 2022-12-12T21:21:50Z
  notafter: 2023-12-12T21:21:50Z
  signaturealgorithm: SHA256WithRSA
  aki: ""
  ski: ""
  rawpem: |
    -----BEGIN CERTIFICATE-----
    MIIDUDCCAjgCCQD4SIQOM9D96TANBgkqhkiG9w0BAQsFADBqMQswCQYDVQQGEwJY
    WDELMAkGA1UECAwCWFgxCzAJBgNVBAcMAlhYMRYwFAYDVQQKDA10bHMtY2VydC10
    ZXN0MQswCQYDVQQLDAJYWDEcMBoGA1UEAwwTdGxzLWNlcnQtdGVzdC5sb2NhbDAe
    Fw0yMjEyMTIyMTIxNTBaFw0yMzEyMTIyMTIxNTBaMGoxCzAJBgNVBAYTAlhYMQsw
    CQYDVQQIDAJYWDELMAkGA1UEBwwCWFgxFjAUBgNVBAoMDXRscy1jZXJ0LXRlc3Qx
    CzAJBgNVBAsMAlhYMRwwGgYDVQQDDBN0bHMtY2VydC10ZXN0LmxvY2FsMIIBIjAN
    BgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvI34kcstaD4P3y8xl9+O0AGVfSq1
    fYQQnI4mQSM1aY6Gf7TOO20deoBo2Cxggkf9pnEpTemGZ6YP9M+hXKoCW4nldXRG
    mzNDsGOL4AT1I5d9tcpy56asHGM+AObGRYNCYv5Vt020z62ZZYf5pnduXTIdjRn7
    5C0b5mkgkGCQEm2PGRZL3kV8vvcoROsjhemfBd+yWUYgxGW86XTQCeCT3d6rZD2k
    Oe7WEtf+tKY4uiT/EEFXI5udKfAw8O8Eb2g/aAfVs6Nv+JUR9uyykNQS1wCjOjZB
    B+gvVs/LwAAHQBiSnb2Cj27rTs2MdTALNH5t/GK5K3U0DLf3n2p6+M2psQIDAQAB
    MA0GCSqGSIb3DQEBCwUAA4IBAQAinzdxD9cVQ97URgpuH39Ke7T97+S3AK8uC2XB
    jAOcchnPpdMcY86Tff98HaqwIH6JP0a5abEdPQQ7iyLEMsc+CEN2AA4xsWI55ufC
    gwPGM4StYQvfgjTrGorl0qiiYb9qHJLOdTiGPFbiABvEQTT3PoCynJIoZsA1gbvc
    U2pEL291LD1rABHu5dtUGg3Agjbi1KdiVSv7vyRbNRHkH+C+olk/go2LeDktAsUk
    JpjC9FLXcmY/fXcJxQ+8a/iy2Tn3zbRtKgex4n5SJn6Eg3rvi3t7Zk/TF/EQnayy
    ibsii9gJwb0taf7m0ezpDNkciMvcRDjICRfZy0J4DrB/LkgO
    -----END CERTIFICATE-----
```


### Client TLS1.3
```shell
$ echo -n "GET / HTTP/1.0\nHost: tls-cert-test.local\n\n" | openssl s_client -connect localhost:8003 -quiet -tls1_3
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, OU = XX, CN = tls-cert-test.local
verify error:num=18:self signed certificate
verify return:1
depth=0 C = XX, ST = XX, L = XX, O = tls-cert-test, OU = XX, CN = tls-cert-test.local
verify return:1
HTTP/1.0 200 OK
Date: Mon, 12 Dec 2022 21:28:42 GMT
Content-Type: text/plain; charset=utf-8

protocol: TLS1.3
certificate:
  subject:
    commonname: tls-cert-test.local
    serialnumber: ""
    country: XX
    organization: tls-cert-test
    organizationalunit: XX
    locality: XX
    province: XX
    streetaddress: ""
    postalcode: ""
    names:
    - XX
    - XX
    - XX
    - tls-cert-test
    - XX
    - tls-cert-test.local
  issuer:
    commonname: tls-cert-test.local
    serialnumber: ""
    country: XX
    organization: tls-cert-test
    organizationalunit: XX
    locality: XX
    province: XX
    streetaddress: ""
    postalcode: ""
    names:
    - XX
    - XX
    - XX
    - tls-cert-test
    - XX
    - tls-cert-test.local
  serialnumber: "17890694716263038441"
  sans: []
  notbefore: 2022-12-12T21:21:50Z
  notafter: 2023-12-12T21:21:50Z
  signaturealgorithm: SHA256WithRSA
  aki: ""
  ski: ""
  rawpem: |
    -----BEGIN CERTIFICATE-----
    MIIDUDCCAjgCCQD4SIQOM9D96TANBgkqhkiG9w0BAQsFADBqMQswCQYDVQQGEwJY
    WDELMAkGA1UECAwCWFgxCzAJBgNVBAcMAlhYMRYwFAYDVQQKDA10bHMtY2VydC10
    ZXN0MQswCQYDVQQLDAJYWDEcMBoGA1UEAwwTdGxzLWNlcnQtdGVzdC5sb2NhbDAe
    Fw0yMjEyMTIyMTIxNTBaFw0yMzEyMTIyMTIxNTBaMGoxCzAJBgNVBAYTAlhYMQsw
    CQYDVQQIDAJYWDELMAkGA1UEBwwCWFgxFjAUBgNVBAoMDXRscy1jZXJ0LXRlc3Qx
    CzAJBgNVBAsMAlhYMRwwGgYDVQQDDBN0bHMtY2VydC10ZXN0LmxvY2FsMIIBIjAN
    BgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvI34kcstaD4P3y8xl9+O0AGVfSq1
    fYQQnI4mQSM1aY6Gf7TOO20deoBo2Cxggkf9pnEpTemGZ6YP9M+hXKoCW4nldXRG
    mzNDsGOL4AT1I5d9tcpy56asHGM+AObGRYNCYv5Vt020z62ZZYf5pnduXTIdjRn7
    5C0b5mkgkGCQEm2PGRZL3kV8vvcoROsjhemfBd+yWUYgxGW86XTQCeCT3d6rZD2k
    Oe7WEtf+tKY4uiT/EEFXI5udKfAw8O8Eb2g/aAfVs6Nv+JUR9uyykNQS1wCjOjZB
    B+gvVs/LwAAHQBiSnb2Cj27rTs2MdTALNH5t/GK5K3U0DLf3n2p6+M2psQIDAQAB
    MA0GCSqGSIb3DQEBCwUAA4IBAQAinzdxD9cVQ97URgpuH39Ke7T97+S3AK8uC2XB
    jAOcchnPpdMcY86Tff98HaqwIH6JP0a5abEdPQQ7iyLEMsc+CEN2AA4xsWI55ufC
    gwPGM4StYQvfgjTrGorl0qiiYb9qHJLOdTiGPFbiABvEQTT3PoCynJIoZsA1gbvc
    U2pEL291LD1rABHu5dtUGg3Agjbi1KdiVSv7vyRbNRHkH+C+olk/go2LeDktAsUk
    JpjC9FLXcmY/fXcJxQ+8a/iy2Tn3zbRtKgex4n5SJn6Eg3rvi3t7Zk/TF/EQnayy
    ibsii9gJwb0taf7m0ezpDNkciMvcRDjICRfZy0J4DrB/LkgO
    -----END CERTIFICATE-----
```

### Errors
Sending request to port 8000/TLSv1 using a client that only supports TLSv1.2.

```shell
$ echo -n "GET / HTTP/1.0\nHost: tls-cert-test.local\n\n" | openssl s_client -connect localhost:8000 -quiet -tls1_2
8255137024:error:1404B42E:SSL routines:ST_CONNECT:tlsv1 alert protocol version:/AppleInternal/Library/BuildRoots/810eba08-405a-11ed-86e9-6af958a02716/Library/Caches/com.apple.xbs/Sources/libressl/libressl-3.3/ssl/tls13_lib.c:151:
```

The event is logged to the server's standard output.

```shell
2022/12/12 18:28:58 (0.0.0.0:8000/TLS1.0) http: TLS handshake error from [::1]:59470: tls: received record with version 303 when expecting version 301
```

Request on port 8003/TLSv1.3 with TLSv1.1 client.

```shell
$ echo -n "GET / HTTP/1.0\nHost: tls-cert-test.local\n\n" | openssl s_client -connect localhost:8003 -quiet -tls1_1
8255137024:error:1404B42E:SSL routines:ST_CONNECT:tlsv1 alert protocol version:/AppleInternal/Library/BuildRoots/810eba08-405a-11ed-86e9-6af958a02716/Library/Caches/com.apple.xbs/Sources/libressl/libressl-3.3/ssl/tls13_lib.c:129:SSL alert number 70
```

```shell
2022/12/12 18:29:42 (0.0.0.0:8003/TLS1.3) http: TLS handshake error from [::1]:59476: tls: client offered only unsupported versions: [302 301]
```

Check client trust

```shell
$ curl -sv https://localhost:8002
*   Trying 127.0.0.1:8002...
* Connected to localhost (127.0.0.1) port 8002 (#0)
* ALPN: offers h2
* ALPN: offers http/1.1
*  CAfile: /etc/ssl/cert.pem
*  CApath: none
* (304) (OUT), TLS handshake, Client hello (1):
* (304) (IN), TLS handshake, Server hello (2):
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (OUT), TLS alert, unknown CA (560):
* SSL certificate problem: self signed certificate
* Closing connection 0
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (OUT), TLS alert, unknown CA (560):
```

```shell
2022/12/12 18:30:05 (0.0.0.0:8002/TLS1.2) http: TLS handshake error from 127.0.0.1:59482: remote error: tls: unknown certificate authority
```

## Contributing
Bug reports and pull requests are welcome on GitHub at https://github.com/dassump/tls-cert-test.