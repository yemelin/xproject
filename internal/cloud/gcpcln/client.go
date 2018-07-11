// Package gcpcln periodically fetches csv from gcp bucket and writes in into db
package gcpcln

import (
	"context"
	"encoding/csv"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// mime content type
const contentTypeCsv = "text/csv"

// Client for GCP predictions which stores google api as fileds
type Client struct {
	strgCln *storage.Client
	ctx     context.Context
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
func (c *Client) CsvObjectsList(bktName, prefix string) (objs []string, err error) {

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
			objs = append(objs, o.Name)
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

// Fetch periodically fetches data into db from GCP
func (c *Client) Fetch(bktName, objName string, dt time.Duration) {
	for {
		// select last date from db

		c.CsvObjectContent(bktName, objName)
		// TODO: use CsvObjectContent
		// TODO: parse raw csv content
		// TODO: use pgcln here to write parsed csv into db

		time.Sleep(dt * time.Hour)
	}
}
