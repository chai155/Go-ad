package main

import (
	"testing"
)

func Test_md5hash(t *testing.T) {
	testUrl := "http://adjust.com"
	actual := getMD5Hash(testUrl)
	expected := "b53f3f2ec2e7e01d9e1130baac274a90"
	if actual != expected {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}

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
