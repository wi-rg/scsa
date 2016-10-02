package helper

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
)

func FetchUrl(url string) string {

	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s", err)
		return err.Error()
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			return err.Error()
		}
		return string(contents)
	}
}

func GetNs(t html.Token) (ok bool, abbrev string, ns string, typeof string) {
	// Iterate over all of the Token's attributes until we find an "xmlns"
	for _, a := range t.Attr {
		if a.Key[0:5] == "xmlns" {
			abbrev = a.Key
			ns = a.Val
			ok = true
		}

		if a.Key == "typeof" {
			typeof = a.Val
			ok = true
		}

	}

	// "bare" return will return the variables (ok, div) as defined in
	// the function definition
	return
}

func GetSpanProperty(t html.Token) (ok bool, span string) {
	// Iterate over all of the Token's attributes until we find an "xmlns"
	for _, a := range t.Attr {
		if a.Key == "property" {
			span = a.Val
			ok = true
		}
	}

	return
}
