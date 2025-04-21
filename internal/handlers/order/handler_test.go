package orderhandlers_test

import (
	orderhandler "2kitchen/internal/handlers/order"
	"2kitchen/internal/models"
	orderrepositories "2kitchen/internal/repositories/order"
	orderroutes "2kitchen/internal/routes/order"
	orderservices "2kitchen/internal/services/order"
	"bytes"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func setupTestApp() *fiber.App {
	// Создаем SQLite в памяти
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("Error initializing in-memory DB:", err)
	}
	defer db.Close()

	// Создаём репозиторий с этим db
	repo, err := orderrepositories.NewOrderRepository(":memory:")
	if err != nil {
		log.Fatal("Error initializing orders repository:", err)
	}

	// Создаём сервис с репозиторием
	service := orderservices.NewOrderService(repo)

	// Создаём хэндлер с сервисом
	handler := orderhandler.NewOrderHandler(service)

	// Настроим маршруты
	app := fiber.New()
	orderroutes.SetupOrderRoutes(app, handler)

	// Добавим тестовые данные в базу
	addTestOrders(repo)

	return app
}

func addTestOrders(repo *orderrepositories.OrderRepository) {
	orders := []models.CreateOrder{
		{Restaurant: 1, Items: []models.OrderItem{}},
		{Restaurant: 1, Items: []models.OrderItem{}},
	}

	for _, order := range orders {
		if _, err := repo.CreateOrder(order); err != nil {
			log.Fatal("Error adding test orders:", err)
		}
	}
}

func TestAllOrders(t *testing.T) {
	app := setupTestApp()
	// Создаём тестовый запрос
	req := httptest.NewRequest("GET", "/orders", nil)
	resp, err := app.Test(req)

	// Проверяем ошибки
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Проверяем тело ответа
	var orders []models.Order
	err = json.NewDecoder(resp.Body).Decode(&orders)
	require.NoError(t, err)

	// Проверяем, что данные корректные
	require.Len(t, orders, 2)
	require.Equal(t, 1, orders[0].Restaurant)
}

func TestCreateOrder(t *testing.T) {
	app := setupTestApp()
	newOrder := models.CreateOrder{
		Restaurant: 1,
		Items:      []models.OrderItem{},
	}
	body, err := json.Marshal(newOrder)
	if err != nil {
		t.Fatalf("could not marshal newOrder: %v", err)
	}
	req := httptest.NewRequest("POST", "/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	var res map[string]any
	err = json.NewDecoder(resp.Body).Decode(&res)
	require.NoError(t, err)

	idFloat, ok := res["id"].(float64)
	require.True(t, ok, "expected id to be a number, got %#v", res["id"])

	id := int(idFloat)
	require.True(t, id > 0, "expected ID > 0, got %d", id)
	require.Equal(t, 3, id, "expected new ID to be 3, got %d", id)
}
