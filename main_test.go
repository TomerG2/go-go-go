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

func TestQuoteAPI_Channels(t *testing.T) {
	doneChan := make(chan string)
	go func() {
		user := make(chan string)
		subsQueue := make(chan string)
		quotes := make(chan int)
		go getUserChannel("a1", user)
		go getUserSubChannel(user, subsQueue)
		go generateQuoteChannel(subsQueue, quotes)
		doneChan <- "I’m all done!"
	}()
	<-doneChan                    // block until go routine signals work is done
	<-time.After(time.Second * 5) //Receiving from channel after 5 sec
}
