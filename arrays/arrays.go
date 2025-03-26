package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}
	names := [4]string{"James", "Esther", "Peter", "Mary"}
	names[1] = "Vick"
	fmt.Println(names, len(names))
	fmt.Println(ages, len(ages))

	// slices (uses array under the hood)
	var scores = []int{100, 50, 60}
	fmt.Println(scores, len(scores))

}
