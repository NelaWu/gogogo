package s3client

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
	"io"
	"log"
)

var client *s3.Client
var bucket string

func InitS3() {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(viper.GetString("s3.region")),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			viper.GetString("s3.access_key"),
			viper.GetString("s3.secret_key"),
			"",
		)),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           viper.GetString("s3.endpoint"),
					SigningRegion: region,
				}, nil
			}),
		),
	)
	if err != nil {
		log.Fatalf("S3 初始化失敗: %v", err)
	}

	client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	bucket = viper.GetString("s3.bucket")

	// 檢查 bucket 是否存在，如果不存在就創建
	_, err = client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		log.Printf("檢查 bucket 失敗: %v，嘗試創建新的 bucket", err)

		// 如果 bucket 不存在，創建它
		_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(bucket),
		})
		if err != nil {
			log.Printf("創建 bucket 失敗: %v", err)
		} else {
			log.Printf("成功創建 bucket: %s", bucket)
		}
	}
}

func UploadToS3(key string, body io.Reader, contentType string) error {
	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})
	return err
}

func ListObjects() (*s3.ListObjectsV2Output, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	}
	ctx := context.TODO()
	result, err := client.ListObjectsV2(ctx, input)
	return result, err
}
