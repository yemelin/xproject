package gcpclient

import (
	"fmt"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

type Object struct {
	Name    string
	Created time.Time
}

// TODO: testing
func FetchProjectBuckets(ctx context.Context, client *storage.Client,
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

// TODO: testing
func FetchBucketObjects(ctx context.Context, client *storage.Client,
	bucketName string) (objects []Object) {

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
		fmt.Println(o.Name, o.Created) // FIXME: cut
		objects = append(objects, Object{o.Name, o.Created})
	}

	return objects
}

// FIXME: delete from slice or insert into new slice?
func SelectObjectsWithPrefix(objects []Object, prefix string) []Object {
	for i, o := range objects {
		if !strings.HasPrefix(o.Name, prefix) {
			objects = append(objects[:i], objects[i+1:]...) // FIXME: memory and order problem?
		}
	}
	fmt.Println(objects) // FIXME: cut

	return objects
}

// FIXME: delete from slice or insert into new slice?
func SelectObjectsWithFromToTime(objects []Object, from, to time.Time) (result []Object) {
	if to.Before(from) {
		log.Fatalln("SelectObjectsWithFromToTime", "to < from")
	}
	for _, o := range objects {
		if o.Created.After(from) && o.Created.Before(to) {
			result = append(result, o)
		}
	}

	return result
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
