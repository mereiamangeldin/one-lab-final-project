package postgre

import (
	model2 "github.com/mereiamangeldin/One-lab-Homework-1/internal/model"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserProducts(id uint) ([]model2.Product, error) {
	var products []model2.Product
	var ids []uint
	r.db.Raw("select product_id from transactions where user_id = ?", id).Pluck("id", &ids)
	err := r.db.Where("id IN (?)", ids).Find(&products)
	if err.Error != nil {
		return nil, err.Error
	}

	return products, nil
}

func (r *UserRepository) GetUserFavoriteProducts(userID uint) ([]model2.Product, error) {
	var products []model2.Product
	var ids []uint
	r.db.Raw("select product_id from favorites where user_id = ? and liked = true", userID).Pluck("id", &ids)
	err := r.db.Where("id IN (?)", ids).Find(&products)
	if err.Error != nil {
		return nil, err.Error
	}

	return products, nil
}

func (r *UserRepository) UpdatePassword(id uint, pass model2.ChangePassword) error {
	var user model2.User
	err := r.db.Where("password = ?", pass.CurrentPassword).First(&user)
	if err.Error != nil {
		return err.Error
	}
	user.Password = pass.NewPassword
	r.db.Save(&user)
	return nil
}

func (r *UserRepository) GetById(id uint) (model2.UserCreateResp, error) {
	var userResp model2.UserCreateResp
	err := r.db.Table("users").Where("deleted_at IS NULL").First(&userResp, id)
	if err.Error != nil {
		return model2.UserCreateResp{}, err.Error
	}
	return userResp, nil
}

func (r *UserRepository) GetUserById(id uint) (model2.User, error) {
	var user model2.User
	err := r.db.Table("users").Where("deleted_at IS NULL").First(&user, id)
	if err.Error != nil {
		return model2.User{}, err.Error
	}
	return user, nil
}

func (r *UserRepository) Delete(id uint) error {
	var user model2.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = r.db.Delete(&user)
	return result.Error
}

func (r *UserRepository) DepositBalance(id uint) error {
	var user model2.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	user.Balance += 10000
	r.db.Save(&user)
	return nil
}

func (r *UserRepository) Update(id uint, user model2.User) error {
	result := r.db.Where("id = ?", id).Updates(user)
	return result.Error
}

func (r *UserRepository) BuyProduct(userID uint, productID uint, amount float64) error {
	var transaction model2.Transaction
	transaction.ProductID = productID
	transaction.UserID = userID
	transaction.TakenAt = time.Now()
	transaction.Amount = amount
	err := r.db.Create(&transaction).Error
	if err != nil {
		return err
	}
	return nil

}
