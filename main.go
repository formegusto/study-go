package main

import (
	"fmt"

	"github.com/formegusto/study-go/gorout"
)

func main() {
	/*
	channel 생성
	make(chan [type])
	- 생성된 채널을 Goroutine 실행함수로 보내서 연결시키는 것 이다.
	*/
	// c := make(chan bool)
	c := make(chan string)

	// people := [2]string{"forme", "gusto"}
	people := [5]string{"forme", "gusto", "no", "tae", "heon"}
	for _, person := range people {
		// go gorout.IsSexy(person, c)
		go gorout.IsSexyString(person, c)
	}
	// await 같은건데,,
	// 그니까 같은 채널을 쓰고 있고
	// 먼저 무언가가 반환 될 때까지 main은 기다리는 거
	// 그러니까 첫 번째든 두 번째 함수든 먼저 보내는 쪽의 결과가
	// 붙는다.
	// result := <- c
	// fmt.Println(result)
	fmt.Println("Waiting for msg")
	for i:=0;i<len(people);i++ {
		fmt.Println("Received this msg", i, ":",<-c)
	}

	// deadlock!
	// fmt.Println(<-c)
}