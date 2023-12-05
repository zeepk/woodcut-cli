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

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

type RawSkill struct {
	SkillId   int `json:"skillId"`
	Rank      int `json:"rank"`
	Level     int `json:"level"`
	Xp        int `json:"xp"`
	DayGain   int `json:"dayGain"`
	WeekGain  int `json:"weekGain"`
	MonthGain int `json:"monthGain"`
	YearGain  int `json:"yearGain"`
}

type Skill struct {
	SkillId int `json:"skillId"`
	Xp      int `json:"xp"`
	DayGain int `json:"dayGain"`
}

type GainsResponse struct {
	Success bool       `json:"success"`
	Price   string     `json:"message"`
	Skills  []RawSkill `json:"skills"`
}

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
		var skills = FetchHiscoreData(username)
		PrintHiscoreData(skills)
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

var GainsURL string = "https://www.woodcut.dev/api/rs3/player?username="
var Skills []string = []string{"Overall", "Attack", "Defence", "Strength", "Hitpoints", "Ranged", "Prayer", "Magic", "Cooking", "Woodcutting", "Fletching", "Fishing", "Firemaking", "Crafting", "Smithing", "Mining", "Herblore", "Agility", "Thieving", "Slayer", "Farming", "Runecrafting", "Hunter", "Construction", "Summoning", "Dungeoneering", "Divination", "Invention", "Archaeology", "Necromancy"}

func FetchHiscoreData(username string) []Skill {
	client := &http.Client{}
	var formattedUsername = strings.ReplaceAll(username, " ", "+")
	var url string = GainsURL + formattedUsername
	fmt.Println("Fetching data for " + formattedUsername + "...")
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

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// unmarshal json
	var responseObject GainsResponse
	json.Unmarshal(body, &responseObject)

	var skills []Skill
	for _, skill := range responseObject.Skills {
		skills = append(skills, Skill{SkillId: skill.SkillId, Xp: skill.Xp, DayGain: skill.DayGain})
	}

	return skills
}

func PrintHiscoreData(skills []Skill) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	for i := 0; i < len(skills); i++ {
		var xpString = humanize.Comma(int64(skills[i].Xp))
		var dayGainString string
		if skills[i].DayGain <= 0 {
			dayGainString = ""
		} else {
			dayGainString = "+" + humanize.Comma(int64(skills[i].DayGain)) + "xp"
		}

		fmt.Fprintln(w, Skills[i], "\t", xpString+"xp", "\t", dayGainString)
	}
	w.Flush()
	return
}
