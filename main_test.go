package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestExample_LinearSolution(t *testing.T) {
	getUser("a1")
	getUserSub("a1")
	generateQuote("a1")
}

var wg sync.WaitGroup

func TestExample_WaitGroup(t *testing.T) {
	wg.Add(1)
	go getUser("a1")
	wg.Add(1)
	go getUserSub("a1")
	wg.Add(1)
	go generateQuote("a1")

	wg.Wait()
}

func getUser(s string) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("user found !")
	wg.Done()
}

func getUserSub(s string) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("subscription found !")
	wg.Done()
}

func generateQuote(s string) {
	time.Sleep(time.Millisecond * 500)
	fmt.Println("quote generated !")
	wg.Done()
}
