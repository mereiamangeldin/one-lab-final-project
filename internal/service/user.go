package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	model2 "github.com/mereiamangeldin/One-lab-Homework-1/internal/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/repository"
	"net/http"
)

const transactionUrl = "http://transaction:8000/transactions"

type UserService struct {
	Repo *repository.Repository
}

func (s *UserService) createTransaction(transaction model2.Transaction) error {
	transactionJson, err := json.Marshal(transaction)
	if err != nil {
		return err
	}
	resp, err := http.Post(transactionUrl, "application/json", bytes.NewBuffer(transactionJson))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("remote service returned error: %s", resp.Status)
	}
	return nil
}

func (s *UserService) GetUserProducts(id uint) ([]model2.Product, error) {
	return s.Repo.User.GetUserProducts(id)
}

func (s *UserService) GetUserFavoriteProducts(userID uint) ([]model2.Product, error) {
	return s.Repo.User.GetUserFavoriteProducts(userID)
}

func ProductsContains(ids []uint, id uint) bool {
	for _, i := range ids {
		if i == id {
			return true
		}
	}
	return false
}

func (s *UserService) GetLikedUserProducts(products []model2.Product, userID uint) ([]model2.Product, error) {
	favProducts, err := s.GetUserFavoriteProducts(userID)
	if err != nil {
		return nil, err
	}
	var ids []uint
	for _, product := range favProducts {
		ids = append(ids, product.ID)
	}
	for i, product := range products {
		if ProductsContains(ids, product.ID) {
			products[i].Liked = true

		}
	}
	return products, nil
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetById(id uint) (model2.UserCreateResp, error) {
	return s.Repo.User.GetById(id)
}
func (s *UserService) BuyProduct(userID uint, productID uint) error {
	user, err := s.Repo.User.GetUserById(userID)
	if err != nil {
		return err
	}
	product, err := s.Repo.Product.GetProductById(productID)
	if err != nil {
		return err
	}
	if user.Balance < product.Price {
		return errors.New("not enough balance")
	}
	err = s.Repo.User.BuyProduct(userID, productID, product.Price)
	if err != nil {
		return err
	}

	user.Balance -= product.Price
	s.Repo.User.Update(userID, user)
	return nil
}

func (s *UserService) Delete(id uint) error {
	return s.Repo.User.Delete(id)
}

func (s *UserService) DepositBalance(id uint) error {
	return s.Repo.User.DepositBalance(id)
}

func (s *UserService) Update(id uint, user model2.User) error {
	return s.Repo.User.Update(id, user)
}
