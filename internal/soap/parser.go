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
