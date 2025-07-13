package orderrepository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop/src/domain/order"
	db "shop/src/infra/db/core"
	models "shop/src/infra/order/models"
)

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) order.Repository {
	return &repository{db: db}
}

func (r *repository) Save(o order.Order) (order.Order, error) {
	newOrder := &models.OrderORM{
		OrderID:      o.ID,
		OrderNum:     o.OrderNum,
		Status:       o.Status,
		UserID:       o.UserId,
		ProductItems: make([]*models.ProductItemORM, len(o.ProductItems)),
	}

	for i, item := range o.ProductItems {
		newOrder.ProductItems[i] = &models.ProductItemORM{
			ProductID: item.Product.ID,
			Quantity:  item.Quantity,
			OrderID:   o.ID,
		}
	}

	err := r.db.WithSession(func(tx *gorm.DB) error {
		if err := tx.Create(newOrder).Error; err != nil {
			return err
		}
		// Загрузим обратно с подгруженными продуктами
		return tx.
			Preload("ProductItems.Product").
			Preload("ProductItems").
			First(newOrder, "order_id = ?", newOrder.OrderID).Error
	})
	if err != nil {
		return order.Order{}, err
	}

	return models.FromORM(*newOrder), nil
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
