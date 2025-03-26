package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}
	names := [4]string{"James", "Esther", "Peter", "Mary"}
	fmt.Println(names, len(names))
	fmt.Println(ages, len(ages))
}
