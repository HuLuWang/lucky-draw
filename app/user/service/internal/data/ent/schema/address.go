package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id"),
		field.String("username"),
		field.String("mobile"),
		field.String("address"),
		field.Time("created_at").Default(time.Now).
			SchemaType(map[string]string{dialect.MySQL: "datetime"}),
		field.Time("updated_at").Default(time.Now).
			SchemaType(map[string]string{dialect.MySQL: "datetime"}),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("address").Unique(),
	}
}
