package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/zssky/log"
)

const (
	DB_TYPE_SQLITE3 = "sqlite3"
	DB_TYPE_MYSQL   = "mysql"
	DB_TYPE_MSSQL   = "mssql"
	DB_TYPE_OCI8    = "oci8"
	DB_TYPE_DB2     = "db2-cli"
	DB_TYPE_ADO     = "adodb"
	DB_TYPE_ODBC    = "odbc"
)

var (
	// ErrNoDB is returned by functions to specity the requested,
	// database implation not exist.
	ErrNoDB = errors.New("db implation doesn't exist")
)

// Factory is an factory method to create CobraDB object.
type Factory func(host string, port int, user string, password string, dbname string) (Impl, error)

var (
	//factories has the factories for the CobraDB objects.
	factories = make(map[string]Factory)
)

// RegisterFactory register a factory for an implation for a Server.
// If an implation with the name already exists, it log.Fatals out.
// call this in the 'init' function in you db implation module.
func RegisterFactory(name string, factory Factory) {
	if factories[name] != nil {
		log.Fatalf("Duplicate DB.Factory register for %v", name)
	}

	factories[name] = factory
}

// OpenDB returns a Server using the provided implementation.
// database host, database port, login user, login password, database name
func OpenDB(implation string, host string, port int, user, password, dbname string) (*CobraDB, error) {
	factory, ok := factories[implation]
	if !ok {
		return nil, ErrNoDB
	}

	impl, err := factory(host, port, user, password, dbname)
	if err != nil {
		return nil, err
	}

	return &CobraDB{impl}, nil
}

// Impl is an interface for sql.DB and extend it, Add for extension method
type Impl interface {
	// Begin starts a transaction. The default isolation level is dependent on
	Begin() (*sql.Tx, error)

	// BeginTx starts a transaction.
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	// Close closes the database, releasing any open resources.
	Close() error

	// Conn returns a single connection by either opening a new connection
	// or returning an existing connection from the connection pool. Conn will
	// block until either a connection is returned or ctx is canceled.
	// Queries run on the same Conn will be run in the same database session.
	//
	// Every Conn must be returned to the database pool after use by
	// calling Conn.Close.
	Conn(ctx context.Context) (*sql.Conn, error)

	// Driver returns the database's underlying driver.
	Driver() driver.Driver

	// Exec executes a query without returning any rows.
	// The args are for any placeholder parameters in the query.
	Exec(query string, args ...interface{}) (sql.Result, error)

	// ExecContext executes a query without returning any rows.
	// The args are for any placeholder parameters in the query.
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// Ping verifies a connection to the database is still alive,
	// establishing a connection if necessary.
	Ping() error

	// PingContext verifies a connection to the database is still alive,
	// establishing a connection if necessary.
	PingContext(ctx context.Context) error

	// Query executes a query that returns rows, typically a SELECT.
	// The args are for any placeholder parameters in the query.
	Query(query string, args ...interface{}) (*sql.Rows, error)

	// QueryContext executes a query that returns rows, typically a SELECT.
	// The args are for any placeholder parameters in the query.
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	// QueryRow executes a query that is expected to return at most one row.
	// QueryRow always returns a non-nil value. Errors are deferred until
	// Row's Scan method is called.
	// If the query selects no rows, the *Row's Scan will return ErrNoRows.
	// Otherwise, the *Row's Scan scans the first selected row and discards
	// the rest.
	QueryRow(query string, args ...interface{}) *sql.Row

	// QueryRowContext executes a query that is expected to return at most one row.
	// QueryRowContext always returns a non-nil value. Errors are deferred until
	// Row's Scan method is called.
	// If the query selects no rows, the *Row's Scan will return ErrNoRows.
	// Otherwise, the *Row's Scan scans the first selected row and discards
	// the rest.
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//
	// Expired connections may be closed lazily before reuse.
	//
	// If d <= 0, connections are reused forever.
	SetConnMaxLifetime(d time.Duration)

	// SetMaxIdleConns sets the maximum number of connections in the idle
	// connection pool.
	//
	// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns
	// then the new MaxIdleConns will be reduced to match the MaxOpenConns limit
	//
	// If n <= 0, no idle connections are retained.
	SetMaxIdleConns(n int)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	//
	// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than
	// MaxIdleConns, then MaxIdleConns will be reduced to match the new
	// MaxOpenConns limit
	//
	// If n <= 0, then there is no limit on the number of open connections.
	// The default is 0 (unlimited).
	SetMaxOpenConns(n int)

	// Stats returns database statistics.
	Stats() sql.DBStats

	// DBType return database type
	DBType() string

	// ShowTables - Show db's table list
	ShowTables() ([]string, error)
}

// CobraDB is a wrapper type that can have extra methods
type CobraDB struct {
	Impl
}
