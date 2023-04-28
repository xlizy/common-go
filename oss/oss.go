package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/xlizy/common-go/zlog"
	"io"
)

type RootConfig struct {
	Config OSSConfig `yaml:"oss"`
}

type OSSConfig struct {
	Endpoint          string `yaml:"endpoint"`
	AccessKeyId       string `yaml:"accessKeyId"`
	AccessKeySecret   string `yaml:"accessKeySecret"`
	DefaultBucketName string `yaml:"defaultBucketName"`
}

var defaultBucketName = ""

var client *oss.Client

func InitOSS(rc RootConfig) {
	cfg := rc.Config
	defaultBucketName = cfg.DefaultBucketName
	_client, err := oss.New(cfg.Endpoint, cfg.AccessKeyId, cfg.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	client = _client
}

func UploadForFile(key string, reader io.Reader, bucketName string) error {
	if bucketName == "" {
		bucketName = defaultBucketName
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		zlog.Error("初始化OSS-bucket异常", err)
	}
	return bucket.PutObject(key, reader)
}
