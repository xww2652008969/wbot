package sendapi

import (
	"io"
	"net/http"
)

func (sendapi *SendAPI) httpget(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", sendapi.httpurl+path, nil)
	if err != nil {
		return nil, err
	}
	if sendapi.httpAuthorization != "" {
		req.Header.Set("Authorization", sendapi.httpAuthorization)
	}
	res, err := sendapi.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	return res, err
}
func (sendapi *SendAPI) httppost(path string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", sendapi.httpurl+path, payload)
	if err != nil {
		return nil, err
	}
	if sendapi.httpAuthorization != "" {
		req.Header.Set("Authorization", sendapi.httpAuthorization)
	}
	res, err := sendapi.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	return res, err
}
