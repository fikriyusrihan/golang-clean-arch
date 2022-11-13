package domain

import "time"

type Review struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ISBN      string    `json:"isbn"`
	Writer    string    `json:"writer"`
	Content   string    `json:"content"`
	Rating    uint8     `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}
