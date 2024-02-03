package models

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"

	"github.com/uptrace/bun"
)

type UserModel struct {
	bun.BaseModel `bun:"table:users,alias:users"`

	ID        int64  `bun:"id,pk,autoincrement" json:"id"`
	FirstName string `bun:"first_name,notnull" json:"first_name"`
	LastName  string `bun:"last_name,notnull" json:"last_name"`
	Email     string `bun:"email,notnull,unique" json:"email"`
	Password  string `bun:"password,notnull" json:"password"`

	CreatedAt time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,default:current_timestamp"`
	DeletedAt time.Time `bun:"deleted_at,soft_delete,nullzero,default:null"`
}

func CreateUsersTable(client *bun.DB) error {
	_, err := client.NewCreateTable().
		Model((*UserModel)(nil)).Exec(context.TODO())

	log.Error().Msgf("%s", err.Error())
	return err
}
