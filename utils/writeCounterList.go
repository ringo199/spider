package utils

import (
	"log"
)

type WriteCounterList struct {
	List []*WriteCounter
}

func (wcl *WriteCounterList) SetWriteCounter(wc *WriteCounter) error {
	wcl.List = append(wcl.List, wc)
	return nil
}

func (wcl *WriteCounterList) IsEmpty() bool {
	return len(wcl.List) == 0
}

func (wcl *WriteCounterList) FilterWc() error {
	var k_wcl []int
	for k, v := range wcl.List {
		if v.IsFinish {
			k_wcl = append(k_wcl, k)
		}
	}
	if len(k_wcl) == 0 {
		return nil
	}

	for i, j := 0, len(k_wcl)-1; i < j; i, j = i+1, j-1 {
		k_wcl[i], k_wcl[j] = k_wcl[j], k_wcl[i]
	}

	for _, k_wc := range k_wcl {
		wcl.List = append(wcl.List[:k_wc], wcl.List[k_wc+1:]...)
	}

	return nil
}

type WaitObject struct {
	Url      string
	FilePath string
}

type DownloadList struct {
	WaitList []WaitObject
	Limit    int
	Wcl      *WriteCounterList
}

func (dl *DownloadList) AddWaitList(wo WaitObject) error {
	dl.WaitList = append(dl.WaitList, wo)
	return nil
}

func (dl *DownloadList) Check() bool {
	if len(dl.Wcl.List) < dl.Limit {
		return true
	} else {
		return false
	}
}

func (dl *DownloadList) IsEmpty() bool {
	return len(dl.WaitList) == 0
}

func (dl *DownloadList) StartDownload() error {
	if dl.Check() {
		if len(dl.WaitList) > 0 {
			wo := dl.WaitList[0]
			dl.WaitList = dl.WaitList[1:len(dl.WaitList)]

			go func() {
				err := Download(wo.Url, wo.FilePath, dl.Wcl)
				if err != nil {
					log.Fatal(err)
				}
			}()
		}
	}
	return nil
}

func (dl *DownloadList) CheckOver() bool {
	if dl.Wcl.IsEmpty() && dl.IsEmpty() {
		return true
	}
	return false
}
