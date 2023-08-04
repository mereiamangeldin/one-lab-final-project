package postgre

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/model"
	"gorm.io/gorm"
	"strings"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProducts() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products)
	if err.Error != nil {
		return nil, err.Error
	}
	return products, nil
}
func (r *ProductRepository) GetProductsByTitle(title string) ([]model.Product, error) {
	var products []model.Product
	err := r.db.Where("LOWER(name) like ?", "%"+strings.ToLower(title)+"%").Find(&products)
	if err.Error != nil {
		return nil, err.Error
	}
	return products, nil
}

func (r *ProductRepository) GetCategories() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories)
	if err.Error != nil {
		return nil, err.Error
	}
	return categories, nil
}

func (r *ProductRepository) GetCategoryProducts(id uint) ([]model.Product, error) {
	var products []model.Product
	err := r.db.Where("category_id = ?", id).Find(&products)
	if err.Error != nil {
		return nil, err.Error
	}
	return products, nil
}

func (r *ProductRepository) CreateProduct(product model.Product) (uint, error) {
	res := r.db.Create(&product)
	return product.ID, res.Error
}

func (r *ProductRepository) LikeAction(userID uint, productID uint) error {
	var favorite model.Favorite
	err := r.db.Where("user_id = ? and product_id = ?", userID, productID).First(&favorite).Error
	if err == nil {
		favorite.Liked = !favorite.Liked
		r.db.Save(&favorite)
	} else {
		favorite.UserID = userID
		favorite.ProductID = productID
		favorite.Liked = true
		r.db.Create(&favorite)
	}
	return nil
}

func (r *ProductRepository) UpdateProduct(id uint, productReq model.Product) error {
	var product model.Product
	result := r.db.Model(&product).Where("id = ?", id).Updates(productReq)
	return result.Error
}

func (r *ProductRepository) GetProductById(id uint) (model.Product, error) {
	var product model.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return model.Product{}, result.Error
	}
	return product, nil
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	var product model.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return result.Error
	}
	result = r.db.Delete(&product)
	return result.Error
}
