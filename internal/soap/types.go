package soap

type envelop struct {
	header string `xml:"header"`
	body   string `xml:"body"`
}
