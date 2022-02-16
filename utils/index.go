package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func GetRequest(apiUrl string, params *map[string]string, header *map[string]string) (*http.Response, error) {
	client := &http.Client{}
	data := url.Values{}
	if params != nil {
		for k, v := range *params {
			data.Set(k, v)
		}
	}
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = data.Encode() // URL encode
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if header != nil {
		for k, v := range *header {
			req.Header.Add(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ReadAll(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

func ReadDir(path string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func CreateFile(paths []string, basePath string) error {
	_, err := os.Stat(basePath)
	if os.IsNotExist(err) {
		os.MkdirAll(basePath, os.FileMode(0755))
	}
	for _, path := range paths {
		_, err := os.Stat(basePath + path)
		if os.IsNotExist(err) {
			os.Create(basePath + path)
		}
	}
	return nil
}
