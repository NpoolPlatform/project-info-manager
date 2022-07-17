// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coinproductinfo"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CoinProductInfoUpdate is the builder for updating CoinProductInfo entities.
type CoinProductInfoUpdate struct {
	config
	hooks    []Hook
	mutation *CoinProductInfoMutation
}

// Where appends a list predicates to the CoinProductInfoUpdate builder.
func (cpiu *CoinProductInfoUpdate) Where(ps ...predicate.CoinProductInfo) *CoinProductInfoUpdate {
	cpiu.mutation.Where(ps...)
	return cpiu
}

// SetCreateAt sets the "create_at" field.
func (cpiu *CoinProductInfoUpdate) SetCreateAt(u uint32) *CoinProductInfoUpdate {
	cpiu.mutation.ResetCreateAt()
	cpiu.mutation.SetCreateAt(u)
	return cpiu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cpiu *CoinProductInfoUpdate) SetNillableCreateAt(u *uint32) *CoinProductInfoUpdate {
	if u != nil {
		cpiu.SetCreateAt(*u)
	}
	return cpiu
}

// AddCreateAt adds u to the "create_at" field.
func (cpiu *CoinProductInfoUpdate) AddCreateAt(u int32) *CoinProductInfoUpdate {
	cpiu.mutation.AddCreateAt(u)
	return cpiu
}

// SetUpdateAt sets the "update_at" field.
func (cpiu *CoinProductInfoUpdate) SetUpdateAt(u uint32) *CoinProductInfoUpdate {
	cpiu.mutation.ResetUpdateAt()
	cpiu.mutation.SetUpdateAt(u)
	return cpiu
}

// AddUpdateAt adds u to the "update_at" field.
func (cpiu *CoinProductInfoUpdate) AddUpdateAt(u int32) *CoinProductInfoUpdate {
	cpiu.mutation.AddUpdateAt(u)
	return cpiu
}

// SetDeleteAt sets the "delete_at" field.
func (cpiu *CoinProductInfoUpdate) SetDeleteAt(u uint32) *CoinProductInfoUpdate {
	cpiu.mutation.ResetDeleteAt()
	cpiu.mutation.SetDeleteAt(u)
	return cpiu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cpiu *CoinProductInfoUpdate) SetNillableDeleteAt(u *uint32) *CoinProductInfoUpdate {
	if u != nil {
		cpiu.SetDeleteAt(*u)
	}
	return cpiu
}

// AddDeleteAt adds u to the "delete_at" field.
func (cpiu *CoinProductInfoUpdate) AddDeleteAt(u int32) *CoinProductInfoUpdate {
	cpiu.mutation.AddDeleteAt(u)
	return cpiu
}

// SetAppID sets the "app_id" field.
func (cpiu *CoinProductInfoUpdate) SetAppID(u uuid.UUID) *CoinProductInfoUpdate {
	cpiu.mutation.SetAppID(u)
	return cpiu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cpiu *CoinProductInfoUpdate) SetCoinTypeID(u uuid.UUID) *CoinProductInfoUpdate {
	cpiu.mutation.SetCoinTypeID(u)
	return cpiu
}

// SetProductPage sets the "product_page" field.
func (cpiu *CoinProductInfoUpdate) SetProductPage(s string) *CoinProductInfoUpdate {
	cpiu.mutation.SetProductPage(s)
	return cpiu
}

// Mutation returns the CoinProductInfoMutation object of the builder.
func (cpiu *CoinProductInfoUpdate) Mutation() *CoinProductInfoMutation {
	return cpiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cpiu *CoinProductInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cpiu.defaults(); err != nil {
		return 0, err
	}
	if len(cpiu.hooks) == 0 {
		affected, err = cpiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinProductInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cpiu.mutation = mutation
			affected, err = cpiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cpiu.hooks) - 1; i >= 0; i-- {
			if cpiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cpiu *CoinProductInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := cpiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cpiu *CoinProductInfoUpdate) Exec(ctx context.Context) error {
	_, err := cpiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpiu *CoinProductInfoUpdate) ExecX(ctx context.Context) {
	if err := cpiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpiu *CoinProductInfoUpdate) defaults() error {
	if _, ok := cpiu.mutation.UpdateAt(); !ok {
		if coinproductinfo.UpdateDefaultUpdateAt == nil {
			return fmt.Errorf("ent: uninitialized coinproductinfo.UpdateDefaultUpdateAt (forgotten import ent/runtime?)")
		}
		v := coinproductinfo.UpdateDefaultUpdateAt()
		cpiu.mutation.SetUpdateAt(v)
	}
	return nil
}

func (cpiu *CoinProductInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinproductinfo.Table,
			Columns: coinproductinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinproductinfo.FieldID,
			},
		},
	}
	if ps := cpiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpiu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinproductinfo.FieldCreateAt,
		})
	}
	if value, ok := cpiu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinproductinfo.FieldCreateAt,
		})
	}
	if value, ok := cpiu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinproductinfo.FieldUpdateAt,
		})
	}
	if value, ok := cpiu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinproductinfo.FieldUpdateAt,
		})
	}
	if value, ok := cpiu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinproductinfo.FieldDeleteAt,
		})
	}
	if value, ok := cpiu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinproductinfo.FieldDeleteAt,
		})
	}
	if value, ok := cpiu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinproductinfo.FieldAppID,
		})
	}
	if value, ok := cpiu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinproductinfo.FieldCoinTypeID,
		})
	}
	if value, ok := cpiu.mutation.ProductPage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinproductinfo.FieldProductPage,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cpiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinproductinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CoinProductInfoUpdateOne is the builder for updating a single CoinProductInfo entity.
type CoinProductInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CoinProductInfoMutation
}

// SetCreateAt sets the "create_at" field.
func (cpiuo *CoinProductInfoUpdateOne) SetCreateAt(u uint32) *CoinProductInfoUpdateOne {
	cpiuo.mutation.ResetCreateAt()
	cpiuo.mutation.SetCreateAt(u)
	return cpiuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cpiuo *CoinProductInfoUpdateOne) SetNillableCreateAt(u *uint32) *CoinProductInfoUpdateOne {
	if u != nil {
		cpiuo.SetCreateAt(*u)
	}
	return cpiuo
}

// AddCreateAt adds u to the "create_at" field.
func (cpiuo *CoinProductInfoUpdateOne) AddCreateAt(u int32) *CoinProductInfoUpdateOne {
	cpiuo.mutation.AddCreateAt(u)
	return cpiuo
}

// SetUpdateAt sets the "update_at" field.
func (cpiuo *CoinProductInfoUpdateOne) SetUpdateAt(u uint32) *CoinProductInfoUpdateOne {
	cpiuo.mutation.ResetUpdateAt()
	cpiuo.mutation.SetUpdateAt(u)
	return cpiuo
}

// AddUpdateAt adds u to the "update_at" field.
func (cpiuo *CoinProductInfoUpdateOne) AddUpdateAt(u int32) *CoinProductInfoUpdateOne {
	cpiuo.mutation.AddUpdateAt(u)
	return cpiuo
}

// SetDeleteAt sets the "delete_at" field.
func (cpiuo *CoinProductInfoUpdateOne) SetDeleteAt(u uint32) *CoinProductInfoUpdateOne {
	cpiuo.mutation.ResetDeleteAt()
	cpiuo.mutation.SetDeleteAt(u)
	return cpiuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cpiuo *CoinProductInfoUpdateOne) SetNillableDeleteAt(u *uint32) *CoinProductInfoUpdateOne {
	if u != nil {
		cpiuo.SetDeleteAt(*u)
	}
	return cpiuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (cpiuo *CoinProductInfoUpdateOne) AddDeleteAt(u int32) *CoinProductInfoUpdateOne {
	cpiuo.mutation.AddDeleteAt(u)
	return cpiuo
}

// SetAppID sets the "app_id" field.
func (cpiuo *CoinProductInfoUpdateOne) SetAppID(u uuid.UUID) *CoinProductInfoUpdateOne {
	cpiuo.mutation.SetAppID(u)
	return cpiuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cpiuo *CoinProductInfoUpdateOne) SetCoinTypeID(u uuid.UUID) *CoinProductInfoUpdateOne {
	cpiuo.mutation.SetCoinTypeID(u)
	return cpiuo
}

// SetProductPage sets the "product_page" field.
func (cpiuo *CoinProductInfoUpdateOne) SetProductPage(s string) *CoinProductInfoUpdateOne {
	cpiuo.mutation.SetProductPage(s)
	return cpiuo
}

// Mutation returns the CoinProductInfoMutation object of the builder.
func (cpiuo *CoinProductInfoUpdateOne) Mutation() *CoinProductInfoMutation {
	return cpiuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cpiuo *CoinProductInfoUpdateOne) Select(field string, fields ...string) *CoinProductInfoUpdateOne {
	cpiuo.fields = append([]string{field}, fields...)
	return cpiuo
}

// Save executes the query and returns the updated CoinProductInfo entity.
func (cpiuo *CoinProductInfoUpdateOne) Save(ctx context.Context) (*CoinProductInfo, error) {
	var (
		err  error
		node *CoinProductInfo
	)
	if err := cpiuo.defaults(); err != nil {
		return nil, err
	}
	if len(cpiuo.hooks) == 0 {
		node, err = cpiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinProductInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cpiuo.mutation = mutation
			node, err = cpiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cpiuo.hooks) - 1; i >= 0; i-- {
			if cpiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cpiuo *CoinProductInfoUpdateOne) SaveX(ctx context.Context) *CoinProductInfo {
	node, err := cpiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cpiuo *CoinProductInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := cpiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpiuo *CoinProductInfoUpdateOne) ExecX(ctx context.Context) {
	if err := cpiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpiuo *CoinProductInfoUpdateOne) defaults() error {
	if _, ok := cpiuo.mutation.UpdateAt(); !ok {
		if coinproductinfo.UpdateDefaultUpdateAt == nil {
			return fmt.Errorf("ent: uninitialized coinproductinfo.UpdateDefaultUpdateAt (forgotten import ent/runtime?)")
		}
		v := coinproductinfo.UpdateDefaultUpdateAt()
		cpiuo.mutation.SetUpdateAt(v)
	}
	return nil
}

func (cpiuo *CoinProductInfoUpdateOne) sqlSave(ctx context.Context) (_node *CoinProductInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinproductinfo.Table,
			Columns: coinproductinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coinproductinfo.FieldID,
			},
		},
	}
	id, ok := cpiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CoinProductInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cpiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinproductinfo.FieldID)
		for _, f := range fields {
			if !coinproductinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coinproductinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cpiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpiuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinproductinfo.FieldCreateAt,
		})
	}
	if value, ok := cpiuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinproductinfo.FieldCreateAt,
		})
	}
	if value, ok := cpiuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinproductinfo.FieldUpdateAt,
		})
	}
	if value, ok := cpiuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinproductinfo.FieldUpdateAt,
		})
	}
	if value, ok := cpiuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinproductinfo.FieldDeleteAt,
		})
	}
	if value, ok := cpiuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinproductinfo.FieldDeleteAt,
		})
	}
	if value, ok := cpiuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinproductinfo.FieldAppID,
		})
	}
	if value, ok := cpiuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinproductinfo.FieldCoinTypeID,
		})
	}
	if value, ok := cpiuo.mutation.ProductPage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinproductinfo.FieldProductPage,
		})
	}
	_node = &CoinProductInfo{config: cpiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cpiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinproductinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}