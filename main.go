package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudflare/cfssl/certinfo"
	"gopkg.in/yaml.v2"
)

var (
	app         = "tls-cert-test"
	version     = "unknown"
	description = "Simple HTTP server to test TLS protocol versions and certificates"
	site        = "https://github.com/dassump/tls-cert-test"

	ip   = flag.String("ip", "0.0.0.0", "IP address to listen")
	port = flag.Uint("port", 8000, fmt.Sprintf("First TCP port to listen, %d sequential ports required", len(protocols)))
	pem  = flag.String("pem", "certificate.pem", "Certificate PEM file")
	key  = flag.String("key", "certificate.key", "Certificate KEY file")

	protocols = map[uint16]protocol{
		tls.VersionTLS10: TLS10,
		tls.VersionTLS11: TLS11,
		tls.VersionTLS12: TLS12,
		tls.VersionTLS13: TLS13,
	}

	chain       tls.Certificate
	certificate *x509.Certificate
	serialized  []byte
	err         error
)

type protocol uint

const (
	TLS10 protocol = iota
	TLS11
	TLS12
	TLS13
)

func (p protocol) String() string {
	switch p {
	case TLS10:
		return "TLS1.0"
	case TLS11:
		return "TLS1.1"
	case TLS12:
		return "TLS1.2"
	case TLS13:
		return "TLS1.3"
	default:
		return "unknown"
	}
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s (%s)\n\n%s\n%s\n\n", app, version, description, site)
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {
	logger := log.New(os.Stderr, "(CERTIFICATE) ", log.Lmsgprefix+log.LstdFlags)

	chain, err = tls.LoadX509KeyPair(*pem, *key)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("loaded from %s and %s with chain size %d", *pem, *key, len(chain.Certificate))

	certificate, err = x509.ParseCertificate(chain.Certificate[0])
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("parsed %s, issued by %s, valid until %s", certificate.Subject.CommonName, certificate.Issuer.CommonName, certificate.NotAfter)

	if _, err = certificate.Verify(x509.VerifyOptions{}); err != nil {
		logger.Print(err)
	} else {
		logger.Print("everything seems to be ok")
	}

	info := certinfo.ParseCertificate(certificate)

	serialized, err = yaml.Marshal(map[string]any{"certificate": info})
	if err != nil {
		logger.Fatal(err)
	}

	for ver, proto := range protocols {
		go server(fmt.Sprintf("%s:%d", *ip, *port+uint(proto)), ver)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}

func server(addr string, ver uint16) {
	logger := log.New(os.Stderr, fmt.Sprintf("(%s/%s) ", addr, protocols[ver]), log.Lmsgprefix+log.LstdFlags)

	server := &http.Server{
		Addr: addr,
		TLSConfig: &tls.Config{
			MinVersion:         ver,
			MaxVersion:         ver,
			Certificates:       []tls.Certificate{chain},
			InsecureSkipVerify: true,
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf("successful request from %s: %s %s", r.RemoteAddr, r.Method, r.RequestURI)
			fmt.Fprintf(w, "protocol: %s\n%s\n", protocols[ver], serialized)
		}),
		ErrorLog: logger,
	}

	logger.Print("server started")
	logger.Fatal(server.ListenAndServeTLS("", ""))
}
