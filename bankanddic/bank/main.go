package main

import (
	"fmt"
	"github.com/firstgo/bankanddic/bank/account"
)

func main() {
	useraccount := account.NewAccount("user")
	useraccount.Deposit(30)
	//err := useraccount.Withdraw(1)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(useraccount.ShowBalance(), useraccount.ShowName())
	fmt.Println(useraccount)
}
