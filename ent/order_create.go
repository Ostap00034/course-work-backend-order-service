// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Ostap00034/course-work-backend-order-service/ent/order"
	"github.com/google/uuid"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (oc *OrderCreate) SetTitle(s string) *OrderCreate {
	oc.mutation.SetTitle(s)
	return oc
}

// SetDescription sets the "description" field.
func (oc *OrderCreate) SetDescription(s string) *OrderCreate {
	oc.mutation.SetDescription(s)
	return oc
}

// SetPrice sets the "price" field.
func (oc *OrderCreate) SetPrice(f float32) *OrderCreate {
	oc.mutation.SetPrice(f)
	return oc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (oc *OrderCreate) SetNillablePrice(f *float32) *OrderCreate {
	if f != nil {
		oc.SetPrice(*f)
	}
	return oc
}

// SetAddress sets the "address" field.
func (oc *OrderCreate) SetAddress(s string) *OrderCreate {
	oc.mutation.SetAddress(s)
	return oc
}

// SetLongitude sets the "longitude" field.
func (oc *OrderCreate) SetLongitude(s string) *OrderCreate {
	oc.mutation.SetLongitude(s)
	return oc
}

// SetLatitude sets the "latitude" field.
func (oc *OrderCreate) SetLatitude(s string) *OrderCreate {
	oc.mutation.SetLatitude(s)
	return oc
}

// SetCategoryID sets the "category_id" field.
func (oc *OrderCreate) SetCategoryID(u uuid.UUID) *OrderCreate {
	oc.mutation.SetCategoryID(u)
	return oc
}

// SetClientID sets the "client_id" field.
func (oc *OrderCreate) SetClientID(u uuid.UUID) *OrderCreate {
	oc.mutation.SetClientID(u)
	return oc
}

// SetMasterID sets the "master_id" field.
func (oc *OrderCreate) SetMasterID(u uuid.UUID) *OrderCreate {
	oc.mutation.SetMasterID(u)
	return oc
}

// SetNillableMasterID sets the "master_id" field if the given value is not nil.
func (oc *OrderCreate) SetNillableMasterID(u *uuid.UUID) *OrderCreate {
	if u != nil {
		oc.SetMasterID(*u)
	}
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrderCreate) SetStatus(o order.Status) *OrderCreate {
	oc.mutation.SetStatus(o)
	return oc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oc *OrderCreate) SetNillableStatus(o *order.Status) *OrderCreate {
	if o != nil {
		oc.SetStatus(*o)
	}
	return oc
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrderCreate) SetCreatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCreatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrderCreate) SetUpdatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUpdatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetID sets the "id" field.
func (oc *OrderCreate) SetID(u uuid.UUID) *OrderCreate {
	oc.mutation.SetID(u)
	return oc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oc *OrderCreate) SetNillableID(u *uuid.UUID) *OrderCreate {
	if u != nil {
		oc.SetID(*u)
	}
	return oc
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.Price(); !ok {
		v := order.DefaultPrice
		oc.mutation.SetPrice(v)
	}
	if _, ok := oc.mutation.Status(); !ok {
		v := order.DefaultStatus
		oc.mutation.SetStatus(v)
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := order.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := order.DefaultUpdatedAt()
		oc.mutation.SetUpdatedAt(v)
	}
	if _, ok := oc.mutation.ID(); !ok {
		v := order.DefaultID()
		oc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Order.title"`)}
	}
	if v, ok := oc.mutation.Title(); ok {
		if err := order.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Order.title": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Order.description"`)}
	}
	if v, ok := oc.mutation.Description(); ok {
		if err := order.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Order.description": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Order.price"`)}
	}
	if _, ok := oc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Order.address"`)}
	}
	if v, ok := oc.mutation.Address(); ok {
		if err := order.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Order.address": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Longitude(); !ok {
		return &ValidationError{Name: "longitude", err: errors.New(`ent: missing required field "Order.longitude"`)}
	}
	if v, ok := oc.mutation.Longitude(); ok {
		if err := order.LongitudeValidator(v); err != nil {
			return &ValidationError{Name: "longitude", err: fmt.Errorf(`ent: validator failed for field "Order.longitude": %w`, err)}
		}
	}
	if _, ok := oc.mutation.Latitude(); !ok {
		return &ValidationError{Name: "latitude", err: errors.New(`ent: missing required field "Order.latitude"`)}
	}
	if v, ok := oc.mutation.Latitude(); ok {
		if err := order.LatitudeValidator(v); err != nil {
			return &ValidationError{Name: "latitude", err: fmt.Errorf(`ent: validator failed for field "Order.latitude": %w`, err)}
		}
	}
	if _, ok := oc.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category_id", err: errors.New(`ent: missing required field "Order.category_id"`)}
	}
	if _, ok := oc.mutation.ClientID(); !ok {
		return &ValidationError{Name: "client_id", err: errors.New(`ent: missing required field "Order.client_id"`)}
	}
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Order.status"`)}
	}
	if v, ok := oc.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Order.created_at"`)}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Order.updated_at"`)}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(order.Table, sqlgraph.NewFieldSpec(order.FieldID, field.TypeUUID))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := oc.mutation.Title(); ok {
		_spec.SetField(order.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := oc.mutation.Description(); ok {
		_spec.SetField(order.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := oc.mutation.Price(); ok {
		_spec.SetField(order.FieldPrice, field.TypeFloat32, value)
		_node.Price = value
	}
	if value, ok := oc.mutation.Address(); ok {
		_spec.SetField(order.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := oc.mutation.Longitude(); ok {
		_spec.SetField(order.FieldLongitude, field.TypeString, value)
		_node.Longitude = value
	}
	if value, ok := oc.mutation.Latitude(); ok {
		_spec.SetField(order.FieldLatitude, field.TypeString, value)
		_node.Latitude = value
	}
	if value, ok := oc.mutation.CategoryID(); ok {
		_spec.SetField(order.FieldCategoryID, field.TypeUUID, value)
		_node.CategoryID = value
	}
	if value, ok := oc.mutation.ClientID(); ok {
		_spec.SetField(order.FieldClientID, field.TypeUUID, value)
		_node.ClientID = value
	}
	if value, ok := oc.mutation.MasterID(); ok {
		_spec.SetField(order.FieldMasterID, field.TypeUUID, value)
		_node.MasterID = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	err      error
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
