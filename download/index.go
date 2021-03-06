package download

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/ringo199/spider/constant"
	"github.com/ringo199/spider/session"
	"github.com/ringo199/spider/utils"
)

var space struct{}
var sessionSet map[string]struct{} = make(map[string]struct{})

func InitFileAndWaitDownload(dl *DownloadMgr) error {
	utils.SendlogMsg("开始读取未完成的下载")
	sessionInfoList, _ := session.ReadSession()
	if num := len(sessionInfoList); num != 0 {
		utils.SendlogMsg(
			fmt.Sprintf("%d个文件未下载完成", num))
	}

	for _, sessionInfo := range sessionInfoList {
		sessionSet[sessionInfo.Url] = space
		err := dl.AddWaitList(WaitObject{
			Url:      sessionInfo.Url,
			FilePath: sessionInfo.FilePath,
			TmpName:  sessionInfo.TmpName,
		})
		if err != nil {
			return err
		}
	}
	for _, fileTypePath := range constant.FileTypePaths {
		for _, platformType := range constant.PlatformTypePaths {
			utils.CreateFile(constant.ASoulPaths,
				constant.InputBasePath+
					fileTypePath+
					platformType)

			err := WaitDownload(
				constant.ASoulPaths,
				fileTypePath,
				platformType,
				dl,
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func WaitDownload(paths []string, fileTypePath string, platformType string, dl *DownloadMgr) error {
	// fmt.Println("开始读取下载列表...")
	for _, memPath := range paths {
		// utils.SendlogMsg(
		// 	fmt.Sprintf("开始读取%s%s%s的下载列表...", fileTypePath, platformType, memPath))
		body, err := utils.ReadAll(
			constant.InputBasePath +
				fileTypePath +
				platformType +
				memPath)
		if err != nil {
			return err
		}
		// fmt.Println("下载列表读取完成")
		allPaths := strings.Split(string(body), "\n")
		var tmpPaths []string
		for _, path := range allPaths {
			if _, exist := sessionSet[path]; !exist {
				if path != "" {
					fileName := filepath.Base(path)
					parseFileName, err := url.QueryUnescape(fileName)
					if err != nil {
						return err
					}
					fileInfo, _ := os.Stat(constant.OutputBasePath + fileTypePath + platformType + memPath + "/" + parseFileName)
					if fileInfo == nil {
						tmpPaths = append(tmpPaths, path)
					}
				}
			}
		}
		if len(tmpPaths) <= 0 {
			// utils.SendlogMsg(
			// 	fmt.Sprintf("%s%s%s没有文件需要下载", fileTypePath, platformType, memPath))
		} else {
			// utils.SendlogMsg(
			// 	fmt.Sprintf("%s%s%s加入了%d个文件进入下载队列", fileTypePath, platformType, memPath, len(tmpPaths)))
			for _, path := range tmpPaths {
				filePath := constant.OutputBasePath + fileTypePath + platformType + memPath + "/" + filepath.Base(path)

				err = dl.AddWaitList(WaitObject{
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
