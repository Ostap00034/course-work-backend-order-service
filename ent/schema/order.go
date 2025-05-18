package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Order — модель заказа.
type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("title").NotEmpty().Comment("Название"),
		field.String("description").
			NotEmpty().
			Comment("Описание заказа"),
		field.Float32("price").Default(0).Comment("Цена"),
		field.String("address").NotEmpty().Comment("Адрес заказа"),
		field.String("longitude").NotEmpty().Comment("Долгота"),
		field.String("latitude").NotEmpty().Comment("Широта"),
		field.UUID("category_id", uuid.UUID{}).Comment("ID категории"),
		field.UUID("client_id", uuid.UUID{}).Comment("ID автора"),
		field.UUID("master_id", uuid.UUID{}).Optional().Comment("ID исполнителя"),
		field.Enum("status").Values("active", "in_progress", "cancel", "done").Default("active"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Order) Edges() []ent.Edge {
	return nil
}
