package download

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/ringo199/spider/constant"
	"github.com/ringo199/spider/utils"
)

func DownloadFn(paths []string, fileType string, nameType string, dl *utils.DownloadList) error {
	// fmt.Println("开始读取下载列表...")
	for _, inputPath := range paths {
		// fmt.Printf("开始读取%s的下载列表...\n", inputPath)
		body, err := utils.ReadAll(
			constant.InputBasePath +
				fileType +
				nameType +
				inputPath)
		if err != nil {
			return err
		}
		// fmt.Println("下载列表读取完成")
		allPaths := strings.Split(string(body), "\n")
		var tmpPaths []string
		for _, path := range allPaths {
			if path != "" {
				fileName := filepath.Base(path)
				parseFileName, err := url.QueryUnescape(fileName)
				if err != nil {
					return err
				}
				fileInfo, _ := os.Stat(constant.OutputBasePath + fileType + nameType + inputPath + "/" + parseFileName)
				if fileInfo == nil {
					tmpPaths = append(tmpPaths, path)
				}
			}
		}
		if len(tmpPaths) <= 0 {
			// fmt.Printf("%s没有文件需要下载\n", inputPath)
		} else {
			// fmt.Printf("%s加入了%d个文件进入下载队列\n", inputPath, len(tmpPaths))
			for _, path := range tmpPaths {
				filePath := constant.OutputBasePath + fileType + nameType + inputPath + "/" + filepath.Base(path)

				err = dl.AddWaitList(utils.WaitObject{
					Url:      path,
					FilePath: filePath,
				})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
