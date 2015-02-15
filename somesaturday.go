package main

import (
    "log"
    "net/http"
    "github.com/spf13/viper"
)


var config = struct {
	BasePath             string
	HostName             string
	SalesEmail           string
	Verbose              bool
}{}


//Learning to add viper
func ParseConfig(configFile string) {
	viper.SetConfigFile(configFile)
       viper.AddConfigPath("/deploy")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Viper failed: ", err)
	}
	config.BasePath = viper.GetString("BasePath")
	config.HostName = viper.GetString("HostName")
	config.SalesEmail = viper.GetString("SalesEmail")
	config.Verbose = viper.GetBool("Verbose")
}


func main() {
    ParseConfig("deploy/config.yaml") // context.go
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("pages/"))))
    http.Handle("/Images/", http.StripPrefix("/Images/", http.FileServer(http.Dir("Images"))))
    http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
    http.Handle("/Lightbox/", http.StripPrefix("/Lightbox/", http.FileServer(http.Dir("Lightbox"))))
    if err := http.ListenAndServe(config.HostName, nil); err != nil {
    log.Fatal("ListenAndServe: ", err)
    }
}