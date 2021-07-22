package storage

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
)

var storageClient *storage.Client

func InitStorage() {
	var err error
	ctx := context.Background()
	if storageClient, err = storage.NewClient(ctx); err != nil {
		fmt.Println("connection to bucket failed")
		os.Exit(1)
	}
}

func Upload(bucketName string, fileName string, dataType string, data []byte) {
	ctx := context.Background()
	obj := storageClient.Bucket(bucketName).Object(fileName)
	w := obj.NewWriter(ctx)
	meta := make(map[string]string)
	meta["dataType"] = dataType
	w.ObjectAttrs.Metadata = meta
	defer w.Close()
	if _, err := w.Write(data); err != nil {
		fmt.Println("Error while uploading file to storage", err)
		return
	}
}

func Delete(bucketName string, fileName string) error {
	ctx := context.Background()
	if err := storageClient.Bucket(bucketName).Object(fileName).Delete(ctx); err != nil {
		fmt.Println("Error while deleting file to storage", err)
		return err
	}
	return nil
}
