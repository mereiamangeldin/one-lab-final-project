package pg

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/entity"
	"strings"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                username, -- 1 
			                first_name, -- 2
			                last_name, -- 3
			                hashed_password -- 4
			                )
			VALUES ($1, $2, $3, $4) RETURNING id
			`, usersTable)
	queryBalance := fmt.Sprintf(`
			INSERT INTO %s (
			                user_id, -- 1 
			                amount -- 2
			                )
			VALUES ($1, $2)
			`, balanceTable)
	fmt.Println(u)
	var userID int64
	err := p.Pool.QueryRow(ctx, query, u.Username, u.FirstName, u.LastName, u.Password).Scan(&userID)
	if err != nil {
		return err
	}

	_, err = p.Pool.Exec(ctx, queryBalance, userID, 10000)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetUser(ctx context.Context, username string) (*entity.User, error) {
	user := new(entity.User)

	query := fmt.Sprintf("SELECT id, username, first_name, last_name, hashed_password FROM %s WHERE username = $1", usersTable)

	err := pgxscan.Get(ctx, p.Pool, user, query, strings.TrimSpace(username))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Postgres) GetUserBalance(ctx context.Context, userID int64) (*entity.Balance, error) {
	balance := new(entity.Balance)

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", balanceTable)

	err := pgxscan.Get(ctx, p.Pool, balance, query, userID)
	if err != nil {
		return &entity.Balance{}, err
	}

	return balance, nil
}

func (p *Postgres) UpdateUserBalance(ctx context.Context, balance entity.Balance) error {

	query := fmt.Sprintf("UPDATE %s SET amount = $1 WHERE user_id = $2", balanceTable)

	_, err := p.Pool.Exec(ctx, query, balance.Amount, balance.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetUserItems(ctx context.Context, userID int64) ([]entity.Item, error) {
	var itemsID []int64
	query := fmt.Sprintf("SELECT item_id from %s group by user_id where user_id = $1", transactionTable)
	err := pgxscan.Select(ctx, p.Pool, &itemsID, query)
	if err != nil {
		return nil, err
	}
	query = fmt.Sprintf("SELECT * FROM %s WHERE id = ANY(%s)", itemsTable, itemsID)
	var items []entity.Item
	err = pgxscan.Select(ctx, p.Pool, &items, query)
	if err != nil {
		return nil, err
	}
	return items, nil
}
