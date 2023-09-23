package discovery

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcs-unity/onvif/pkg/onvif"
)

func discover(w http.ResponseWriter, r *http.Request) {
	w.Write(onvif.Probe())
}

func Load(h *mux.Router) {
	h.HandleFunc("/discover", discover).Methods(http.MethodGet)
}
