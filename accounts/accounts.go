package accounts

// Account struct
type account struct {
	owner string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *account {
	a := account{owner: owner, balance: 0}

	return &a
}