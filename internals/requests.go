package internals

import (
	log2 "DatabaseHandler/internals/log"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetRequest(api *Api) *[]byte {
	log2.INFO("Get request started")
	req := create("GET", api)

	headers(req, api.Headers)

	res := call(req)

	body := read(res)

	return body
}

func PostRequest() {

}

func create(t string, a *Api) *http.Request {
	req, err := http.NewRequest(t, a.Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	log2.DEBUG("Request created", "request", req)
	return req

}

func headers(r *http.Request, m *map[string]string) {
	var logHeaders []string
	for k, v := range *m {
		r.Header.Add(k, v)
		if !strings.Contains(strings.ToLower(k), "key") {
			logHeaders = append(logHeaders, k)
		}
	}
	log2.DEBUG("Headers Added", "headers", logHeaders)
}

func call(r *http.Request) *http.Response {
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	log2.DEBUG("Client Called")
	//defer res.Body.Close()
	return res
}

func read(r *http.Response) *[]byte {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log2.DEBUG("Response Readed to *[]byte")
	defer r.Body.Close()
	return &body
}
