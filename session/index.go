package session

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ringo199/spider/constant"
	"github.com/ringo199/spider/utils"
)

type SessionInfo struct {
	Url      string
	FilePath string
	TmpName  string
}

func CreateSession(tmpName string, sessionInfo SessionInfo) error {
	sessionPath := constant.TmpBasePath + tmpName + ".session"

	f, err := os.Create(sessionPath)
	if err != nil {
		return err
	}
	str := fmt.Sprintf("%s\n%s", sessionInfo.Url, sessionInfo.FilePath)
	writer := bufio.NewWriter(f)
	writer.WriteString(str)
	writer.Flush()
	return nil
}

func ReadSession() ([]*SessionInfo, error) {
	files, err := utils.ReadDir(constant.TmpBasePath)
	if err != nil {
		return nil, err
	}
	var SessionInfoList []*SessionInfo
	for _, v := range files {
		if filepath.Ext(v.Name()) == ".session" {
			b, err := utils.ReadAll(constant.TmpBasePath + v.Name())
			if err != nil {
				return nil, err
			}
			fileObject := strings.Split(string(b), "\n")
			SessionInfoList = append(SessionInfoList, &SessionInfo{
				fileObject[0],
				fileObject[1],
				strings.TrimSuffix(v.Name(), ".session"),
			})
		}
	}
	return SessionInfoList, nil
}

func DeleteSession(tmpName string) error {
	sessionPath := constant.TmpBasePath + tmpName + ".session"
	if err := os.Remove(sessionPath); err != nil {
		return err
	}
	return nil
}
