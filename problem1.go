// Problem 1 https://projecteuler.net/problem=1
// Project Euler
// jpbanczek@gmail.com

package main

import "fmt"

func main() {

	result := 0

	for i := 1; i <= 999; i++ {

		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}

	fmt.Println("Result: ", result)
}
