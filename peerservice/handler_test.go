package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getReaderFromPayload(payload Payload) io.Reader {
	body, _ := json.Marshal(payload)
	return bytes.NewReader(body)
}

func TestHandlers(t *testing.T) {
	Initialize()

	tests := []struct {
		name    string
		method  string
		url     string
		body    io.Reader
		handler func(w http.ResponseWriter, r *http.Request)

		expectedCode int
	}{
		{
			name:         "indexGET",
			method:       "GET",
			url:          "/",
			body:         nil,
			handler:      HandlerGET,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:   "indexPOST",
			method: "POST",
			url:    "/",
			body: getReaderFromPayload(Payload{
				Authentication: "none",
				SubmissionID:   1,
				Reviewers:      2,
			}),
			handler:      HandlerPOST,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:   "indexPOSTwAuth",
			method: "POST",
			url:    "/",
			body: getReaderFromPayload(Payload{
				Authentication: "5243980712315079823517089",
				SubmissionID:   1,
				Reviewers:      2,
			}),
			handler:      HandlerPOST,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		r, _ := http.NewRequest(test.method, test.url, test.body)
		w := httptest.NewRecorder()

		test.handler(w, r)

		t.Run(test.name, func(t *testing.T) {
			if w.Body.String() == "" {
				t.Logf("error, response body was empty")
				t.Fail()
			}

			// Check status code
			if w.Code != test.expectedCode {
				t.Logf("expected: %v, got: %v\n", test.expectedCode, w.Code)
				t.Fail()
			}
		})
	}
}
