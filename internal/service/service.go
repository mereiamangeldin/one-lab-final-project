package service

import (
	"context"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/entity"
)

type Service interface {
	CreateUser(ctx context.Context, u *entity.User) error
	Login(ctx context.Context, username, password string) (string, error)
	VerifyToken(token string) (int64, error)
	CreateItem(ctx context.Context, i *entity.Item) error
	GetItems(ctx context.Context) ([]entity.Item, error)
	GetItemByID(ctx context.Context, id int64) (entity.Item, error)
	DeleteItem(ctx context.Context, id int64) error
	UpdateItem(ctx context.Context, id int64, item entity.Item) error
}
