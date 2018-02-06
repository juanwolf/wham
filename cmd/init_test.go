package cmd

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

const BACKUP_FOLDER = "~/.wham.bkp"

func preTest() {
	fileInfo, err := os.Stat("~/.wham")
	if fileInfo != nil {
		if fileInfo.IsDir() {
			os.Rename("~/.wham", BACKUP_FOLDER)
		}
	}
}

func postTest() {
	fileInfo, err := os.Stat(BACKUP_FOLDER)
	if fileInfo != nil {
		if fileInfo.IsDir() {
			os.Remove("~/.wham")
			os.Rename(BACKUP_FOLDER, "~/.wham")
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

func TestInitTwice(t *testing.T) {
	_init()
	err := _init()
	if err != nil {
		if strings.Contains(err.Error(), "already") {
			t.Error("'already' expected in error message, got", err.Error(), "instead")
		}
	}
}

func TestInit(t *testing.T) {
	err := _init()
	if err != nil {
		t.Error("An error occured, none were expected. Error: ", err.Error())
	}
	fileInfo, err := os.Stat("~/.wham")
	if fileInfo == nil {
		t.Error("Could not stat the directory ~/.wham, Not found, or not accessible")
	}
	now := time.Now()
	onCallDB := fmt.Sprintf("~/.wham/oncall_%d_%d.csv", now.Month(), now.Year())
	fileInfo, err = os.Stat(onCallDB)
	if fileInfo == nil {
		t.Error("Could not stat on the DB file:", onCallDB, "is not found or not accessible")
	}
}
