package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "" {
			fmt.Println("Message cannot be empty. Please try again.")
			continue
		}

		url := "http://localhost:8080"
		resp, err := http.Post(url, "text/plain", strings.NewReader(message))
		if err != nil {
			fmt.Println("Failed to send request:", err)
			return
		}
		defer resp.Body.Close()

		headers := resp.Header
		date := headers.Get("Date")
		fmt.Println(date)
	}
}
