package array

import "fmt"

func Basic() {
	// [size]type{...datas}
	names := [5]string{"forme", "gusto", "hi"}
	names[3] = "alala"
	names[4] = "hello" 

	fmt.Println(names)

	/*
	slice
	- Array
	- 제한적인 크기가 없는 Array 
	- append를 통해 추가
	*/
	names2 := []string{"forme"}
	fmt.Println(names2)
	names2 = append(names2, "gusto", "hi")
	fmt.Println(names2)
}