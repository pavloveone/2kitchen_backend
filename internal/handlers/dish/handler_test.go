package dishhandlers_test

import (
	dishhandlers "2kitchen/internal/handlers/dish"
	"2kitchen/internal/models"
	dishrepositories "2kitchen/internal/repositories/dish"
	dishroutes "2kitchen/internal/routes/dish"
	dishservices "2kitchen/internal/services/dish"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

var testDB *pgxpool.Pool
var ctx = context.Background()

func TestMain(m *testing.M) {
	var err error

	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		log.Fatal("TEST_DATABASE_URL not set")
	}

	testDB, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to test DB: %v", err)
	}

	_, err = testDB.Exec(ctx, `TRUNCATE TABLE dishes RESTART IDENTITY CASCADE`)
	if err != nil {
		log.Fatalf("Failed to truncate tables: %v", err)
	}

	code := m.Run()

	testDB.Close()
	os.Exit(code)
}

func setupTestApp() *fiber.App {

	repo, err := dishrepositories.NewDishRepository(ctx, testDB)
	if err != nil {
		log.Fatal("Error initializing dishes repository:", err)
	}

	service := dishservices.NewDishService(repo)
	handler := dishhandlers.NewDishHandler(service, ctx)

	app := fiber.New()
	dishroutes.SetupDishRoutes(app, handler)

	addTestDishes(ctx, repo)

	return app
}

func addTestDishes(ctx context.Context, repo *dishrepositories.DishRepository) {
	dishes := []models.ModificationDish{
		{
			ID:          1,
			Name:        "Паста Карбонара",
			Price:       1200,
			Description: "Спагетти, бекон, сливки, яйца, пармезан",
			Protein:     25,
			Fat:         32,
			Carbs:       45,
			Calories:    568,
			Restaurant:  1,
		},
		{
			ID:          2,
			Name:        "Стейк Рибай",
			Price:       2400,
			Description: "Говяжий стейк с овощами гриль",
			Protein:     38,
			Fat:         28,
			Carbs:       5,
			Calories:    424,
			Restaurant:  2,
		},
	}

	for _, dish := range dishes {
		if _, err := repo.AddDish(ctx, dish); err != nil {
			log.Fatal("Error adding test dishes:", err)
		}
	}
}

func TestAllDishes(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest("GET", "/dishes", nil)
	resp, err := app.Test(req)

	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)

	var dishes []models.Dish
	err = json.NewDecoder(resp.Body).Decode(&dishes)
	require.NoError(t, err)

	require.Len(t, dishes, 2)
	require.Equal(t, 1, dishes[0].Restaurant)
	require.Equal(t, 2, dishes[1].Restaurant)
}

func TestRestaurantDishes(t *testing.T) {
	app := setupTestApp()

	ids := [2]int{1, 2}

	for _, id := range ids {
		target := fmt.Sprintf("/dishes/%d", id)
		req := httptest.NewRequest("GET", target, nil)
		resp, err := app.Test(req)

		require.NoError(t, err)
		require.Equal(t, fiber.StatusOK, resp.StatusCode)

		var dishes []models.Dish
		err = json.NewDecoder(resp.Body).Decode(&dishes)

		require.NoError(t, err)

		for _, dish := range dishes {
			require.Equal(t, id, dish.Restaurant)
		}
	}
}
