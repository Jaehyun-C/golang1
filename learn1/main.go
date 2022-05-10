package main

import (
	"fmt"

	"github.com/learn1/accounts"
)

func main() {
	account := accounts.NewAccount("choe")
	account.Deposit(100)
	err := account.Withdraw(200)

	if err != nil{
		//log.Fatalln(err)
		fmt.Println(err)
	}
	
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println("-----------------------")
	fmt.Println(account)

}


