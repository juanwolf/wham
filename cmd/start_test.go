package cmd

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
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
}

func TestStartTwice(t *testing.T) {
	stop() // Just in case
	start()
	err := start()
	if err == nil {
		t.Error("A startError were expected, but none returned")
	}
	if !strings.Contains(err.message, "already started") {
		t.Error("already started were expected in the error message but found", err.message)
	}

}
