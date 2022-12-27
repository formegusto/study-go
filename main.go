package main

import (
	"fmt"

	"github.com/formegusto/study-go/accounts"
)


func main() {
	// account := banking.Account{owner: "forme"}
	// fmt.Println(account)

	account := accounts.NewAccount("forme")
	fmt.Println(account)
	fmt.Println(*account)

	// 접근 허용 되지 않음
	// account.owner = ""
	// account.balance = 1
}