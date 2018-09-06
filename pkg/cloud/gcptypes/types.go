package gcptypes

// gcptypes package need fo represent gcp types (csv objects metadata and content)

import (
	"time"
)

// FileMetadata represents csv objects (files) from gcp
type FileMetadata struct {
	ID        int
	Name      string
	Bucket    string
	Created   time.Time
	AccountID int
}

// FilesMetadata is a list of Object structures
type FilesMetadata []FileMetadata

// ServiceBill represent a composition of main service fields from gcp billing csv file
// for cost calculation tasks
type ServiceBill struct {
	LineItem    string
	StartTime   time.Time
	EndTime     time.Time
	Cost        float64
	Currency    string
	ProjectID   string
	Description string
}

// ServicesBills is a set of ServiceBill
type ServicesBills []*ServiceBill

// Report is a composition of file metadata and bills for this file
type Report struct {
	Metadata FileMetadata
	Bills    ServicesBills
}

// Reports is a list of Report ptrs
type Reports []*Report

// After filters object, select only objects which after t
func (f FilesMetadata) After(t time.Time) (res FilesMetadata) {
	for _, v := range f {
		if v.Created.After(t) {
			res = append(res, v)
		}
	}

	return res
}
