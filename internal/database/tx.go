package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PgTxMock struct {
}

func (t *PgTxMock) Commit(ctx context.Context) error {
	return nil
}

func (t *PgTxMock) Begin(ctx context.Context) (pgx.Tx, error) {
	return nil, nil
}

func (t *PgTxMock) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return nil
}

func (t *PgTxMock) LargeObjects() pgx.LargeObjects {
	var a pgx.LargeObjects
	return a
}

func (t *PgTxMock) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return nil, nil
}

func (t *PgTxMock) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return nil
}

func (t *PgTxMock) Conn() *pgx.Conn {
	return nil
}

func (t *PgTxMock) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	return nil, nil
}

func (t *PgTxMock) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	var a pgconn.CommandTag
	return a, nil
}

func (t *PgTxMock) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return 0, nil
}

func (t *PgTxMock) Rollback(ctx context.Context) error {
	return nil
}
