package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func quoteHandler(w http.ResponseWriter, req *http.Request) {
	getUser("a1")
	getUserSub("a1")
	generateQuote("a1")
	fmt.Fprintf(w, "Done quote ready\n")
}

func main() {
	http.HandleFunc("/quote", quoteHandler)
	http.ListenAndServe(":8090", nil)
}

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

// Old example
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
