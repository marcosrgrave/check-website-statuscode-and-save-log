package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"go-read-write-files/src"
)

func main() {
	urls, err := src.ReadFile("websites.txt")
	if err != nil {
		panic(err)
	}

	for _, url := range urls {
		fmt.Println("Checking: ", url)

		statusCode, err := checkURL200StatusCode(url)
		if err != nil {
			panic(err)
		}

		err = registerURLStatusCodeInLogFile(url, statusCode)
		if err != nil {
			panic(err)
		}
	}

	logs, err := ioutil.ReadFile("log.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(logs))
}

func checkURL200StatusCode(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode >= 300 {
		return 0, &src.StatusCodeError{}
	}

	return resp.StatusCode, nil
}

func registerURLStatusCodeInLogFile(url string, statusCode int) error {
	logMessage := fmt.Sprintf("[%s] URL: %s - Status Code: %d", time.Now().UTC().Format("2006-01-02 15:04:05"), url, statusCode)
	err := src.WriteToLogFile("log.txt", logMessage)
	if err != nil {
		return err
	}
	return nil
}
