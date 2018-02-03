package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the timer",
	Long:  `Start to record your oncall overtime`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

const tmpFile = "/tmp/wham.lock"

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func start() {
	_, err := os.Stat(tmpFile)
	if !os.IsNotExist(err) {
		fmt.Println("Wham already started!")
		os.Exit(1)
	}

	startTime := time.Now()
	startTimeBytes, err := startTime.MarshalText()
	check(err)
	err = ioutil.WriteFile(tmpFile, startTimeBytes, 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
