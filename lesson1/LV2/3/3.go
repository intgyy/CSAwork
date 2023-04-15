package main

import "fmt"

func main() {
	var x float64
	var y float64

	fmt.Scanf("%f %f\n", &x, &y)
	//fmt.Scanf("%f\n", &y)
	if x > 0 && y > 0 {
		fmt.Println("Q1")
	} else if x > 0 && y < 0 {
		fmt.Println("Q4")
	} else if x < 0 && y > 0 {
		fmt.Println("Q2")
	} else if x < 0 && y < 0 {
		fmt.Println("Q3")
	} else if x == 0 && y == 0 {
		fmt.Println("Origem")
	} else if x == 0 {
		fmt.Println("Eixo Y")
	} else if y == 0 {
		fmt.Println("Eixo X")
	} else {
		fmt.Println("输错了")
	}

}
