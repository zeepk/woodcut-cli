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
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

// geCmd represents the ge command
var geCmd = &cobra.Command{
	Use:   "ge",
	Short: "Check Grand Exchange data",
	Long:  `Grand Exchange data (price)`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		var itemName string = strings.Join(args, " ")
		var itemPrice = FetchItemPrice(itemName)
		// var itemID, itemPrice = FetchItemID(itemName)
		// var itemData = FetchGEData(itemID)
		// itemData.Price = itemPrice
		// PrintGEData(itemData)
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		var price = humanize.Comma(int64(itemPrice))

		fmt.Fprintln(w, "Item:", "\t", itemName)
		fmt.Fprintln(w, "Price:", "\t", price+"gp")
		w.Flush()
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Println("Time elapsed:", elapsed)
	},
}

func init() {
	rootCmd.AddCommand(geCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// geCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// geCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var GrandExchangeURL string = "https://services.runescape.com/m=itemdb_rs/api/catalogue/detail.json?item="
var ItemIDURL string = "https://api.weirdgloop.org/exchange/history/rs/latest?name="

// structs
type ItemIDResponse struct {
	ID    string `json:"id"`
	Price int    `json:"price"`
}

// type ItemCurrentPrice struct {
// 	Price string `json:"price"`
// }
type ItemDataResponse struct {
	Icon      string `json:"icon"`
	IconLarge string `json:"icon_large"`
	Name      string `json:"name"`
	Price     int
	// Current   ItemCurrentPrice `json:"current"`
}

func FetchItemPrice(itemName string) int {
	// initialize http client
	client := &http.Client{}

	// generate URL
	var url string = ItemIDURL + itemName

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
		fmt.Println("Error: item not found in the items database")
		os.Exit(1)
	}

	// TODO: learn more about IO and buffers
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// unmarshal the json into a map of items (even though we only expect 1)
	var responseObject = make(map[string]ItemIDResponse)
	json.Unmarshal(body, &responseObject)

	// iterate through the map of items (of which there should only be 1)
	// and return the item ID
	var itemPrice int
	for _, item := range responseObject {
		itemPrice = item.Price
		return itemPrice
	}

	fmt.Println("Error: item not found in the items database")
	return -1
}
func FetchItemID(itemName string) (int, int) {
	// initialize http client
	client := &http.Client{}

	// generate URL
	var url string = ItemIDURL + itemName

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
		fmt.Println("Error: item not found in the items database")
		os.Exit(1)
	}

	// TODO: learn more about IO and buffers
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// unmarshal the json into a map of items (even though we only expect 1)
	var responseObject = make(map[string]ItemIDResponse)
	json.Unmarshal(body, &responseObject)

	// iterate through the map of items (of which there should only be 1)
	// and return the item ID
	var itemID int
	var itemPrice int
	for _, item := range responseObject {
		itemID, err = strconv.Atoi(item.ID)
		itemPrice = item.Price
		return itemID, itemPrice
	}

	fmt.Println("Error: item not found in the items database")
	return -1, -1
}

func PrintGEData(item ItemDataResponse) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	var price = humanize.Comma(int64(item.Price))

	fmt.Fprintln(w, "Item:", "\t", item.Name)
	fmt.Fprintln(w, "Price:", "\t", price+"gp")
	w.Flush()
	return
}

func FetchGEData(itemID int) ItemDataResponse {
	client := &http.Client{}
	var url string = GrandExchangeURL + strconv.Itoa(itemID)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Print(err.Error())
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print(err.Error())
	}
	if resp.StatusCode == 404 {
		fmt.Println("Error: item not found on the Grand Exchange")
		os.Exit(1)
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var responseObject = make(map[string]ItemDataResponse)
	json.Unmarshal(body, &responseObject)
	return responseObject["item"]
}
