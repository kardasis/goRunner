package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func putRunToS3(run Run) error {
	byteArray, err := json.Marshal(run)
	if err != nil {
		fmt.Println(err)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)

	bucket := aws.String(os.Getenv("BUCKET_NAME"))
	filename := aws.String(fmt.Sprintf("run-%s.json", run.StartTime))

	body := bytes.NewReader(byteArray)
	input := &s3.PutObjectInput{
		Bucket: bucket,
		Key:    filename,
		Body:   body,
	}

	output, err := client.PutObject(context.TODO(), input)
	if err != nil {
		return err
	}
	fmt.Println(output)

	return nil
}
