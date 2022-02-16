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
	CurSpeed uint64
	Total    uint64
	AllTotal uint64
	Percent  float64

	CurSpeedFormatData string
	FormatData         string
	AllFormatData      string
}

func (wc *WriteCounter) getFormatData(n uint64, str *string) error {
	*str = ""
	if n>>GB > 0 {
		*str += fmt.Sprintf("%dG", n>>GB)
		n = (1<<GB - 1) & n
	}
	if n>>MB > 0 {
		*str += fmt.Sprintf("%dM", n>>MB)
		n = (1<<MB - 1) & n
	}
	if n>>KB > 0 {
		*str += fmt.Sprintf("%dK", n>>KB)
		n = (1<<KB - 1) & n
	}
	if n>>B > 0 {
		*str += fmt.Sprintf("%dB ", n>>B)
	}

	return nil
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.CurSpeed = uint64(n)
	wc.Total += wc.CurSpeed

	wc.getFormatData(wc.CurSpeed, &wc.CurSpeedFormatData)
	wc.getFormatData(wc.Total, &wc.FormatData)
	wc.getFormatData(wc.AllTotal, &wc.AllFormatData)
	wc.Percent = float64(wc.Total*100/wc.AllTotal) / 100
	return n, nil
}
