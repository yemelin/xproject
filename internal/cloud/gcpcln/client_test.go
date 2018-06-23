package gcpcln

import (
	"testing"
)

func Test_Client(t *testing.T) {
	// ctx := context.Background()
	// projectID := os.Getenv("APP_PROJECT_ID")
	// client, err := storage.NewClient(ctx)
	// if err != nil {
	// 	t.Error("Failed to create client:", err)
	// }
	//
	// it := client.Buckets(ctx, projectID)
	// for {
	// 	b, err := it.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	t.Log(b)
	// }
}

//
// func main() {
// 	ctx := context.Background()
//
// 	client, err := storage.NewClient(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	objects, err := storageutils.FetchBucketObjects(ctx, client, "churomann-bucket")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	var serviceBills billingutils.ServicesBills
// 	serviceBills.FillByObjects(ctx, client, &objects)
//
// 	var sum float64
// 	for _, sb := range serviceBills {
// 		sum += sb.Cost
// 	}
// 	fmt.Println("Full cost from bucket:", sum)
//
// 	sum = 0
// 	objects, err = objects.SelectInTimeRange(
// 		time.Date(2018, time.May, 29, 0, 0, 0, 0, time.Local),
// 		time.Now())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	serviceBills.FillByObjects(ctx, client, &objects)
// 	for _, sb := range serviceBills {
// 		sum += sb.Cost
// 	}
// 	fmt.Println("Full cost in time period:", sum)
// }
