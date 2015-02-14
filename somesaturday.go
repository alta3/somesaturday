package main

import (
    "log"
    "net/http"
)

func main() {
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("pages/"))))
    http.Handle("/Images/", http.StripPrefix("/Images/", http.FileServer(http.Dir("Images"))))
    http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
    http.Handle("/Lightbox/", http.StripPrefix("/Lightbox/", http.FileServer(http.Dir("Lightbox"))))
    if err := http.ListenAndServe("192.168.241.22:80", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}