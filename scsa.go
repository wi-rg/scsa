package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	helper "github.com/wi-rg/scsa/helper"
	html "golang.org/x/net/html"
	"os"
	"strings"
)

func main() {

	var cmdRun = &cobra.Command{
		Use:   "run [file]",
		Short: "Scrap semantic contents from list of URLs in [file]",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Scraping from file: " + strings.Join(args, " "))

			fileInput, _ := os.Open(args[0])

			scanner := bufio.NewScanner(fileInput)
			for scanner.Scan() {
				fmt.Println(scanner.Text())

				// see helper/helper.go
				contents := helper.FetchUrl(scanner.Text())
				fmt.Printf("%s\n", string(contents))
				b := bytes.NewBufferString(contents)
				z := html.NewTokenizer(b)

				for {
					tt := z.Next()

					switch {
					case tt == html.ErrorToken:
						// End of the document, we're done
						return
					case tt == html.StartTagToken:

						t := z.Token()

						// Check if the token is an <div> tag
						isDiv := t.Data == "div"
						if !isDiv {
							continue
						}

						// Extract the div value, if there is one
						ok, abbrev, xmlns, typeof := helper.GetNs(t)
						if !ok {
							continue
						}

						fmt.Println(abbrev)
						fmt.Println(xmlns)
						fmt.Println(typeof)

					}

				}

			}

			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}

		},
	}

	var cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Shows current running scsa version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("scsa - version 1.0.0")
		},
	}
	var rootCmd = &cobra.Command{Use: "scsa"}
	rootCmd.AddCommand(cmdRun, cmdVersion)
	rootCmd.Execute()
}
