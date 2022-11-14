package domain

import "time"

type Book struct {
	ISBN        string    `json:"isbn" gorm:"primaryKey" `
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	PageCount   uint      `json:"page_count"`
	CoverUrl    string    `json:"cover_url"`
	CreatedAt   time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RequestBook struct {
	ISBN        string    `json:"isbn" form:"isbn" param:"isbn"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	PageCount   uint      `json:"page_count"`
	CoverUrl    string    `json:"cover_url"`
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
	ISBN        string           `json:"isbn"`
	Title       string           `json:"title"`
	Author      string           `json:"author"`
	Description string           `json:"description"`
	PageCount   uint             `json:"page_count"`
	CoverUrl    string           `json:"cover_url"`
	Reviews     []ResponseReview `json:"reviews"`
}
