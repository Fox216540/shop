package productrepository

import "github.com/google/uuid"

type ProductORM struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	Name        string    `gorm:"not null;unique"` // Unique product name
	Img         string    `gorm:"not null"`        // Product image URL
	Price       float64   `gorm:"not null"`        // Product price
	Category    string    `gorm:"not null"`        // Product category
	Description string    `gorm:"not null"`        // Product description
	Stock       int       `gorm:"not null"`        // Product stock quantity
}
