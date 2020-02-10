package main

import (
	"testing"
)

func TestExample_Intro(t *testing.T) {
	Say("hey")
	Say("there")
}

func TestRecycling_LinearSolution(t *testing.T) {
	garbage := []string{"plastic bottle", "plastic bottle", "apple", "banana"}
	plasticOnly := FilterNonPlastic(garbage)
	RecyclePlastic(plasticOnly)
}
