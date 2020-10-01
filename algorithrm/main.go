package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	nPtr := flag.Int64("n", 2000000, "Text to parse.")
	flag.Parse()

	var n int64 = *nPtr
	var arr [2000001]bool
	var i, j int64
	var sumNotPrimes int64 = 0
	for i = 2; i < int64(math.Ceil(math.Sqrt(float64(n)))); i++ {
		var k int64 = 0
		for {
			j = i*i + k*i
			if j > n {
				break
			}
			if arr[j] == true {
				k++
				continue
			}
			arr[j] = true
			sumNotPrimes += j
			k++
		}
	}

	// fmt.Println(n * (n + 1) / 2)
	// fmt.Println(sumNotPrimes)
	resl := n*(n+1)/2 - sumNotPrimes

	fmt.Println(resl)
}
