package config

type Minio struct {
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`                            // 对象存储服务的URL
	AccessKeyID     string `mapstructure:"access_key_id" json:"access_key_id" yaml:"access_key_id"`             // Access Key 是唯一标识账户的用户ID
	SecretAccessKey string `mapstructure:"secret_access_key" json:"secret_access_key" yaml:"secret_access_key"` // 对应账户的密码
	UseSSL          bool   `mapstructure:"use_ssl" json:"use_ssl" yaml:"use_ssl"`                               // true 为使用 https
	Bucket          string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Location        string `mapstructure:"location" json:"location" yaml:"location"`
}
