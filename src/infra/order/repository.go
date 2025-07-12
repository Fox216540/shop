package orderrepository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop/src/domain/order"
	"shop/src/infra/db"
)

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) order.Repository {
	return &repository{db: db}
}

func (r *repository) Save(o order.Order) (order.Order, error) {
	newOrder := &OrderORM{
		OrderID:  o.ID,
		OrderNum: o.OrderNum,
		Status:   o.Status,
	}

	for _, item := range o.ProductItems {
		newOrder.ProductItems = append(newOrder.ProductItems, ProductItemORM{
			OrderID:   o.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Create(newOrder).Error
	})
	if err != nil {
		return order.Order{}, err // Возвращаем ошибку, если не удалось сохранить заказ
	}
	return fromORM(*newOrder), nil
}

func (r *repository) Remove(ID, userID uuid.UUID) error {
	err := r.db.WithSession(func(tx *gorm.DB) error {
		result := tx.Where("id = ? AND user_id = ?", ID, userID).Delete(&order.Order{})
		if result.RowsAffected == 0 {
			// TODO: Поменять на кастомную ошибку
			return fmt.Errorf("order not found")
		}
		return result.Error // Возвращаем ошибку, если не удалось удалить заказ
	})

	if err != nil {
		return err // Возвращаем ошибку, если не удалось удалить заказ
	}
	return nil
}

func (r *repository) GetByID(ID uuid.UUID) (order.Order, error) {
	var o order.Order
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Where("id = ?", ID).First(&o).Error
	})
	if err != nil {
		return order.Order{}, err // Возвращаем ошибку, если не удалось найти заказ
	}
	return o, nil
}

func (r *repository) OrdersByUserID(userID uuid.UUID) ([]order.Order, error) {
	var orders []order.Order
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Where("user_id = ?", userID).Find(&orders).Error
	})
	if err != nil {
		return nil, err // Возвращаем ошибку, если не удалось найти заказы пользователя
	}
	return orders, nil
}
