package entity

import "time"

type Wallet struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Amount    uint64    `json:"amount" `
	UserID    uint64    `gorm:"primaryKey" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
