package models

type Account struct {
    AccountID      int     `json:"account_id"`
    InitialBalance float64 `json:"initial_balance"`
    Balance        float64 `json:"balance"`
}
