package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

var Cfg *Config

type Config struct {
	Database struct {
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	} `yaml:"database"`
	Etcd struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"etcd"`
	Server struct {
		User struct {
			Name string `yaml:"name"`
			Port int    `yaml:"port"`
		} `yaml:"user"`
		Auth struct {
			Name string `yaml:"name"`
			Port int    `yaml:"port"`
		} `yaml:"auth"`
		Video struct {
			Name string `yaml:"name"`
			Port int    `yaml:"port"`
		} `yaml:"video"`
	} `yaml:"server"`
	File struct {
		Dir string `yaml:"dir"`
	} `yaml:"file"`
}

func init() {
	_, filename, _, _ := runtime.Caller(0)
	configDir := path.Dir(filename)
	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(Cfg)
}
