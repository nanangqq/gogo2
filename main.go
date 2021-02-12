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
}
