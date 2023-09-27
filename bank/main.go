package main

import (
	"fmt"
	"github.com/firstgo/bank/account"
	"log"
)

func main() {
	useraccount := account.NewAccount("user")
	useraccount.Deposit(30)
	err := useraccount.Withdraw(100)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(useraccount.ShowBalance())
}
