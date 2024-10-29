/*
Copyright © 2024 NAM HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
// var (
//
//	l bool
//	w bool
//	c bool
//
// )
type commandFlags struct {
	l bool
	w bool
	c bool
}

var flags commandFlags

var rootCmd = &cobra.Command{
	Use:   "wc",
	Args:  cobra.MinimumNArgs(1),
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		if isDir, _ := isDirectory(filePath); isDir {
			fmt.Printf("wc: %s: read: Is a dirctory\n", filePath)
			os.Exit(1)
		}

		f, err := os.Open(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		// var output string
		if flags.l {
			fmt.Printf("%8d %s\n", getLineCount(f), filePath)
		}
		if flags.w {
			fmt.Printf("%8d %s\n", getWordCount(f), filePath)
		}
		if flags.c {
			fmt.Printf("%8d %s\n", getCharCount(f), filePath)
		}
		if (commandFlags{} == flags) {
			l := getLineCount(f)
			w := getWordCount(f)
			c := getCharCount(f)
			fmt.Printf("%8d%8d%8d %s\n", l, w, c, filePath)
		}
	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&flags.l, "line", "l", false, "Get line count of files")
	rootCmd.Flags().BoolVarP(&flags.w, "word", "w", false, "Get word count of files")
	rootCmd.Flags().BoolVarP(&flags.c, "char", "c", false, "Get word count of files")
}

func getLineCount(f *os.File) int {
	f.Seek(0, 0)
	r := bufio.NewReader(f)
	lineCount := 0
	for {
		_, isPrefix, err := r.ReadLine()
		if isPrefix {
			continue
		}
		if err != nil {
			break
		}
		lineCount++
	}
	return lineCount
}

func getWordCount(f *os.File) int {
	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}

func getCharCount(f *os.File) int {
	f.Seek(0, 0)
	reader := bufio.NewReader(f)
	charCount := 0
	for {
		_, err := reader.ReadByte()
		if err != nil {
			break
		}
		charCount++
		// fmt.Println(string(r), charCount)
	}
	return charCount
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
