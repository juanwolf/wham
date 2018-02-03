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
		stop()
	},
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

func stop() {
	startTimeBytes, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Wham not started!")
			os.Exit(1)
		}
	}
	err = os.Remove(tmpFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	startTime, err := time.Parse(time.RFC3339, string(startTimeBytes))

	if err != nil {
		fmt.Println(tmpFile, `corrupted, please delete this file running:
rm`, tmpFile)
		os.Exit(1)
	}
	now := time.Now()
	delta := now.Sub(startTime)
	fmt.Println("You worked for", int64(delta.Minutes()), "minutes")
}
