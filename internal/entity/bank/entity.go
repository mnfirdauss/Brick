package bank

type Account struct {
	AccountNumber int    `json:"account_number"`
	AccountName   string `json:"account_name"`
	BankName      string `json:"bank_name"`
}
