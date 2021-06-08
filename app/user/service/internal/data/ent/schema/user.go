package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("nickname"),
		field.String("avatar"),
		field.String("password_hash"),
		field.Time("created_at").Default(time.Now).
			SchemaType(map[string]string{dialect.MySQL: "datetime"}),
		field.Time("updated_at").Default(time.Now).
			SchemaType(map[string]string{dialect.MySQL: "datetime"}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("address", Address.Type),
	}
}
