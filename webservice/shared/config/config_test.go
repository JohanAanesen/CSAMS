package config_test

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/config"
	"testing"
)

func TestLoad(t *testing.T) {
	_, err := config.Load("../../config/config.json")

	if err != nil {
		t.Logf("error: %v\n", err)
		t.Fail()
	}

	_, err = config.Load("nofile")
	if err == nil {
		t.Logf("expected error, got none")
		t.Fail()
	}

	_, err = config.Load("../../config/database.sql")
	if err == nil {
		t.Logf("expected json.Marshal error, got none")
		t.Fail()
	}
}

func TestInitialize(t *testing.T) {
	cfg := config.Initialize("../../config/config.json")

	if cfg == nil {
		t.Error("Return pointer is nil, expected not-nil")
		t.Fail()
	}

	cfg = config.Initialize("nofile")
	if cfg != nil {
		t.Error("expected nil return value, got something")
		t.Fail()
	}
}
