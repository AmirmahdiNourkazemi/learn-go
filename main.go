package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var c, python, java bool
var i, j int = 1, 2

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Hello World")
	fmt.Println("The time is", time.Now())
	fmt.Println("Random number", rand.Intn(100))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(21))
	var i int
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no!"
	fmt.Println(i, c, python, java)

	k := 5
	fmt.Printf("The value of k is %T.\n %d\n", k, k)

	var f float64 = math.Sqrt(2)
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	h := false
	fmt.Printf("The value of h is %v and it's type is %T \n", h, h)

	const Truth = true
	fmt.Println(Truth)

	loop()
	forContinue()
	while()
	ifCondition(-2)
	fmt.Println(switchCaseCondition(time.Sunday))
	fmt.Println(switchCasewithoutCondition())
}
