package main

func calcFlips(pancakes []bool) int {
	//Happy pancakes on the bottom are good.
	//Find the bottom most blank pancake and work on making it happy
	numPancakesToWorkOn := findNumOfBottomMostBlank(pancakes)

	if numPancakesToWorkOn == 0 {
		return 0
	}

	numFlips := 0

	//Any happy pancakes on top will become blank when we flip
	//so first need to flip them to blank
	numHappy := findNumHappyPancakesOnTop(pancakes)
	if numHappy > 0 {
		flip(pancakes, numHappy)
		numFlips++
	}

	//The blank pancakes on top will replace the blank pancake
	//that was deeper in the stack with a happy pancake
	flip(pancakes, numPancakesToWorkOn)
	numFlips++

	//Repeat the process excluding the pancakes that are known
	//to be happy
	return numFlips + calcFlips(pancakes[:numPancakesToWorkOn-numHappy])
}

func findNumOfBottomMostBlank(pancakes []bool) int {
	for i := len(pancakes) - 1; i >= 0; i-- {
		if !pancakes[i] {
			return i + 1
		}
	}

	return 0
}

func findNumHappyPancakesOnTop(pancakes []bool) int {
	for i := 0; i < len(pancakes); i++ {
		if !pancakes[i] {
			return i
		}
	}

	return len(pancakes)
}

func flip(pancakes []bool, top int) {
	newState := make([]bool, top)

	for i := 0; i < top; i++ {
		newState[i] = !pancakes[top-i-1]
	}

	for i := 0; i < top; i++ {
		pancakes[i] = newState[i]
	}
}
