package search

import (
    "encoding/json"
    "os"
)

// constant for the feeds data file name
const datafile = "github.com/eroatta/lets_go/search_engine/data/feeds.json"

// Feed contains information we need to process a feed
type Feed struct {
    Name string `json:"site"`
    URI string `json:"link"`
    Type string `json:"type"`
}

// reads and unmarshalls the feeds data file
func RetrieveFeeds() ([]*Feed, error) {
    // open the file
    file, err := os.Open(datafile)
    if err != nil {
        return nil, err
    }

    // schedule the file to be closed once the function returns
    defer file.Close()

    // decode the file into a slice of pointers to Feed values
    var feeds []*Feed
    err = json.NewDecoder(file).Decode(&feeds)

    // we don't need to check for errors, the caller can do this
    return feeds, err
}