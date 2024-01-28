package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func (s *Connection) Connect() error {

	sqlDb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%v)/%s",
		s.Username, s.Password, s.Host, s.Port, s.Database))
	if err != nil {
		log.Error().Msgf("Failed to connect to SQL Database: %s", err.Error())
		return err
	}

	// Create a Bun db on top of it.
	s.Client = bun.NewDB(sqlDb, mysqldialect.New())

	// Print all queries to stdout.
	s.Client.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(s.Debug)))

	return nil
}
