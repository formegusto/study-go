package gorout

import (
	"fmt"
	"time"
)

func SexyCount(person string) {
	for i:=0;i<10;i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func IsSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	/*
	return 대신 channel로 값을 보내주는 방식을 사용한다.
	*/
	fmt.Println(person, "---")
	c <- true
}

func IsSexyString(person string, c chan string) {
	time.Sleep(time.Second * 5)
	
	c <- person + " is sexy."
}