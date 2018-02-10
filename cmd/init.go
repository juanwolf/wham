package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

type initError struct {
	message string
}

func (e *initError) Error() string {
	return fmt.Sprintf("Error - %s", e.message)
}

// WhamDirectory return the homedirectory where wham will store csvs
var WhamDirectory = fmt.Sprintf("%s/.wham", getHomeDir())

func getHomeDir() string {
	dir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return dir
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise your home directory",
	Long: `Initialise your home directory. All the file are stored under .wham
`,
	Run: func(cmd *cobra.Command, args []string) {
		_init()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func _init() error {
	fileInfo, err := os.Stat(WhamDirectory)
	if err == nil && fileInfo.IsDir() {
		return &initError{"Wham already initialised!"}
	}

	err = os.Mkdir(WhamDirectory, os.ModePerm)
	if err != nil {
		errorMsg := fmt.Sprintf("Impossible to create %s directory, error: %s", WhamDirectory, err.Error())
		return &initError{errorMsg}
	}

	configFilePath := fmt.Sprintf("%s/config.yaml", WhamDirectory)
	_, err = os.Create(configFilePath)
	if err != nil {
		errorMsg := fmt.Sprintf("Impossible to create %s", configFilePath)
		return &initError{errorMsg}
	}

	return nil
}
