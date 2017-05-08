package main

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusGone)
	log.Printf("Referer: %s\n", r.Header.Get("Referer"))
}

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   os.Args[1],
		MaxSize:    250, // megabytes
		MaxBackups: 3,
		MaxAge:     1, //days
	})
	fmt.Printf("Starting http-sinkhole at 8080 port")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
