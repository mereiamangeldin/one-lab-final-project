package pg

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/entity"
)

func (p *Postgres) CreateItem(ctx context.Context, i *entity.Item) error {
	query := fmt.Sprintf(`
			INSERT INTO %s (
			                name, -- 1 
			                description, -- 2
			                cost -- 3
			                )
			VALUES ($1, $2, $3)
			`, itemsTable)

	_, err := p.Pool.Exec(ctx, query, i.Name, i.Description, i.Cost)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetItems(ctx context.Context) ([]entity.Item, error) {
	var items []entity.Item

	query := fmt.Sprintf("SELECT * FROM %s", itemsTable)

	err := pgxscan.Select(ctx, p.Pool, &items, query)

	if err != nil {
		return nil, err
	}
	return items, nil
}

func (p *Postgres) GetItemByID(ctx context.Context, id int64) (entity.Item, error) {
	var item entity.Item

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", itemsTable)

	err := pgxscan.Get(ctx, p.Pool, &item, query, id)

	if err != nil {
		return entity.Item{}, err
	}
	return item, nil
}

func (p *Postgres) DeleteItem(ctx context.Context, id int64) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", itemsTable)

	_, err := p.Pool.Exec(ctx, query, id)

	if err != nil {
		return err
	}
	return nil
}

func (p *Postgres) UpdateItem(ctx context.Context, id int64, item entity.Item) error {

	query := fmt.Sprintf("UPDATE %s SET name = $1, description = $2, cost = $3 WHERE id = $4", itemsTable)

	_, err := p.Pool.Exec(ctx, query, item.Name, item.Description, item.Cost, id)

	if err != nil {
		return err
	}
	return nil
}
