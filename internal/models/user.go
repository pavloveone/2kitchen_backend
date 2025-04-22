package models

type User struct {
	ID         int
	Username   string
	Password   string
	FirstName  string
	LastName   string
	MiddleName string
	Email      string
	CreatedOn  string
}

type CreateUserRequest struct {
	Username   string `json:"username" validate:"required,min=3,max=32"`
	Password   string `json:"password" validate:"required,min=8,max=64"`
	FirstName  string `json:"firstName" validate:"required,alpha,min=2,max=32"`
	LastName   string `json:"lastName" validate:"required,alpha,min=2,max=32"`
	MiddleName string `json:"middleName" validate:"omitempty,alpha,min=2,max=32"`
	Email      string `json:"email" validate:"required,email"`
}

type UserResponse struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName,omitempty"`
	Email      string `json:"email"`
	CreatedOn  string `json:"createdOn"`
}

type LogInUser struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}
