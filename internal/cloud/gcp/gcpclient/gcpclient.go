package gcpclient

import (
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

// Template?
func GetProjectBuckets(ctx context.Context, client *storage.Client,
	projectID string) (buckets []string) {

	it := client.Buckets(ctx, projectID)
	for {
		b, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("GetBucket() failed")
			return
		}
		buckets = append(buckets, b.Name)
	}

	return buckets
}

func GetBucketObjects(ctx context.Context, client *storage.Client,
	bucketName string) (objects []string) {

	it := client.Bucket(bucketName).Objects(ctx, nil)
	for {
		o, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("GetBucketObjects() failed")
			return
		}
		fmt.Println(o.Name, o.Created)
		objects = append(objects, o.Name)
	}

	return objects
}

// func ()  {
//
// }

// func ()  {
//
// }

// func main() {
// 	ctx := context.Background()
//
// 	// Sets your Google Cloud Platform project ID.
// 	projectID := os.Getenv("APP_PROJECT_ID")
//
// 	// Creates a client.
// 	client, err := storage.NewClient(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}
//
// 	bucket, err := client.Buckets(ctx, projectID).Next()
// 	if err != nil {
// 		log.Fatal("fatal")
// 	}
//
// 	rc, err := client.Bucket(bucket.Name).Object("test-2018-05-23.csv").NewReader(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rc.Close()
//
// 	csvReader := csv.NewReader(rc)
// 	res, err := csvReader.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(res[1][0])
// }
