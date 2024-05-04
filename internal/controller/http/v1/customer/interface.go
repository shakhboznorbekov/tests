package customer

import (
	"context"
	"github.com/tests/internal/pkg"
	"github.com/tests/internal/repository/postgres/customer"
)

type Customer interface {
	AdminGetList(ctx context.Context, filter customer.Filter) ([]customer.AdminGetListResponse, int, *pkg.Error)
	AdminGetById(ctx context.Context, id string) (customer.AdminGetDetail, *pkg.Error)
	AdminCreate(ctx context.Context, request customer.AdminCreateRequest) (customer.AdminCreateResponse, *pkg.Error)
	AdminUpdate(ctx context.Context, request customer.AdminUpdateRequest) *pkg.Error
	AdminDelete(ctx context.Context, id, username string) *pkg.Error
}
