package main

import (
	"fmt"
	"log"
	"os"
	"xproject/internal/cloud/gcp/gcpclient"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

func main() {

	ctx := context.Background()
	projectID := os.Getenv("APP_PROJECT_ID")
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	buckets := gcpclient.GetProjectBuckets(ctx, client, projectID)
	fmt.Println(buckets)
	objects := gcpclient.GetBucketObjects(ctx, client, buckets[0])
	for _, o := range objects {
		fmt.Println(o)
	}

}
