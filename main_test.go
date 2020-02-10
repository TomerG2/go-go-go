package main

import (
	"testing"
	"time"
)

func TestExample_Intro(t *testing.T) {
	go Say("hey")
	go Say("there")
	time.Sleep(time.Second)
}

func TestRecycling_LinearSolution(t *testing.T) {
	garbage := []string{"plastic bottle", "plastic bottle", "apple", "banana"}
	plasticOnly := FilterNonPlastic(garbage)
	RecyclePlastic(plasticOnly)
}

func TestRecycling_ConcurrentSolution1(t *testing.T) {
	garbage := []string{"plastic bottle", "plastic bottle", "apple", "banana"}
	plasticOnly := FilterNonPlastic(garbage)
	go RecyclePlastic(plasticOnly)
}
