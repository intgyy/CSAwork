package main

import "fmt"

func main() {
	var n = 1
	for n != 0 {
		fmt.Scanf("%d\n", &n)
		for i := 1; i <= n; i++ {
			fmt.Printf("%d ", i)
		}
		if n == 0 {
			break
		}
		fmt.Printf("\n")
	}
}
