package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
	"xproject/internal/cloud/gcp/storageutil"

	"cloud.google.com/go/storage"
)

// TODO: think about  structure
// bill per service
type ServiceBill struct {
	Description string // NOTE: field from csv table: Now - 17 (Description)
	StartTime   time.Time
	EndTime     time.Time
	Cost        float64
	Currency    string
}

// NOTE: check copy object may be need ptr
func (sb *ServiceBill) setAttributes(src []string) {
	sb.Description = src[17]
	t, err := time.Parse(time.RFC3339, src[2])
	if err != nil {
		log.Fatal(err)
	}
	sb.StartTime = t
	t, err = time.Parse(time.RFC3339, src[3])
	if err != nil {
		log.Fatal(err)
	}
	sb.EndTime = t
	cost, err := strconv.ParseFloat(src[11], 64)
	if err != nil {
		log.Fatal(err)
	}
	sb.Cost = cost
	sb.Currency = src[12]
}

type ServicesBills []ServiceBill

func main() {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	objects, err := storageutil.FetchBucketObjects(ctx, client, "churomann-bucket")
	if err != nil {
		log.Fatal(err)
	}
	csvReader := objects[0].NewCSVReader(ctx, client)
	row, err := csvReader.Read()
	for i, _ := range row {
		fmt.Println(i, "-", row[i])
	}
	fmt.Println()
	row, err = csvReader.Read()

	for i, _ := range row {
		fmt.Println(i, "-", row[i])
	}

	var sb ServiceBill
	sb.setAttributes(row)
	fmt.Println(sb)

	// var servicesBills ServicesBills

	// for {
	// 	row, err := csvReader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	//
	// 	servicesBills = append(servicesBills, ServiceBill{row[17]})
	//
	// }
}
