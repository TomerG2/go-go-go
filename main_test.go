package main

import (
	"testing"
)

func TestExample_1(t *testing.T) {
	Say("hey")
	Say("there")
}

func TestExample_2(t *testing.T) {
	findUserById("a1")
	checkSubStatus("a1")
	getQuotes("a1")
}
