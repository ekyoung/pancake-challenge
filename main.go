package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	strTestCases, _ := reader.ReadString('\n')
	numTestCases, _ := strconv.Atoi(strings.TrimSpace(strTestCases))

	testCases := make([][]bool, numTestCases)
	for i := 0; i < numTestCases; i++ {
		input, _ := reader.ReadString('\n')
		testCases[i] = parseInput(strings.TrimSpace(input))
	}

	for i, pancakes := range testCases {
		flips := calcFlips(pancakes)
		fmt.Printf("Case #%v: %v\r\n", i+1, flips)
	}
}

func parseInput(input string) []bool {
	pancakes := make([]bool, len(input))

	for i, side := range input {
		pancakes[i] = side == '+'
	}

	return pancakes
}
