package test

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/ringo199/spider/constant"
	"github.com/ringo199/spider/utils"
)

func CleanErrFile() {
	baseDir := constant.OutputBasePath +
		constant.AudioBasePath +
		constant.BilibiliPath +
		"asoul/"
	files, err := utils.ReadDir(baseDir)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("logs/deleteAudios.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, path := range files {
		file, err := os.Stat(baseDir + path.Name())
		if err != nil {
			panic(err)
		}
		u := url.QueryEscape(path.Name())
		u = strings.ReplaceAll(u, "+", "%20")
		apiurl := "https://asoul-rec.herokuapp.com/ASOUL-REC/AAC%E5%BD%95%E6%92%AD%E9%9F%B3%E8%BD%A8/" + u

		header := map[string]string{
			"referrer":   "https://asoul-rec.herokuapp.com",
			"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36",
		}
		resp, err := utils.GetRequest(apiurl, nil, &header)
		if err != nil {
			panic(err)
		}

		off := file.Size() - resp.ContentLength
		str := fmt.Sprintf("%s\n%d,%d off: %d\n",
			path.Name(), file.Size(), resp.ContentLength, off)
		fmt.Printf("%s", str)

		defer resp.Body.Close()

		f.WriteString(str)
		if off != 0 {
			os.Remove(baseDir + path.Name())
		} else {
			os.Rename(baseDir+path.Name(), filepath.Join(baseDir, "../", "tmp", path.Name()))
		}
	}
}
