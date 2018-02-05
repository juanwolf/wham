package cmd

import (
	"strings"
	"testing"
)

func TestInitTwice(t *testing.T) {
	_init()
	err := _init()
	if err != nil {
		if strings.Contains(err.Error(), "already") {
			t.Error("'already' expected in error message, got", err.Error(), "instead")
		}
	}
}
