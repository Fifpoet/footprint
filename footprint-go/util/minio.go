package utils

import (
	"context"
	"github.com/fifpoet/footprint/global"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
	"mime/multipart"
	"net/url"
	"time"
)

func CreateBucket(c *minio.Client, bucket string, location string) error {
	ctx := context.Background()
	b, _ := c.BucketExists(ctx, bucket)
	if b {
		return nil
	}
	err := c.MakeBucket(ctx, bucket, minio.MakeBucketOptions{
		Region:        location,
		ObjectLocking: false,
	})
	return err
}

func UploadFile(c *minio.Client, bucket string, location string, obj string, reader multipart.File, size int64) (string, error) {
	ctx := context.Background()
	err := CreateBucket(c, bucket, location)
	if err != nil {
		global.FP_LOG.Error("bucket create error", zap.Error(err))
		return "", err
	}
	info, err := c.PutObject(ctx, bucket, obj, reader, size, minio.PutObjectOptions{
		ContentType: "image/jpg",
	})
	if err != nil {
		global.FP_LOG.Error("upload failed", zap.Error(err))
		return "", err
	}
	global.FP_LOG.Info("upload success ", zap.String("target file", obj), zap.Int("file size", int(info.Size)))
	//生成URL
	fileUrl, err := c.PresignedGetObject(ctx, bucket, obj, time.Hour*24*7, url.Values{})
	if err != nil {
		global.FP_LOG.Error("generate url failed", zap.Error(err))
		return "", err
	}
	return global.FP_CONFIG.Minio.Endpoint + fileUrl.Path, nil
}
