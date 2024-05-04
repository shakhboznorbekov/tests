package transactionview

import (
	"context"
	"github.com/tests/internal/pkg"
	"github.com/tests/internal/repository/postgres/transactionview"
)

type TransactionView interface {
	AdminGetList(ctx context.Context, filter transactionview.Filter) ([]transactionview.AdminGetListResponse, int, *pkg.Error)
	AdminGetById(ctx context.Context, id string) (transactionview.AdminGetDetail, *pkg.Error)
	AdminCreate(ctx context.Context, request transactionview.AdminCreateRequest) (transactionview.AdminCreateResponse, *pkg.Error)
	AdminUpdate(ctx context.Context, request transactionview.AdminUpdateRequest) *pkg.Error
	AdminDelete(ctx context.Context, id, username string) *pkg.Error
}
