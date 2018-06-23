package gcpcln

import (
	"context"
	"encoding/csv"
	"errors"

	"cloud.google.com/go/storage"
)

// Common client for GCP predictions which stores google api as fileds
type Client struct {
	strgCln *storage.Client
	ctx     context.Context
}

// Predictor should be common interface for AWS and GCP
type Predictor interface {
}

// Creating new client
func NewClient(ctx context.Context) (c *Client, err error) {
	strgCln, err := storage.NewClient(ctx)
	if err != nil {
		return nil, errors.New("NewClient: " + err.Error())
	}
	c.ctx = ctx
	c.strgCln = strgCln
	return c, nil
}

// fetch data from bucket object by bucket name and object name
func (c *Client) fetchObjectCSVData(bktName, objName string) ([][]string, error) {
	r, err := c.strgCln.Bucket(bktName).Object(objName).NewReader(c.ctx)
	if err != nil {
		return nil, errors.New("fetchObjectCSVData: " + err.Error())
	}
	records, err := csv.NewReader(r).ReadAll()
	if err != nil {
		return nil, errors.New("fetchObjectCSVData: " + err.Error())
	}
	return records, nil
}
