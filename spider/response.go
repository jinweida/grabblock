package spider

type Response struct {
	StatusCode int
	Body       []byte
	Request    *Request
}
