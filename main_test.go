package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestQuoteAPI_Procedural(t *testing.T) {
	getUser("a1")
	getUserSub("a1")
	generateQuote("a1")
}

func TestQuoteAPI_Goroutines(t *testing.T) {
	getUser("a1")
	getUserSub("a1")
	generateQuote("a1")
}

var wg sync.WaitGroup

func TestQuoteAPI_WaitGroup(t *testing.T) {
	wg.Add(1)
	go getUserWait("a1")
	wg.Add(1)
	go getUserSubWait("a1")
	wg.Add(1)
	go generateQuoteWait("a1")

	wg.Wait()
}

// Helper functions
func getUser(s string) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("user found !")
}

func getUserSub(s string) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("subscription found !")
}

func generateQuote(s string) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("quote generated !")
}

func getUserWait(s string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 500)
	fmt.Println("user found !")
}

func getUserSubWait(s string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 500)
	fmt.Println("subscription found !")
}

func generateQuoteWait(s string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 500)
	fmt.Println("quote generated !")
}
