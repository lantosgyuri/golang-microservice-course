package app

import (
	"net/http"
)

// StartApp start sets up the server
func StartApp() {
	http.HandleFunc("/", func(response http.ResponseWriter, req *http.Request) {
		response.Write([]byte("IT WORKS"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
