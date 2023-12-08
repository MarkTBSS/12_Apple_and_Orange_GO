package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32,
	apples []int32, oranges []int32) {
	// Write your code here
	var appleCount int32
	appleCount = 0
	//appleCount := 0
	var orangeCount int32
	orangeCount = 0
	// orangeCount := 0
	// Iteration
	for i := 0; i < len(apples); i++ {
		var applePosition = a + apples[i]
		if applePosition >= s && applePosition <= t {
			appleCount++
		}
	}
	// For Range
	for _, value := range oranges {
		var orangePosition = b + value
		if orangePosition >= s && orangePosition <= t {
			orangeCount++
		}
	}
	fmt.Println(appleCount)
	fmt.Println(orangeCount)
}

func main() {
	countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
}
