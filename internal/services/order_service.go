package services

import (
	"2kitchen/internal/models"
	"2kitchen/internal/repositories"
)

type OrderService struct {
	repo *repositories.OrderRepository
}

func NewOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) AllOrders() ([]models.Order, error) {
	return s.repo.AllOrders()
}

func (s *OrderService) CreateOrder(order models.CreateOrder) (int, error) {
	return s.repo.CreateOrder(order)
}
