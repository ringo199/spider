package download

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"

	"github.com/ringo199/spider/constant"
	"github.com/ringo199/spider/utils"
)

const (
	PREPARE     = "PREPARE"
	WAITING     = "WAITING"
	DOWNLOADING = "DOWNLOADING"
	FINISH      = "FINISH"
	FINISHED    = "FINISHED"
)

type DownloadingObject struct {
	Url      string
	FilePath string
	Wc       *WriteCounter
	TmpName  string
	Status   string
}

func (dlo *DownloadingObject) Init(
	Url string,
	FilePath string,
) {
	dlo.Url = Url
	dlo.Status = PREPARE

	var err error
	dlo.FilePath, err = url.PathUnescape(FilePath)
	if err != nil {
		dlo.FilePath = FilePath
	}
}

func (dlo *DownloadingObject) download() error {
	tmpPath := constant.TmpBasePath + dlo.TmpName
	dlo.Wc.FilePath = tmpPath
	referrer, err := filepath.Abs(dlo.Url)
	if err != nil {
		return err
	}
	header := map[string]string{
		"referrer":   referrer,
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36",
	}
	resp, err := utils.GetRequest(dlo.Url, nil, &header)
	if err != nil {
		return err
	}

	err = utils.CreateTmpDir()
	if err != nil {
		return err
	}
	out, err := os.Create(tmpPath)
	if err != nil {
		return err
	}
	dlo.Wc.AllTotal = uint64(resp.ContentLength)
	dlo.Status = DOWNLOADING
	_, err = io.Copy(out, io.TeeReader(resp.Body, dlo.Wc))
	if err != nil {
		return err
	}
	dlo.Status = FINISH
	defer resp.Body.Close()
	defer out.Close()

	return nil
}

func (dlo *DownloadingObject) startDownload() error {
	var err error
	dlo.TmpName, err = utils.RandomFilename16Char()
	if err != nil {
		return err
	}
	dlo.Wc = &WriteCounter{}
	dlo.Status = WAITING
	go func() {
		err = dlo.download()
		if err != nil {
			fmt.Println(err)
		}
	}()
	return nil
}

func (dlo *DownloadingObject) downloadFinish() error {
	err := utils.CreateDir(dlo.FilePath)
	if err != nil {
		return err
	}
	os.Rename(constant.TmpBasePath+dlo.TmpName, dlo.FilePath)

	dlo.Status = FINISHED
	return nil
}

func (dlo *DownloadingObject) Update() {
	switch dlo.Status {
	case PREPARE:
		dlo.startDownload()
	case WAITING:
	case DOWNLOADING:
	case FINISH:
		dlo.downloadFinish()
	case FINISHED:
	}
}
