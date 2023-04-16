package dto

type CreateTxReq struct {
	Name   string `json:"name" validate:"required,min=6"`
	Amount uint64 `json:"amount"`
}

type TopUpReq struct {
	BankName string `json:"bank_name" validate:"required,min=6"`
	Amount   uint64 `json:"amount"  validate:"required,min=1000"`
}
