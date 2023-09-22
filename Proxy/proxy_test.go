package Proxy

import (
	"fmt"
	"testing"
)

func TestContentFetcherProxy(t *testing.T) {
	// Create a ContentFetcherProxy instance.
	proxy := NewContentFetcherProxy()

	// Test fetching content for a new URL.
	url1 := "https://example.com/page1"
	result1 := proxy.FetchContent(url1)

	// Ensure that the result is as expected.
	expectedResult1 := fmt.Sprintf("%s is being fetched!", url1)
	if result1 != expectedResult1 {
		t.Errorf("Expected result for URL1: %s, but got: %s", expectedResult1, result1)
	}

	// Test fetching content for the same URL again.
	result2 := proxy.FetchContent(url1)

	// Ensure that the result is as expected (should return cached result).
	if result2 != fmt.Sprintf("%s has been already fetched!", url1) {
		t.Errorf("Expected cached result for URL1, but got: %s", result2)
	}

	// Test fetching content for a different URL.
	url2 := "https://example.com/page2"
	result3 := proxy.FetchContent(url2)

	// Ensure that the result is as expected.
	expectedResult2 := fmt.Sprintf("%s is being fetched!", url2)
	if result3 != expectedResult2 {
		t.Errorf("Expected result for URL2: %s, but got: %s", expectedResult2, result3)
	}

	// Test fetching content for the same URL again (should return cached result).
	result4 := proxy.FetchContent(url2)

	// Ensure that the result is as expected (should return cached result).
	if result4 != fmt.Sprintf("%s has been already fetched!", url2) {
		t.Errorf("Expected cached result for URL2, but got: %s", result4)
	}
}
