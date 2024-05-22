package response

import "net/http"

type response struct {
	writer     http.ResponseWriter
	statusCode int
	body       []byte
}

func New(writer http.ResponseWriter) *response {
	return &response{
		writer: writer,
	}
}

func NewInternalError(writer http.ResponseWriter) {
	r := &response{
		writer: writer,
	}

	r.writer.WriteHeader(r.statusCode)
}

func (r *response) SetStatusCode(statusCode int) *response {
	r.statusCode = statusCode

	return r
}

func (r *response) SetBody(data []byte) *response {
	r.body = data

	return r
}

func (r *response) Build() {
	r.writer.WriteHeader(r.statusCode)
	r.writer.Write(r.body)
}
