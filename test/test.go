package test

import "github.com/ringo199/spider/filter"

func Test() {
	list := filter.AsDBInfoList{}

	list.ParseJson()

	for _, staff := range filter.AsStaffLetter {
		pathList := list.GetOnlyStaffTitleList(staff)
		if err := list.MoveOnlyStaffFile(pathList, staff); err != nil {
			panic(err)
		}
	}
}
