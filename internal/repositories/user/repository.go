package userrepositories

import (
	"2kitchen/internal/auth"
	"2kitchen/internal/models"
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(ctx context.Context, db *pgxpool.Pool) (*UserRepository, error) {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users ( 
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		middle_name TEXT,
		email TEXT NOT NULL UNIQUE,
		created_on TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(ctx, createTableQuery)
	if err != nil {
		return nil, err
	}
	return &UserRepository{db: db}, nil
}

func (r *UserRepository) AllUsers(ctx context.Context) ([]models.UserResponse, error) {
	query := `SELECT id, username, first_name, last_name, middle_name, email, created_on FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.UserResponse, 0)
	for rows.Next() {
		var user models.UserResponse
		err := rows.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.MiddleName, &user.Email, &user.CreatedOn)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) UserById(ctx context.Context, id int) (models.UserResponse, error) {
	query := `SELECT id, username, first_name, last_name, middle_name, email, created_on from users WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var user models.UserResponse
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.MiddleName, &user.Email, &user.CreatedOn)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.UserResponse{}, errors.New("user not found")
		}
		return models.UserResponse{}, nil
	}
	return user, nil
}

func (r *UserRepository) AddUser(ctx context.Context, newUser models.CreateUserRequest) (int, error) {
	hashPass, err := auth.HashPassword(newUser.Password)
	if err != nil {
		return 0, err
	}
	query := `
	INSERT INTO users (username, password, first_name, last_name, middle_name, email)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`
	var id int
	err = r.db.QueryRow(ctx, query, newUser.Username, hashPass, newUser.FirstName, newUser.LastName, newUser.MiddleName, newUser.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) LogIn(ctx context.Context, loginUser models.LogInUser) (models.LoginResponse, error) {
	query := `SELECT id, username, password, first_name, last_name, middle_name, email, created_on FROM users WHERE username = $1`
	row := r.db.QueryRow(ctx, query, loginUser.Username)

	var user models.UserResponse
	var hashedPass string
	err := row.Scan(&user.ID, &user.Username, &hashedPass, &user.FirstName, &user.LastName, &user.MiddleName, &user.Email, &user.CreatedOn)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.LoginResponse{}, errors.New("user not found")
		}
		return models.LoginResponse{}, err
	}
	if ok := auth.CheckPasswordHash(loginUser.Password, hashedPass); !ok {
		return models.LoginResponse{}, errors.New("an error occurred while checking password")
	}
	access, refresh, err := auth.GenerateTokens(user.ID)
	if err != nil {
		return models.LoginResponse{}, nil
	}
	return models.LoginResponse{
		User:         user,
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
