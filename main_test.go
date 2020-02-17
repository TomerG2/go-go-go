package main

import (
	"sync"
	"testing"
	"time"
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
	wg.Add(1)
	go getUserWait("a1", &wg)
	wg.Add(1)
	go getUserSubWait("a1", &wg)
	wg.Add(1)
	go generateQuoteWait("a1", &wg)

	wg.Wait()
}
