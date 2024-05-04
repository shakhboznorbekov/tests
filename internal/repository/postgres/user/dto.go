package user

import (
	"github.com/uptrace/bun"
	"time"
)

type Filter struct {
	Limit     *int
	Offset    *int
	FirstName *string
}

type AdminGetListResponse struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    string `json:"status"`
}

type AdminGetDetail struct {
	bun.BaseModel `bun:"table:users"`

	Id        string  `json:"id" bun:"id"`
	FirstName string  `json:"first_name" bun:"first_name"`
	LastName  string  `json:"last_name" bun:"last_name"`
	Username  string  `json:"username" bun:"username"`
	Password  *string `json:"-" bun:"password"`
	Status    string  `json:"status" bun:"status"`
	Gmail     string  `json:"gmail" bun:"gmail"`
}

type AdminCreateRequest struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	Status    string `json:"status" form:"status"`
	Gmail     string `json:"gmail" form:"gmail"`
}

type AdminCreateResponse struct {
	bun.BaseModel `bun:"table:users"`
	Id            string    `json:"id" bun:"id"`
	FirstName     string    `json:"first_name" bun:"first_name"`
	LastName      string    `json:"last_name" bun:"last_name"`
	Username      string    `json:"username" bun:"username"`
	Password      string    `json:"password" bun:"password"`
	Status        string    `json:"status" bun:"status"`
	Gmail         string    `json:"gmail" bun:"gmail"`
	CreatedAt     time.Time `json:"-" bun:"created_at"`
	CreatedBy     *string   `json:"-" bun:"created_by"`
}

type AdminUpdateRequest struct {
	Id        string  `json:"id" form:"id"`
	FirstName *string `json:"first_name" form:"first_name"`
	LastName  *string `json:"last_name" form:"last_name"`
	Username  *string `json:"username" form:"username"`
	Password  *string `json:"password" form:"password"`
	Status    *string `json:"status" form:"status"`
	Gmail     *string `json:"gmail" form:"gmail"`
}
