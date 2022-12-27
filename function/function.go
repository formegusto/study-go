package function

import (
	"fmt"
	"strings"
)

// func Multiply(a int, b int) int {
// 	return a * b
// }

// 축약형
/*
a와 b모두 integer이다.
*/
func Multiply(a ,b int) int {
	return a * b
}

func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func RepeatMe(words ...string) {
	fmt.Println(words)
}

func LakedLenAndUpper(name string) (length int, uppercase string) {
	length, uppercase = len(name), strings.ToUpper(name)
	// 이것이 없어도 동작한다.
	// return length, uppercase
	return
}

func DeferLenAndUpper(name string)  (length int, uppercase string) {
	defer fmt.Println("LenAndUpper func() is Done")
	length, uppercase = len(name), strings.ToUpper(name)
	return
}