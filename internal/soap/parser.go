package soap

import "encoding/xml"

func writeXml(data any) ([]byte, error) {
	b, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ParseXml(data []byte, pointer any) error {
	return xml.Unmarshal(data, pointer)
}

func SoapBody(SoapAction string) IXmlTag {
	return &Envelop{
		Xmlns:  "http://www.w3.org/2003/05/soap-envelope",
		Header: header{Action: action{Action: SoapAction}},
		Body:   body{},
	}
}
