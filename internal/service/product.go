package service

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/repository"
)

type ProductService struct {
	Repo *repository.Repository
}

func (s *ProductService) GetProducts() ([]model.Product, error) {
	return s.Repo.Product.GetProducts()
}

func (s *ProductService) LikeAction(userID uint, productID uint) error {
	return s.Repo.Product.LikeAction(userID, productID)
}

func (s *ProductService) GetProductsByTitle(title string) ([]model.Product, error) {
	return s.Repo.Product.GetProductsByTitle(title)
}

func (s *ProductService) GetCategoryProducts(id uint) ([]model.Product, error) {
	return s.Repo.Product.GetCategoryProducts(id)
}

func (s *ProductService) GetCategories() ([]model.Category, error) {
	return s.Repo.Product.GetCategories()
}

func (s *ProductService) CreateProduct(book model.Product) (uint, error) {
	return s.Repo.Product.CreateProduct(book)
}

func (s *ProductService) UpdateProduct(id uint, book model.Product) error {
	return s.Repo.Product.UpdateProduct(id, book)
}

func (s *ProductService) GetProductById(id uint) (model.Product, error) {
	return s.Repo.Product.GetProductById(id)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.Repo.Product.DeleteProduct(id)
}

func NewBookService(repo *repository.Repository) *ProductService {
	return &ProductService{Repo: repo}
}
