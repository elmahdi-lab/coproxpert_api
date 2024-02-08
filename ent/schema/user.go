// file: ent/schema/user.go

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		mixin.BaseMixin{}.AddCreatedAt(),
		mixin.BaseMixin{}.AddUpdatedAt(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("credential", Credential.Type).
			Unique(), // Each user has one unique credential.
	}
}
