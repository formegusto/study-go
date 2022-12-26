package variables

import "fmt"

// 변수는 전역으로 꺼내어 사용할 수 없다.
// allName := "forme"
func ConsAndVar() {
	// typesafe
	// 상수
	const name string = "th"

	// bad
	// const isName bool = "th"

	// bad
	// name = "u'r"

	fmt.Println(name)

	// var changeableName string = "forme"
	// 두개는 같은 것 이다.
	// 타입 추론이 들어간다.
	changeableName := "forme"
	changeableName += " gusto"
	fmt.Println(changeableName)

	isName := false
	// isName = "forme"
	isName = true
	fmt.Println(isName)
}