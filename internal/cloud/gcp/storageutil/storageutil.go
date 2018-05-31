package storageutil

import (
	"encoding/csv"
	"errors"
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

// FIXME: One client or every time new client?

// TODO: testing (need test api?)
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
		objects = append(objects, Object{o.Name, o.Created})
	}

	return objects
}

// TODO: testing
func NewBucketObjectReader(ctx context.Context, client *storage.Client,
	bucketName, objectName string) *csv.Reader {
	objectReader, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return csv.NewReader(objectReader) // TODO:  have no closed yet
}

// FIXME: delete from slice or insert into new slice?
func SelectObjectsWithPrefix(objects []Object, prefix string) []Object {
	for i, o := range objects {
		if !strings.HasPrefix(o.Name, prefix) {
			objects = append(objects[:i], objects[i+1:]...) // FIXME: memory and order problem?
		}
	}

	return objects
}

// FIXME: delete from slice or insert into new slice?
func SelectObjectsWithFromToTime(objects []Object, from, to time.Time) (result []Object, err error) {
	if to.Before(from) {
		return result, errors.New("error: to before from")
	}
	for _, o := range objects {
		if o.Created.After(from) && o.Created.Before(to) {
			result = append(result, o)
		}
	}

	return result, nil
}
