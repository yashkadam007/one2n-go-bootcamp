/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type commandFlags struct {
	o string
	i bool
}

var flags commandFlags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grep",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var linesContainingText []string
		searchText := args[0]
		if len(args) == 1 {
			var output []string
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), args[0]) {
					output = append(output, scanner.Text())
				}
			}
			for _, v := range output {
				fmt.Println(v)
			}
			return
		}
		filePath := args[1]
		if isDir, _ := isDirectory(filePath); isDir {
			root := filePath
			var paths []string
			err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() {
					paths = append(paths, path)
				}
				return nil
			})
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range paths {
				fmt.Println(v)
			}
			return
		}
		linesContainingText, err := SearchInFile(searchText, filePath)
		if flags.o != "" {
			writeOutputToFile(flags.o, linesContainingText)
			return
		}
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		for _, v := range linesContainingText {
			fmt.Println(v)
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

func writeOutputToFile(filePath string, result []string) {
	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, v := range result {
		_, err := writer.WriteString(v + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	writer.Flush()
}

func SearchInFile(searchText, filePath string) ([]string, error) {
	var result []string
	if isDir, _ := isDirectory(filePath); isDir {
		err := fmt.Errorf("grep: " + filePath + ": Is a directory" + "\n")
		return nil, err
	}
	f, err := os.Open(filePath)
	if err != nil {
		err = fmt.Errorf("grep:" + strings.Replace(err.Error(), "open", "", 1) + "\n")
		return nil, err
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if flags.i {
			isPresent := caseInsensitiveSearch(scanner.Text(), searchText)
			if isPresent {
				result = append(result, scanner.Text())
			}
		} else {

			isPresent := caseSensitiveSearch(scanner.Text(), searchText)
			if isPresent {
				result = append(result, scanner.Text())
			}
		}
	}
	return result, nil
}

func caseSensitiveSearch(line string, searchText string) bool {
	return strings.Contains(line, searchText)
}

func caseInsensitiveSearch(line string, searchText string) bool {
	return strings.Contains(strings.ToLower(line), strings.ToLower(searchText))
}

func isDirectory(filePath string) (bool, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.grep.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&flags.o, "output", "o", "", "output to file")
	rootCmd.Flags().BoolVarP(&flags.i, "i", "i", false, "perform case insensitive search")
}
