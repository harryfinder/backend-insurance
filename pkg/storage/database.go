package storage

import "context"

type QueryResult interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

type CommandTag interface {
	String() string
}

type Transaction interface {
	QueryResult
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	QueryRow(ctx context.Context, sql string, args ...interface{}) QueryResult
	Query(ctx context.Context, sql string, args ...interface{}) (QueryResult, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (CommandTag, error)
}

type Database interface {
	Query(ctx context.Context, sql string, args ...interface{}) (QueryResult, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) QueryResult
	Exec(ctx context.Context, sql string, arguments ...interface{}) (CommandTag, error)
	Begin(ctx context.Context) (Transaction, error)
	Close()
}
