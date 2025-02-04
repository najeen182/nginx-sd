package config

import "github.com/spf13/viper"

type ACME struct {
	Email string `toml:"email" mapstructure:"email"`
	CA    string `toml:"ca" mapstructure:"ca"`
}

type Config struct {
	Acme         ACME         `toml:"acme" mapstructure:"acme"`
	ConfigSource ConfigSource `toml:"configSource" mapstructure:"configSource"`
}

type S3 struct {
	AccessKey    string `toml:"accessKey" mapstructure:"accessKey"`
	AccessSecret string `toml:"accessSecret" mapstructure:"accessSecret"`
	EndPoint     string `toml:"endPoint" mapstructure:"endPoint"`
	BucketName   string `toml:"bucketName" mapstructure:"bucketName"`
}

type ConfigSource struct {
	Source string `toml:"string" mapstructure:"string"`
	Path   string `toml:"path" mapstructure:"string"`
	S3     *S3    `toml:"s3" mapstructure:"s3"`
}

var RuntimeViper = viper.New()
var (
	AppName    = "Nginx SD"
	AppVersion = "0.0.0"
)
