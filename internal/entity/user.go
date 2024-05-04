package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	Id        string     `json:"id" bun:"id"`
	FirstName string     `json:"first_name" bun:"first_name"`
	LastName  string     `json:"last_name" bun:"last_name"`
	Username  string     `json:"username" bun:"username"`
	Password  *string    `json:"password" bun:"password"`
	Status    bool       `json:"status" bun:"status"`
	Gmail     string     `json:"gmail" bun:"gmail"`
	CreatedAt time.Time  `json:"-" bun:"created_at"`
	CreatedBy *string    `json:"-" bun:"created_by"`
	UpdatedAt *time.Time `json:"-" bun:"updated_at"`
	UpdatedBy *string    `json:"-" bun:"updated_by"`
	DeletedAt *time.Time `json:"-" bun:"deleted_at"`
	DeletedBy *string    `json:"-" bun:"deleted_by"`
}

// Role
//1. Admin
//2. User
