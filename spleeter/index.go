package spleeter

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ringo199/spider/utils"
)

func Spleeter() {
	outPath := "outputPath/spleeter/inputPath"
	inputPath := "outputPath/audios/bilibili/diana/output/"

	files, err := utils.ReadDir(inputPath)
	if err != nil {
		panic(err)
	}

	const COUNT = 4
	var tmpStr string
	var tmpIndex int
	var tmpCount int
	var tmpMap map[int][]string = make(map[int][]string)

	for k, v := range files {
		fileName := strings.Split(v.Name(), "_")[0]
		if k == 0 {
			tmpStr = fileName
		}

		if tmpStr != fileName {
			tmpStr = fileName
			tmpCount++
			if tmpCount > COUNT {
				tmpCount = 0
				tmpIndex++
			}
		}

		if tmpMap[tmpIndex] == nil {
			tmpMap[tmpIndex] = []string{v.Name()}
		} else {
			tmpMap[tmpIndex] = append(tmpMap[tmpIndex], v.Name())
		}
	}

	for k, l := range tmpMap {
		f, err := utils.OpenFile(filepath.Join(outPath, fmt.Sprintf("%d", k)))
		if err != nil {
			panic(err)
		}
		err = f.Truncate(0)
		if err != nil {
			panic(err)
		}
		f.Seek(0, 0)
		defer f.Close()

		for _, v := range l {
			f.WriteString(v + "\n")
		}
	}
}
