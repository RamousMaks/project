package main

import "net/http"

func main() {
	http.HandleFunc("/project", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("String"))
	})

	http.ListenAndServe(":9000", nil)
}
