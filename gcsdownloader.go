// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Sample objects creates, list, deletes objects and runs
// other similar operations on them by using the Google Storage API.
// More documentation is available at
// https://cloud.google.com/storage/docs/json_api/v1/.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"


	"golang.org/x/net/context"

	"cloud.google.com/go/storage"
)

func main() {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		os.Exit(1)
	}
	fullPath := os.Getenv("GCS_FULL_PATH")
	if fullPath == "" {
		fmt.Fprintf(os.Stderr, "GCS_FULL_PATH environment variable must be set.\n")
		os.Exit(1)
	}

	names := strings.Split(fullPath, ":")
	if len(names) < 2 {
		usage("wrong path")
	}
	bucket, object := names[0], names[1]

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	data, err := read(client, bucket, object)
	if err != nil {
		log.Fatalf("Cannot read object: %v", err)
	}
	ioutil.WriteFile(object, data, 0644)

}

func read(client *storage.Client, bucket, object string) ([]byte, error) {
	ctx := context.Background()
	// [START download_file]
	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
	// [END download_file]
}


const helptext = `usage: gcsdownloader

ENV VARS:
	- GCS_FULL_PATH=bucket:name
`

func usage(msg string) {
	if msg != "" {
		fmt.Fprintln(os.Stderr, msg)
	}
	fmt.Fprintln(os.Stderr, helptext)
	os.Exit(2)
}
