package transactionview

import (
	"time"

	"github.com/uptrace/bun"
)

type Filter struct {
	Limit        *int
	Offset       *int
	CustomerName *string
	ItemName     *string
	Id           *string
}

type AdminGetListResponse struct {
	Id           string   `json:"id"`
	CustomerName *string  `json:"customer_name"`
	ItemName     *string  `json:"item_name"`
	QTY          *int     `json:"qty"`
	Price        *float64 `json:"price"`
	Amount       *float64 `json:"amount"`
	CustomerID   *string  `json:"customer_id"`
	ItemID       *string  `json:"item_id"`
}

type AdminGetDetail struct {
	bun.BaseModel `bun:"table:transactionviews"`

	Id           string   `json:"id" bun:"id"`
	CustomerName *string  `json:"customer_name" bun:"customer_name"`
	ItemName     *string  `json:"item_name" bun:"item_name"`
	QTY          *int     `json:"qty" bun:"qty"`
	Price        *float64 `json:"price" bun:"price"`
	Amount       *float64 `json:"amount" bun:"amount"`
	CustomerID   *string  `json:"customer_id" bun:"customer_id"`
	ItemID       *string  `json:"item_id" bun:"item_id"`
}

type AdminCreateRequest struct {
	CustomerName *string  `json:"customer_name" form:"customer_name"`
	ItemName     *string  `json:"item_name" form:"item_name"`
	QTY          *int     `json:"qty" form:"qty"`
	Price        *float64 `json:"price" form:"price"`
	Amount       *float64 `json:"amount" form:"amount"`
	CustomerID   *string  `json:"customer_id" form:"customer_id"`
	ItemID       *string  `json:"item_id" form:"item_id"`
}

type AdminCreateResponse struct {
	bun.BaseModel `bun:"table:transactionviews"`
	Id            string `json:"id" bun:"id"`

	CustomerName *string   `json:"customer_name" bun:"customer_name"`
	ItemName     *string   `json:"item_name" bun:"item_name"`
	QTY          *int      `json:"qty" bun:"qty"`
	Price        *float64  `json:"price" bun:"price"`
	Amount       *float64  `json:"amount" bun:"amount"`
	CustomerID   *string   `json:"customer_id" bun:"customer_id"`
	ItemID       *string   `json:"item_id" bun:"item_id"`
	CreatedAt    time.Time `json:"-" bun:"created_at"`
	CreatedBy    *string   `json:"-" bun:"created_by"`
}

type AdminUpdateRequest struct {
	Id           string   `json:"id" form:"id"`
	CustomerName *string  `json:"customer_name" form:"customer_name"`
	ItemName     *string  `json:"item_name" form:"item_name"`
	QTY          *int     `json:"qty" form:"qty"`
	Price        *float64 `json:"price" form:"price"`
	Amount       *float64 `json:"amount" form:"amount"`
	CustomerID   *string  `json:"customer_id" form:"customer_id"`
	ItemID       *string  `json:"item_id" form:"item_id"`
}
