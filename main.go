package main

import (
	"strings"
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

func FilterNonPlastic(garbage []string) []string {
	fmt.Printf("Garbage: %s \n", strings.Join(garbage, ", "))
	var plasticOnly []string
	for _, g := range garbage {
		time.Sleep(time.Millisecond * 100)
		if !strings.Contains(g, "plastic") {
			continue
		}
		plasticOnly = append(plasticOnly, g)
	}
	fmt.Printf("FilterNonPlastic: %s \n", strings.Join(plasticOnly, ", "))
	return plasticOnly
}

func RecyclePlastic(plastics []string) []string {
	var rawPlastic []string
	for _, _ = range plastics {
		time.Sleep(time.Millisecond * 100)
		rawPlastic = append(rawPlastic, "raw plastic")
	}
	fmt.Printf("FilterNonPlastic: %s \n", strings.Join(rawPlastic, ", "))
	return plastics
}
