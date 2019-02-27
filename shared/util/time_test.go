package util_test

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/shared/util"
	"testing"
	"time"
)

func TestDatetimeLocalToRFC3339(t *testing.T) {
	input := "2019-02-27T10:41"

	output, err := util.DatetimeLocalToRFC3339(input)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if output.IsZero() {
		t.Error("Time should not be zero")
		t.Fail()
	}

	emptyInput := ""

	output, err = util.DatetimeLocalToRFC3339(emptyInput)
	if err == nil {
		t.Error("Expected error!")
		t.Fail()
	}

	if !output.IsZero() {
		t.Error("Expected output to be zero!")
		t.Fail()
	}

	shortInput := "2019-02-27"

	output, err = util.DatetimeLocalToRFC3339(shortInput)
	if err == nil {
		t.Error("Expected error!")
		t.Fail()
	}

	if !output.IsZero() {
		t.Error("Expected output to be zero!")
		t.Fail()
	}
}

func TestDatetimeLocalToRFC33392(t *testing.T) {
	input, _ := time.Parse(time.RFC3339, "2010-10-10T10:10:10.371Z")

	output := util.GoToHTMLDatetimeLocal(input)
	if output == "" {
		t.Error("Output was empty, expected not empty")
		t.Fail()
	}

	zeroInput, _ := time.Parse(time.RFC3339, "")

	output = util.GoToHTMLDatetimeLocal(zeroInput)
	if output != "" {
		t.Error("Expected empty output, got something")
		t.Fail()
	}

	lowInput, _ := time.Parse(time.RFC3339, "2010-01-01T01:01:01.371Z")

	output = util.GoToHTMLDatetimeLocal(lowInput)
	if output == "" {
		t.Error("Output was empty, expected not empty")
		t.Fail()
	}
}
