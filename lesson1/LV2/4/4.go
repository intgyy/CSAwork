package main

import "fmt"

func main() {
	var n float32
	fmt.Scanf("%f", &n)
	if 0 <= n && n <= 25 {
		fmt.Println("Intervalo [0,25]")
	} else if 25 < n && n <= 50 {
		fmt.Println("Intervalo (25,50]")
	} else if 50 < n && n <= 75 {
		fmt.Println("Intervalo (50,75]")
	} else if 75 < n && n <= 100 {
		fmt.Println("Intervalo (75,100]")
	} else {
		fmt.Println("Fora de intervalo")
	}
}
