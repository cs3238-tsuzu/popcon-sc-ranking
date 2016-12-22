package main

import (
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("POPCON_SC_RANKING_TOKEN")
	addr := os.Getenv("POPCON_SC_RANKING_ADDR")
	db := os.Getenv("POPCON_SC_RANKING_DB")

	mux := http.NewServeMux()

	mux.Handle("/contest/", http.StripPrefix("/contest", func() http.HandlerFunc {
		mux := http.NewServeMux()

		mux.HandleFunc("/create", func(rw http.ResponseWriter, req *http.Request) {
			cid, err := parseFormInt64(req, "cid")

			if err != nil {
				resprondJSON(rw, "Unable to get 'cid'", nil)

				return
			}

		})

		return mux.ServeHTTP
	}()),
	)
}
