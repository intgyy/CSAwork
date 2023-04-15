package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(maxNum)
	for {
		fmt.Println("输入你想的数字")
		_, err := fmt.Scanf("%d\n", &n)
		if err != nil {
			fmt.Println("输入错误")
			continue
		}
		if n > secret {
			fmt.Println("你输的大了")
		} else if n < secret {
			fmt.Println("你输的小了")
		} else {
			fmt.Println("你输对了")
			break
		}

	}
	fmt.Println("猜的数是", secret)
}
