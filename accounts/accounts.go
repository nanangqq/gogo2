package accounts

// type Account struct {
// account struct
type account struct {
	// Owner   string
	// Balance int
	owner   string
	balance int
}

// NewAccount creates account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}
