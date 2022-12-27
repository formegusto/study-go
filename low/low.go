package low

import "fmt"

func Basic() {
	a := 2
	b := a
	a = 10

	// 10, 2
	fmt.Println(a, b)

	// memory address
	fmt.Println(&a, &b)
	var c *int = &a 
	(*c) = 20
	fmt.Println(a)
	// 오우,, 걍 C언어네

	d := &a
	(*d) = 2000
	fmt.Println(*d,a)
}