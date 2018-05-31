package gcpclient

import (
	"log"

	"golang.org/x/net/context"

	"cloud.google.com/go/storage"
)

// TODO: way to hide ctx and client?
type Client struct {
	ctx           context.Context
	storageClient *storage.Client
}

// TODO: ALL
type Predictor interface {
	predict()
}

// TODO: constructor or singleton?
func (c Client) init() {
	c.ctx = context.Background()
	storageClient, err := storage.NewClient(c.ctx)
	if err != nil {
		log.Fatal("Client init error\n", err)
	}
	c.storageClient = storageClient
}
