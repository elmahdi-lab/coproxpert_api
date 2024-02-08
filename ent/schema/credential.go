// file: ent/schema/credential.go

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"ithumans.com/coproxpert/ent/schema/mixin"
)

type Credential struct {
	ent.Schema
}

func (Credential) Fields() []ent.Field {
	return []ent.Field{
		mixin.BaseMixin{}.AddUuid(),
		field.String("password").Sensitive(),
		field.JSON("scope", map[string]interface{}{}),
		field.JSON("rights", map[string]interface{}{}),
		mixin.BaseMixin{}.AddCreatedAt(),
		mixin.BaseMixin{}.AddUpdatedAt(),
	}
}

// Edges of the Credential.
func (Credential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("credential").
			Unique(), // Each credential belongs to one unique user.
	}
}
