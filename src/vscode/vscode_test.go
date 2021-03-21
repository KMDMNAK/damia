package vscode

import (
	"testing"
)

func TestUtils(t *testing.T) {
	hex := convertHex("a")
	if hex != "61" {
		t.Fail()
	}
}

func TestGetContainerInfo(t *testing.T) {
	c := getContaierInfo("test")
	if c != "7b22636f6e7461696e65724e616d65223a2274657374227d" {
		t.Fail()
		t.Log(c)
	}
}

func TestCheckDocker(t *testing.T) {
	err := checkDocker()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
