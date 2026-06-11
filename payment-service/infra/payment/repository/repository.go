package repository

import (
	domain "github.com/Fox216540/shop/payment-service/domain/payment"
	db "github.com/Fox216540/shop/payment-service/infra/db/core"
	orm "github.com/Fox216540/shop/payment-service/infra/payment/repository/models"
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"
)

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) domain.Repository {
	return &repository{db: db}
}

func (r *repository) Save(p domain.Payment) error {
	newPayment := &orm.PaymentORM{
		IDYoKassa:   p.ID,
		OrderID:     p.OrderID,
		Value:       p.Amount.Value,
		Currency:    p.Amount.Currency,
		Method:      p.Method,
		Description: p.Description,
		Status:      p.Status,
	}

	err := r.db.WithSession(func(tx *gorm.DB) error {
		if err := tx.Create(newPayment).Error; err != nil {
			return NewInvalidSavePayment(pkgerrors.WithStack(err))
		}
		return nil
	})
	return err
}
