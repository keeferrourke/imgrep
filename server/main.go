package serv

import (
	/* Standard library packages */

	/* Third party */
	// imports as "cli", pinned to v1; cliv2 is going to be drastically
	// different and pinning to v1 avoids issues with unstable API changes
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	cli "gopkg.in/urfave/cli.v1"

	/* Local packages */
	"github.com/gorilla/mux"
	"github.com/keeferrourke/imgrep/storage"
)

var PORT string

type ResultRow struct {
	Filename string `json:"filename"`
	Bytes    []byte `json:"bytes"`
}

func StartServer(c *cli.Context) {
	r := mux.NewRouter()
	r.HandleFunc("/imgrep/search", func(w http.ResponseWriter, r *http.Request) {
		keyword := r.FormValue("keyword")
		filenames, err := storage.Get(keyword)
		if err != nil {
			log.Fatalf(err.Error())
		}

		results := []*ResultRow{}
		for _, file := range filenames {
			f, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatalf(err.Error())
			}

			results = append(results, &ResultRow{
				Filename: file,
				Bytes:    f,
			})
		}
		resp := map[string][]*ResultRow{}
		resp["files"] = results

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates, err := template.ParseFiles("./index.html")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	s := http.StripPrefix("/assets/", http.FileServer(http.Dir(os.Getenv("GOPATH")+"/src/github.com/keeferrourke/imgrep/assets")))
	log.Println(os.Getenv("GOPATH") + "/src/github.com/keeferrourke/imgrep/assets")
	r.PathPrefix("/assets/").Handler(s)

	http.ListenAndServe(":"+PORT, r)
}
