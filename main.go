package main

import (
	"fmt"

	"github.com/formegusto/study-go/function"
	"github.com/formegusto/study-go/variables"
)

func main() {
	fmt.Println("Hello World")

	variables.ConsAndVar()

	fmt.Println(function.Multiply(2,2))
	totalLength, upperName := function.LenAndUpper("forme")
	fmt.Println(totalLength, upperName)
	totalLength2, _ := function.LenAndUpper("gusto")
	fmt.Println(totalLength2)

	function.RepeatMe("forme", "gusto", "hello", "World")

	totalLength3, upperName3 := function.LakedLenAndUpper("forme")
	fmt.Println(totalLength3, upperName3)

	totalLength4, upperName4 := function.DeferLenAndUpper("gusto")
	fmt.Println(totalLength4, upperName4)
}