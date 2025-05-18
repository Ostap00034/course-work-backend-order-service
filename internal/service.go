package order

import (
	"context"

	"github.com/Ostap00034/course-work-backend-order-service/ent"
	"github.com/google/uuid"
)

type Service interface {
	Get(ctx context.Context, id uuid.UUID) (*ent.Order, error)
	GetAll(ctx context.Context, categories_ids []uuid.UUID, status string, client_id uuid.UUID, master_id uuid.UUID) ([]*ent.Order, error)
	GetAllActive(ctx context.Context, categories_ids []uuid.UUID, client_id, master_id uuid.UUID) ([]*ent.Order, error)
	Create(ctx context.Context, title, description, address, longitude, latitude, status string, price float32, category_id uuid.UUID, client_id, master_id uuid.UUID) (*ent.Order, error)
	Update(ctx context.Context, id uuid.UUID, title, description, address, longitude, latitude, status string, price float32, category_id uuid.UUID, client_id, master_id uuid.UUID) (*ent.Order, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo Repoistory
}

func NewService(r Repoistory) Service {
	return &service{repo: r}
}

func (s *service) Get(ctx context.Context, id uuid.UUID) (*ent.Order, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) GetAll(ctx context.Context, categories_ids []uuid.UUID, status string, client_id, master_id uuid.UUID) ([]*ent.Order, error) {
	return s.repo.GetAll(ctx, categories_ids, status, client_id, master_id)
}

func (s *service) GetAllActive(ctx context.Context, categories_ids []uuid.UUID, client_id, master_id uuid.UUID) ([]*ent.Order, error) {
	return s.repo.GetAllActive(ctx, categories_ids)
}

func (s *service) Create(ctx context.Context, title, description, address, longitude, latitude, status string, price float32, category_id uuid.UUID, client_id uuid.UUID, master_id uuid.UUID) (*ent.Order, error) {
	return s.repo.Create(ctx, title, description, address, longitude, latitude, status, price, category_id, client_id, master_id)
}

func (s *service) Update(ctx context.Context, id uuid.UUID, title, description, address, longitude, latitude, status string, price float32, category_id uuid.UUID, client_id, master_id uuid.UUID) (*ent.Order, error) {
	return s.repo.Update(ctx, id, title, description, address, longitude, latitude, status, price, category_id, client_id, master_id)
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
