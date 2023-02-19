package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var lastSeen = make(map[string]time.Time)


func main() {
    port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":"+port, nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	// Send status code 200 and an empty response body
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(""))


    // Get the x-real-ip and x-service-name headers from the request
    xRealIP := r.Header.Get("x-real-ip")
    xServiceName := r.Header.Get("x-service-name")

    // Check if x-real-ip has been seen in the last 5 minutes
    if !checkSeen(xRealIP) {

		locateAndReport(xRealIP, xServiceName)

    }

}


func checkSeen(ip string) bool {
    lastTime, ok := lastSeen[ip]
    if ok && time.Since(lastTime) <= 5*time.Minute {
        // The IP has been seen in the last 5 minutes
        return true
    }
    // Update the last seen time for the IP
    lastSeen[ip] = time.Now()
    return false
}

func locateAndReport(ip string, service string) {
    url := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,country,regionName,city", ip)
    resp, err := http.Get(url)
    if err != nil {
        // Handle the error
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        // Handle the error
        return
    }

    var result map[string]interface{}
    if err := json.Unmarshal(body, &result); err != nil {
        // Handle the error
        return
    }

    if status, ok := result["status"].(string); ok && status == "success" {
        country, _ := result["country"].(string)
        regionName, _ := result["regionName"].(string)
        city, _ := result["city"].(string)

        dateString := time.Now().UTC().Format(time.RFC1123)
        logString := fmt.Sprintf("[%s] [%s] [%s] [%s] [%s] [%s]\n",
            service, dateString, country, regionName, city, ip)


        f, err := os.OpenFile("status.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            // Handle the error
            return
        }
        defer f.Close()
        if _, err := f.WriteString(logString); err != nil {
            // Handle the error
            return
        }
    }
}


