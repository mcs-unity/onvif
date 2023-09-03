package version

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func version(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(os.Getenv("VERSION")))
}

func Load(h *mux.Router) {
	h.HandleFunc("/", version).Methods(http.MethodGet)
}
