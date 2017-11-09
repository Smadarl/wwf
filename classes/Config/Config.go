package Config

import (
	"github.com/go-ini/ini"
	"fmt"
//	"utils/file"
//	"utils/user"
//	"classes/Logger"
)

const (
	CONFIG_FILE = "/etc/wwf/config.ini"
	defaultRequestResetTime float64 = 60
	defaultRateLimit int = 5000
	defaultSessionTimeout int64 = 30
)

var (
	config *CodeConfig
)

type CodeConfig struct {
	Database DB `ini:"db"`
	Google   Google `ini:"googleauth"`
	RabbitMQ RabbitMQ `ini:"rabbitmq"`
	Facebook Facebook `ini:"facebook"`
}

type DB struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Log      string `ini:"log"`
	Security bool `ini:"security"`
	Path     string `ini:"path"`
}

type Facebook struct {
    AppID       string `ini:"appId"`
    AppSecret   string `ini:"appSecret"`
    CallbackURL string `ini:"callbackUrl"`
}

type Google struct {
	ClientId     string `ini:"client_id"`
	ClientSecret string `ini:"client_secret"`
	RedirectUrl  string `ini:"redirect_url"`
	AuthUrl      string `ini:"auth_url"`
	TokenUrl     string `ini:"token_url"`
}

type RabbitMQ struct {
	Username string `ini:"username"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
}

func init() {
	config = new(CodeConfig)
	err := ini.MapTo(config, CONFIG_FILE)
	if err != nil {
		fmt.Println("Failed to load config file: " + err.Error())
	}

}

func GetConfig() *CodeConfig {
    return config
}

func GetDatabase() DB {
	return config.Database
}

func GetGoogleAuth() Google {
	return config.Google
}

func GetRabbitMQ() RabbitMQ {
	return config.RabbitMQ
}

func GetFacebook() Facebook {
	return config.Facebook
}
