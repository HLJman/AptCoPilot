package route

import (
	"net/http"

	"github.com/HLJman/AptCoPilot/route/handler"

	goji "goji.io"
	"goji.io/pat"
)

func Register(root *goji.Mux) {
	root.Use(corsHandler)
	subMux := goji.SubMux()
	root.Handle(pat.Get("/api/*"), subMux)
	subMux.HandleFunc(pat.Get("/properties"), handler.Properties)
	subMux.HandleFunc(pat.Get("/properties/:id"), handler.Property)
}

func corsHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "x-pingpong")
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
