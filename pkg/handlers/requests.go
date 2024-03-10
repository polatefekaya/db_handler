package handlers

import (
	"io"
	"log"
	"net/http"
)

func GetRequest(url string) *[]byte {
	req := create("GET", url)

	headers(req, map[string]string{
		"X-Header": "value",
	})

	res := call(req)

	body := read(res)

	return body
}

func PostRequest() {

}

func create(t string, u string) *http.Request {
	req, err := http.NewRequest(t, u, nil)
	if err != nil {
		log.Fatal(err)
	}
	return req
}

func headers(r *http.Request, m map[string]string) {
	for k, v := range m {
		r.Header.Add(k, v)
	}
}

func call(r *http.Request) *http.Response {
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	return res
}

func read(r *http.Response) *[]byte {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	return &body
}
