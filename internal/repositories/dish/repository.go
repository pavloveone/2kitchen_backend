package dishrepositories

import (
	"2kitchen/internal/models"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DishRepository struct {
	db *pgxpool.Pool
}

func NewDishRepository(ctx context.Context, db *pgxpool.Pool) (*DishRepository, error) {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS dishes (
		id SERIAL PRIMARY KEY, 
		restaurant INTEGER,
		name TEXT,
		description TEXT,
		price DOUBLE PRECISION,
		image TEXT,
		protein DOUBLE PRECISION,
		fat DOUBLE PRECISION,
		carbs DOUBLE PRECISION,
		calories DOUBLE PRECISION
	);
	`
	_, err := db.Exec(ctx, createTableQuery)
	if err != nil {
		return nil, err
	}

	return &DishRepository{db: db}, nil
}

func (r *DishRepository) AllDishes(ctx context.Context) ([]models.Dish, error) {
	query := `SELECT id, restaurant, name, description, price, image, protein, fat, carbs, calories FROM dishes`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes []models.Dish
	for rows.Next() {
		var dish models.Dish
		err := rows.Scan(&dish.ID, &dish.Restaurant, &dish.Name, &dish.Description, &dish.Price, &dish.Image, &dish.Protein, &dish.Fat, &dish.Carbs, &dish.Calories)
		if err != nil {
			return nil, err
		}
		dishes = append(dishes, dish)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dishes, nil
}

func (r *DishRepository) RestaurantDishes(ctx context.Context, restId int) ([]models.Dish, error) {
	query := `SELECT id, restaurant, name, description, price, image, protein, fat, carbs, calories FROM dishes WHERE restaurant = $1`
	rows, err := r.db.Query(ctx, query, restId)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	dishes := make([]models.Dish, 0)
	for rows.Next() {
		var dish models.Dish
		err := rows.Scan(&dish.ID, &dish.Restaurant, &dish.Name, &dish.Description, &dish.Price, &dish.Image, &dish.Protein, &dish.Fat, &dish.Carbs, &dish.Calories)
		if err != nil {
			return nil, err
		}
		dishes = append(dishes, dish)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dishes, nil
}

func (r *DishRepository) DishById(ctx context.Context, restId, dishId int) (models.Dish, error) {
	query := "SELECT id, restaurant, name, description, price, image, protein, fat, carbs, calories FROM dishes WHERE restaurant = $1 AND id = $2"
	row := r.db.QueryRow(ctx, query, restId, dishId)

	var dish models.Dish
	err := row.Scan(&dish.ID, &dish.Restaurant, &dish.Name, &dish.Description, &dish.Price, &dish.Image, &dish.Protein, &dish.Fat, &dish.Carbs, &dish.Calories)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Dish{}, errors.New("dish not found")
		}
		return models.Dish{}, err
	}

	return dish, nil
}

func (r *DishRepository) AddDish(ctx context.Context, newDish models.ModificationDish) (int, error) {
	query := `
		INSERT INTO dishes (restaurant, name, description, price, image, protein, fat, carbs, calories)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	var id int
	err := r.db.QueryRow(ctx, query,
		newDish.Restaurant,
		newDish.Name,
		newDish.Description,
		newDish.Price,
		newDish.Image,
		newDish.Protein,
		newDish.Fat,
		newDish.Carbs,
		newDish.Calories,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *DishRepository) RemoveDish(ctx context.Context, dish models.ModificationDish) error {
	query := "DELETE FROM dishes WHERE restaurant = $1 AND id = $2"

	_, err := r.db.Exec(ctx, query, dish.Restaurant, dish.ID)
	if err != nil {
		return fmt.Errorf("failed to delete dish: %w", err)
	}

	return nil
}
