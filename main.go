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
		time.Sleep(time.Millisecond * 100)
	}
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
