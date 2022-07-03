/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var ShowHistory bool

// vosCmd represents the vos command
var vosCmd = &cobra.Command{
	Use:   "vos",
	Short: "Voice of Seren clans",
	Long:  `Returns Voice of Seren data for Prifddinas`,
	Run: func(cmd *cobra.Command, args []string) {
		if ShowHistory {
			resp := FetchVOSHistory()
			PrintVOSHistory(resp)
			os.Exit(0)
		}
		resp := FetchVOS()
		PrintVOS(resp)
	},
}

func init() {
	// setting manual help command without shorthand so we can use -h for ShowHistory
	vosCmd.PersistentFlags().BoolP("help", "", false, "help info for this command")
	vosCmd.Flags().BoolVarP(&ShowHistory, "history", "h", false, "Show recent history of active VOS clans")
	rootCmd.AddCommand(vosCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var VOSUrl string = "https://api.weirdgloop.org/runescape/vos"
var VOSHistoryUrl string = "https://api.weirdgloop.org/runescape/vos/history"

// structs
type VOSResponse struct {
	Timestamp string `json:"timestamp"`
	District1 string `json:"district1"`
	District2 string `json:"district2"`
}
type VOSHistoryResponse struct {
	History []VOSResponse `json:"data"`
}

func FetchVOS() VOSResponse {
	// initialize http client
	client := &http.Client{}

	// generate URL
	var url string = VOSUrl

	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	// execute request
	resp, err := client.Do(req)
	// defer ~= await the line above
	defer resp.Body.Close()
	if err != nil {
		fmt.Print(err.Error())
	}
	if resp.StatusCode == 404 {
		fmt.Println("Error: unable to fetch VoS clans")
		os.Exit(1)
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var responseObject VOSResponse
	json.Unmarshal(body, &responseObject)

	return responseObject
}

func PrintVOS(resp VOSResponse) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	var timestampString string = strings.Replace(resp.Timestamp, "T", " at ", -1)

	fmt.Fprintln(w, "Clans:", "\t", resp.District1, " and ", resp.District2)
	fmt.Fprintln(w, "As of:", "\t", strings.Replace(timestampString, ":00.000Z", " game time (GMT)", -1))
	w.Flush()
	return
}

func FetchVOSHistory() VOSHistoryResponse {
	// initialize http client
	client := &http.Client{}

	// generate URL
	var url string = VOSHistoryUrl

	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	// execute request
	resp, err := client.Do(req)
	// defer ~= await the line above
	defer resp.Body.Close()
	if err != nil {
		fmt.Print(err.Error())
	}
	if resp.StatusCode == 404 {
		fmt.Println("Error: unable to fetch VoS history")
		os.Exit(1)
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var responseObject VOSHistoryResponse
	json.Unmarshal(body, &responseObject)

	return responseObject
}

func PrintVOSHistory(resp VOSHistoryResponse) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	for _, item := range resp.History {
		PrintVOS(item)
		fmt.Println("----------")
	}
	w.Flush()
	return
}
