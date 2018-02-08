package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetjsonDetails(w http.ResponseWriter, r *http.Request) {

	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(file)

}

func main() {
	var r = mux.NewRouter()
	fs := http.FileServer(http.Dir("./block/"))
	//http.Handle("/", fs)
	r.PathPrefix("/block/").Handler(http.StripPrefix("/block/", fs))

	r.HandleFunc("/Getjson", GetjsonDetails)

	server := &http.Server{
		Addr:    ":8088",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}
