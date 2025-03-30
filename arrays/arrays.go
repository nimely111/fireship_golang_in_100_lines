package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20, 25, 30}
	// var ages = [3]int{20, 25, 30}
	// names := [4]string{"James", "Esther", "Peter", "Mary"}
	// names[1] = "Vick"
	// fmt.Println(names, len(names))
	// fmt.Println(ages, len(ages))

	// // slices (uses array under the hood)
	// var scores = []int{100, 50, 60}
	// scores[2] = 25
	// scores = append(scores, 45)
	// fmt.Println(scores, len(scores))

	// // slice ranges
	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]
	// rangeOne = append(rangeOne, "Jestina")
	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	/* Define an array of size 4 that stores deployment options */
	// DeploymentOptions := [4]string{"R-pi", "AWS", "GCP", "Azure"}

	/* Loop through the deployment options array */
	// for i := 0; i < len(DeploymentOptions); i++ {
	// 	option := DeploymentOptions[i]
	// 	fmt.Println(i, option)
	// }

	/* Define an array and let the compiler count its size */
	DeploymentOptions := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	/* Loop through the deployment options array */
	for _, option := range DeploymentOptions {
		fmt.Println(option)
	}

}
