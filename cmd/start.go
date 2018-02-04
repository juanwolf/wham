package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type startError struct {
	message string
}

func (e *startError) Error() {
	fmt.Println("Error -", e.message)
	os.Exit(1)
}

const tmpFile = "/tmp/wham.lock"

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the timer",
	Long:  `Start to record your oncall overtime`,
	Run: func(cmd *cobra.Command, args []string) {
		err := start()
		if err != nil {
			err.Error()
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start() *startError {
	_, err := os.Stat(tmpFile)
	if !os.IsNotExist(err) {
		return &startError{"Wham already started!"}
	}

	startTime := time.Now()
	startTimeBytes, err := startTime.MarshalText()
	if err != nil {
		return &startError{err.Error()}
	}
	err = ioutil.WriteFile(tmpFile, startTimeBytes, 0644)
	if err != nil {
		return &startError{err.Error()}
	}
	return nil
}
