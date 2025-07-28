package orderrepositories

import (
	"2kitchen/internal/models"
	"context"
	"encoding/json"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(ctx context.Context, db *pgxpool.Pool) (*OrderRepository, error) {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		restaurant INTEGER NOT NULL,
		items JSONB NOT NULL,
		order_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		status TEXT DEFAULT 'pending',
		payment_status TEXT DEFAULT 'unpaid'
		);
	`
	_, err := db.Exec(ctx, createTableQuery)
	if err != nil {
		return nil, err
	}

	return &OrderRepository{db: db}, nil
}

func (r *OrderRepository) AllOrders(ctx context.Context) ([]models.Order, error) {
	query := `SELECT id, restaurant, items, status, order_time, payment_status FROM orders`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]models.Order, 0)
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.Restaurant, &order.Items, &order.Status, &order.OrderTime, &order.PaymentStatus)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order models.CreateOrder) (int, error) {
	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		log.Fatal(err)
	}
	query := `
		INSERT INTO orders (restaurant, items)
		VALUES ($1, $2)
		RETURNING id
	`
	var id int
	err = r.db.QueryRow(ctx, query, order.Restaurant, string(itemsJSON)).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
