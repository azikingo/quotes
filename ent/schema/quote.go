package schema

import "C"
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type Quote struct {
	ent.Schema
}

func (Quote) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("quote").Optional(),
		field.Float32("author").Nillable().Optional(),
		field.Time("created_at").Default(time.Now()),
	}
}

func (Quote) Edges() []ent.Edge {
	return nil
}
