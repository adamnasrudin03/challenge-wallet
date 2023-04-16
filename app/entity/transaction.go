package entity

import "time"

type Transaction struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Type      string    `gorm:"not null;default:'OUT'" json:"type"`
	Amount    uint64    `json:"amount" `
	UserID    uint64    `json:"user_id"`
	User      User      `json:"user,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
