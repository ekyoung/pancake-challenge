package main

import "testing"

func TestH(t *testing.T) {
	pancakes := parseInput("+")

	flips := calcFlips(pancakes)

	if flips != 0 {
		t.Fail()
	}
}

func TestB(t *testing.T) {
	pancakes := parseInput("-")

	flips := calcFlips(pancakes)

	if flips != 1 {
		t.Fail()
	}
}

func TestBH(t *testing.T) {
	pancakes := parseInput("-+")

	flips := calcFlips(pancakes)

	if flips != 1 {
		t.Fail()
	}
}

func TestHB(t *testing.T) {
	pancakes := parseInput("+-")

	flips := calcFlips(pancakes)

	if flips != 2 {
		t.Fail()
	}
}

func TestHHH(t *testing.T) {
	pancakes := parseInput("+++")

	flips := calcFlips(pancakes)

	if flips != 0 {
		t.Fail()
	}
}

func TestBBHB(t *testing.T) {
	pancakes := parseInput("--+-")

	flips := calcFlips(pancakes)

	if flips != 3 {
		t.Fail()
	}
}
