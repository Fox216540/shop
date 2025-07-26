package core

import "gorm.io/gorm"

type Database struct {
	DB *gorm.DB
}

// Глобальный объект Database
var database *Database

func InitDatabase(db *gorm.DB) {
	if db == nil {
		panic("InitDatabase: db is nil")
	}
	database = &Database{DB: db}
}

func GetDatabase() *Database {
	return database
}

func (d *Database) WithSession(fn func(tx *gorm.DB) error) error {
	return d.DB.Transaction(fn)
}
