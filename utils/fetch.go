package utils

import (
	"fmt"
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
