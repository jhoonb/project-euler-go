// Problem 2 https://projecteuler.net/problem=2
// Project Euler
// jpbanczek@gmail.com

package main

import "fmt"

func main() {

	t1, t2 := 0, 1
	result := 0

	for {

		if t2%2 == 0 {
			result += t2
		}

		if t2 >= 4000000 {
			break
		}
		t1, t2 = (t2), (t2 + t1)

	}
	fmt.Println("Result: ", result)
}
