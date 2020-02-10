package main

import (
	"time"
)

import (
	"fmt"
)

func Say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func findUserById(uid string) string {
	time.Sleep(time.Second)
	fmt.Println("user found")
	return "John"
}

func checkSubStatus(uid string) bool {
	time.Sleep(time.Second)
	fmt.Println("subscription found")
	return true
}

func getQuotes(uid string) int {
	time.Sleep(time.Second)
	fmt.Println("quote found")
	return 3000
}
