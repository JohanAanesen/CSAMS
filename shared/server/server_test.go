package server_test

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/server"
	"testing"
)

func Init() *server.Server {
	return &server.Server{
		Hostname:  "localhost",
		UseHTTP:   true,
		UseHTTPS:  false,
		HTTPPort:  8089,
		HTTPSPort: 4433,
		CertFile:  "",
		KeyFile:   "",
	}
}

func TestRun(t *testing.T) {

}

func TestServer_HTTPAddress(t *testing.T) {

}

func TestServer_HTTPSAddress(t *testing.T) {

}

func TestServer_StartHTTP(t *testing.T) {

}

func TestServer_StartHTTPS(t *testing.T) {

}
