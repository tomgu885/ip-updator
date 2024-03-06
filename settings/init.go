package settings

import (
	"fmt"
	"github.com/spf13/viper"
)

type Server struct {
	Port int
}

type Client struct {
	Server    string `json:"server" yaml:"server"`
	Name      string `json:"name" yaml:"name"`
	LocalPort int    `json:"local_port" yaml:"localPort"`
}

type Global struct {
	Mode   string `json:"mode" yaml:"mode"`
	Secret string `json:"secret" yaml:"secret"`
}

type Config struct {
	Global Global `json:"global" yaml:"global"`
	Server Server `json:"server" yaml:"server"`
	Client Client `json:"client" yaml:"client"`
}

var (
	conf Config
)

func GetGlobal() *Global {
	return &conf.Global
}

func GetServer() *Server {
	return &conf.Server
}

func GetClient() *Client {
	return &conf.Client
}

func Setup(confPath string) (err error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(confPath)

	if errLoad := viper.ReadInConfig(); errLoad != nil {
		fmt.Println("failed to load config:", confPath, "err:", errLoad.Error())
		return errLoad
	}

	if errUn := viper.Unmarshal(&conf); errUn != nil {
		fmt.Println("failed to unmarshal config", errUn.Error())
		return errUn
	}

	return
}
