# cert

[![Build Status](https://travis-ci.org/genkiroid/cert.svg?branch=master)](https://travis-ci.org/genkiroid/cert)

Get server's certificate information.

## Installation

For Mac it provide Homebrew integration.
Please install like following.

```sh
$ brew tap genkiroid/homebrew-cert
$ brew install cert
```

For other platforms, Precompiled binaries for released versions are available in the [releases](https://github.com/genkiroid/cert/releases) page.

Or `go get`.

```sh
$ go get github.com/genkiroid/cert/...
```

## Usage

Give domain names as arguments.

```sh
$ cert github.com google.co.jp
DomainName:    github.com
IP:            140.82.121.3
Issuer:        DigiCert TLS Hybrid ECC SHA384 2020 CA1
NotBefore:     2021-03-12 01:00:00 +0100 CET
NotAfter:      2022-03-24 00:59:59 +0100 CET
CommonName:    github.com
MinTLSVersion: TLS 1.2
SANs:          [github.com www.github.com]
Error:         

DomainName:    google.co.jp
IP:            216.58.213.195
Issuer:        GTS CA 1O1
NotBefore:     2021-03-11 16:02:15 +0100 CET
NotAfter:      2021-06-03 17:02:14 +0200 CEST
CommonName:    *.google.co.jp
MinTLSVersion: TLS 1.0
SANs:          [*.google.co.jp google.co.jp]
Error:     

```

You can specify port number.
So you can get server certificate information of not only web server but also *mail server and others*.

```sh
$ cert github.com google.co.jp:443 imap.gmail.com:993
DomainName:    github.com
IP:            140.82.121.3
Issuer:        DigiCert TLS Hybrid ECC SHA384 2020 CA1
NotBefore:     2021-03-12 01:00:00 +0100 CET
NotAfter:      2022-03-24 00:59:59 +0100 CET
CommonName:    github.com
MinTLSVersion: TLS 1.2
SANs:          [github.com www.github.com]
Error:         

DomainName:    google.co.jp
IP:            216.58.213.195
Issuer:        GTS CA 1O1
NotBefore:     2021-03-11 16:02:15 +0100 CET
NotAfter:      2021-06-03 17:02:14 +0200 CEST
CommonName:    *.google.co.jp
MinTLSVersion: TLS 1.0
SANs:          [*.google.co.jp google.co.jp]
Error:         

DomainName:    imap.gmail.com
IP:            142.251.9.109
Issuer:        GTS CA 1O1
NotBefore:     2021-02-23 16:41:18 +0100 CET
NotAfter:      2021-05-18 17:41:17 +0200 CEST
CommonName:    imap.gmail.com
MinTLSVersion: TLS 1.0
SANs:          [imap.gmail.com]
Error:         

```

## Options

```sh
$ cert --help
Usage of cert:
  -c string
        Specify cipher suite. Refer to https://golang.org/pkg/crypto/tls/#pkg-constants for supported cipher suites.
  -cipher string
        Specify cipher suite. Refer to https://golang.org/pkg/crypto/tls/#pkg-constants for supported cipher suites.
  -f string
        Output format. md: as markdown, json: as JSON.  (default "simple table")
  -format string
        Output format. md: as markdown, json: as JSON.  (default "simple table")
  -k    Skip verification of server's certificate chain and host name.
  -s int
        Timeout seconds. (default 3)
  -skip-verify
        Skip verification of server's certificate chain and host name.
  -t string
        Output format as Go template string or Go template file path.
  -template string
        Output format as Go template string or Go template file path.
  -timeout int
        Timeout seconds. (default 3)
  -u    Use UTC to represent NotBefore and NotAfter.
  -utc
        Use UTC to represent NotBefore and NotAfter.
  -v    Show version.
  -version
        Show version.
```

### Output as JSON

Use `cert -f json`.

```sh
$ cert -f json github.com | jq .
[
  {
    "domainName": "github.com",
    "ip": "140.82.121.3",
    "issuer": "DigiCert TLS Hybrid ECC SHA384 2020 CA1",
    "commonName": "github.com",
    "sans": [
      "github.com",
      "www.github.com"
    ],
    "notBefore": "2021-03-12 01:00:00 +0100 CET",
    "notAfter": "2022-03-24 00:59:59 +0100 CET",
    "minimumTLSVersion": "TLS 1.2",
    "error": ""
  }
]
```

### Output as Markdown

Use `cert -f md`.

```sh
$ cert -f md github.com
DomainName | IP | Issuer | NotBefore | NotAfter | CN | SANs | Error
--- | --- | --- | --- | --- | --- | --- | ---
github.com | 192.30.255.113 | DigiCert SHA2 Extended Validation Server CA | 2016-03-10 09:00:00 +0900 JST | 2018-05-17 21:00:00 +0900 JST | github.com | github.com<br/>www.github.com<br/> |
```

DomainName | IP | Issuer | NotBefore | NotAfter | CN | SANs | Error
--- | --- | --- | --- | --- | --- | --- | ---
github.com | 192.30.255.113 | DigiCert SHA2 Extended Validation Server CA | 2016-03-10 09:00:00 +0900 JST | 2018-05-17 21:00:00 +0900 JST | github.com | github.com<br/>www.github.com<br/> |

### Specify output format by Go template

Use `cert -t`.

By direct string.

```sh
$ cert -t "{{range .}}{{.Issuer}}{{end}}" github.com
DigiCert SHA2 Extended Validation Server CA
```

By template file.

```sh
$ cat /tmp/cert_templ
{{range .}}{{range .CertChain}}Issuer: {{.Issuer.CommonName}}
{{end}}{{end}}
$
$ cert -t /tmp/cert_templ github.com
Issuer: DigiCert SHA2 Extended Validation Server CA
Issuer: DigiCert High Assurance EV Root CA

```

### Specify cipher suite

see https://github.com/genkiroid/cert/issues/13

You can specify cipher suite.
As a result, you can get the information of each certificate.

Note that the issuers are different in the following example.

```sh
# Get information of the certificate using RSA signature/public key algorithm.
$ cert -cipher TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305 cloudflaressl.com
DomainName: cloudflaressl.com
IP:         104.20.47.142
Issuer:     COMODO RSA Domain Validation Secure Server CA 2
NotBefore:  2019-08-23 09:00:00 +0900 JST
NotAfter:   2020-03-01 08:59:59 +0900 JST
CommonName: ssl509631.cloudflaressl.com
SANs:       [ssl509631.cloudflaressl.com *.cloudflaressl.com cloudflaressl.com]
Error:

# Get information of the certificate using ECDSA signature/public key algorithm.
$ cert -cipher TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305 cloudflaressl.com
DomainName: cloudflaressl.com
IP:         104.20.48.142
Issuer:     COMODO ECC Domain Validation Secure Server CA 2
NotBefore:  2019-08-23 09:00:00 +0900 JST
NotAfter:   2020-03-01 08:59:59 +0900 JST
CommonName: ssl509632.cloudflaressl.com
SANs:       [ssl509632.cloudflaressl.com *.cloudflaressl.com cloudflaressl.com]
Error:

```

**If you specify a cipher suite, the maximum TLS version used is limited to TLS1.2. This is because if the server supports TLS1.3, the specified cipher suite is ignored and communication is performed using TLS1.3. This eliminates the meaning of specifying a cipher suite and confuses us. This specification will change when the cipher suite for Tls1.3 becomes configurable in Go.**

## License

[MIT](https://github.com/genkiroid/cert/blob/master/LICENSE)

## Author

[genkiroid](https://github.com/genkiroid)

