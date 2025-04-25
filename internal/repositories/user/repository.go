package userrepositories

import (
	"2kitchen/internal/auth"
	"2kitchen/internal/models"
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(dbPath string) (*UserRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		middle_name TEXT,
		email TEXT NOT NULL UNIQUE,
		created_on DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}
	return &UserRepository{db: db}, nil
}

func (r *UserRepository) AllUsers() ([]models.UserResponse, error) {
	query := `SELECT id, username, first_name, last_name, middle_name, email, created_on from users`
	rows, err := r.db.Query(query)
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

func (r *UserRepository) UserById(id int) (models.UserResponse, error) {
	query := `SELECT id, username, first_name, last_name, middle_name, email, created_on from users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var user models.UserResponse
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.MiddleName, &user.Email, &user.CreatedOn)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserResponse{}, errors.New("user not found")
		}
		return models.UserResponse{}, nil
	}
	return user, nil
}

func (r *UserRepository) AddUser(newUser models.CreateUserRequest) (int, error) {
	hashPass, err := auth.HashPassword(newUser.Password)
	if err != nil {
		return 0, err
	}
	query := `
	INSERT INTO users (username, password, first_name, last_name, middle_name, email)
	VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(query, newUser.Username, hashPass, newUser.FirstName, newUser.LastName, newUser.MiddleName, newUser.Email)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastInsertId), nil
}

func (r *UserRepository) LogIn(loginUser models.LogInUser) (models.LoginResponse, error) {
	query := `SELECT id, username, password, first_name, last_name, middle_name, email, created_on FROM users WHERE username = ?`
	row := r.db.QueryRow(query, loginUser.Username)

	var user models.UserResponse
	var hashedPass string
	err := row.Scan(&user.ID, &user.Username, &hashedPass, &user.FirstName, &user.LastName, &user.MiddleName, &user.Email, &user.CreatedOn)
	if err != nil {
		if err == sql.ErrNoRows {
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
