package main

import "fmt"

func countApplesAndOranges(s int, t int, a int, b int, apples []int, oranges []int) []int {
	numberOfApples := 0
	for _, apple := range apples {
		appleLocation := a + apple
		if appleLocation >= s && appleLocation <= t {
			numberOfApples++
		}
	}
	numberOfOranges := 0
	for _, orange := range oranges {
		orangeLocation := b + orange
		if orangeLocation >= s && orangeLocation <= t {
			numberOfOranges++
		}
	}
	fmt.Println("Apples", numberOfApples, "Oranges", numberOfOranges)
	return []int{numberOfApples, numberOfOranges}
}

func main() {
	// -------------------- (20 locations right most orange +5 of orange tree 15)
	// ----A-|   |---O-----  trees and house
	//   a  aa o          o  fruits fallen from trees
	leftSideHouseLocation := 7
	rightSideHouseLocation := 11
	appleTreeLocation := 5
	orangeTreeLocation := 15
	appleLocationsRelativeAroundAppleTree := []int{-2, 2, 1}
	orangeLocationsRelativeAroundOrangeTree := []int{5, -6}
	countApplesAndOranges(leftSideHouseLocation, rightSideHouseLocation, appleTreeLocation, orangeTreeLocation, appleLocationsRelativeAroundAppleTree, orangeLocationsRelativeAroundOrangeTree)
}
