package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"time"
)

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) AddUuid() ent.Field {
	return field.UUID("id", uuid.UUID{}).
		Default(uuid.New).
		StorageKey("id").
		Unique()
}

// AddCreatedAt adds the created_at field to the schema.
func (BaseMixin) AddCreatedAt() ent.Field {
	return field.Time("created_at").Default(time.Now)
}

// AddUpdatedAt adds the updated_at field to the schema.
func (BaseMixin) AddUpdatedAt() ent.Field {
	return field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now)
}
