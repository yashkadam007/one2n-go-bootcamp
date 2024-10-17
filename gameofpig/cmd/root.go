/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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
		p1s, p2s, err := validateArguments(args)
		if err != nil {
			fmt.Println(err)
		}
		totalGames := 10
		p1Win := 0
		for i := 0; i < totalGames; i++ {
			winner := playPig(p1s, p2s)
			if winner == "p1" {
				p1Win += 1
			}
		}
		fmt.Println("Player 1 won", p1Win, "games")
	},
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
		fmt.Println("Player 1 final Score is ", p1Score, "and Player 2 final score is ", p2Score)
		return "p1"
	} else {
		fmt.Println("Player 1 final Score is ", p1Score, "and Player 2 final score is ", p2Score)
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
		fmt.Println("The die rolled", roll)
		if roll == 1 {
			return 0
		}
		score += roll
	}
	return score
}

func validateArguments(args []string) (p1Strategy int, p2Strategy int, err error) {
	p1Strategy, err1 := strconv.Atoi(args[0])
	p2Strategy, err2 := strconv.Atoi(args[1])
	if err1 != nil && err2 != nil {
		return 0, 0, errors.New("invalid arguments")
	}
	return p1Strategy, p2Strategy, nil
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
