package main

import "fmt"

func NoBranch(a int) {
	fmt.Println(a)
}

func OneBranch(a int) {
	if a > 5 {
		fmt.Println(a)
	}
	fmt.Println(a + 5)
}

func TwoBranch(a, b int) {
	if a > 5 || b > 10 {
		fmt.Println(a)
		fmt.Println(b)
	} else if b == 14 {
		fmt.Println(b + 5)
	} else {
		fmt.Println(b + 22)
	}
}

func main() {
	fmt.Println("Hello, world.")
	TwoBranch(6, 110)
	// TwoBranch(7, 110)
	// TwoBranch(13, 110)
	// TwoBranch(14, 110)
}
