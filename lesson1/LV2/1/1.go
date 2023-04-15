package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("输入失败")
	}
	fmt.Printf("%d minutos", 2*n)
}
