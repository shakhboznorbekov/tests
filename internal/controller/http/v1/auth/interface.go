package auth

import (
	"context"
	"github.com/tests/internal/auth"
	"github.com/tests/internal/entity"
	"github.com/tests/internal/pkg"
	"github.com/tests/internal/repository/postgres/user"
)

type Auth interface {
	GenerateToken(ctx context.Context, data auth.GenerateToken) (string, error)
	IsValidToken(ctx context.Context, token string) (entity.User, error)
	GetTokenData(ctx context.Context, token string) (auth.TokenData, error)
}

type User interface {
	GetByFirstName(ctx context.Context, firstName string) (user.AdminGetDetail, *pkg.Error)
}
