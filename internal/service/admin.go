package service

import (
	"context"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/entity"
)

func (m *Manager) CreateItem(ctx context.Context, i *entity.Item) error {
	err := m.Repository.CreateItem(ctx, i)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetItems(ctx context.Context) ([]entity.Item, error) {
	return m.Repository.GetItems(ctx)
}

func (m *Manager) GetItemByID(ctx context.Context, id int64) (entity.Item, error) {
	return m.Repository.GetItemByID(ctx, id)
}

func (m *Manager) DeleteItem(ctx context.Context, id int64) error {
	return m.Repository.DeleteItem(ctx, id)
}

func (m *Manager) UpdateItem(ctx context.Context, id int64, item entity.Item) error {
	return m.Repository.UpdateItem(ctx, id, item)
}
