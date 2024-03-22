package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/xlizy/common-go/zlog"
	"io"
)

type RootConfig struct {
	Config ossConfig `yaml:"oss"`
}

type ossConfig struct {
	Endpoint          string `yaml:"endpoint"`
	AccessKeyId       string `yaml:"accessKeyId"`
	AccessKeySecret   string `yaml:"accessKeySecret"`
	DefaultBucketName string `yaml:"defaultBucketName"`
}

var _defaultBucketName = ""

var _client *oss.Client

func NewConfig() *RootConfig {
	return &RootConfig{}
}

func InitOSS(rc *RootConfig) {
	cfg := rc.Config
	_defaultBucketName = cfg.DefaultBucketName
	c, err := oss.New(cfg.Endpoint, cfg.AccessKeyId, cfg.AccessKeySecret)
	if err != nil {
		zlog.Error("初始化OSS异常:{}", err.Error())
		panic(err)
	}
	_client = c
}

func UploadForFile(key string, reader io.Reader, bucketName string) error {
	if bucketName == "" {
		bucketName = _defaultBucketName
	}
	bucket, err := _client.Bucket(bucketName)
	if err != nil {
		zlog.Error("初始化OSS-bucket异常", err)
	}
	return bucket.PutObject(key, reader)
}

func DeleteOSSFile(key, bucketName string) {
	if bucketName == "" {
		bucketName = _defaultBucketName
	}
	bucket, err := _client.Bucket(bucketName)
	if err != nil {
		zlog.Error("初始化OSS-bucket异常", err)
	}
	_ = bucket.DeleteObject(key)
}
