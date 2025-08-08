package models

import "gorm.io/gorm"

type Book struct {
	ID          uint   `json:"id" gorm:"primaryKey,autoIncrement"`
	Author      *string `json:"author"`
	Title       *string `json:"title"`
	Publisher   *string `json:"publisher"`
}

func (b *Book) TableName() string {
	return "books"
}
func MigrateBooks(db *gorm.DB) error {
	if err := db.AutoMigrate(&Book{}); err != nil {
		return err
	}
	return nil
}