package config

import (
	"fmt"
	"github.com/spf13/viper"
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
	} `yaml:"server"`
}

func init() {
	//viper.AddConfigPath("src/config")
	//viper.SetConfigName("config")
	//viper.SetConfigType("yml")
	viper.SetConfigFile("C:\\Users\\67561\\GolandProjects\\rpc-douyin\\src\\config\\config.yml")
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
