package gcpclient

import (
	"fmt"
	"log"
	"xproject/internal/cloud/gcp/storageutil"

	"golang.org/x/net/context"

	"cloud.google.com/go/storage"
)

type Client struct {
	ctx           context.Context
	storageClient *storage.Client
}

type Client interface {
	int64 thisMounth()
	estimated
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

func main() {

	reader := storageutil.GetBucketObjectCSVReader(ctx, client, "churomann-bucket", "test-2018-05-23.csv")

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("error 2", err)
	}

	fmt.Println(records)
}
