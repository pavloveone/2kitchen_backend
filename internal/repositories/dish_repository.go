package repositories

import (
	"2kitchen/internal/models"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DishRepository struct {
	db *sql.DB
}

func NewDishRepository(dbPath string) (*DishRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS dishes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		restaurant INTEGER,
		name TEXT,
		description TEXT,
		price REAL,
		image TEXT,
		protein REAL,
		fat REAL,
		carbs REAL,
		calories REAL
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	return &DishRepository{db: db}, nil
}

func (r *DishRepository) AllDishes() ([]models.Dish, error) {
	query := "SELECT id, restaurant, name, description, price, image, protein, fat, carbs, calories FROM dishes"
	rows, err := r.db.Query(query)
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

func (r *DishRepository) RestaurantDishes(restId int) ([]models.Dish, error) {
	query := "SELECT id, restaurant, name, description, price, image, protein, fat, carbs, calories FROM dishes WHERE restaurant = ?"
	rows, err := r.db.Query(query, restId)
	if err != nil {
		fmt.Println("Error executing query:", err)
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

func (r *DishRepository) DishById(restId, dishId int) (models.Dish, error) {
	query := "SELECT id, restaurant, name, description, price, image, protein, fat, carbs, calories FROM dishes WHERE restaurant = ? AND id = ?"
	row := r.db.QueryRow(query, restId, dishId)

	var dish models.Dish
	err := row.Scan(&dish.ID, &dish.Name, &dish.Description, &dish.Price, &dish.Image, &dish.Protein, &dish.Fat, &dish.Carbs, &dish.Calories)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Dish{}, errors.New("dish not found")
		}
		return models.Dish{}, err
	}

	return dish, nil
}

func (r *DishRepository) AddDish(newDish models.CreateDish) (int, error) {
	query := `
		INSERT INTO dishes (restaurant, name, description, price, image, protein, fat, carbs, calories)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(query, newDish.Restaurant, newDish.Name, newDish.Description, newDish.Price, newDish.Image, newDish.Protein, newDish.Fat, newDish.Carbs, newDish.Calories)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertId), nil
}
