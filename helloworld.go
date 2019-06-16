package main

import "fmt"

func main() {
	fmt.Println("In main")
	foo()
	fmt.Println("something more")
}

func foo() {
	fmt.Println("I am in foo")
}
