package main

import "fmt"

func println(){
	fmt.Println("hello world")
}

func printf(){
	const name, age = "SJ", 34
	fmt.Printf("%s is %d years old.\n", name, age)
}

func main() {
	println()
	printf()
}