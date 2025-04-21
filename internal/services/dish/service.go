package dishservices

import (
	"2kitchen/internal/models"
	dishrepositories "2kitchen/internal/repositories/dish"
)

type DishService struct {
	repo *dishrepositories.DishRepository
}

func NewDishService(repo *dishrepositories.DishRepository) *DishService {
	return &DishService{repo: repo}
}

func (s *DishService) GetAllDishes() ([]models.Dish, error) {
	return s.repo.AllDishes()
}

func (s *DishService) GetRestDishes(id int) ([]models.Dish, error) {
	return s.repo.RestaurantDishes(id)
}

func (s *DishService) DishById(rest, dish int) (models.Dish, error) {
	return s.repo.DishById(rest, dish)
}

func (s *DishService) AddDish(newDish models.ModificationDish) (int, error) {
	return s.repo.AddDish(newDish)
}

func (s *DishService) RemoveDish(dish models.ModificationDish) error {
	return s.repo.RemoveDish(dish)
}
