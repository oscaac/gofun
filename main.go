package main

import "fmt"

var y = 100 + 32 //declare with var, assign with =
var z int
var blah string = "shaken, not stirred"
type hotdog int
var a hotdog

func main() {
	fmt.Println("Hello World!")
	foo()

	for i := 0; i < 100; i++ {
		if i%5 == 0 {
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

	a = 43
	fmt.Println("a =", a)
	fmt.Printf("%T\n", a)

	y = int(a)
	fmt.Println("y =", y)
	fmt.Printf("%T\n", y)

	s := "Oscar I Hernandez"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	var arrayx [5] int;
	arrayx[3] = 21
	fmt.Println(arrayx)
	fmt.Println(len(arrayx))

	//slice
	//x := []type{values}
	slicex := []int{1,2,3,4,5}
	fmt.Println(slicex)
	fmt.Println(slicex[3])
	fmt.Println(len(slicex))

	for i,v := range slicex {
		fmt.Println(i, v)
	}

	//slicing a slice
	fmt.Println(slicex[1:])
	fmt.Println(slicex[1:3])

	slicex = append(slicex, 6,7,8)
	fmt.Println(slicex)

	slicey := []int{0,9,8,7}

	// use ... at the end of slice to pass in each value individually
	slicex = append(slicex, slicey...)
	fmt.Println(slicex)

	//delete from slice
	slicex = append(slicex[:2], slicex[4:]...)
	fmt.Println(slicex)

	for i, v := range slicex {
		fmt.Println(i, v)
	}

	makeslicex := make([]int, 10, 100)
	fmt.Println(makeslicex)
	fmt.Println(len(makeslicex))
	fmt.Println(cap(makeslicex))

	sliceo := []string{"oscar", "isaac", "hernandez"}
	sliceg := []string{"mary", "grace", "hernandez"}
	sliceus := [][]string{sliceo, sliceg}
	fmt.Println(sliceus)

	m := map[string]int{
			"Oscar":31,
			"Grace":21,
	}
	fmt.Println(m)
	fmt.Println(m["Oscar"])
	fmt.Println(m["Doesn't Exist"])		//returns default value

	if 	v, ok := m["doesn't exist"] ; ok {
		fmt.Println(v)
	}

	//add to map
	m["Oreo"] = 4

	for k, v := range m {
		fmt.Println(k, v)
	}
	delete(m, "Oscar") //will not error if doesn't exist in map
	fmt.Println(m)

}

func foo() {
	fmt.Println("I'm in foo")
}
