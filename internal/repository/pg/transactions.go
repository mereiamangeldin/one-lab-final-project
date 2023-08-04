package pg

import (
	"context"
	"fmt"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/entity"
)

func (p *Postgres) CreateTransaction(ctx context.Context, transactions entity.Transactions) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                user_id, -- 1 
			                item_id, -- 2
			                amount, -- 3
			                created_at-- 4
			                )
			VALUES ($1, $2, $3, $4)
			`, transactionTable)

	_, err := p.Pool.Exec(ctx, query, transactions.UserID, transactions.ItemID, transactions.Amount, transactions.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
