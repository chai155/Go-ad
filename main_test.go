package main

import (
	"testing"
)

func Test_validateParallelArgs(t *testing.T) {
	invalidArg := 0
	actual := validateParallelArg(invalidArg)
	expected := false
	if actual != expected {
		t.Errorf("expected '%t' but got '%t'", expected, actual)
	}

	negativeArg := -1
	actual = validateParallelArg(negativeArg)
	expected = false
	if actual != expected {
		t.Errorf("expected '%t' but got '%t'", expected, actual)
	}

	validArg := 3
	actual = validateParallelArg(validArg)
	expected = true
	if actual != expected {
		t.Errorf("expected '%t' but got '%t'", expected, actual)
	}
}

func Test_validateDomainArgs(t *testing.T) {
	noArg := []string{}
	actual := validateDomainArgs(noArg)
	expected := false
	if actual != expected {
		t.Errorf("expected '%t' but got '%t'", expected, actual)
	}

	validArg := []string{"google.com", "yahoo.com"}
	actual = validateDomainArgs(validArg)
	expected = true
	if actual != expected {
		t.Errorf("expected '%t' but got '%t'", expected, actual)
	}
}

func Test_validateURL(t *testing.T) {
	invalidArg := "google.com"
	actual := validateURL(invalidArg)
	expected := false
	if actual != expected {
		t.Errorf("expected '%t' but got '%t'", expected, actual)
	}

	validArg := "http://google.com"
	actual = validateURL(validArg)
	expected = true
	if actual != expected {
		t.Errorf("expected '%t' but got '%t'", expected, actual)
	}
}

func Test_sendHTTPRequests(t *testing.T) {
	invalidURL := "http://sdfgsdg"
	actual, _ := getHashFromHTTPRequest(invalidURL)
	expected := 0
	hashLen := len(actual)
	if hashLen != expected {
		t.Errorf("expected '%d' but got '%d'", expected, hashLen)
	}

	validUrl := "http://google.com"
	actual, _ = getHashFromHTTPRequest(validUrl)
	// MD5 processes an arbitrary-length message into a fixed-length output of 128 bits, typically represented as a sequence of 32 hexadecimal digits
	expected = 32
	hashLen = len(actual)
	if hashLen != expected {
		t.Errorf("expected '%d' but got '%d'", expected, hashLen)
	}
}
