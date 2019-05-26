package main

import "fmt"
 
type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, `says, "Good morning Pascal"`)
}

func main() {
	x := []int{2, 3, 4, 4}
	fmt.Println(x)

	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	p1 := person{
		"miss",
		"moneypenny",
	}
	p1.speak()
}
