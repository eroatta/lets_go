package search

import (
    "log"
    "sync"
)

// a map or registered matchers for searching
var matchers = make(map[string]Matcher)

// Run performs the search logic
func Run(searchTerm string) {
    // retrieve the list of feeds to search through
    feeds, err := RetrieveFeeds()
    if err != nil {
        log.Fatal(err)
    }

    // create an unbuffered channel to receive match results to display
    results := make(chan *Result)

    // setup a wait group so we can process all the feeds
    var waitGroup sync.WaitGroup

    // set the number of goroutines we need to wait for while
    // they process the individual feeds
    waitGroup.Add(len(feeds))

    // launch a goroutine for each feed to find the results
    for _, feed := range feeds {
        // retrieve a matcher for the search
        matcher, exists := matchers[feed.Type]
        if !exists {
            matcher = matchers["default"]
        }

        // launch a goroutine to perform the search
        go func (matcher Matcher, feed *Feed) {
            Match(matcher, feed, searchTerm, results)
            waitGroup.Done()
        }(matcher, feed)
    }

    // launch a goroutine to monitor when all the work is done
    go func() {
        // wait for everything to be processed
        waitGroup.Wait()

        // close the channel once we are done to signal to the Display
        // function that we can exit the program
        close(results)
    }()

    // start displaying results as they are available and return after
    // the final result is displayed
    Display(results)
}

// registers a matcher to be used by the program
func Register(feedType string, matcher Matcher) {
    if _, exists := matchers[feedType]; exists {
        log.Fatalln(feedType, "Matcher already registered")
    }

    log.Println("Register", feedType, "matcher")
    matchers[feedType] = matcher
}