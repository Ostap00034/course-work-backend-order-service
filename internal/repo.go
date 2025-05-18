package order

import (
	"context"
	"errors"

	"github.com/Ostap00034/course-work-backend-order-service/ent"
	"github.com/Ostap00034/course-work-backend-order-service/ent/order"
	"github.com/google/uuid"
)

var (
	ErrOrderNotFound      = errors.New("заказ не найден")
	ErrGetOrderFailed     = errors.New("ошибка получения заказа")
	ErrGetAllOrderFailed  = errors.New("ошибка получения всех заказов")
	ErrOrderAlreadyExists = errors.New("заказ с таким названием уже существует")
	ErrCreateOrderFailed  = errors.New("ошибка при создании заказа")
	ErrUpdateOrderFailed  = errors.New("ошибка при обновлении заказа")
	ErrInvalidId          = errors.New("неправильный формат UUID")
)

type Repoistory interface {
	Get(ctx context.Context, id uuid.UUID) (*ent.Order, error)
	GetAll(ctx context.Context, categories_ids []uuid.UUID, status string, client_id, master_id uuid.UUID) ([]*ent.Order, error)
	GetAllActive(ctx context.Context, categories_ids []uuid.UUID) ([]*ent.Order, error)
	Create(ctx context.Context, title, description, address, longitude, latitude, status string, price float32, category_id uuid.UUID, client_id uuid.UUID, master_id uuid.UUID) (*ent.Order, error)
	Update(ctx context.Context, id uuid.UUID, title, description, address, longitude, latitude, status string, price float32, category_id uuid.UUID, client_id uuid.UUID, master_id uuid.UUID) (*ent.Order, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type repo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) Repoistory {
	return &repo{client: client}
}

func (r *repo) Get(ctx context.Context, id uuid.UUID) (*ent.Order, error) {
	order, err := r.client.Order.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrOrderNotFound
		}
		return nil, ErrGetOrderFailed
	}

	return order, nil
}

func (r *repo) GetAll(
	ctx context.Context,
	categories_ids []uuid.UUID,
	status string,
	client_id,
	master_id uuid.UUID,
) ([]*ent.Order, error) {
	q := r.client.Order.Query()

	if len(categories_ids) > 0 {
		q = q.Where(order.CategoryIDIn(categories_ids...))
	}

	if status != "" {
		q = q.Where(order.StatusEQ(order.Status(status)))
	}

	if client_id != uuid.Nil {
		q = q.Where(order.ClientIDEQ(client_id))
	}

	if master_id != uuid.Nil {
		q = q.Where(order.MasterIDEQ(master_id))
	}

	orders, err := q.All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrGetAllOrderFailed
		}
		return nil, err
	}

	return orders, nil
}

func (r *repo) GetAllActive(ctx context.Context, categories_ids []uuid.UUID) ([]*ent.Order, error) {
	q := r.client.Order.Query()

	if len(categories_ids) > 0 {
		q = q.Where(order.CategoryIDIn(categories_ids...))
	}

	q = q.Where(order.StatusEQ(order.StatusActive))

	orders, err := q.All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrGetAllOrderFailed
		}
		return nil, err
	}
	return orders, nil
}

func (r *repo) Create(ctx context.Context, title, description, address, longitude, latitude, status string, price float32, category_id uuid.UUID, client_id uuid.UUID, master_id uuid.UUID) (*ent.Order, error) {
	order, err := r.client.Order.Create().
		SetTitle(title).
		SetDescription(description).
		SetAddress(address).
		SetLongitude(longitude).
		SetLatitude(latitude).
		SetStatus(order.Status(status)).
		SetPrice(price).SetCategoryID(category_id).
		SetClientID(client_id).
		SetStatus(order.StatusActive).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, ErrOrderAlreadyExists
		}
		return nil, ErrCreateOrderFailed
	}

	return order, nil
}

func (r *repo) Update(ctx context.Context, id uuid.UUID, title, description, address, longitude, latitude, status string, price float32, category_id uuid.UUID, client_id uuid.UUID, master_id uuid.UUID) (*ent.Order, error) {
	builder := r.client.Order.UpdateOneID(id)

	if title != "" {
		builder = builder.SetTitle(title)
	}
	if description != "" {
		builder = builder.SetDescription(description)
	}
	if address != "" {
		builder = builder.SetAddress(address)
	}
	if longitude != "" {
		builder = builder.SetLongitude(longitude)
	}
	if latitude != "" {
		builder = builder.SetLatitude(latitude)
	}
	if status != "" {
		builder = builder.SetStatus(order.Status(status))
	}
	if price != 0 {
		builder = builder.SetPrice(price)
	}

	if category_id != uuid.Nil {
		builder = builder.SetCategoryID(category_id)
	}

	if master_id != uuid.Nil {
		builder = builder.SetMasterID(master_id)
	}

	updated, err := builder.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrOrderNotFound
		}
		return nil, ErrUpdateOrderFailed
	}

	return updated, nil
}

func (r *repo) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.client.Order.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrOrderNotFound
		} else {
			return ErrGetAllOrderFailed
		}
	}

	return nil
}
