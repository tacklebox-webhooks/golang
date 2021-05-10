package apis

type HttpRequest struct {
	method  string
	baseUrl string
	path    string
	data    map[string]interface{}
	attempt int
}
