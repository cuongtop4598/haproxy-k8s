package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	url := "http://192.168.39.203:32036"

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		go func() {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println("Request creation error:", err)
				return
			}

			req.Host = "nginx.local"

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Request error:", err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Read error:", err)
				return
			}

			fmt.Printf("Status: %s, Body length: %d\n", resp.Status, len(body))
		}()
	}
}
