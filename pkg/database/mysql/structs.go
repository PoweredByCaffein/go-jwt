package mysql

import "github.com/uptrace/bun"

type Connection struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Debug    bool

	Client *bun.DB
}
