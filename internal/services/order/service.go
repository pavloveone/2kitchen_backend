package orderservices

import (
	"2kitchen/internal/models"
	orderrepositories "2kitchen/internal/repositories/order"
	"context"
)

type OrderService struct {
	repo *orderrepositories.OrderRepository
}

func NewOrderService(repo *orderrepositories.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) AllOrders(ctx context.Context) ([]models.Order, error) {
	return s.repo.AllOrders(ctx)
}

func (s *OrderService) CreateOrder(ctx context.Context, order models.CreateOrder) (int, error) {
	return s.repo.CreateOrder(ctx, order)
}
