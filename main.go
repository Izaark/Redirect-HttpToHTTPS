package main

import (
	"github.com/tasks/Redirect-HttpToHTTPS/config"
	"log"
	"net/http"
	"os"
)

//Init: Initialize before the main function for loading environment file
func init() {
	err := config.FunInitEnvironment()
	if err != nil {
		log.Fatal("ERROR init: couldn't initialize environment-> ", err.Error())
	}
}

//handlerPage: HomePage in default Port 80
func handlerPage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)
}

//handlerRedirectToHttps: redirect to https in specific url example: https: https://github.com
func handlerRedirectToHttps(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, os.Getenv("REDIRECT_HTTPS_URL")+r.RequestURI, http.StatusMovedPermanently)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerPage)

	//Goroutine for listen Transport Layer Security https Port
	go http.ListenAndServeTLS(":"+os.Getenv("HTTPS_ENV_API_PORT"), "cert.pem", "key.pem", nil)
	http.ListenAndServe(":"+os.Getenv("HTTP_ENV_API_PORT"), http.HandlerFunc(handlerRedirectToHttps))
}
