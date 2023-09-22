package Proxy

import "fmt"

// IContentFetcher is the interface for fetching content from a URL.
type IContentFetcher interface {
	FetchContent(url string) string
}

// RealContentFetcher is the real implementation of the content fetcher.
type RealContentFetcher struct{}

// FetchContent retrieves content from the specified URL.
func (rcf *RealContentFetcher) FetchContent(url string) string {
	return fmt.Sprintf("%s is being fetched!", url)
}

// ContentFetcherProxy is the proxy for the content fetcher.
type ContentFetcherProxy struct {
	cache map[string]IContentFetcher
}

// NewContentFetcherProxy creates a new instance of ContentFetcherProxy.
func NewContentFetcherProxy() IContentFetcher {
	return &ContentFetcherProxy{cache: make(map[string]IContentFetcher)}
}

// FetchContent fetches content from the URL, either by retrieving it from the cache or
// creating a new RealContentFetcher if not cached.
func (cfp *ContentFetcherProxy) FetchContent(url string) string {
	if cfp.cache[url] != nil {
		return fmt.Sprintf("%s has been already fetched!", url)
	}
	cfp.cache[url] = &RealContentFetcher{}
	return cfp.cache[url].FetchContent(url)
}
