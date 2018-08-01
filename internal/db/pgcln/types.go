package pgcln

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
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
	AddAccount(GcpAccount) error
	ListCsvFiles() (GcpCsvFiles, error)
	AddCsvFile(GcpCsvFile) error
	ListAllBills() (ServiceBills, error)
	ListBillsByTime(time.Time, time.Time) (ServiceBills, error)
	ListBillsByService(string) (ServiceBills, error)
	ListBillsByProject(string) (ServiceBills, error)
	GetLastBill() (ServiceBill, error)
	AddBill(ServiceBill) error
}

// GcpAccount contains information about GCP user account
type GcpAccount struct {
	ID             int
	GcpAccountInfo string
}

// GcpCsvFile contains information about CSV files with billing reports
type GcpCsvFile struct {
	ID          int
	Name        string
	Bucket      string
	TimeCreated time.Time
	AccountID   int
}

// ServiceBill contains relevant information from billing report
type ServiceBill struct {
	ID           int
	LineItem     string
	StartTime    time.Time
	EndTime      time.Time
	Cost         float64
	Currency     string
	ProjectID    string
	Description  string
	GcpCsvFileID int
}

// GcpAccounts is a set of GcpAccount
type GcpAccounts []*GcpAccount

// GcpCsvFiles is a set of GcpCsvFile
type GcpCsvFiles []*GcpCsvFile

// ServiceBills is a set of ServiceBill
type ServiceBills []*ServiceBill
