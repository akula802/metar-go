package main

import (
    "fmt"
    "net/http"
    "io"
    "flag"
)

func main() {

    // Command-line flags. See: https://gobyexample.com/command-line-flags
    station_id := flag.String("station", "KBJC", "Station ID") 

    // Parse the flags
    flag.Parse()

    // Do some safety checking here on the supplied station_id

    // Construct the query URL
    var metar_query_base_url string = fmt.Sprintf("https://www.aviationweather.gov/api/data/metar?ids=%v", *station_id)

    // Make the GET request
    resp, err := http.Get(metar_query_base_url)
    if err != nil {
        fmt.Println("Unable to complete the GET request.")
    }

    // Read the response
    body, err := io.ReadAll(resp.Body)
    
    // Close the response body
    defer resp.Body.Close()

    // Print the result
    fmt.Println(string(body))
}

