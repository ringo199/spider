package filter

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/ringo199/spider/constant"
	"github.com/ringo199/spider/utils"
)

type AsDBInfo struct {
	Date  string   `json:date`
	Title string   `json:title`
	Staff []string `json:staff`
}

type AsDBInfoList struct {
	List []AsDBInfo
}

var AsStaffLetter = []string{"A", "B", "C", "D", "E"}

var asNameMap map[string]string = map[string]string{
	"A": "向晚",
	"B": "贝拉",
	"C": "珈乐",
	"D": "嘉然",
	"E": "乃琳",
}

var asFullMap map[string]string = map[string]string{
	"A": "ava",
	"B": "bella",
	"C": "carol",
	"D": "diana",
	"E": "eileen",
}

func (list *AsDBInfoList) ParseJson() {
	body1, err := utils.ReadAll("outputPath/asdb/2021.json")
	if err != nil {
		panic(err)
	}
	body2, err := utils.ReadAll("outputPath/asdb/2022.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body1, &list.List)

	if err != nil {
		panic(err)
	}

	list2 := &AsDBInfoList{}

	err = json.Unmarshal(body2, &list2.List)

	if err != nil {
		panic(err)
	}

	list.List = append(list.List, list2.List...)
}

func (list *AsDBInfoList) GetOnlyStaffTitleList(staff string) []string {
	var pathList []string
	path := constant.OutputBasePath +
		constant.SubtitleBasePath +
		constant.BilibiliPath +
		"asoul/"
	audios, _ := utils.ReadDir(path)
	for _, info := range list.List {
		if len(info.Staff) == 1 {
			if info.Staff[0] == staff {
				for _, audio := range audios {
					date, _ := utils.ParseDate(info.Date)
					if strings.Contains(audio.Name(), date) {
						if strings.Contains(audio.Name(), asNameMap[staff]) {
							pathList = append(pathList, path+audio.Name())
						}
					}
				}
			}
		}
	}

	return pathList
}

func (list *AsDBInfoList) MoveOnlyStaffFile(pathList []string, staff string) error {
	for _, path := range pathList {
		name := filepath.Base(path)
		newPath := filepath.Join(path, "../../", asFullMap[staff], name)
		utils.CreateDir(newPath)
		os.Rename(path, newPath)
	}
	return nil
}
