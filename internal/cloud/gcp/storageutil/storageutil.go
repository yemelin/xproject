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

// make reader?
type Object struct {
	Name    string
	Created time.Time
}

type Objects []Object

// TODO: testing (need test api?)
// fetching project bucket names
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

// fetching objects names with created time from bucket
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
// csv reader for files from bucket
func GetBucketObjectCSVReader(ctx context.Context, client *storage.Client,
	bucketName, objectName string) *csv.Reader {
	objectReader, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return csv.NewReader(objectReader) // TODO:  have no closed yet
}

// func (object Object) NewReader(context.Context) {
// 	storage.Reader
// }

// FIXME: *Objects?
// Select objects from objects list where name has prefix
func (objects Objects) SelectWithPrefix(prefix string) (result Objects) {
	for _, o := range objects {
		if strings.HasPrefix(o.Name, prefix) {
			result = append(result, o)
		}
	}

	return result
}

// Select objects from objects list where from < created time < to
func (objects Objects) SelectInTimeRange(from, to time.Time) (result Objects, err error) {
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
