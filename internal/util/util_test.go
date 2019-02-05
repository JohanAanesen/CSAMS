package util

import (
	"os"
	"testing"
)

func TestGetPort(t *testing.T) {
	port := GetPort()

	if port == "" {
		t.Errorf("No port return from 'GetPort()', excpected something, got nothing.")
	}

	penv := os.Getenv("PORT")

	if port != penv {
		t.Errorf("Port returned from 'GetPort()' is not the same as environment port. Expected: %v, Got: %v", penv, port)
	}
}

func TestLoadMenuConfig(t *testing.T) {
	menu := LoadMenuConfig("../../configs/menu/dashboard.json")

	if len(menu.Items) == 0 {
		t.Errorf("No menu items loaded from \"dashbaord.json\". Expected more then 0, got: %v", len(menu.Items))
	}

	for _, item := range menu.Items {
		if item.Name == "" || item.Href == "" {
			t.Errorf("No data inside 'Name' or 'Href' in MenuItem: %v, %v", item.Name, item.Href)
		}
	}
}
