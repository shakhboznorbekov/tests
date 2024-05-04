package item

import (
	"context"
	"github.com/tests/internal/pkg"
	"github.com/tests/internal/repository/postgres/item"
)

type Item interface {
	AdminGetList(ctx context.Context, filter item.Filter) ([]item.AdminGetListResponse, int, *pkg.Error)
	AdminGetById(ctx context.Context, id string) (item.AdminGetDetail, *pkg.Error)
	AdminCreate(ctx context.Context, request item.AdminCreateRequest) (item.AdminCreateResponse, *pkg.Error)
	AdminUpdate(ctx context.Context, request item.AdminUpdateRequest) *pkg.Error
	AdminDelete(ctx context.Context, id, username string) *pkg.Error
}
