package main

import (
	"fmt"
)

/*
Exercise: 
In this exercise you'll use Go's concurrency features to parallelize a web crawler.

Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use! 

--------

ORIGINAL

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

*/


type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel: make multiple goroutines, but calling one in the range ended after only one call
	// TODO: Don't fetch the same URL twice
	//     seems like add them to a mutex-controlled set, but if set declared in this method it gets over-written, and there are no args to pass a set,
	//     or should goroutines use a channel
	// This implementation doesn't do either
	
	fmt.Println("\nfetching: " + url)
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("    error: ", err)
		return
	}
	fmt.Printf("    found urls: %s, and body: %q\n", urls, body)
	for _, u := range urls {
		if _, exists := set[u]; !exists {
		    set[u] = struct{}{}
		    fmt.Println("        added to set, crawling ",  u)
		    Crawl(u, depth-1, fetcher)
		} else {
			fmt.Println("        already crawled ", u)
		}
	}

	fmt.Println("unique urls: ", set)
	return
}

set := make(map[string]struct{})

func Run_tour_10_ex_crawler() {

	Crawl("https://golang.org/", 4, fetcher)
}

// helper struct and methods

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
