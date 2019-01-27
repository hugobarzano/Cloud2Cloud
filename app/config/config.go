package config

import (
	"github.com/fulldump/goconfig"
	"github.com/gorilla/sessions"
	"sync"
)

var config *AppConfiguration
var once sync.Once


type MongodbConfig struct {
	URI string `json:"uri"`
}

type AppConfiguration struct {
	Mongodb          MongodbConfig          `json:"mongodb"`
}



var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)



func InitConfig() {
	config = &AppConfiguration{

		Mongodb: MongodbConfig{
			URI: "mongodb://localhost:27017/app",
		},
	}
	goconfig.Read(config)
}

func Get() *AppConfiguration {
	once.Do(func() {
		if config == nil {
			InitConfig()
		}
	})
	return config
}
