package service

import (
	"net/http"
	"log"
)

func StartWebServer(port string) {
	log.Println("Starting HTTP service at" + port)

	r := NewRouter()
	http.Handle("/", r)

	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		log.Println("An error occured starting HTTPlistener at port", port)
		log.Println("Error: " + err.Error())
	}
}