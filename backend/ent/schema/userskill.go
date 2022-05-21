package schema

import (
	"entgo.io/ent"
)

// UserSkill holds the schema definition for the UserSkill entity.
type UserSkill struct {
	ent.Schema
}

// Fields of the UserSkill.
func (UserSkill) Fields() []ent.Field {
	return nil
}

// Edges of the UserSkill.
func (UserSkill) Edges() []ent.Edge {
	return nil
}
