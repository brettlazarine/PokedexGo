package api

import (
	"io"
	"net/http"
)

type Api struct{}

func (a *Api) Get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	

	return body, nil
}
