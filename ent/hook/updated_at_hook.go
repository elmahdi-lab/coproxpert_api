package hook

import (
	"context"
	"entgo.io/ent"
	"time"
)

func UpdateTimestamp(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		if !m.Op().Is(ent.OpUpdateOne) {
			return next.Mutate(ctx, m)
		}

		if _, exists := m.Field("updated_at"); exists {
			err := m.SetField("updated_at", time.Now())
			if err != nil {
				return nil, err
			}
		}

		return next.Mutate(ctx, m)
	})
}
