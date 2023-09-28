package core

import (
	"fmt"
	"github.com/fifpoet/footprint/global"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

func InitMinio() *minio.Client {
	conf := global.FP_CONFIG.Minio
	minioClient, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKeyID, conf.SecretAccessKey, ""),
		Secure: conf.UseSSL,
	})
	if err != nil {
		global.FP_LOG.Error("failed to conn to minio", zap.Error(err))
	}
	fmt.Println("====5-minIO====: minIO link success")
	return minioClient
}
