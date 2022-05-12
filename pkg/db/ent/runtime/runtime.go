// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coindescription"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/schema"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	coindescriptionMixin := schema.CoinDescription{}.Mixin()
	coindescription.Policy = privacy.NewPolicies(coindescriptionMixin[0], schema.CoinDescription{})
	coindescription.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := coindescription.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	coindescriptionMixinFields0 := coindescriptionMixin[0].Fields()
	_ = coindescriptionMixinFields0
	coindescriptionFields := schema.CoinDescription{}.Fields()
	_ = coindescriptionFields
	// coindescriptionDescCreateAt is the schema descriptor for create_at field.
	coindescriptionDescCreateAt := coindescriptionMixinFields0[0].Descriptor()
	// coindescription.DefaultCreateAt holds the default value on creation for the create_at field.
	coindescription.DefaultCreateAt = coindescriptionDescCreateAt.Default.(func() uint32)
	// coindescriptionDescUpdateAt is the schema descriptor for update_at field.
	coindescriptionDescUpdateAt := coindescriptionMixinFields0[1].Descriptor()
	// coindescription.DefaultUpdateAt holds the default value on creation for the update_at field.
	coindescription.DefaultUpdateAt = coindescriptionDescUpdateAt.Default.(func() uint32)
	// coindescription.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	coindescription.UpdateDefaultUpdateAt = coindescriptionDescUpdateAt.UpdateDefault.(func() uint32)
	// coindescriptionDescDeleteAt is the schema descriptor for delete_at field.
	coindescriptionDescDeleteAt := coindescriptionMixinFields0[2].Descriptor()
	// coindescription.DefaultDeleteAt holds the default value on creation for the delete_at field.
	coindescription.DefaultDeleteAt = coindescriptionDescDeleteAt.Default.(func() uint32)
	// coindescriptionDescMessage is the schema descriptor for message field.
	coindescriptionDescMessage := coindescriptionFields[4].Descriptor()
	// coindescription.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	coindescription.MessageValidator = coindescriptionDescMessage.Validators[0].(func(string) error)
	// coindescriptionDescID is the schema descriptor for id field.
	coindescriptionDescID := coindescriptionFields[0].Descriptor()
	// coindescription.DefaultID holds the default value on creation for the id field.
	coindescription.DefaultID = coindescriptionDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.10.1"                                         // Version of ent codegen.
	Sum     = "h1:dM5h4Zk6yHGIgw4dCqVzGw3nWgpGYJiV4/kyHEF6PFo=" // Sum of ent codegen.
)