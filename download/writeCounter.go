package download

import (
	"fmt"
)

const (
	B uint64 = iota * 10
	KB
	MB
	GB
)

type WriteCounter struct {
	FilePath string
	Total    uint64
	AllTotal uint64
	Percent  float64

	FormatData    string
	AllFormatData string
	IsFinish      bool
}

func (wc *WriteCounter) getFormatData(n uint64, str *string) error {
	*str = ""
	if n>>GB > 0 {
		*str += fmt.Sprintf("%dG ", n>>GB)
		n = (1<<GB - 1) & n
	}
	if n>>MB > 0 {
		*str += fmt.Sprintf("%dM ", n>>MB)
		n = (1<<MB - 1) & n
	}
	if n>>KB > 0 {
		*str += fmt.Sprintf("%dK ", n>>KB)
		n = (1<<KB - 1) & n
	}
	if n>>B > 0 {
		*str += fmt.Sprintf("%dB ", n>>B)
	}

	return nil
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.getFormatData(wc.Total, &wc.FormatData)
	wc.getFormatData(wc.AllTotal, &wc.AllFormatData)
	wc.Percent = float64(wc.Total*100/wc.AllTotal) / 100
	return n, nil
}

func (wc *WriteCounter) Finish() {
	wc.IsFinish = true
}
