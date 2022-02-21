package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/ringo199/spider/constant"
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

func CreateTmpDir() error {
	dir := constant.TmpBasePath
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, os.FileMode(0755))
	}
	return nil
}

func CreateDir(path string) error {
	dir := filepath.Dir(path)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, os.FileMode(0755))
	}
	return nil
}

func RandomFilename16Char() (s string, err error) {
	b := make([]byte, 8)
	_, err = rand.Read(b)
	if err != nil {
		return
	}
	s = fmt.Sprintf("%x", b)
	return
}

func GetMd5(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func GetSha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func ParseDate(d string) (string, error) {
	ds := strings.Split(d, "-")
	year, err := strconv.Atoi(ds[0])
	if err != nil {
		return "", err
	}
	year += 2000
	pre_month, err := strconv.Atoi(ds[1])
	if err != nil {
		return "", err
	}

	month := time.Month(pre_month)
	date, err := strconv.Atoi(ds[2])
	if err != nil {
		return "", err
	}
	dt := time.Date(year, month, date, 0, 0, 0, 0, time.Local)
	return dt.Format("2006.01.02"), nil
}

func RunCmd(cmd_file string, cmd_args []string) (string, error) {
	cmd := exec.Command(cmd_file, cmd_args...)
	fmt.Println(cmd.String())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	defer stdout.Close()
	if err = cmd.Start(); err != nil {
		return "", err
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil {
		return "", err
	} else {
		return string(opBytes), nil
	}
}

func OpenFile(path string) (*os.File, error) {
	err := CreateDir(path)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	err = CreateDir(dstName)
	if err != nil {
		return
	}
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
