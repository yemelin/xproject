package main

import (
	"context"
	"fmt"
	"log"
	"xproject/internal/cloud/gcp/utils/storageutils"
	"xproject/internal/cloud/gcp/utils/storageutils/billingutils"

	"cloud.google.com/go/storage"
)

func main() {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	objects, err := storageutils.FetchBucketObjects(ctx, client, "churomann-bucket")
	if err != nil {
		log.Fatal(err)
	}

	var serviceBills billingutils.ServicesBills
	serviceBills.FillByObject(ctx, client, &objects[0])

	var sum float64
	for _, sb := range serviceBills {
		sum += sb.Cost
		fmt.Println(sb.Cost)
	}
	fmt.Println(sum)
}
