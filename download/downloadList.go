package download

type WaitObject struct {
	Url      string
	FilePath string
}

type DownloadMgr struct {
	WaitList        []WaitObject
	DownloadingList []*DownloadingObject
	Limit           int
}

func (dl *DownloadMgr) AddWaitList(wo WaitObject) error {
	dl.WaitList = append(dl.WaitList, wo)
	return nil
}

func (dl *DownloadMgr) Check() bool {
	if len(dl.DownloadingList) < dl.Limit {
		return true
	} else {
		return false
	}
}

func (dl *DownloadMgr) Update() error {
	if dl.Check() {
		if len(dl.WaitList) > 0 {
			wo := dl.WaitList[0]
			dl.WaitList = dl.WaitList[1:len(dl.WaitList)]
			dlo := &DownloadingObject{}
			dlo.Init(wo.Url, wo.FilePath)
			dl.DownloadingList = append(dl.DownloadingList, dlo)
		}
	}
	for _, dlo := range dl.DownloadingList {
		dlo.Update()
	}
	dl.filterFinish()
	return nil
}

func (dl *DownloadMgr) CheckOver() bool {
	return dl.IsEmpty()
}

func (dl *DownloadMgr) GetWaitListSize() int {
	return len(dl.WaitList)
}

func (dl *DownloadMgr) GetDownloadingListSize() int {
	return len(dl.DownloadingList)
}

func (dl *DownloadMgr) IsEmpty() bool {
	return dl.GetWaitListSize() == 0 && dl.GetDownloadingListSize() == 0
}

func (dl *DownloadMgr) filterFinish() error {
	var k_dl []int
	for k, v := range dl.DownloadingList {
		if v.Status == FINISHED {
			k_dl = append(k_dl, k)
		}
	}
	if len(k_dl) == 0 {
		return nil
	}

	for i, j := 0, len(k_dl)-1; i < j; i, j = i+1, j-1 {
		k_dl[i], k_dl[j] = k_dl[j], k_dl[i]
	}

	for _, k_d := range k_dl {
		dl.DownloadingList = append(dl.DownloadingList[:k_d], dl.DownloadingList[k_d+1:]...)
	}

	return nil
}
