package serv

import (
	/* Standard library packages */

	/* Third party */
	// imports as "cli", pinned to v1; cliv2 is going to be drastically
	// different and pinning to v1 avoids issues with unstable API changes
	"encoding/json"
	"log"
	"net/http"

	cli "gopkg.in/urfave/cli.v1"

	/* Local packages */
	"github.com/gorilla/mux"
	"github.com/keeferrourke/imgrep/storage"
)

var PORT string

func StartServer(c *cli.Context) {
	r := mux.NewRouter()
	r.HandleFunc("/imgrep/search", func(w http.ResponseWriter, r *http.Request) {
		keyword := r.PostFormValue("keyword")
		filenames, err := storage.Get(keyword)
		if err != nil {
			log.Fatalf(err.Error())
		}

		resp := map[string][]string{}
		resp["filenames"] = filenames

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":"+PORT, r)
}
