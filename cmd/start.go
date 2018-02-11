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

func (e *startError) Error() string {
	return fmt.Sprintf("Error - %s", e.message)
}

func (e *startError) Terminate() {
}

const tmpFile = "/tmp/wham.lock"

// DBFilePrefix contains the prefix for the "db" files created
const DBFilePrefix = "oncall"

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the timer",
	Long:  `Start to record your oncall overtime`,
	Run: func(cmd *cobra.Command, args []string) {
		err := start()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

// Create the CSV 'db' for the current month
func createCSV() (*os.File, error) {
	thisMonth := time.Now().Month()
	thisYear := time.Now().Year()
	dbFilePath := fmt.Sprintf("%s/%s_%d_%d.csv", WhamDirectory, DBFilePrefix, thisMonth, thisYear)
	dbFile, err := os.Create(dbFilePath)
	if err != nil {
		return dbFile, &startError{err.Error()}
	}
	return dbFile, nil
}

func start() error {
	_, err := os.Stat(WhamDirectory)
	if err != nil {
		return &startError{`Wham not initialized! Run instead:
wham init`}
	}

	_, err = os.Stat(tmpFile)
	if err == nil {
		return &startError{"Wham already started!"}
	}

	err = nil

	_, sErr := createCSV()
	if sErr != nil {
		return sErr
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
