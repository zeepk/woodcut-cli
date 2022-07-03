/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Check stats for users",
	Long:  `Print skills (level, xp, rank) and minigames (score, rank) for RuneScape users`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// args are collected as an Array, so need to join with a space
		var username string = strings.Join(args, " ")
		// if username is more than 12 characters
		if len(username) > 12 {
			fmt.Println("Error: username must be less than 12 characters")
			os.Exit(1)
		}
		var responseString = FetchHiscoreData(username)
		PrintHiscoreData(responseString)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var HiscoresURL string = "https://secure.runescape.com/m=hiscore/index_lite.ws?player="
var Skills []string = []string{"Overall", "Attack", "Defence", "Strength", "Hitpoints", "Ranged", "Prayer", "Magic", "Cooking", "Woodcutting", "Fletching", "Fishing", "Firemaking", "Crafting", "Smithing", "Mining", "Herblore", "Agility", "Thieving", "Slayer", "Farming", "Runecrafting", "Hunter", "Construction", "Summoning", "Dungeoneering", "Divination", "Invention"}

func FetchHiscoreData(username string) string {
	client := &http.Client{}
	var url string = HiscoresURL + username
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print(err.Error())
	}
	if resp.StatusCode == 404 {
		fmt.Println("Error: username not found in official hiscores")
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseString := buf.String()
	return responseString
}

func PrintHiscoreData(responseString string) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	var responseLines []string = strings.Split(responseString, "\n")
	for i := 0; i < len(Skills); i++ {
		var responseLineItems []string = strings.Split(responseLines[i], ",")
		// [2] is the skill xp, base 10, 64bit int
		xp, err := strconv.ParseInt(responseLineItems[2], 10, 64)
		if err != nil {
			fmt.Print(err.Error())
			continue
		}

		var xpString = humanize.Comma(xp)
		fmt.Fprintln(w, Skills[i], "\t", xpString+"xp")
	}
	w.Flush()
	return
}
