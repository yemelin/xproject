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

	csvReader := objects[0].NewCSVReader(ctx, client)

	var serviceBills billingutils.ServicesBills
	serviceBills.Fill(csvReader)
	fmt.Println(serviceBills)

}
