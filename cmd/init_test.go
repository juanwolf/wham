package cmd

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var BackupDirectory = fmt.Sprintf("%s/.wham.bkp", getHomeDir())

func preTest() {
	fileInfo, err := os.Stat(WhamDirectory)
	if err == nil {
		if fileInfo.IsDir() {
			os.Rename(WhamDirectory, BackupDirectory)
		}
	}
}

func postTest() {
	fileInfo, err := os.Stat(BackupDirectory)
	if err == nil {
		if fileInfo.IsDir() {
			os.Remove(WhamDirectory)
			os.Rename(BackupDirectory, WhamDirectory)
		}
	}
}

// Redefining general testing behavior
// Credit: https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go
func TestMain(m *testing.M) {
	preTest()
	retCode := m.Run()
	postTest()
	os.Exit(retCode)
}

func TestInit(t *testing.T) {
	err := _init()
	if err != nil {
		t.Error("An error occured, none were expected. Error: ", err.Error())
	}
	_, err = os.Stat(WhamDirectory)
	if err != nil {
		t.Error("Could not stat the directory", WhamDirectory, err.Error())
	}
}

func TestInitTwice(t *testing.T) {
	_init()
	err := _init()
	if err != nil {
		if !strings.Contains(err.Error(), "already") {
			t.Error("'already' expected in error message, got", err.Error(), "instead")
		}
	}
}
