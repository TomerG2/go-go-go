package main

import (
	"net/http"
	"sync"
	"testing"
)

func TestQuoteAPI_Sync(t *testing.T) {
	getUser("a1")
	getUserSub("a1")
	generateQuote("a1")
}

func TestQuoteAPI_Goroutines(t *testing.T) {
	go getUser("a1")
	getUserSub("a1")
	generateQuote("a1")
}

func TestQuoteAPI_WaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	// &wg passing a pointer to wg
	wg.Add(1)
	go getUserWait("a1", &wg)
	wg.Add(1)
	go getUserSubWait("a1", &wg)
	wg.Add(1)
	go generateQuoteWait("a1", &wg)

	wg.Wait()
}

func TestQuoteAPI_Load(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go callAPI("http://localhost:8090/quote", &wg)
	}
	wg.Wait()
}

func TestQuoteAPI_Load_v2(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go callAPI("http://localhost:8090/quote/v2", &wg)
	}
	wg.Wait()
}

func TestQuoteAPI_Load_v3(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go callAPI("http://localhost:8090/quote/v3", &wg)
	}
	wg.Wait()
}

func callAPI(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	http.Get(url)
}
