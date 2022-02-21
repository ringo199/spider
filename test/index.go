package test

import (
	"github.com/ringo199/spider/filter"
	"github.com/ringo199/spider/mockingbird"
	"github.com/ringo199/spider/spleeter"
)

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

func TestSpleeter() {
	spleeter.Spleeter()
}

func TestPreGenerate() {
	mockingbird.PreGenerate()
}
