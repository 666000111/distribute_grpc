package model

type Book struct {
	ID int64 `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Content string `gorm:"column:content" json:"content"`
}
