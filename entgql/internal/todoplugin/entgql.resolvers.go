// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package todoplugin

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"entgo.io/contrib/entgql/internal/todoplugin/ent"
	"entgo.io/contrib/entgql/internal/todoplugin/ent/todo"
)

func (r *masterUserResolver) Age(ctx context.Context, obj *ent.User) (float64, error) {
	return float64(obj.Age), nil
}

func (r *masterUserResolver) Amount(ctx context.Context, obj *ent.User) (float64, error) {
	return float64(obj.Amount), nil
}

func (r *todoResolver) Children(ctx context.Context, obj *ent.Todo, orderBy *ent.TodoOrder, after *ent.Cursor, before *ent.Cursor, first *int, last *int) (*ent.TodoConnection, error) {
	// TODO(giautm): auto-generate this method.
	return r.client.Todo.Query().
		Where(todo.HasParentWith(todo.ID(obj.ID))).
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
		)
}

// MasterUser returns MasterUserResolver implementation.
func (r *Resolver) MasterUser() MasterUserResolver { return &masterUserResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type masterUserResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
