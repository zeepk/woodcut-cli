/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var IsCool bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "woodcut",
	Short: "CLI tool for RuneScape stats and other things",
	Long: `CLI tool for RuneScape stats and other things.

	Not sure what to put here yet.`,
	// Uncomment if using root command with --flags
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Print: " + strings.Join(args, " "))
	// 	if IsCool {
	// 		fmt.Println("Is IsCool, yay")
	// 	} else {
	// 		fmt.Println("Not IsCool, aw man")
	// 	}
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.woodcut.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&IsCool, "cool", "c", false, "is cool?")
}
