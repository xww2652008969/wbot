package utils

import (
	"io"
	"net/http"
)

type Request struct {
}

func Httpget(url string, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return res, err
	}
	return res, err
}
func Httppost(url string, header map[string]string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	//out, err := isgzip(res)
	if err != nil {
		return res, err
	}
	return res, err
}

//func isgzip(res *http.Response) (*http.Response, error) {
//	gzipFlag := false
//	for k, v := range res.Header {
//		if strings.ToLower(k) == "content-encoding" && strings.ToLower(v[0]) == "gzip" {
//			gzipFlag = true
//		}
//	}
//	if gzipFlag {
//		gr, err := gzip.NewReader(res.Body)
//		if err != nil {
//			return nil, errors.New("解析有误")
//		}
//		res.Body = gr
//		return res, nil
//	}
//	return res, nil
//}
