package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("title"),
		field.String("description"),
		field.Time("start_at"),
		field.Time("end_at"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("posts").
			Unique().
			Field("user_id"),
	}
}
