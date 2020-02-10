package main

import (
	"fmt"
	"sync"
	"time"
)

// configuration
var funcWaitTime = time.Millisecond * 300

// Simple functions
func getUser(s string) {
	time.Sleep(funcWaitTime)
	fmt.Println("user found !")
}

func getUserSub(s string) {
	time.Sleep(funcWaitTime)
	fmt.Println("subscription found !")
}

func generateQuote(s string) {
	time.Sleep(funcWaitTime)
	fmt.Println("quote generated !")
}

// Wait group functions
func getUserWait(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(funcWaitTime)
	fmt.Println("user found !")
}

func getUserSubWait(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(funcWaitTime)
	fmt.Println("subscription found !")
}

func generateQuoteWait(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(funcWaitTime)
	fmt.Println("quote generated !")
}
