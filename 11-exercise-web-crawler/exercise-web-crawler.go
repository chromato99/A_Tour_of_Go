package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, record fetchRecord) {
	var wg sync.WaitGroup
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	record.Visit(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if !record.CheckVisited(u) {
			wg.Add(1)
			go func(u string) {
				defer wg.Done()
				Crawl(u, depth-1, fetcher, record)
			}(u)
		}
	}
	wg.Wait()
	return
}

func main() {
	record := fetchRecord{urls: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, record)
}

// type for store visited urls
type fetchRecord struct {
	mu   sync.Mutex
	urls map[string]bool
}

// function to log visit
func (r fetchRecord) Visit(url string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.urls[url] = true
}

// function to check if it has been visited before
func (r fetchRecord) CheckVisited(url string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.urls[url]
	return ok
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
