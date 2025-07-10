package models

type Transaction struct {
    SourceAccountID      int     `json:"source_account_id"`
    DestinationAccountID int     `json:"destination_account_id"`
    Amount               float64 `json:"amount"`
}
