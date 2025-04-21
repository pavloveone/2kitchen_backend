package orderservices

import (
	"2kitchen/internal/models"
	orderrepositories "2kitchen/internal/repositories/order"
)

type OrderService struct {
	repo *orderrepositories.OrderRepository
}

func NewOrderService(repo *orderrepositories.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) AllOrders() ([]models.Order, error) {
	return s.repo.AllOrders()
}

func (s *OrderService) CreateOrder(order models.CreateOrder) (int, error) {
	return s.repo.CreateOrder(order)
}
