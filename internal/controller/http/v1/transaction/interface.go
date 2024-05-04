package transaction

import (
	"context"
	"github.com/tests/internal/pkg"
	"github.com/tests/internal/repository/postgres/transaction"
)

type Transaction interface {
	AdminGetList(ctx context.Context, filter transaction.Filter) ([]transaction.AdminGetListResponse, int, *pkg.Error)
	AdminGetById(ctx context.Context, id string) (transaction.AdminGetDetail, *pkg.Error)
	AdminCreate(ctx context.Context, request transaction.AdminCreateRequest) (transaction.AdminCreateResponse, *pkg.Error)
	AdminUpdate(ctx context.Context, request transaction.AdminUpdateRequest) *pkg.Error
	AdminDelete(ctx context.Context, id, username string) *pkg.Error
}
