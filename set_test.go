package setenv

import (
	"os"
	"testing"
)

func TestSetEnv(t *testing.T) {
	SetEnv()

	value, ok := os.LookupEnv("REMOTE_ADDRESS")

	if ok == false {
		t.Log("REMOTE_ADDRESS environment variable should be set")
		t.Fail()
	}

	if value != "http://localhost" {
		t.Log("Incorrect environment variable value set")
		t.Fail()
	}
}

func TestReadFile(t *testing.T) {
	_, err := readFile(".env")

	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestReadFile2(t *testing.T) {
	_, err := readFile(".ajah")

	if err == nil {
		t.Log(err.Error())
		t.Fail()
	}
}
