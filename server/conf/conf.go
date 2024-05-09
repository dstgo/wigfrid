package conf

import (
	"github.com/ginx-contribs/dbx"
	"log/slog"
	"time"
)

type App struct {
	GRPC  GRPC        `mapstructure:"grpc"`
	Log   Log         `mapstructure:"log"`
	Redis Redis       `mapstructure:"redis"`
	DB    dbx.Options `mapstructure:"db"`

	Version   string `mapstructure:"_"`
	BuildTime string `mapstructure:"_"`
}

type Log struct {
	Filename string     `mapstructure:"filename"`
	Prompt   string     `mapstructure:"-"`
	Level    slog.Level `mapstructure:"level"`
	Format   string     `mapstructure:"format"`
	Source   bool       `mapstructure:"source"`
	Color    bool       `mapstructure:"color"`
}

type Redis struct {
	Address      string        `mapstructure:"address"`
	Password     string        `mapstructure:"password"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
}

type GRPC struct {
	Address       string `mapstructure:"address"`
	MaxRecv       int    `mapstructure:"maxRecv"`
	MaxSend       int    `mapstructure:"maxSend"`
	ReadBuffer    int    `mapstructure:"readBuffer"`
	WriteBuffer   int    `mapstructure:"writeBuffer"`
	MaxHeaderSize uint32 `mapstructure:"maxHeaderSize"`
	TLS           struct {
		Cert string `mapstructure:"cert"`
		Key  string `mapstructure:"key"`
	} `mapstructure:"tls"`
}
