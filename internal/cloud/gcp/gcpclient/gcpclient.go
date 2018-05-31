package gcpclient

import (
	"fmt"
	"log"
	"xproject/internal/cloud/gcp/storageutil"

	"golang.org/x/net/context"

	"cloud.google.com/go/storage"
)

// TODO: constructor or singleton?
type Client struct {
	ctx    *context.Context
	client *storage.Client
}

type Client interface {
}

func init() {

}

func main() {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal("my fatal error", err)
	}

	reader := storageutil.NewBucketObjectReader(ctx, client, "churomann-bucket", "test-2018-05-23.csv")

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("error 2", err)
	}

	fmt.Println(records)
}
