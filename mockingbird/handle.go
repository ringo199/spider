package mockingbird

import (
	"path/filepath"
	"strings"

	"github.com/ringo199/spider/utils"
)

func PreGenerate() error {
	audiosPath := "outputPath/audios/bilibili/tmp"
	spleeterPath := "outputPath/spleeter/diana"
	mockingBirdPath := "outputPath/mockingBird/dataset/aidatatang_200zh"

	fileNameMap := make(map[string]string)
	audiosFiles, err := utils.ReadDir(audiosPath)
	if err != nil {
		return err
	}
	for _, f_info := range audiosFiles {
		k := f_info.Name()
		k = strings.TrimSuffix(k, ".aac")
		fileNameMap[k] = utils.GetSha1(k)
	}
	files, err := utils.ReadDir(filepath.Join(spleeterPath, "output"))
	if err != nil {
		return err
	}
	for _, f_info := range files {
		f_name := f_info.Name()
		for k, fnm := range fileNameMap {
			if strings.Contains(f_name, k) {
				sub := strings.Split(f_name, "_")[1]
				f_name = fnm + "_" + sub
			}
		}
		utils.CopyFile(
			filepath.Join(mockingBirdPath, "corpus", "train", "diana", f_name+".wav"),
			filepath.Join(spleeterPath, "output", f_info.Name(), "vocals.wav"),
		)
	}
	transcript, err := utils.OpenFile(
		filepath.Join(mockingBirdPath, "transcript", "aidatatang_200_zh_transcript.txt"))

	if err != nil {
		return err
	}
	err = transcript.Truncate(0)
	if err != nil {
		return err
	}
	_, err = transcript.Seek(0, 0)
	if err != nil {
		return err
	}
	defer transcript.Close()

	transcriptStr := ""
	b, err := utils.ReadAll(filepath.Join(spleeterPath, "transcript.txt"))
	if err != nil {
		return err
	}
	transcriptStr += string(b)
	b, err = utils.ReadAll(filepath.Join(spleeterPath, "transcript_short.txt"))
	if err != nil {
		return err
	}
	transcriptStr += string(b)

	for k, fnm := range fileNameMap {
		transcriptStr = strings.ReplaceAll(transcriptStr, k, fnm)
	}
	transcript.WriteString(transcriptStr)
	return nil
}
