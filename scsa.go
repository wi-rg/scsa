package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	fetch "github.com/wi-rg/scsa/utils"
	"os"
	"strings"
)

func main() {

	var cmdRun = &cobra.Command{
		Use:   "run [file.ifa]",
		Short: "Execute IfA commands in [file.ifa]",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Reading and executing: " + strings.Join(args, " "))

			fileInput, _ := os.Open(args[0])

			scanner := bufio.NewScanner(fileInput)
			for scanner.Scan() {
				fmt.Println(scanner.Text()) // Println will add back the final '\n'

				// see src/wit/net/fetch.go
				contents := fetch.FetchUrl(scanner.Text())
				fmt.Printf("%s\n", string(contents))

			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}

		},
	}

	var cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Shows current running hayy version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("scsa - version 1.0.0")
		},
	}
	var rootCmd = &cobra.Command{Use: "scsa"}
	rootCmd.AddCommand(cmdRun, cmdVersion)
	rootCmd.Execute()
}
