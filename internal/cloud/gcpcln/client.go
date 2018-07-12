// Package gcpcln periodically fetches csv from gcp bucket and writes in into db
package gcpcln

import (
	"context"
	"encoding/csv"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"github.com/pavlov-tony/xproject/pkg/cloud/gcpparser"
	"google.golang.org/api/iterator"
)

// mime content type
const contentTypeCsv = "text/csv"

// Client for GCP predictions which stores google api as fileds
type Client struct {
	strgCln *storage.Client
	ctx     context.Context
}

type Object struct {
	Name    string
	Bucket  string
	Created time.Time
}

type Objects []Object

type Report struct {
	Object Object
	Bills  gcpparser.ServicesBills
}

type Reports []Report

func (objs Objects) after(t time.Time) (res Objects) {
	for _, o := range objs {
		if o.Created.After(t) {
			res = append(res, o)
		}
	}

	return res
}

// NewClient creates new client for GCP
func NewClient(ctx context.Context) (*Client, error) {
	c := new(Client)
	strgCln, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("new client: %v", err)
	}
	c.ctx = ctx
	c.strgCln = strgCln

	return c, nil
}

// BucketsList fetches bucket list from project
func (c *Client) BucketsList(projectID string) (buckets []string, err error) {

	it := c.strgCln.Buckets(c.ctx, projectID)
	for {
		b, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("fetch project buckets: %v", err)
		}
		buckets = append(buckets, b.Name)
	}

	return buckets, nil
}

// CsvObjectsList fetches csv ojects list from bucket with prefix
func (c *Client) CsvObjsList(bktName, prefix string) (objs Objects, err error) {

	it := c.strgCln.Bucket(bktName).Objects(c.ctx, &storage.Query{Prefix: prefix})
	for {
		o, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("fetch bucket objects: %v", err)
		}
		if o.ContentType == contentTypeCsv {
			fmt.Println(o)
			objs = append(objs, Object{Name: o.Name, Created: o.Created})
		}
	}

	return objs, nil
}

// CsvObjectContent fetches data from bucket object by bucket name and object name
func (c *Client) CsvObjectContent(bktName, objName string) ([][]string, error) {
	r, err := c.strgCln.Bucket(bktName).Object(objName).NewReader(c.ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch object csv data: %v", err)
	}
	records, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("fetch object csv data: %v", err)
	}

	return records, nil
}

func (c *Client) MakeReport(obj Object) (Report, error) {
	return
}

func (c *Client) MakeReports(objs Objects) (Reports, error) {
	return
}

// Fetch periodically fetches data into db from GCP
func (c *Client) Fetch(bktName, prefix string, dt time.Duration) {
	for {
		// TODO: select last report from db
		// lastReport := dbclient.getLastReport()

		// select last date from db
		objs := c.CsvObjsList(bktName, prefix)
		// objs := objs.after(lastReport.Object.Created)

		reps := c.MakeReports(objs)

		// TODO: use pgcln here to write parsed csv into db

		time.Sleep(dt * time.Hour)
	}
}

// select client
// select Buckets
// choose bucket
// fetch from bucket obj with prefix since
