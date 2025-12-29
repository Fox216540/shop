package core

import "gorm.io/gorm"

type Database struct {
	DB *gorm.DB
}

// NewDatabase Глобальный объект Database
func NewDatabase(db *gorm.DB) *Database {
	if db == nil {
		panic("NewDatabase: db is nil")
	}
	return &Database{DB: db}
}

// WithSession Транзакция
func (d *Database) WithSession(fn func(tx *gorm.DB) error) error {
	return d.DB.Transaction(fn)
}
