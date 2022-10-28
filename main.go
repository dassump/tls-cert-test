package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	app         = "tls-cert-test"
	version     = "unknown"
	description = "Simple HTTP server to test TLS protocol versions and certificates"
	site        = "https://github.com/dassump/tls-cert-test"
	info        = "%s (%s)\n\n%s\n%s\n\n"
	usage       = "Usage of %s:\n"

	pem  = flag.String("pem", "certificate.pem", "Certificate PEM file")
	key  = flag.String("key", "certificate.key", "Certificate KEY file")
	port = flag.Uint("port", 8000, "first TCP port to listen, 4 sequential ports required")

	protocol = map[uint16]string{
		tls.VersionTLS10: "TLS1.0",
		tls.VersionTLS11: "TLS1.1",
		tls.VersionTLS12: "TLS1.2",
		tls.VersionTLS13: "TLS1.3",
	}
)

const (
	tls10 = iota
	tls11
	tls12
	tls13
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), info, app, version, description, site)
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {
	go server(fmt.Sprintf(":%d", *port+tls10), *pem, *key, tls.VersionTLS10)
	go server(fmt.Sprintf(":%d", *port+tls11), *pem, *key, tls.VersionTLS11)
	go server(fmt.Sprintf(":%d", *port+tls12), *pem, *key, tls.VersionTLS12)
	go server(fmt.Sprintf(":%d", *port+tls13), *pem, *key, tls.VersionTLS13)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}

func server(addr, pem, key string, version uint16) {
	logger := log.New(os.Stderr, fmt.Sprintf("(%s/%s) ", addr, protocol[version]), log.Lmsgprefix+log.LstdFlags)
	logger.Printf("Starting with certificate %s and key %s", pem, key)

	server := &http.Server{
		Addr:      addr,
		TLSConfig: &tls.Config{MinVersion: version, MaxVersion: version, InsecureSkipVerify: true},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w,
				"Protocol-Version: %s | Cerificate-PEM: %s | Certificate-KEY: %s",
				protocol[version], pem, key,
			)
		}),
		ErrorLog: logger,
	}

	logger.Fatal(server.ListenAndServeTLS(pem, key))
}
