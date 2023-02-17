package main

import "fmt"

func split(sum int) (first, second int) {
	first = sum * 4/9
	second = sum - first
	return
}

func main() {
	fmt.Println(split(17))
}