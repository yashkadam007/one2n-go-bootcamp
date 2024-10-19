/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pig",
	Args:  cobra.ExactArgs(2),
	Short: "Pig is a two-player game played with a 6 sided die",
	Long: `Simulate a game of pig between two computer agents and decide the strategy that each of the agent uses
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// validate args to be int
		p1s, p2s := args[0], args[1]

		if strings.Contains(args[0], "-") && strings.Contains(args[1], "-") {
			variableVsvariableStrategey(p1s, p2s)
			return
		}
		p1HoldScore, err := strconv.Atoi(p1s)
		if err != nil || p1HoldScore > 100 || p1HoldScore < 1 {
			fmt.Println("invalid strategy for Player 1")
		}
		if strings.Contains(p2s, "-") {
			fixedVsVariableStrategy(p1HoldScore, p2s)
			return
		}
		p2HoldScore, err := strconv.Atoi(p2s)
		if err != nil || p2HoldScore > 100 || p2HoldScore < 1 {
			fmt.Println("invalid strategy for player 2")
		}
		p1Win, totalGames := fixedVsfixedStrategy(p1HoldScore, p2HoldScore)

		losses := totalGames - p1Win
		fmt.Printf(
			"Holding at %d vs Holding at %d: wins: %d/%d (%.2f%%), losses: %d/%d (%.2f%%)\n",
			p1HoldScore, p2HoldScore, p1Win, totalGames, (float64(p1Win)/float64(totalGames))*100,
			losses, totalGames, (float64(losses)/float64(totalGames))*100,
		)
	},
}

func variableVsvariableStrategey(p1s, p2s string) {
	p1start, p1end := getRange(p1s)
	p2start, p2end := getRange(p2s)
	for i := p1start; i <= p1end; i++ {
		p1TotalWin := 0
		totalGames := 0
		for j := p2start; j <= p2end; j++ {
			if i == j {
				continue
			}
			p1Win, totalGame := fixedVsfixedStrategy(i, j)
			p1TotalWin += p1Win
			totalGames += totalGame
		}
		losses := totalGames - p1TotalWin
		fmt.Printf("Result: Wins, losess staying at k = %d: %d/%d (%.2f%%), %d/%d (%.2f%%)\n",
			i, p1TotalWin, totalGames, (float32(p1TotalWin)/float32(totalGames))*100,
			losses, totalGames, (float32(losses)/float32(totalGames))*100,
		)
	}
}

func fixedVsVariableStrategy(p1s int, p2s string) {
	start, end := getRange(p2s)
	for i := start; i <= end; i++ {
		if p1s == i {
			continue
		}
		p1Win, totalGames := fixedVsfixedStrategy(p1s, i)
		losses := totalGames - p1Win
		fmt.Printf(
			"Holding at %d vs Holding at %d: wins: %d/%d (%.2f%%), losses: %d/%d (%.2f%%)\n",
			p1s, i, p1Win, totalGames, (float64(p1Win)/float64(totalGames))*100,
			losses, totalGames, (float64(losses)/float64(totalGames))*100,
		)

	}
}

func getRange(rangeArg string) (int, int) {
	rangeArray := strings.Split(rangeArg, "-")
	if len(rangeArray) > 2 {
		fmt.Println("incorrect range, should be of format 1-100")
	}
	start, err := strconv.Atoi(rangeArray[0])
	if err != nil || start < 1 || start > 100 {
		fmt.Println("invalid strategy range")
	}
	end, err := strconv.Atoi(rangeArray[1])
	if err != nil || end < 1 || end > 100 {
		fmt.Println("invalid strategy range")
	}
	if start > end {
		fmt.Println("invalid strategy range")
	}
	return start, end
}

func fixedVsfixedStrategy(p1s, p2s int) (int, int) {
	totalGames := 10
	p1Win := 0
	for i := 0; i < totalGames; i++ {
		winner := playPig(p1s, p2s)
		if winner == "p1" {
			p1Win += 1
		}
	}
	return p1Win, totalGames
}

func playPig(p1s, p2s int) string {
	p1Score := 0
	p2Score := 0
	p1Turn := true
	for p1Score < 100 && p2Score < 100 {
		if p1Turn {
			p1CurrScore := playTurn(p1s)
			if p1CurrScore == 0 {
				p1Turn = false
				continue
			} else {
				p1Score += p1CurrScore
				p1Turn = false
				continue
			}
		} else {
			p2CurrScore := playTurn(p2s)
			if p2CurrScore == 0 {
				p1Turn = true
				continue
			} else {
				p2Score += p2CurrScore
				p1Turn = true
				continue
			}
		}
	}
	if p1Score >= 100 {
		return "p1"
	} else {
		return "p2"
	}
}

func rollDie() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

func playTurn(playerStrategy int) int {
	score := 0
	for score < playerStrategy {
		roll := rollDie()
		if roll == 1 {
			return 0
		}
		score += roll
	}
	return score
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gameofpig.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
