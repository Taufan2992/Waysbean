package repositories

import (
	"fmt"
	"waysbean/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetCart(ID int) (models.Cart, error)
	CreateCart(Cart models.Cart) (models.Cart, error)
	UpdateCart(Cart models.Cart) (models.Cart, error)
	FindCartId(CartId int) ([]models.Cart, error)
	UpdateCartTrans(models.Cart) (models.Cart, error)
	DeleteCart(cart models.Cart, ID int) (models.Cart, error)
	CartProduct(ProductID int) (models.Cart, error)
	CartProductId(ProductID int) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) FindCartId(UserID int) ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Product").Find(&cart, "user_id = ?", UserID).Find(&cart, "transaction_id IS NULL").Error
	// .Find(&cart, "user_id = ?, transaction_id = ?", UserID, nil)
	return cart, err
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Product").First(&cart, ID).Error

	return cart, err
}

func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

func (r *repository) UpdateCartTrans(cart models.Cart) (models.Cart, error) {
	err := r.db.Debug().Save(&cart).Error

	return cart, err
}

func (r *repository) DeleteCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}

func (r *repository) CartProduct(ProductID int) (models.Cart, error) {
	var cart models.Cart
	fmt.Println(ProductID)
	err := r.db.Preload("Product").Debug().Find(&cart, "product_id = ?", ProductID).Error
	return cart, err
}

func (r *repository) CartProductId(ProductID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Product").Find(&cart, "product_id = ?", ProductID).Error

	return cart, err
}
