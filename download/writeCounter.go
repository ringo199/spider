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
	FilePath      string
	LastTransSize uint64
	Total         uint64
	AllTotal      uint64
	Percent       float64

	LastTransSizeFormatData string
	FormatData              string
	AllFormatData           string
}

func (wc *WriteCounter) GetFormatData(n uint64, str *string) error {
	*str = ""
	if n>>GB > 0 {
		*str += fmt.Sprintf("%dg ", n>>GB)
		n = (1<<GB - 1) & n
	}
	if n>>MB > 0 {
		*str += fmt.Sprintf("%dm ", n>>MB)
		n = (1<<MB - 1) & n
	}
	if n>>KB > 0 {
		*str += fmt.Sprintf("%dk ", n>>KB)
		n = (1<<KB - 1) & n
	}
	if n>>B > 0 {
		*str += fmt.Sprintf("%db", n>>B)
	}

	return nil
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.LastTransSize = uint64(n)
	wc.Total += wc.LastTransSize

	wc.GetFormatData(wc.LastTransSize, &wc.LastTransSizeFormatData)
	wc.GetFormatData(wc.Total, &wc.FormatData)
	if wc.AllTotal != 0 {
		wc.Percent = float64(wc.Total*100/wc.AllTotal) / 100
	} else {
		if wc.Total != 0 {
			wc.Percent = 100
		} else {
			wc.Percent = 0
		}
	}
	return n, nil
}
