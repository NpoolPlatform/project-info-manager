package rule

import (
	"context"

	"entgo.io/ent/entql"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/privacy"
)

func FilterTimeRule() privacy.QueryMutationRule {
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		f.Where(entql.FieldEQ("delete_at", 0))
		return privacy.Skip
	})
}