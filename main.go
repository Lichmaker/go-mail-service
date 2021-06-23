package main

import (
	"io"
	"log"
	"mail-service/pkg/sender"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		email := r.PostFormValue("email")
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")

		json := sender.SendSingleMail(email, title, body)
		io.WriteString(rw, json)
	}).Methods("POST")
	log.Fatal(http.ListenAndServe(":3001", router))

}
