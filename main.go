package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func mainRouteHandler(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	if r.Method == "" {
	}
	io.WriteString(w, "pong")
	fmt.Println("sended abiba")
	fmt.Println(r.Method)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", mainRouteHandler)

	err := http.ListenAndServe(":8080", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("serser closed")
	} else if err != nil {
		fmt.Println("error ocurred starting server ", err)
		os.Exit(1)
	}

	/* r := mux.NewRoutHandleFunc
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var a string
		json.NewDecoder(r.Body).Decode(&a)
		fmt.Println(a)
		if a == "ping" {
			json.NewEncoder(w).Encode("pong")
			return
		}
		json.NewEncoder(w).Encode("")

	}).Methods("GET") */

	//log.Fatal(http.ListenAndServe(":8000", r))
}
