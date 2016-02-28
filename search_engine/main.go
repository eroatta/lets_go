package main

import (
    "os"
    "log"

    _ "github.com/eroatta/lets_go/search_engine/matchers"
    "github.com/eroatta/lets_go/search_engine/search"
)

// called prior to the main function
func init() {
    // change the device for logging to stdout
    log.SetOutput(os.Stdout)
}

// main is the entry point for the program
func main() {
    // perform the search for the specified term
    search.Run("president")
}