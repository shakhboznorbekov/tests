package customer

import (
	"time"

	"github.com/uptrace/bun"
)

type Filter struct {
	Limit  *int
	Offset *int
}

type AdminGetListResponse struct {
	Id           string   `json:"id"`
	CustomerName *string  `json:"customer_name"`
	Balance      *float64 `json:"balanse"`
}

type AdminGetDetail struct {
	bun.BaseModel `bun:"table:customers"`

	Id           string   `json:"id" bun:"id"`
	CustomerName *string  `json:"customer_name" bun:"customer_name"`
	Balance      *float64 `json:"balance" bun:"balance"`
}

type AdminCreateRequest struct {
	CustomerName *string  `json:"customer_name" form:"customer_name"`
	Balance      *float64 `json:"balance" form:"balance"`
}

type AdminCreateResponse struct {
	bun.BaseModel `bun:"table:customers"`
	Id            string    `json:"id" bun:"id"`
	CustomerName  *string   `json:"customer_name" bun:"customer_name"`
	Balance       *float64  `json:"balance" bun:"balance"`
	CreatedAt     time.Time `json:"-" bun:"created_at"`
	CreatedBy     *string   `json:"-" bun:"created_by"`
}

type AdminUpdateRequest struct {
	Id           string   `json:"id" form:"id"`
	CustomerName *string  `json:"customer_name" form:"customer_name"`
	Balance      *float64 `json:"balance" form:"balance"`
}
