package strc

import "fmt"

/*
object와 비슷하고 map보다 유연함
*/
type person struct {
	name 	string
	age 	int
	favFood []string
}

func Basic() {
	favFood := []string{"pizza", "chicken"}

	// 1. 순서 대로 넣는 방법
	forme := person{"nth", 27, favFood}
	fmt.Println(forme)

	// 2. 명시적으로 넣는 방법
	gusto := person{name: "nth", age: 27, favFood: favFood}
	fmt.Println(gusto)
}

/*
Go는
class, object의 개념이없다.
method도 struct가 가지고 있다.
함수가 일급객체라는 소리
-- 간결한 언어! 좋은 C언어 느낌인듯

Go에서는 이렇게 문법이 간결하기 때문에
대부분의 기능이 struct로 부터 와서 struct를 잘 알고 있는 것 이좋다.
*/