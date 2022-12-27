package main

import (
	"fmt"

	"github.com/formegusto/study-go/condition"
)

func main() {
	fmt.Println(condition.CanIDrink(16))
	fmt.Println(condition.VECanIDrink(16))
	fmt.Println(condition.SwitchCanIDrink(16))
}