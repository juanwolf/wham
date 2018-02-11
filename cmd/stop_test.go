package cmd

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestStopWhenNotStarted(t *testing.T) {
	initBootstrap(t, func(t *testing.T) {
		_, err := stop()
		if se, ok := err.(*stopError); ok {
			if se.message != "Wham not started!" {
				t.Error("Wham not started! error expected but got", se.message)
			}
		}
	})
}

func TestStopWhenLockFileCorrupted(t *testing.T) {
	initBootstrap(t, func(t *testing.T) {

		ioutil.WriteFile(tmpFile, []byte("test"), 0644)
		_, err := stop()
		if se, ok := err.(*stopError); ok {
			if !strings.Contains(se.message, "corrupted") {
				t.Error("'corrupted' were expected in the error message but got", se.message)
			}
		}
	})
}

func TestStop(t *testing.T) {
	initBootstrap(t, func(t *testing.T) {
		err := _init()
		if err != nil {
			t.Error("Init failed. Error:", err.Error())
		}
		resExpected := "You worked for 0 minutes"
		err = start()
		if err != nil {
			t.Error("Start failed with this error:", err)
		}
		res, err := stop()
		if err != nil {
			t.Error("Stop failed with this error", err)
		}
		if res != resExpected {
			t.Error(resExpected, "were expected but got", res)
		}
	})
}
