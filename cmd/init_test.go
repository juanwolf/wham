package cmd

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var BackupDirectory = fmt.Sprintf("%s/.wham.bkp", getHomeDir())

type testFn func(*testing.T)

func preTest() error {
	fileInfo, err := os.Stat(WhamDirectory)
	os.Remove(BackupDirectory)
	if err == nil {
		if fileInfo.IsDir() {
			return os.Rename(WhamDirectory, BackupDirectory)
		}
	}
	return nil
}

func postTest() error {
	fileInfo, err := os.Stat(BackupDirectory)
	if err == nil {
		if fileInfo.IsDir() {
			err = nil
			err = os.RemoveAll(WhamDirectory)
			if err != nil {
				return err
			}

			err = os.Rename(BackupDirectory, WhamDirectory)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func initBootstrap(t *testing.T, testFunction testFn) {
	err := preTest()
	if err != nil {
		t.Error("pretest error: ", err.Error())
	} else {
		testFunction(t)
	}
	err = postTest()
	if err != nil {
		t.Error("postTest error:", err.Error())
	}
}

func TestInit(t *testing.T) {
	initBootstrap(t, func(t *testing.T) {
		err := _init()
		if err != nil {
			t.Error("An error occured, none were expected. Error: ", err.Error())
		}
		_, err = os.Stat(WhamDirectory)
		if err != nil {
			t.Error("Could not stat the directory", WhamDirectory, err.Error())
		}
		err = postTest()
		if err != nil {
			t.Error("An error occured during posTest: ", err.Error())
		}
	})
}

func TestInitTwice(t *testing.T) {
	initBootstrap(t, func(t *testing.T) {
		_init()
		err := _init()
		if err != nil {
			if !strings.Contains(err.Error(), "already") {
				t.Error("'already' expected in error message, got", err.Error(), "instead")
			}
		}
	})
}
