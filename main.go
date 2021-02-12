package main

import (
	"fmt"
	"gogo2/accounts"
)

// import (
// 	"fmt"
// 	"gogo1/util"
// )

// func main() {
// 	fmt.Println("hi")
// 	util.SayHi()
// }

func main() {
	account := accounts.NewAccount("yoo")
	fmt.Println(account)

	account.Deposit(10)
	fmt.Println(account)
	fmt.Println(account.Balance())

	// account.Withdraw(20)
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
