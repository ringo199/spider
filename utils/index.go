package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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

func Download(apiUrl string, filePath string, wcl *WriteCounterList) error {

	filePath, err := url.QueryUnescape(filePath)
	if err != nil {
		return err
	}
	counter := &WriteCounter{
		FilePath: filePath,
	}
	wcl.SetWriteCounter(counter)
	header := map[string]string{
		"referrer":   "https://asoul-rec.herokuapp.com/",
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36",
	}
	resp, err := GetRequest(apiUrl, nil, &header)
	if err != nil {
		return err
	}

	dir := filepath.Dir(filePath)
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, os.FileMode(0755))
	}
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	counter.AllTotal = uint64(resp.ContentLength)
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}
	counter.Finish()
	defer resp.Body.Close()
	defer out.Close()

	return nil
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
