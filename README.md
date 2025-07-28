# 2kitchen Backend

Backend service for **2kitchen** – a digital solution for restaurants. It allows guests to order dishes, call a waiter, or request a bill, while restaurants manage menus and monitor orders via an admin dashboard.

🔗 Frontend: [2kitchen_frontend GitHub Repo](https://github.com/pavloveone/2kitchen_frontend)

---

## 🛠 Technologies

Go, Fiber, PostgreSQL, Docker, jwt-go, testify, godotenv

---

## 📦 API Overview

### 🍽 Dishes (`/dishes`)
- `GET /dishes` – Get all dishes  
- `GET /dishes/:restId` – Get dishes by restaurant  
- `GET /dishes/:restId/:id` – Get dish by ID  
- `POST /dishes` – Add a new dish  
- `DELETE /dishes` – Remove a dish  

### 🧾 Orders (`/orders`)
- `GET /orders` – Get all orders  
- `POST /orders` – Create an order  

### 👤 Users (`/users`)
> Auth is fully implemented, but frontend integration is planned.

- `GET /users` – Get all users  
- `GET /users/:id` – Get user by ID  
- `POST /users` – Register user  
- `POST /users/login` – Log in and get token  

---

## 🔐 Auth

- Password hashing & validation  
- JWT token generation  
- Auth middleware for route protection  

---

## 🧪 Testing

Unit tests are written for:  
- Dishes  
- Orders  

Run tests:
```bash
TEST_DATABASE_URL=postgres://kitchen_user:kitchen_pass@localhost:5432/kitchen_test?sslmode=disable go test ./...
```

---

## 🐳 Run with Docker
```bash
docker-compose down -v # additional
docker compose up --build # docker-compose up -d
go run ./..
```

Docker setup includes:
- Backend container (Go app)
- PostgreSQL container with preconfigured user and database

Ensure your `.env` contains:

```env
DATABASE_URL=postgres://kitchen_user:kitchen_pass@db:5432/kitchen_db
```

---

## 📌 Roadmap
- [x] Switch to PostgreSQL from SQLite
- [x] Dockerize the backend
- [x] Add authorization middleware
- [x] Implement user authentication
- [x] Write unit tests
- [ ] Connect user service to frontend
- [ ] Add Swagger documentation
- [ ] Add integration tests
