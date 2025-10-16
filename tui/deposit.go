package tui

type Transaction struct {
	Date        string
	Description string
	Amount      float32
}

type User struct {
	Id           int
	Name         string
	Balance      string
	Transactions []Transaction
}

type DepositModel struct {
	User User
}
