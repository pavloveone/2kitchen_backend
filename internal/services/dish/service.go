package dishservices

import (
	"2kitchen/internal/models"
	dishrepositories "2kitchen/internal/repositories/dish"
	"context"
)

type DishService struct {
	repo *dishrepositories.DishRepository
}

func NewDishService(repo *dishrepositories.DishRepository) *DishService {
	return &DishService{repo: repo}
}

func (s *DishService) GetAllDishes(ctx context.Context) ([]models.Dish, error) {
	return s.repo.AllDishes(ctx)
}

func (s *DishService) GetRestDishes(ctx context.Context, id int) ([]models.Dish, error) {
	return s.repo.RestaurantDishes(ctx, id)
}

func (s *DishService) DishById(ctx context.Context, rest, dish int) (models.Dish, error) {
	return s.repo.DishById(ctx, rest, dish)
}

func (s *DishService) AddDish(ctx context.Context, newDish models.ModificationDish) (int, error) {
	return s.repo.AddDish(ctx, newDish)
}

func (s *DishService) RemoveDish(ctx context.Context, dish models.ModificationDish) error {
	return s.repo.RemoveDish(ctx, dish)
}
