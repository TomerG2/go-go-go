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

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go callAPI("http://localhost:8090/quote", &wg, t)
	}
	wg.Wait()
}

func TestQuoteAPI_Load_v2(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go callAPI("http://localhost:8090/quote/v2", &wg, t)
	}
	wg.Wait()
}

func TestQuoteAPI_Load_v3(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go callAPI("http://localhost:8090/quote/v3", &wg, t)
	}
	wg.Wait()
}

func TestQuoteAPI_Load_v3_20k(t *testing.T) {
	var wg sync.WaitGroup

	for j := 0; j < 20000; j += 10000 {
		for i := 0; i < 10000; i++ {
			wg.Add(1)
			go callAPI("http://localhost:8090/quote/v3", &wg, t)
		}
		wg.Wait()
	}
}

func callAPI(url string, wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		t.Fatalf("qoute failed [err=%s]", err.Error())
	}
	defer res.Body.Close()
}
