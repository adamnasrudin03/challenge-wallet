package dto

type CreateTxReq struct {
	Name     string `json:"name" validate:"required,min=6"`
	Amount   uint64 `json:"amount"`
	Quantity uint64 `json:"quantity" validate:"required,min=1"`
}

type TopUpReq struct {
	BankName string `json:"bank_name" validate:"required,min=6"`
	Amount   uint64 `json:"amount"  validate:"required,min=1000"`
}
