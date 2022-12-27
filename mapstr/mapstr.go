package mapstr

import "fmt"

func Basic() {
	// map[key type]value type{}
	forme := map[string]string{"name":"forme", "age": "12"}
	fmt.Println(forme)
	/*
	여러가지 데이터 타입이 혼합된, 흔히 object라고 부르는 타입을 
	Go에서 정의하는 방법은 struct 이다.
	*/
	for key, value := range forme {
		fmt.Println(key, value, forme[key])
	}
}