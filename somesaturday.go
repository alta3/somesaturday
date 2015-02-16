package main

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type Config struct {
	BasePath   string
	HostName   string
	SalesEmail string
	Verbose    bool
	Hours      string
}

var config Config

func ParseConfig(configFile string) {
	var config Config
	viper.SetConfigName(configFile)
	viper.AddConfigPath("/opt/deploy/")
	viper.AddConfigPath("$HOME/go/src/github.com/golang/somesaturday.com/deploy/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Viper failed: ", err)
	}
	config.BasePath = viper.GetString("BasePath")
	config.HostName = viper.GetString("HostName")
	config.SalesEmail = viper.GetString("SalesEmail")
	config.Verbose = viper.GetBool("Verbose")
}

func main() {
	ParseConfig("config")
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("pages/"))))
	http.Handle("/Images/", http.StripPrefix("/Images/", http.FileServer(http.Dir("Images"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.Handle("/Lightbox/", http.StripPrefix("/Lightbox/", http.FileServer(http.Dir("Lightbox"))))
	if err := http.ListenAndServe(config.HostName, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
