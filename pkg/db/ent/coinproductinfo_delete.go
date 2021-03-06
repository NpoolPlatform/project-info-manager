// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coinproductinfo"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/predicate"
)

// CoinProductInfoDelete is the builder for deleting a CoinProductInfo entity.
type CoinProductInfoDelete struct {
	config
	hooks    []Hook
	mutation *CoinProductInfoMutation
}

// Where appends a list predicates to the CoinProductInfoDelete builder.
func (cpid *CoinProductInfoDelete) Where(ps ...predicate.CoinProductInfo) *CoinProductInfoDelete {
	cpid.mutation.Where(ps...)
	return cpid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cpid *CoinProductInfoDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cpid.hooks) == 0 {
		affected, err = cpid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinProductInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cpid.mutation = mutation
			affected, err = cpid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cpid.hooks) - 1; i >= 0; i-- {
			if cpid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpid *CoinProductInfoDelete) ExecX(ctx context.Context) int {
	n, err := cpid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cpid *CoinProductInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: coinproductinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinproductinfo.FieldID,
			},
		},
	}
	if ps := cpid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cpid.driver, _spec)
}

// CoinProductInfoDeleteOne is the builder for deleting a single CoinProductInfo entity.
type CoinProductInfoDeleteOne struct {
	cpid *CoinProductInfoDelete
}

// Exec executes the deletion query.
func (cpido *CoinProductInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := cpido.cpid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{coinproductinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cpido *CoinProductInfoDeleteOne) ExecX(ctx context.Context) {
	cpido.cpid.ExecX(ctx)
}
