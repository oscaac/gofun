package main

import "fmt"

var y = 100 + 32 //declare with var, assign with =
var z int
var blah string = "shaken, not stirred"

func main() {
	fmt.Println("Hello World!")
	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	n, err := fmt.Println("String", 44, false)
	fmt.Println(n)
	fmt.Println(err)

	x := 42 //declare and assign
	fmt.Println(x)
	x = 5 //assign (already declared)
	fmt.Println(x)

	fmt.Println(y)
	fmt.Println("z =", z)
	fmt.Printf("%T\n", z)
	fmt.Println("blah =", blah)
	fmt.Printf("%T\n", blah)

}

func foo() {
	fmt.Println("I'm in foo")
}
