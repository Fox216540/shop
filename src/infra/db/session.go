package db

import "gorm.io/gorm"

func (d *Database) WithSession(fn func(tx *gorm.DB) error) error {
	return d.DB.Transaction(fn)
}
