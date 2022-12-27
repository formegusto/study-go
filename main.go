package main

import (
	"fmt"
	"log"

	"github.com/formegusto/study-go/accounts"
)


func main() {
	// account := banking.Account{owner: "forme"}
	// fmt.Println(account)

	account := accounts.NewAccount("forme")
	fmt.Println(account.Balance())
	account.Deposit(10)
	fmt.Println(account.Balance())
	// err := account.Withdraw(20)
	err := account.Withdraw(5)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

	// 접근 허용 되지 않음
	// account.owner = ""
	// account.balance = 1
}