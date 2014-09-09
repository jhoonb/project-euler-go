// Problem 3 https://projecteuler.net/problem=3
// Project Euler
// jpbanczek@gmail.com

package main

import "fmt"

func factor(n int64) (nfac []int64) {

	switch n {
	case 0, 1, 2:
		nfac = append(nfac, n)
		return nfac
	}

	p := int64(2)
	naux := n

	for p <= naux {

		if naux%p == 0 {
			nfac = append(nfac, p)
			naux = naux / p
		} else {
			p++
		}
	}
	return nfac
}

func main() {

	f := factor(600851475143)
	fmt.Println("factor of: ", 600851475143, " | largest prime factor: ", f[len(f)-1])

}
