package repository

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/config"
	model2 "github.com/mereiamangeldin/One-lab-Homework-1/internal/model"
	postgre2 "github.com/mereiamangeldin/One-lab-Homework-1/internal/repository/postgre"
)

type IUserRepository interface {
	GetById(id uint) (model2.UserCreateResp, error)
	GetUserById(id uint) (model2.User, error)
	Delete(id uint) error
	DepositBalance(id uint) error
	BuyProduct(userID uint, productID uint, amount float64) error
	Update(id uint, user model2.User) error
	UpdatePassword(id uint, pass model2.ChangePassword) error
	GetUserProducts(id uint) ([]model2.Product, error)
	GetUserFavoriteProducts(userID uint) ([]model2.Product, error)
}
type IAuthorizationRepository interface {
	CreateUser(user model2.User) (uint, error)
	GetUser(user model2.AuthUser) (model2.User, error)
}
type IProductRepository interface {
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

type Repository struct {
	User    IUserRepository
	Auth    IAuthorizationRepository
	Product IProductRepository
}

func New(cfg *config.Config) (*Repository, error) {
	pgDB, err := postgre2.Dial(cfg.PgURL)
	if err != nil {
		return nil, err
	}
	userRep := postgre2.NewUserRepository(pgDB)
	auth := postgre2.NewAuthorizationRepository(pgDB)
	book := postgre2.NewProductRepository(pgDB)
	return &Repository{User: userRep, Auth: auth, Product: book}, nil
}
