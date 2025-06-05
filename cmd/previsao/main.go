package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cidades/{cep}/temperaturas", func(w http.ResponseWriter, r *http.Request) {
		cep := r.PathValue("cep")
		w.Write([]byte("Hello, World! - " + cep))
	})

	http.ListenAndServe(":3000", mux)
}
