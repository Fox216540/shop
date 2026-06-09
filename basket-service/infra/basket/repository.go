package basket

import (
	"errors"

	domainbasket "github.com/Fox216540/shop/basket-service/domain/basket"
	"github.com/Fox216540/shop/basket-service/domain/productShort"
	"github.com/Fox216540/shop/basket-service/infra/basket/models"
	"github.com/Fox216540/shop/basket-service/infra/db/core"
	"github.com/google/uuid"
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *core.Database
}

func NewRepository(db *core.Database) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetBasket(userID uuid.UUID) (domainbasket.Basket, error) {
	var rows []models.BasketItemORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.
			Where("user_id = ?", userID).
			Find(&rows).Error
	})
	if err != nil {
		return domainbasket.Basket{}, domainbasket.NewInternalError("failed to get basket", pkgerrors.WithStack(err))
	}
	if len(rows) == 0 {
		return domainbasket.Basket{}, domainbasket.NewBasketNotFoundError(nil)
	}

	items := make([]domainbasket.ItemBasket, 0, len(rows))
	for _, row := range rows {
		items = append(items, row.ToDomain())
	}

	return domainbasket.Basket{
		Products: items,
	}, nil
}

func (r *Repository) AddItemToBasket(userID uuid.UUID, product productShort.ProductShort, quantity uint64) (domainbasket.ItemBasket, error) {
	if quantity == 0 {
		return domainbasket.ItemBasket{}, domainbasket.NewInvalidQuantityError(nil)
	}

	basketItem := models.NewBasketItemORM(userID, product, quantity)

	err := r.db.WithSession(func(tx *gorm.DB) error {
		var row models.BasketItemORM
		err := tx.Where("user_id = ? AND product_id = ?", userID, product.ID).First(&row).Error
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			if err := tx.Create(&basketItem).Error; err != nil {
				return err
			}
			return nil
		case err != nil:
			return err
		default:
			row.Name = product.Name
			row.Img = product.Img
			row.PriceAmount = product.Price.Amount
			row.PriceCurrency = product.Price.Currency
			row.Quantity += quantity
			if err := tx.Save(&row).Error; err != nil {
				return err
			}
			basketItem = row
			return nil
		}
	})
	if err != nil {
		return domainbasket.ItemBasket{}, domainbasket.NewInternalError("failed to add item to basket", pkgerrors.WithStack(err))
	}

	return basketItem.ToDomain(), nil
}

func (r *Repository) DeleteBasket(userID uuid.UUID) error {
	err := r.db.WithSession(func(tx *gorm.DB) error {
		result := tx.
			Where("user_id = ?", userID).
			Delete(&models.BasketItemORM{})
		if result.RowsAffected == 0 {
			return domainbasket.NewBasketNotFoundError(nil)
		}
		if result.Error != nil {
			return domainbasket.NewInternalError("failed to delete basket", pkgerrors.WithStack(result.Error))
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) RemoveItemFromBasket(userID, productID uuid.UUID) error {
	err := r.db.WithSession(func(tx *gorm.DB) error {
		result := tx.
			Where("user_id = ? AND product_id = ?", userID, productID).
			Delete(&models.BasketItemORM{})
		if result.RowsAffected == 0 {
			return domainbasket.NewItemNotFoundError(nil)
		}
		if result.Error != nil {
			return domainbasket.NewInternalError("failed to remove item from basket", pkgerrors.WithStack(result.Error))
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
