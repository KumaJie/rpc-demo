package config

import (
	"github.com/spf13/viper"
	"path"
	"runtime"
)

var Cfg *Config

type Config struct {
	Database Database `yaml:"database"`
	Etcd     Etcd     `yaml:"etcd"`
	Server   Server   `yaml:"server"`
	File     File     `yaml:"file"`
}
type Database struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}
type Etcd struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type User struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}
type Auth struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}
type Video struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}
type Favorite struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}
type Comment struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}
type Server struct {
	User     User     `yaml:"user"`
	Auth     Auth     `yaml:"auth"`
	Video    Video    `yaml:"video"`
	Favorite Favorite `yaml:"favorite"`
	Comment  Comment  `yaml:"comment"`
}
type File struct {
	Dir string `yaml:"dir"`
	Max int    `yaml:"max"`
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
}
