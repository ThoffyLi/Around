package gcs_save

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
)

func saveToGCS(ctx context.Context, r io.Reader, bucketName, name string) (*storage.ObjectHandle, *storage.ObjectAttrs, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	// Next check if the bucket exists
	if _, err = bucket.Attrs(ctx); err != nil {
		return nil, nil, err
	}

	obj := bucket.Object(name)
	wc := obj.NewWriter(ctx)
	if _, err := io.Copy(wc, r); err != nil {
		return nil, nil, err
	}
	if err := wc.Close(); err != nil {
		return nil, nil, err
	}

	// set file-reading permission to all users after writing to GCS (default is non-readable)
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return nil, nil, err
	}

	// get attrs of the file we just stored and return
	attrs, err := obj.Attrs(ctx)
	fmt.Printf("Post is saved to GCS: %s\n", attrs.MediaLink)
	return obj, attrs, err

}
