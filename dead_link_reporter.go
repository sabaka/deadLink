package main

import (
	"bufio"
	"fmt"
	"github.com/sabaka/fileHelper"
	"net/http"
	"net/url"
	"os"
)

func main() {

	fmt.Println("Execution Started")
	pathToFile := getPathToFIle()

	fileHelper.DoOnEachLine(pathToFile, checkLink)

	fmt.Println("Execution finished")
	// Wait for any key to prevent console closure
	bufio.NewScanner(os.Stdin).Scan()

}

func checkLink(link string) {
	if validLink(link) {
		_, getErr := http.Get(link)
		if getErr != nil {
			fmt.Printf("ERROR: Following link can't be opened: %s\nError:%v\n", link, getErr)
		}
	}
}

func validLink(link string) bool {
	_, parsingErr := url.ParseRequestURI(link)

	if parsingErr != nil {
		fmt.Printf("WARN: something is wrong with URL: %s\n", link)
		return false
	}
	return true
}

func getPathToFIle() string {
	if len(os.Args) > 1 {
		return os.Args[1:][0]
	} else {
		return "links.lst"
	}
}

