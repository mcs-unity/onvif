package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mcs-unity/onvif/api/discovery"
	"github.com/mcs-unity/onvif/api/version"
)

func Load() http.Handler {
	mux := &mux.Router{}
	version.Load(mux)
	discovery.Load(mux)
	return mux
}
