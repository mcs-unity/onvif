package soap

import "encoding/xml"

func WriteXml(data any) ([]byte, error) {
	b, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ParseXml(data []byte, pointer any) error {
	return xml.Unmarshal(data, pointer)
}

func New(SoapAction string) IEnvelop {
	return &Envelop{
		Xmlns:  "http://www.w3.org/2003/05/soap-envelope",
		Header: header{Action: action{Action: SoapAction}},
		Body:   body{},
	}
}
