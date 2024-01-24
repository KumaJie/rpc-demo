package config

import (
	"github.com/spf13/viper"
	"path"
	"runtime"
	"strings"
)

var Cfg *Config

type Config struct {
	Database Database `yaml:"database"`
	Etcd     Etcd     `yaml:"etcd"`
	Kafka    Kafka    `yaml:"kafka"`
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
type Kafka struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type User struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type Auth struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type Video struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type Favorite struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type Comment struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
type Server struct {
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
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
	viper.SetConfigType("yml")
	// 读取环境变量
	viper.AutomaticEnv()
	// 将环境变量替换为配置信息, DATABASE.NAME -> DATABASE_NAME
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		panic(err)
	}
}
