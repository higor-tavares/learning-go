package main

import "fmt"

func swapp(first, second string) (string, string) {
	return second, first
}

func main() {
	first, second := swapp("hello", "world")
	fmt.Println(first, second)
}