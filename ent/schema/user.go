package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Positive().
			Immutable(),
		field.String("name").
			MaxLen(100).
			NotEmpty(),
		field.String("email").
			MaxLen(100).
			NotEmpty(),
		field.Int("age").
			Positive(),
	}
}
