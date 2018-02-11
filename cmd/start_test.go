package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestStartNotInit(t *testing.T) {
	initBootstrap(t, func(t *testing.T) {
		_, berr := os.Stat(WhamDirectory)
		if berr == nil {
			t.Error("WhamDirectory existing, aborting...")
		}
		err := start()
		if err == nil {
			t.Error("An error were expected but none returned")
		} else if !strings.Contains(err.Error(), "not initialized") {
			t.Error("not initialized were expected in the error message. got", err.Error())
		}
	})
}

func TestStart(t *testing.T) {
	initBootstrap(t, func(t *testing.T) {
		_init()
		stop() // Just in case
		start()
		now := time.Now()

		startTimeBytes, _ := ioutil.ReadFile(tmpFile)
		startTime, _ := time.Parse(time.RFC3339, string(startTimeBytes))
		if int(now.Sub(startTime).Minutes()) != 0 {
			nowText, _ := now.MarshalText()
			t.Error("Expected 0 but got", now.Sub(startTime).Minutes(), "\nstart calculated:", string(nowText),
				"\nstart written", startTime)
		}

		onCallDB := fmt.Sprintf("%s/oncall_%d_%d.csv", WhamDirectory, now.Month(), now.Year())
		fileInfo, err := os.Stat(onCallDB)
		if fileInfo == nil {
			t.Error("Could not stat on the DB file:", onCallDB, ". Error:\n", err.Error())
		}

	})
}

func TestStartTwice(t *testing.T) {
	initBootstrap(t, func(t *testing.T) {
		_init()
		stop() // Just in case
		start()
		err := start()
		if err == nil {
			t.Error("A startError were expected, but none returned")
			return
		}
		if !strings.Contains(err.Error(), "already started") {
			t.Error("already started were expected in the error message but found", err.Error())
		}
	})
}
