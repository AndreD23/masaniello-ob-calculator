package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	/**
	 * Configs
	 */
	var qtdOps = 50                                  // Quantity of operations
	var totalMartingales = 1                         // Total of martingales. 0 = Without martingale, 1 = Martingale 1, 2 = Martingale 2
	var assertsStruct = [3]float32{60.0, 85.0, 95.0} // Asserts for Win, G1, G2. Ex: 60.0 = 60%. Above third value, is Loss

	// Create file
	f, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Run Martingale
	for i := 0; i < qtdOps; i++ {
		// Get random number between 0 and 100
		var random = getRandomNumber(0, 100)

		// Check asserts
		var result = checkAssert(random, totalMartingales, assertsStruct[:])

		// Write file
		writeFile(f, result, totalMartingales)
	}
}

/**
 * Get random number between min and max
 */
func getRandomNumber(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

/*
 * Check asserts
 */
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

/**
 * Write file based on result
 */
func writeFile(f *os.File, result string, martingales int) {
	//fmt.Println("Recebido: ", result, martingales)

	var content string

	if result == "W" {
		content = fmt.Sprintf("W\n")
	}
	if martingales == 1 {
		if result == "G1" {
			content = fmt.Sprintf("L\nW\n")
		}

		if result == "L" {
			content = fmt.Sprintf("L\nL\n")
		}
	}
	if martingales == 2 {
		if result == "G1" {
			content = fmt.Sprintf("L\nW\n")
		}

		if result == "G2" {
			content = fmt.Sprintf("L\nL\nW\n")
		}

		if result == "L" {
			content = fmt.Sprintf("L\nL\nL\n")
		}
	}
	if content == "" {
		content = fmt.Sprintf("L\n")
	}

	_, err := f.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}
