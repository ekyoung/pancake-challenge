package main

func calcFlips(pancakes []bool) int {
	i := 0
	for j := len(pancakes) - 1; j >= 0; j-- {
		if !pancakes[j] {
			i = j + 1
			break
		}
	}

	//fmt.Printf("Top %v of %v pancakes are mixed.\n", i, len(pancakes))

	if i == 0 {
		return 0
	}

	remainingPancakes := pancakes[:i]

	numFlips := setupTop(remainingPancakes)
	flip(remainingPancakes, i)
	numFlips++

	return numFlips + calcFlips(pancakes[:i])
}

func setupTop(pancakes []bool) int {
	//fmt.Printf("Setting up top %v pancakes.\n", len(pancakes))

	numHappy := 0
	for i := 0; i < len(pancakes); i++ {
		if !pancakes[i] {
			numHappy = i
			break
		}
	}

	if numHappy == 0 {
		//fmt.Println("No pancakes need to be set up.")
		return 0
	}

	//fmt.Printf("%v pancakes need to be set up.\n", numHappy)
	flip(pancakes, numHappy)

	return 1
}

func flip(pancakes []bool, top int) {
	//fmt.Printf("Flipping %v of %v pancakes.\n", top, len(pancakes))

	newState := make([]bool, top)

	for i := 0; i < top; i++ {
		newState[i] = !pancakes[top-i-1]
	}

	for i := 0; i < top; i++ {
		pancakes[i] = newState[i]
	}
}
