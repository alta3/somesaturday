package main

import (
	"github.com/spf13/viper"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

type Config struct {
	BasePath   string
	HostName   string
}

var config Config


func Log(handler http.Handler) http.Handler {
     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
     log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
     handler.ServeHTTP(w, r)
   })
}



func ParseConfig(configFile string) Config{
	var config Config
	viper.SetConfigName(configFile)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./deploy/")
	viper.AddConfigPath("$HOME/git/go/somesaturday/deploy/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Viper failed: ", err)
	}
	config.BasePath = viper.GetString("BasePath")
	config.HostName = viper.GetString("HostName")
        return config
}

func errorHandler(status int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		if status == http.StatusNotFound {
			r.URL.Path = "404.html"
			Template().ServeHTTP(w, r)
		}
	})
}

func Template() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {


	if r.URL.Path == "" {
			http.ServeFile(w, r, config.BasePath + "deploy/pages/index.html")
			return
		} else if r.URL.Path == "robots.txt" {
			http.ServeFile(w, r,  config.BasePath + "deploy/robots.txt")
			return
		} else if r.URL.Path == "sitemap.xml" {
			http.ServeFile(w, r, config.BasePath + "deploy/sitemap.xml")
			return
		} else if r.URL.Path == "favicon.ico" {
			http.ServeFile(w, r, config.BasePath + "deploy/Images/favicon.ico")
			return
		}


		lp := path.Join( config.BasePath + "deploy/templates", "layout.html")
		fp := path.Join( config.BasePath + "deploy/templates", r.URL.Path)

                log.Println(lp + " <--- and ---> " + fp)


		info, err := os.Stat(fp)
		if err != nil {
			if os.IsNotExist(err) {
				log.Printf("The 404 Barn file is: %s", r.URL.Path)
				errorHandler(http.StatusNotFound).ServeHTTP(w, r)
				return
			}
		}
		if info.IsDir() {
			http.NotFound(w, r)
			return
		}
		tmpl, err := template.ParseFiles(lp, fp)
		if err != nil {
			log.Printf("Template Error: %s", err)
		}
		err = tmpl.ExecuteTemplate(w, "layout", nil)
		if err != nil {
			log.Printf("Template Error: %s", err)
		}
		return
	})
}

func main() {
	config := ParseConfig("config")

        log.Println("config file loaded")
        log.Println("BasePath: " + config.BasePath + " (The location of the deploy folder)") 
        log.Println("HostName: " + config.HostName + " (The ip address and port of this server)" )


	//templates
	http.Handle("/", http.StripPrefix("/", Template()))


	// static folders
	http.Handle("/Images/", http.StripPrefix("/Images/", http.FileServer(http.Dir( config.BasePath + "deploy/Images"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir( config.BasePath + "deploy/styles"))))
	http.Handle("/Lightbox/", http.StripPrefix("/Lightbox/", http.FileServer(http.Dir( config.BasePath + "deploy/Lightbox"))))
	if err := http.ListenAndServe(config.HostName, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
