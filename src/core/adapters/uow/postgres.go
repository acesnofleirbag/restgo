package uow

import (
	"api/src/core/guard"
	"context"
	"database/sql"
)

type Postgres struct {
    conn *sql.DB
    ctx context.Context
}

func NewPostgres(uri string) Postgres {
    conn := Connect(uri)

    return Postgres {
        conn: conn,
    }
}

func Connect(uri string) *sql.DB {
    conn, err := sql.Open("postgres", uri)
    guard.Err(err)

    return conn
}

func (self *Postgres) Run(query string) *sql.Rows {
    trx, err := self.conn.BeginTx(self.ctx, &sql.TxOptions{ Isolation: sql.LevelReadCommitted })
    guard.Err(err)

    data, err := trx.Query(query)

    if err != nil {
        trx.Rollback()
    } else {
        trx.Commit()
    }

    return data
}
