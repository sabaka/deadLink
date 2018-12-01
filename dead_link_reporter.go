package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	fmt.Println("Execution Started")
	pathToFile := getPathToFIle()

	fileHandler, err := os.Open(pathToFile)
	check(err)
	defer fileHandler.Close()

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		link := scanner.Text()
		_, parsingErr := url.ParseRequestURI(link)

		if parsingErr != nil {
			fmt.Printf("WARN: something is wrong with URL: %s\n", link)
			continue
		}
		_, getErr := http.Get(link)
		if getErr != nil {
			fmt.Printf("ERROR: Following link can't be opened: %s\nError:%v\n", link, getErr)
		}
	}

	fmt.Println("Execution finished")
	bufio.NewScanner(os.Stdin).Scan()


}

func getPathToFIle() string {
	if len(os.Args) > 1 {
		return os.Args[1:][0]
	} else {
		return "links.lst"
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
