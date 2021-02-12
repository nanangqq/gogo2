package accounts

import "errors"

// type Account struct {
// account struct
type account struct {
	// Owner   string
	// Balance int
	owner   string
	balance int
}

var errNoMoney = errors.New("not enough money")

// NewAccount creates account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

//method
// Deposit x amount on your account
// func (a account) Deposit(amount int) { // a: reciever variable,,, making a copy of account
func (a *account) Deposit(amount int) { // a: reciever variable
	// defer fmt.Println(amount, a.balance)
	// fmt.Println(amount)
	a.balance += amount
	// fmt.Println(a.balance)
}

//
func (a account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
