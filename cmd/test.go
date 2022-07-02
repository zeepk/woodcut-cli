/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var Verbose bool
var Source string

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test [username]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	// requires username to be entered, eg. `go run . test username -v`
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
		if Verbose {
			fmt.Println("Is Verbose")
		} else {
			fmt.Println("Not Verbose")
		}
	},
}

func init() {
	testCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "is verbose?")
	rootCmd.AddCommand(testCmd)
}
