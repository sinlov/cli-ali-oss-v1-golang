package clialioss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"errors"
	"fmt"
)

var errorNotInitClient = errors.New("not InitClient")

type SDKAliOss struct {
	client *oss.Client
	bucket *oss.Bucket
}

func (sdk *SDKAliOss) InitClient(endpoint, accessKeyID, accessKeySecret, bucketName string) error {
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}
	sdk.client = client
	sdk.bucket = bucket
	return nil
}

func (sdk *SDKAliOss) SearchBucket(objKey string) (string, error) {
	if sdk.bucket != nil {
		return "", errorNotInitClient
	}
	listObjectsResult, err := sdk.bucket.ListObjects()
	if err != nil {
		return "", err
	}
	for _, object := range listObjectsResult.Objects {
		if object.Key == objKey {
			return objKey, nil
		}
	}
	return "", errors.New(fmt.Sprintf("not found by objKey [ %v ], at bucket [ %v ]", objKey, sdk.bucket.BucketName))
}

func (sdk *SDKAliOss) ListBucket() (oss.ListObjectsResult, error) {
	var out oss.ListObjectsResult
	if sdk.bucket != nil {
		return out, errorNotInitClient
	}
	result, err := sdk.bucket.ListObjects()
	if err != nil {
		return out, err
	}
	return result, nil
}

func (sdk *SDKAliOss) PutBucketLocalFile(objKey, path string) (string, error) {
	if sdk.bucket != nil {
		return "", errorNotInitClient
	}
	err := sdk.bucket.PutObjectFromFile(objKey, path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("put success\nObjKey => %v\nPath => %v", objKey, path), nil
}

func (sdk *SDKAliOss) PopBucketLocalFile(objKey, path string) (string, error) {
	if sdk.bucket != nil {
		return "", errorNotInitClient
	}
	err := sdk.bucket.GetObjectToFile(objKey, path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("pop success\nObjKey => %v\nPath => %v", objKey, path), nil
}

func (sdk *SDKAliOss) DeleteBucketLocalFile(objKey string) (string, error) {
	if sdk.bucket != nil {
		return "", errorNotInitClient
	}
	err := sdk.bucket.DeleteObject(objKey)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("delete success\nObjKey => %v", objKey), nil
}
