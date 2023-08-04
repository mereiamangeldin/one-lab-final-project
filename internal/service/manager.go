package service

import (
	"errors"
	model2 "github.com/mereiamangeldin/One-lab-Homework-1/internal/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/repository"
)

type IUserService interface {
	GetById(id uint) (model2.UserCreateResp, error)
	Delete(id uint) error
	Update(id uint, user model2.User) error
	GetUserProducts(id uint) ([]model2.Product, error)
	GetUserFavoriteProducts(userID uint) ([]model2.Product, error)
	GetLikedUserProducts(products []model2.Product, userID uint) ([]model2.Product, error)
	DepositBalance(id uint) error
	BuyProduct(userID uint, productID uint) error
}
type IAuthorizationService interface {
	CreateUser(user model2.User) (uint, error)
	GenerateToken(user model2.AuthUser) (string, error)
	ParseToken(token string) (uint, error)
	UpdatePassword(id uint, pass model2.ChangePassword) error
}

type IProductService interface {
	GetProducts() ([]model2.Product, error)
	GetProductsByTitle(title string) ([]model2.Product, error)
	GetCategoryProducts(id uint) ([]model2.Product, error)
	CreateProduct(product model2.Product) (uint, error)
	UpdateProduct(id uint, product model2.Product) error
	GetProductById(id uint) (model2.Product, error)
	DeleteProduct(id uint) error
	GetCategories() ([]model2.Category, error)
	LikeAction(userID uint, productID uint) error
}

type Manager struct {
	User    IUserService
	Auth    IAuthorizationService
	Product IProductService
}

func NewManager(repo *repository.Repository) (*Manager, error) {
	if repo == nil {
		return nil, errors.New("No storage given")
	}
	userSrv := NewUserService(repo)
	authSrv := NewAuthorizationService(repo)
	productSrv := NewBookService(repo)
	return &Manager{User: userSrv, Auth: authSrv, Product: productSrv}, nil
}
