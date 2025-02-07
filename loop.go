package main

import (
	"fmt"
)

func loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forContinue() {
	sum := 0

	for sum < 10 {
		sum++
	}

	fmt.Println(sum)
}

func while() {
	sum := 1

	for sum <= 100 {
		sum += sum
	}

	fmt.Println(sum)
}
