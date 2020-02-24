package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/quote", quoteHandler)
	http.HandleFunc("/quote/v2", quoteHandlerConcurrent)
	http.HandleFunc("/quote/v3", quoteHandlerWait)
	http.ListenAndServe(":8090", nil)
}

// Synchronous functions
func quoteHandler(w http.ResponseWriter, req *http.Request) {
	getUser("a1")
	getUserSub("a1")
	generateQuote("a1")
	fmt.Fprintf(w, "Done quote ready\n")
}

func getUser(s string) {
	time.Sleep(syntheticWaitTime)
	fmt.Println("user found !")
}

func getUserSub(s string) {
	time.Sleep(syntheticWaitTime)
	fmt.Println("subscription found !")
}

func generateQuote(s string) {
	time.Sleep(syntheticWaitTime)
	fmt.Println("quote generated !")
}

// Asynchronous functions, No Wait
func quoteHandlerConcurrent(w http.ResponseWriter, req *http.Request) {
	go getUser("a1")
	go getUserSub("a1")
	go generateQuote("a1")
	fmt.Fprintf(w, "Done quote ready\n")
}

// Wait group functions (old example)
func quoteHandlerWait(w http.ResponseWriter, req *http.Request) {
	var wg sync.WaitGroup
	wg.Add(1)
	go getUserWait("a1", &wg)
	wg.Add(1)
	go getUserSubWait("a1", &wg)
	wg.Add(1)
	go generateQuoteWait("a1", &wg)
	wg.Wait()
	fmt.Fprintf(w, "Done quote ready\n")
}

func getUserWait(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(syntheticWaitTime)
	fmt.Println("user found !")
}

func getUserSubWait(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(syntheticWaitTime)
	fmt.Println("subscription found !")
}

func generateQuoteWait(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(syntheticWaitTime)
	fmt.Println("quote generated !")
}

// Mimics db / api processing time
var syntheticWaitTime = time.Millisecond * 300
