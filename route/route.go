package route

import (
	"fmt"
	"net/http"

	"github.com/HLJman/AptCoPilot/route/handler"

	goji "goji.io"
	"goji.io/pat"
)

func Register(root *goji.Mux) {
	root.Use(corsHandler)

	// api
	subMux := goji.SubMux()
	root.Handle(pat.Get("/api/*"), subMux)
	subMux.HandleFunc(pat.Get("/properties"), handler.Properties)
	subMux.HandleFunc(pat.Get("/properties/:id"), handler.Property)

	// UI
	root.Handle(pat.New("/src"), http.FileServer(http.Dir("./assets/src")))
	root.HandleFunc(pat.New("/properties"), func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/index.html")
	})
	root.HandleFunc(pat.New("/contact"), func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets/index.html")
	})
	root.Handle(pat.New("/*"), http.FileServer(http.Dir("./assets")))
}

func corsHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
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
