package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/entity"
	"github.com/mereiamangeldin/One-lab-Homework-1/pkg/util"
	"time"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = m.Repository.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) Login(ctx context.Context, username, password string) (string, error) {
	user, err := m.Repository.GetUser(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found")
		}

		return "", fmt.Errorf("get user err: %w", err)
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return "", fmt.Errorf("incorrect password: %w", err)
	}

	accessToken, err := m.Token.CreateToken(user.ID, m.Config.Token.TimeToLive)
	if err != nil {
		return "", fmt.Errorf("create token err: %w", err)
	}

	return accessToken, nil
}

func (m *Manager) VerifyToken(token string) (int64, error) {
	payload, err := m.Token.ValidateToken(token)
	if err != nil {
		return 0, fmt.Errorf("validate token err: %w", err)
	}

	return payload.UserID, nil
}

func (m *Manager) BuyItem(ctx context.Context, itemID int64, userID int64) error {
	item, err := m.GetItemByID(ctx, itemID)
	if err != nil {
		return err
	}
	userBalance, err := m.Repository.GetUserBalance(ctx, userID)
	if err != nil {
		return err
	}
	if item.Cost > userBalance.Amount {
		return errors.New("balance is not enough")
	}
	transaction := entity.Transactions{
		UserID:    userID,
		ItemID:    itemID,
		Amount:    item.Cost,
		CreatedAt: time.Now(),
	}
	err = m.Repository.CreateTransaction(ctx, transaction)

	if err != nil {
		return err
	}

	userBalance.Amount = userBalance.Amount - item.Cost
	err = m.Repository.UpdateUserBalance(ctx, *userBalance)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) GetUserItems(ctx context.Context, userID int64) ([]entity.Item, error) {
	return m.Repository.GetUserItems(ctx, userID)
}
