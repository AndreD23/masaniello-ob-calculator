package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/**
	 * Configs
	 */
	var qtdOps = 10
	var totalMartingales = 1
	var assertsStruct = [3]float32{60.0, 85.0, 95.0}

	// Run Martingale
	for i := 0; i < qtdOps; i++ {
		// Get random number between 0 and 100
		var random = getRandomNumber(0, 100)
		fmt.Println("Random: ", random)

		// Check asserts
		var result = checkAssert(random, totalMartingales, assertsStruct[:])
		fmt.Println(result)
	}
}

func getRandomNumber(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func checkAssert(numAssert float32, martinGales int, asserts []float32) string {
	if numAssert >= 0 && numAssert <= asserts[0] {
		return "W"
	}

	if (martinGales == 1 || martinGales == 2) && numAssert <= asserts[1] {
		return "G1"
	}

	if martinGales == 2 && numAssert <= asserts[2] {
		return "G2"
	}

	return "L"
}
