// go start
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/m/v2/usecases"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		comics, err := usecases.NewBook("comics")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// comics.Print()

		response, err := json.Marshal(comics)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)

	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
