package item

import (
	"time"

	"github.com/uptrace/bun"
)

type Filter struct {
	Limit  *int
	Offset *int
}

type AdminGetListResponse struct {
	Id       string   `json:"id"`
	ItemName *string  `json:"item_name"`
	Cost     *float64 `json:"cost"`
	Price    *float64 `json:"price"`
	Sort     *int     `json:"sort"`
}

type AdminGetDetail struct {
	bun.BaseModel `bun:"table:items"`

	Id       string   `json:"id" bun:"id"`
	ItemName *string  `json:"item_name" bun:"item_name"`
	Cost     *float64 `json:"cost" bun:"cost"`
	Price    *float64 `json:"price" bun:"price"`
	Sort     *int     `json:"sort" bun:"sort"`
}

type AdminCreateRequest struct {
	ItemName *string  `json:"item_name" form:"item_name"`
	Cost     *float64 `json:"cost" form:"cost"`
	Price    *float64 `json:"price" form:"price"`
	Sort     *int     `json:"sort" form:"sort"`
}

type AdminCreateResponse struct {
	bun.BaseModel `bun:"table:items"`
	Id            string     `json:"id" bun:"id"`
	ItemName      *string    `json:"item_name" bun:"item_name"`
	Cost          *float64   `json:"cost" bun:"cost"`
	Price         *float64   `json:"price" bun:"price"`
	Sort          *int       `json:"sort" bun:"sort"`
	CreatedAt     time.Time `json:"-" bun:"created_at"`
	CreatedBy     *string    `json:"-" bun:"created_by"`
}

type AdminUpdateRequest struct {
	Id       string   `json:"id" form:"id"`
	ItemName *string  `json:"item_name" form:"item_name"`
	Cost     *float64 `json:"cost" form:"cost"`
	Price    *float64 `json:"price" form:"price"`
	Sort     *float64 `json:"sort" form:"sort"`
}
