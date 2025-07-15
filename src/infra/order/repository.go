package orderrepository

import (
	"errors"
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
		UserID:       o.UserID,
		Total:        o.Total,
		ProductItems: make([]*models.ProductItemORM, len(o.ProductItems)),
	}

	for _, item := range o.ProductItems {
		newOrder.ProductItems = append(
			newOrder.ProductItems, &models.ProductItemORM{
				ProductID: item.Product.ID,
				Quantity:  item.Quantity,
				OrderID:   o.ID,
			})
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
		result := tx.
			Where("order_id = ? AND user_id = ?", ID, userID).
			Delete(&models.OrderORM{})
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
	var o models.OrderORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.
			Preload("ProductItems.Product"). // подгружаем Product внутри ProductItems
			Preload("ProductItems").
			Where("id = ?", ID).
			First(&o).Error
	})
	if err != nil {
		return order.Order{}, err // Возвращаем ошибку, если не удалось найти заказ
	}
	return models.FromORM(o), nil
}

func (r *repository) OrdersByUserID(userID uuid.UUID) ([]order.Order, error) {
	var ordersORM []models.OrderORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.
			Preload("ProductItems").         // подгружаем ProductItems
			Preload("ProductItems.Product"). // подгружаем Product внутри ProductItems
			Where("user_id = ?", userID).
			Find(&ordersORM).Error
	})
	if err != nil {
		return nil, err // Возвращаем ошибку, если не удалось найти заказы пользователя
	}
	orders := make([]order.Order, 0, len(ordersORM))
	for _, o := range ordersORM {
		orders = append(orders, models.FromORM(o))
	}
	return orders, nil
}

func (r *repository) CheckOrderNum(orderNum string) error {
	err := r.db.WithSession(func(tx *gorm.DB) error {
		err := tx.Where("order_num = ?", orderNum).First(&models.OrderORM{}).Error
		if err == nil {
			// Если заказ с таким номером найден, возвращаем ошибку
			return fmt.Errorf("order number %s already exists", orderNum) // Если заказ с таким номером найден, возвращаем nil
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		// Если заказ с таким номером не найден, возвращаем nil
		return nil
	})
	return err
}
