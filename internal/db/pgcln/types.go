package pgcln

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/yemelin/xproject/pkg/cloud/gcptypes"
)

// Check if obj and client implement the interfaces
var _ IDB = (*sql.DB)(nil)
var _ IClient = (*Client)(nil)

// IDB is interface of sql.DB
type IDB interface {
	PingContext(ctx context.Context) error
	Ping() error
	Close() error
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	SetConnMaxLifetime(d time.Duration)
	Stats() sql.DBStats
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryRow(query string, args ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Begin() (*sql.Tx, error)
	Driver() driver.Driver
	Conn(ctx context.Context) (*sql.Conn, error)
}

// IClient is interface of Client type
type IClient interface {
	Close() error
	Ping() error
	ListAccounts() (GcpAccounts, error)
	GetLastAccount() (*GcpAccount, error)
	AddAccount(GcpAccount) error
	ListFiles() (gcptypes.FilesMetadata, error)
	GetLastFile() (*gcptypes.FileMetadata, error)
	AddFile(gcptypes.FileMetadata) error
	ListAllBills() (gcptypes.ServicesBills, error)
	ListBillsByTime(time.Time, time.Time) (gcptypes.ServicesBills, error)
	ListBillsByService(string) (gcptypes.ServicesBills, error)
	ListBillsByProject(string) (gcptypes.ServicesBills, error)
	GetLastBill() (*gcptypes.ServiceBill, error)
	AddBill(gcptypes.ServiceBill) error
	AddReport(gcptypes.Report) error
	AddReportsToAccount(gcptypes.Reports, int) error
}

// GcpAccount contains information about GCP user account
type GcpAccount struct {
	ID             int
	GcpAccountInfo string
}

// GcpAccounts is a set of GcpAccount
type GcpAccounts []*GcpAccount
