package onvif

import (
	"encoding/json"

	"github.com/mcs-unity/onvif/pkg/wsdiscovery"
)

func Probe() []byte {
	res, err := wsdiscovery.SendMulticastProbe([]byte(""))
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return []byte(b)
}
