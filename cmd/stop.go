package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := stop()
		se := err.(*stopError)
		if err != nil {
			se.StopExec()
		}
		fmt.Println(result)
	},
}

type stopError struct {
	message string
}

func (e *stopError) Error() string {
	return fmt.Sprintf("Error - %s", e.message)
}

func (e *stopError) StopExec() {
	fmt.Println(e.message)
	os.Exit(1)
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func stop() (string, error) {
	startTimeBytes, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "", &stopError{"Wham not started!"}
		}
	}

	err = os.Remove(tmpFile)
	if err != nil {
		errorMessage := fmt.Sprintf("%s", err)
		return "", &stopError{errorMessage}
	}

	startTime, err := time.Parse(time.RFC3339, string(startTimeBytes))
	if err != nil {
		errorMessage := fmt.Sprintf(`%[1]s corrupted, please delete this file running:
rm %[1]s`, tmpFile)
		return "", &stopError{errorMessage}
	}
	now := time.Now()
	delta := now.Sub(startTime)
	result := fmt.Sprintf("You worked for %d minutes", int64(delta.Minutes()))
	return result, nil
}
