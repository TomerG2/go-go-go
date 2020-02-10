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

// Channel functions
func getUserChannel(s string, userOut chan string) {
	time.Sleep(funcWaitTime)
	userOut <- s
	fmt.Println("user found !")
}

func getUserSubChannel(userIn chan string, subscriptionOut chan string) {
	time.Sleep(funcWaitTime)
	s := <-userIn
	subscriptionOut <- "free"
	subscriptionOut <- "north"
	subscriptionOut <- "veggie"
	fmt.Println("subscriptions found !", s)
}

func generateQuoteChannel(subIn chan string, quotes chan int) {
	time.Sleep(funcWaitTime)
	<-subIn
	quotes <- 3000
	quotes <- 2500
	quotes <- 3200
	fmt.Println("subscriptions found !")
}

func finder(mine [5]string, oreChannel chan string) {
	for _, item := range mine {
		if item == "ore" {
			oreChannel <- item //send item on oreChannel
		}
	}
}
