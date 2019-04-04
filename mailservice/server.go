package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Server struct
type Server struct {
	Hostname  string `json:"hostname"`
	UseHTTP   bool   `json:"useHTTP"`
	UseHTTPS  bool   `json:"useHTTPS"`
	HTTPPort  int    `json:"HTTPPort"`
	HTTPSPort int    `json:"HTTPSPort"`
	CertFile  string `json:"certFile"`
	KeyFile   string `json:"keyFile"`
}

//Run todo
func Run(httpHandler, httpsHandler http.Handler, srv *Server) {
	if srv.UseHTTP && srv.UseHTTPS {
		go func() {
			srv.StartHTTPS(httpsHandler)
		}()

		srv.StartHTTP(httpHandler)
	} else if srv.UseHTTP {
		srv.StartHTTP(httpHandler)
	} else if srv.UseHTTPS {
		srv.StartHTTPS(httpsHandler)
	} else {
		log.Fatalln("config file does not specify a listener to start")
	}
}

// StartHTTP server
func (s Server) StartHTTP(handler http.Handler) {
	srv := &http.Server{
		Addr:         s.HTTPAddress(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
	}

	log.Printf("listening to port %d\n", s.HTTPPort)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("http server fail to start: %v\n", err)
	}
}

// StartHTTPS server
func (s Server) StartHTTPS(handler http.Handler) {
	// See https://blog.cloudflare.com/exposing-go-on-the-internet/
	tlsConfig := &tls.Config{
		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	srv := &http.Server{
		Addr:         s.HTTPSAddress(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      handler,
	}

	log.Printf("listening to port %d\n", s.HTTPSPort)
	err := srv.ListenAndServeTLS(s.CertFile, s.KeyFile)
	if err != nil {
		log.Fatalf("https server fail to start: %v\n", err)
	}
}

// HTTPAddress for server
func (s Server) HTTPAddress() string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPPort)
}

// HTTPSAddress for server
func (s Server) HTTPSAddress() string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPSPort)
}
