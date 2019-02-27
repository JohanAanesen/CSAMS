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
}
