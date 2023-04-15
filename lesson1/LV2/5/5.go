package main

import (
	"fmt"
)

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func main() {
	var m int = 1
	var n int = 1
	var sum int
	for {
		sum = 0

		fmt.Scanf("%d %d\n", &m, &n)
		if m <= 0 || n <= 0 {
			break
		}

		for i := min(m, n); i <= max(m, n); i++ {
			fmt.Printf("%d ", i)
			sum += i
		}

		fmt.Printf("Sum=%d\n", sum)
	}

}
