// file: ent/schema/user.go

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"ithumans.com/coproxpert/ent/schema/mixin"
)

type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.BaseMixin{},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		mixin.BaseMixin{}.AddUuid(),
		field.String("username").Unique(),
		field.String("first_name"),
		field.String("last_name"),
		field.String("password").Sensitive(),
		mixin.BaseMixin{}.AddCreatedAt(),
		mixin.BaseMixin{}.AddUpdatedAt(),
	}
}
