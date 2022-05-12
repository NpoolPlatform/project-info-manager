// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coindescription"
	"github.com/google/uuid"
)

// CoinDescriptionCreate is the builder for creating a CoinDescription entity.
type CoinDescriptionCreate struct {
	config
	mutation *CoinDescriptionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateAt sets the "create_at" field.
func (cdc *CoinDescriptionCreate) SetCreateAt(u uint32) *CoinDescriptionCreate {
	cdc.mutation.SetCreateAt(u)
	return cdc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (cdc *CoinDescriptionCreate) SetNillableCreateAt(u *uint32) *CoinDescriptionCreate {
	if u != nil {
		cdc.SetCreateAt(*u)
	}
	return cdc
}

// SetUpdateAt sets the "update_at" field.
func (cdc *CoinDescriptionCreate) SetUpdateAt(u uint32) *CoinDescriptionCreate {
	cdc.mutation.SetUpdateAt(u)
	return cdc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (cdc *CoinDescriptionCreate) SetNillableUpdateAt(u *uint32) *CoinDescriptionCreate {
	if u != nil {
		cdc.SetUpdateAt(*u)
	}
	return cdc
}

// SetDeleteAt sets the "delete_at" field.
func (cdc *CoinDescriptionCreate) SetDeleteAt(u uint32) *CoinDescriptionCreate {
	cdc.mutation.SetDeleteAt(u)
	return cdc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cdc *CoinDescriptionCreate) SetNillableDeleteAt(u *uint32) *CoinDescriptionCreate {
	if u != nil {
		cdc.SetDeleteAt(*u)
	}
	return cdc
}

// SetAppID sets the "app_id" field.
func (cdc *CoinDescriptionCreate) SetAppID(u uuid.UUID) *CoinDescriptionCreate {
	cdc.mutation.SetAppID(u)
	return cdc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cdc *CoinDescriptionCreate) SetCoinTypeID(u uuid.UUID) *CoinDescriptionCreate {
	cdc.mutation.SetCoinTypeID(u)
	return cdc
}

// SetTitle sets the "title" field.
func (cdc *CoinDescriptionCreate) SetTitle(s string) *CoinDescriptionCreate {
	cdc.mutation.SetTitle(s)
	return cdc
}

// SetMessage sets the "message" field.
func (cdc *CoinDescriptionCreate) SetMessage(s string) *CoinDescriptionCreate {
	cdc.mutation.SetMessage(s)
	return cdc
}

// SetUsedFor sets the "used_for" field.
func (cdc *CoinDescriptionCreate) SetUsedFor(s string) *CoinDescriptionCreate {
	cdc.mutation.SetUsedFor(s)
	return cdc
}

// SetID sets the "id" field.
func (cdc *CoinDescriptionCreate) SetID(u uuid.UUID) *CoinDescriptionCreate {
	cdc.mutation.SetID(u)
	return cdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cdc *CoinDescriptionCreate) SetNillableID(u *uuid.UUID) *CoinDescriptionCreate {
	if u != nil {
		cdc.SetID(*u)
	}
	return cdc
}

// Mutation returns the CoinDescriptionMutation object of the builder.
func (cdc *CoinDescriptionCreate) Mutation() *CoinDescriptionMutation {
	return cdc.mutation
}

// Save creates the CoinDescription in the database.
func (cdc *CoinDescriptionCreate) Save(ctx context.Context) (*CoinDescription, error) {
	var (
		err  error
		node *CoinDescription
	)
	if err := cdc.defaults(); err != nil {
		return nil, err
	}
	if len(cdc.hooks) == 0 {
		if err = cdc.check(); err != nil {
			return nil, err
		}
		node, err = cdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cdc.check(); err != nil {
				return nil, err
			}
			cdc.mutation = mutation
			if node, err = cdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cdc.hooks) - 1; i >= 0; i-- {
			if cdc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cdc *CoinDescriptionCreate) SaveX(ctx context.Context) *CoinDescription {
	v, err := cdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cdc *CoinDescriptionCreate) Exec(ctx context.Context) error {
	_, err := cdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cdc *CoinDescriptionCreate) ExecX(ctx context.Context) {
	if err := cdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cdc *CoinDescriptionCreate) defaults() error {
	if _, ok := cdc.mutation.CreateAt(); !ok {
		if coindescription.DefaultCreateAt == nil {
			return fmt.Errorf("ent: uninitialized coindescription.DefaultCreateAt (forgotten import ent/runtime?)")
		}
		v := coindescription.DefaultCreateAt()
		cdc.mutation.SetCreateAt(v)
	}
	if _, ok := cdc.mutation.UpdateAt(); !ok {
		if coindescription.DefaultUpdateAt == nil {
			return fmt.Errorf("ent: uninitialized coindescription.DefaultUpdateAt (forgotten import ent/runtime?)")
		}
		v := coindescription.DefaultUpdateAt()
		cdc.mutation.SetUpdateAt(v)
	}
	if _, ok := cdc.mutation.DeleteAt(); !ok {
		if coindescription.DefaultDeleteAt == nil {
			return fmt.Errorf("ent: uninitialized coindescription.DefaultDeleteAt (forgotten import ent/runtime?)")
		}
		v := coindescription.DefaultDeleteAt()
		cdc.mutation.SetDeleteAt(v)
	}
	if _, ok := cdc.mutation.ID(); !ok {
		if coindescription.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized coindescription.DefaultID (forgotten import ent/runtime?)")
		}
		v := coindescription.DefaultID()
		cdc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cdc *CoinDescriptionCreate) check() error {
	if _, ok := cdc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "CoinDescription.create_at"`)}
	}
	if _, ok := cdc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "CoinDescription.update_at"`)}
	}
	if _, ok := cdc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "CoinDescription.delete_at"`)}
	}
	if _, ok := cdc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "CoinDescription.app_id"`)}
	}
	if _, ok := cdc.mutation.CoinTypeID(); !ok {
		return &ValidationError{Name: "coin_type_id", err: errors.New(`ent: missing required field "CoinDescription.coin_type_id"`)}
	}
	if _, ok := cdc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "CoinDescription.title"`)}
	}
	if _, ok := cdc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "CoinDescription.message"`)}
	}
	if v, ok := cdc.mutation.Message(); ok {
		if err := coindescription.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "CoinDescription.message": %w`, err)}
		}
	}
	if _, ok := cdc.mutation.UsedFor(); !ok {
		return &ValidationError{Name: "used_for", err: errors.New(`ent: missing required field "CoinDescription.used_for"`)}
	}
	return nil
}

func (cdc *CoinDescriptionCreate) sqlSave(ctx context.Context) (*CoinDescription, error) {
	_node, _spec := cdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (cdc *CoinDescriptionCreate) createSpec() (*CoinDescription, *sqlgraph.CreateSpec) {
	var (
		_node = &CoinDescription{config: cdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: coindescription.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coindescription.FieldID,
			},
		}
	)
	_spec.OnConflict = cdc.conflict
	if id, ok := cdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cdc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := cdc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := cdc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	if value, ok := cdc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coindescription.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := cdc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coindescription.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := cdc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coindescription.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := cdc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coindescription.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := cdc.mutation.UsedFor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coindescription.FieldUsedFor,
		})
		_node.UsedFor = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CoinDescription.Create().
//		SetCreateAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CoinDescriptionUpsert) {
//			SetCreateAt(v+v).
//		}).
//		Exec(ctx)
//
func (cdc *CoinDescriptionCreate) OnConflict(opts ...sql.ConflictOption) *CoinDescriptionUpsertOne {
	cdc.conflict = opts
	return &CoinDescriptionUpsertOne{
		create: cdc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CoinDescription.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cdc *CoinDescriptionCreate) OnConflictColumns(columns ...string) *CoinDescriptionUpsertOne {
	cdc.conflict = append(cdc.conflict, sql.ConflictColumns(columns...))
	return &CoinDescriptionUpsertOne{
		create: cdc,
	}
}

type (
	// CoinDescriptionUpsertOne is the builder for "upsert"-ing
	//  one CoinDescription node.
	CoinDescriptionUpsertOne struct {
		create *CoinDescriptionCreate
	}

	// CoinDescriptionUpsert is the "OnConflict" setter.
	CoinDescriptionUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateAt sets the "create_at" field.
func (u *CoinDescriptionUpsert) SetCreateAt(v uint32) *CoinDescriptionUpsert {
	u.Set(coindescription.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CoinDescriptionUpsert) UpdateCreateAt() *CoinDescriptionUpsert {
	u.SetExcluded(coindescription.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *CoinDescriptionUpsert) AddCreateAt(v uint32) *CoinDescriptionUpsert {
	u.Add(coindescription.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *CoinDescriptionUpsert) SetUpdateAt(v uint32) *CoinDescriptionUpsert {
	u.Set(coindescription.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CoinDescriptionUpsert) UpdateUpdateAt() *CoinDescriptionUpsert {
	u.SetExcluded(coindescription.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *CoinDescriptionUpsert) AddUpdateAt(v uint32) *CoinDescriptionUpsert {
	u.Add(coindescription.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *CoinDescriptionUpsert) SetDeleteAt(v uint32) *CoinDescriptionUpsert {
	u.Set(coindescription.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CoinDescriptionUpsert) UpdateDeleteAt() *CoinDescriptionUpsert {
	u.SetExcluded(coindescription.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *CoinDescriptionUpsert) AddDeleteAt(v uint32) *CoinDescriptionUpsert {
	u.Add(coindescription.FieldDeleteAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *CoinDescriptionUpsert) SetAppID(v uuid.UUID) *CoinDescriptionUpsert {
	u.Set(coindescription.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CoinDescriptionUpsert) UpdateAppID() *CoinDescriptionUpsert {
	u.SetExcluded(coindescription.FieldAppID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CoinDescriptionUpsert) SetCoinTypeID(v uuid.UUID) *CoinDescriptionUpsert {
	u.Set(coindescription.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CoinDescriptionUpsert) UpdateCoinTypeID() *CoinDescriptionUpsert {
	u.SetExcluded(coindescription.FieldCoinTypeID)
	return u
}

// SetTitle sets the "title" field.
func (u *CoinDescriptionUpsert) SetTitle(v string) *CoinDescriptionUpsert {
	u.Set(coindescription.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *CoinDescriptionUpsert) UpdateTitle() *CoinDescriptionUpsert {
	u.SetExcluded(coindescription.FieldTitle)
	return u
}

// SetMessage sets the "message" field.
func (u *CoinDescriptionUpsert) SetMessage(v string) *CoinDescriptionUpsert {
	u.Set(coindescription.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *CoinDescriptionUpsert) UpdateMessage() *CoinDescriptionUpsert {
	u.SetExcluded(coindescription.FieldMessage)
	return u
}

// SetUsedFor sets the "used_for" field.
func (u *CoinDescriptionUpsert) SetUsedFor(v string) *CoinDescriptionUpsert {
	u.Set(coindescription.FieldUsedFor, v)
	return u
}

// UpdateUsedFor sets the "used_for" field to the value that was provided on create.
func (u *CoinDescriptionUpsert) UpdateUsedFor() *CoinDescriptionUpsert {
	u.SetExcluded(coindescription.FieldUsedFor)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CoinDescription.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(coindescription.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CoinDescriptionUpsertOne) UpdateNewValues() *CoinDescriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(coindescription.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.CoinDescription.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CoinDescriptionUpsertOne) Ignore() *CoinDescriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CoinDescriptionUpsertOne) DoNothing() *CoinDescriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CoinDescriptionCreate.OnConflict
// documentation for more info.
func (u *CoinDescriptionUpsertOne) Update(set func(*CoinDescriptionUpsert)) *CoinDescriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CoinDescriptionUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *CoinDescriptionUpsertOne) SetCreateAt(v uint32) *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *CoinDescriptionUpsertOne) AddCreateAt(v uint32) *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CoinDescriptionUpsertOne) UpdateCreateAt() *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *CoinDescriptionUpsertOne) SetUpdateAt(v uint32) *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *CoinDescriptionUpsertOne) AddUpdateAt(v uint32) *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CoinDescriptionUpsertOne) UpdateUpdateAt() *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *CoinDescriptionUpsertOne) SetDeleteAt(v uint32) *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *CoinDescriptionUpsertOne) AddDeleteAt(v uint32) *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CoinDescriptionUpsertOne) UpdateDeleteAt() *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateDeleteAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *CoinDescriptionUpsertOne) SetAppID(v uuid.UUID) *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CoinDescriptionUpsertOne) UpdateAppID() *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateAppID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CoinDescriptionUpsertOne) SetCoinTypeID(v uuid.UUID) *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CoinDescriptionUpsertOne) UpdateCoinTypeID() *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateCoinTypeID()
	})
}

// SetTitle sets the "title" field.
func (u *CoinDescriptionUpsertOne) SetTitle(v string) *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *CoinDescriptionUpsertOne) UpdateTitle() *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateTitle()
	})
}

// SetMessage sets the "message" field.
func (u *CoinDescriptionUpsertOne) SetMessage(v string) *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *CoinDescriptionUpsertOne) UpdateMessage() *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateMessage()
	})
}

// SetUsedFor sets the "used_for" field.
func (u *CoinDescriptionUpsertOne) SetUsedFor(v string) *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetUsedFor(v)
	})
}

// UpdateUsedFor sets the "used_for" field to the value that was provided on create.
func (u *CoinDescriptionUpsertOne) UpdateUsedFor() *CoinDescriptionUpsertOne {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateUsedFor()
	})
}

// Exec executes the query.
func (u *CoinDescriptionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CoinDescriptionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CoinDescriptionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CoinDescriptionUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CoinDescriptionUpsertOne.ID is not supported by MySQL driver. Use CoinDescriptionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CoinDescriptionUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CoinDescriptionCreateBulk is the builder for creating many CoinDescription entities in bulk.
type CoinDescriptionCreateBulk struct {
	config
	builders []*CoinDescriptionCreate
	conflict []sql.ConflictOption
}

// Save creates the CoinDescription entities in the database.
func (cdcb *CoinDescriptionCreateBulk) Save(ctx context.Context) ([]*CoinDescription, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cdcb.builders))
	nodes := make([]*CoinDescription, len(cdcb.builders))
	mutators := make([]Mutator, len(cdcb.builders))
	for i := range cdcb.builders {
		func(i int, root context.Context) {
			builder := cdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CoinDescriptionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cdcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cdcb *CoinDescriptionCreateBulk) SaveX(ctx context.Context) []*CoinDescription {
	v, err := cdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cdcb *CoinDescriptionCreateBulk) Exec(ctx context.Context) error {
	_, err := cdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cdcb *CoinDescriptionCreateBulk) ExecX(ctx context.Context) {
	if err := cdcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CoinDescription.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CoinDescriptionUpsert) {
//			SetCreateAt(v+v).
//		}).
//		Exec(ctx)
//
func (cdcb *CoinDescriptionCreateBulk) OnConflict(opts ...sql.ConflictOption) *CoinDescriptionUpsertBulk {
	cdcb.conflict = opts
	return &CoinDescriptionUpsertBulk{
		create: cdcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CoinDescription.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cdcb *CoinDescriptionCreateBulk) OnConflictColumns(columns ...string) *CoinDescriptionUpsertBulk {
	cdcb.conflict = append(cdcb.conflict, sql.ConflictColumns(columns...))
	return &CoinDescriptionUpsertBulk{
		create: cdcb,
	}
}

// CoinDescriptionUpsertBulk is the builder for "upsert"-ing
// a bulk of CoinDescription nodes.
type CoinDescriptionUpsertBulk struct {
	create *CoinDescriptionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CoinDescription.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(coindescription.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CoinDescriptionUpsertBulk) UpdateNewValues() *CoinDescriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(coindescription.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CoinDescription.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CoinDescriptionUpsertBulk) Ignore() *CoinDescriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CoinDescriptionUpsertBulk) DoNothing() *CoinDescriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CoinDescriptionCreateBulk.OnConflict
// documentation for more info.
func (u *CoinDescriptionUpsertBulk) Update(set func(*CoinDescriptionUpsert)) *CoinDescriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CoinDescriptionUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *CoinDescriptionUpsertBulk) SetCreateAt(v uint32) *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *CoinDescriptionUpsertBulk) AddCreateAt(v uint32) *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *CoinDescriptionUpsertBulk) UpdateCreateAt() *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *CoinDescriptionUpsertBulk) SetUpdateAt(v uint32) *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *CoinDescriptionUpsertBulk) AddUpdateAt(v uint32) *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *CoinDescriptionUpsertBulk) UpdateUpdateAt() *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *CoinDescriptionUpsertBulk) SetDeleteAt(v uint32) *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *CoinDescriptionUpsertBulk) AddDeleteAt(v uint32) *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *CoinDescriptionUpsertBulk) UpdateDeleteAt() *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateDeleteAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *CoinDescriptionUpsertBulk) SetAppID(v uuid.UUID) *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CoinDescriptionUpsertBulk) UpdateAppID() *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateAppID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CoinDescriptionUpsertBulk) SetCoinTypeID(v uuid.UUID) *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CoinDescriptionUpsertBulk) UpdateCoinTypeID() *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateCoinTypeID()
	})
}

// SetTitle sets the "title" field.
func (u *CoinDescriptionUpsertBulk) SetTitle(v string) *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *CoinDescriptionUpsertBulk) UpdateTitle() *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateTitle()
	})
}

// SetMessage sets the "message" field.
func (u *CoinDescriptionUpsertBulk) SetMessage(v string) *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *CoinDescriptionUpsertBulk) UpdateMessage() *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateMessage()
	})
}

// SetUsedFor sets the "used_for" field.
func (u *CoinDescriptionUpsertBulk) SetUsedFor(v string) *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.SetUsedFor(v)
	})
}

// UpdateUsedFor sets the "used_for" field to the value that was provided on create.
func (u *CoinDescriptionUpsertBulk) UpdateUsedFor() *CoinDescriptionUpsertBulk {
	return u.Update(func(s *CoinDescriptionUpsert) {
		s.UpdateUsedFor()
	})
}

// Exec executes the query.
func (u *CoinDescriptionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CoinDescriptionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CoinDescriptionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CoinDescriptionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}