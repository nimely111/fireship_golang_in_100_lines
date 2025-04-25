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
	// DeploymentOptions := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	// /* Loop through the deployment options array */
	// for _, option := range DeploymentOptions {
	// 	fmt.Println(option)
	// }

	/* Define an array containing programming languages */
	languages := [10]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust", "PHP", // Must include the trailing comma
	}

	/* Define slices */
	classics := languages[0:3]  // alternatively languages[:3]
	modern := make([]string, 4) // len(modern) = 4
	modern = languages[3:7]     // include 3 exclude 7
	new := languages[7:9]       // alternatively languages[7:]

	fmt.Printf("classic languagues: %v\n", classics) // classic languagues: [C Lisp C++]
	fmt.Printf("modern languages: %v\n", modern)     // modern languages: [Java Python JavaScript Ruby]
	fmt.Printf("new languages: %v\n", new)

}
