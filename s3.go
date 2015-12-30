package main

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Resource struct {
	S3Conn *s3.S3
	config *configFile
}

type S3Bucket struct {
	Bucket *s3.GetBucketACLOutput `json:"bucket"`
	Name   string                 `json:"name"`
}

// NewS3Resource is an exported method to connect on S3 and return resource.
func NewS3Resource(config *configFile) *S3Resource {
	s3_service := S3Resource{}
	s3_service.S3Conn = s3.New(&aws.Config{
		Region: &config.S3.Region,
	})
	s3_service.config = config
	return &s3_service
}

// Returns bucket name from config file
func (s3s *S3Resource) getBucketName() string {
	return s3s.config.S3.BucketName
}

// GetBucket is a exported method to get all bucket data, ACL and more.
func (s3s *S3Resource) GetBucket(bucketName string) (*S3Bucket, error) {
	params := &s3.GetBucketACLInput{
		Bucket: aws.String(bucketName),
	}
	resp, err := s3s.S3Conn.GetBucketACL(params)
	sBucket := &S3Bucket{Bucket: resp, Name: bucketName}
	return sBucket, err
}

// PutObj is a exported method to send file to S3 storage, receive all payload and set params
// returns a type PutObjectOutput, if has error, returns that.
func (s3s *S3Resource) GetObj(factsheet []byte, filename string) (*s3.PutObjectOutput, error) {
	params := &s3.PutObjectInput{
		Bucket:        aws.String(s3s.getBucketName()),
		Key:           aws.String(filename),
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(factsheet),
		ContentLength: aws.Int64(int64(len(factsheet))),
		ContentType:   aws.String("application/json; charset=utf-8"),
	}
	resp, err := s3s.S3Conn.PutObject(params)
	return resp, err
}
