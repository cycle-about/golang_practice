package main

import (
	"fmt"
	"sync"
)

/*

Exercise: Web Crawler
In this exercise you'll use Go's concurrency features to parallelize a web crawler.
Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.
Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!

*/

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel: goroutines
	// TODO: Don't fetch the same URL twice: hint seems to imply a map of fetched urls, must be concurrent
	// This implementation doesn't do either:
	defer wg.Done()
	fmt.Printf("\nCRAWLING url: %s, depth: %d\n", url, depth)
	if depth <= 0 {
		return
	}
	val := globalMap.Value(url)
	if (val > 0) {
        return
    }
	body, urls, err := fetcher.Fetch(url)
	globalMap.Inc(url)  // add to map of crawled items
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("    found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, wg)
	}
	return
}

/*
Current idea: make a structure with a mutex
Question: why do the instructions say to track visited urls in a map, seems like slice would be sufficient?
*/

// SafeMap is safe to use concurrently.
type SafeMap struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeMap) Inc(key string) {
	c.mu.Lock()  // Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value for the given key.
func (c *SafeMap) Value(key string) int {
	c.mu.Lock()  // Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	val := c.v[key]
	fmt.Println("    value in map:", val)
	return val
}

// Declare a global variable
var globalMap SafeMap

func Run_tour_10_ex_crawler() {
	var wg sync.WaitGroup
	globalMap = SafeMap{v: make(map[string]int)}
	wg.Add(1)
	go Crawl("https://golang.org/", 4, populatedFetcher, &wg)
	wg.Add(1)
	go Crawl("https://golang.org/pkg/os/", 4, populatedFetcher, &wg)
	// check at the end that urls were all only crawled (incremented) once
	wg.Wait() // Wait here for all goroutines to complete
	fmt.Println("\nFinal map:", globalMap)
}

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
var populatedFetcher = fakeFetcher{
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
