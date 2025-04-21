package orderrepositories

import (
	"2kitchen/internal/models"
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(dbPath string) (*OrderRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		restaurant INTEGER,
		items TEXT,
		order_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		status TEXT DEFAULT 'pending',
		payment_status TEXT DEFAULT 'unpaid'
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	return &OrderRepository{db: db}, nil
}

func (r *OrderRepository) AllOrders() ([]models.Order, error) {
	query := `SELECT id, restaurant, items, status, order_time, payment_status FROM orders`
	rows, err := r.db.Query(query)
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

func (r *OrderRepository) CreateOrder(order models.CreateOrder) (int, error) {
	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		log.Fatal(err)
	}
	query := `
		INSERT INTO orders (restaurant, items)
		VALUES (?, ?)
	`
	result, err := r.db.Exec(query, order.Restaurant, string(itemsJSON))
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertId), nil
}
