package domain

import "time"

type Book struct {
	ISBN        string    `gorm:"primaryKey" json:"isbn"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	PageCount   uint      `json:"page_count"`
	CoverUrl    string    `json:"cover_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ResponseBook struct {
	ISBN        string `json:"isbn"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	PageCount   uint   `json:"page_count"`
	CoverUrl    string `json:"cover_url"`
}

type ResponseDetailBook struct {
	ISBN        string   `json:"isbn"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	PageCount   uint     `json:"page_count"`
	CoverUrl    string   `json:"cover_url"`
	Reviews     []Review `json:"reviews"`
}
