package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsValidUrl(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"http://example.com", true},
		{"https://gobyexample.com.ru/testing", true},
		{"invalidurl", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := isValidUrl(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected isValidUrl(%s) to be %v, but got %v", tc.input, tc.expected, actual)
			}
		})
	}
}

func TestRedirectTo(t *testing.T) {
	req, err := http.NewRequest("GET", "/dsdg", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(redirectTo)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %v, but got %v", http.StatusOK, status)
	}
}
