package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
        field.String("name"),
        field.String("email").
						Unique(),
				field.String("password_hash"),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("user_skills", UserSkill.Type),
    }
}
