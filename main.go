package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/dump", dumpRequest)

	log.Println("Server is listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func dumpRequest(w http.ResponseWriter, r *http.Request) {
	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
	log.Printf("%v\n", string(requestDump))
	fmt.Fprintf(w, "%v", string(requestDump))
}
