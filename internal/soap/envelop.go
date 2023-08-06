package soap

func (e *Envelop) AppendToHeader(b ...any) {
	e.Header.Body = append(e.Body.Body, b)
}

func (e *Envelop) AppendToBody(b ...any) {
	e.Header.Body = append(e.Body.Body, b)
}

func (e Envelop) ToString() ([]byte, error) {
	return WriteXml(e)
}
