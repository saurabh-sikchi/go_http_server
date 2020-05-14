package main

import (
	"fmt"
	"net/http"
	"time"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("online"))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().String()
	w.Write([]byte(t))
}

func main() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("https://in.linkedin.com/in/saurabh-sikchi-50b807b8", http.StatusPermanentRedirect)
	mux.Handle("/linkedin", rh)

	mux.HandleFunc("/status", statusHandler)

	mux.HandleFunc("/timestamp", timeHandler)

	fmt.Println("Listening and serving on port 3000:")
	http.ListenAndServe(":3000", mux)
}
