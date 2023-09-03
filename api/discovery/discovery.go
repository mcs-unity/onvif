package discovery

import (
	"net/http"

	"github.com/gorilla/mux"
)

func discover(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("awesome"))
}

func Load(h *mux.Router) {
	h.HandleFunc("/discover", discover).Methods(http.MethodGet)
}
