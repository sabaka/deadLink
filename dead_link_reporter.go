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

	fileHelper.DoOnEachLine(pathToFile,func(line string) {
		_, parsingErr := url.ParseRequestURI(line)

		if parsingErr != nil {
			fmt.Printf("WARN: something is wrong with URL: %s\n", line)
			return
		}
		_, getErr := http.Get(line)
		if getErr != nil {
			fmt.Printf("ERROR: Following link can't be opened: %s\nError:%v\n", line, getErr)
		}
	})

	fmt.Println("Execution finished")
	// Wait for any key to prevent console closure
	bufio.NewScanner(os.Stdin).Scan()

}

func getPathToFIle() string {
	if len(os.Args) > 1 {
		return os.Args[1:][0]
	} else {
		return "links.lst"
	}
}

