package repository

import (
	"context"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, u *entity.User) error
	GetUser(ctx context.Context, username string) (*entity.User, error)
	CreateItem(ctx context.Context, i *entity.Item) error
	GetItems(ctx context.Context) ([]entity.Item, error)
	GetItemByID(ctx context.Context, id int64) (entity.Item, error)
	DeleteItem(ctx context.Context, id int64) error
	UpdateItem(ctx context.Context, id int64, item entity.Item) error
}
