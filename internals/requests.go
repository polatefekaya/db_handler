package internals

import (
	"io"
	"log"
	"net/http"
)

func GetRequest(api *Api) *[]byte {

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
	return req

}

func headers(r *http.Request, m *map[string]string) {
	for k, v := range *m {
		r.Header.Add(k, v)
	}
}

func call(r *http.Request) *http.Response {
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	//defer res.Body.Close()
	return res
}

func read(r *http.Response) *[]byte {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	return &body
}
