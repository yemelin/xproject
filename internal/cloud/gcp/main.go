package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"xproject/internal/cloud/gcp/utils/billingutils"
	"xproject/internal/cloud/gcp/utils/storageutils"

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
	serviceBills.FillByObjects(ctx, client, &objects)

	var sum float64
	for _, sb := range serviceBills {
		sum += sb.Cost
	}
	fmt.Println("Full cost from bucket:", sum)

	sum = 0
	objects, err = objects.SelectInTimeRange(
		time.Date(2018, time.May, 29, 0, 0, 0, 0, time.Local),
		time.Now())
	if err != nil {
		log.Fatal(err)
	}
	serviceBills.FillByObjects(ctx, client, &objects)
	for _, sb := range serviceBills {
		sum += sb.Cost
	}
	fmt.Println("Full cost in time period:", sum)
}
