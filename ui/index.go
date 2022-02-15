package ui

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
	"github.com/ringo199/spider/constant"
	"github.com/ringo199/spider/download"
	"github.com/ringo199/spider/utils"
)

var dl utils.DownloadList = utils.DownloadList{
	Limit: constant.DownloadLimit,
	Wcl:   &utils.WriteCounterList{},
}

type model struct {
	ProgressList []*progress.Model
}

func initialProgress(size int) *[]*progress.Model {
	var progressList []*progress.Model

	for i := 0; i < size; i++ {
		prog := progress.NewModel(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))
		progressList = append(progressList, &prog)
	}

	return &progressList
}

func InitialModel() model {
	progressList := initialProgress(dl.Limit)
	return model{
		ProgressList: *progressList,
	}
}

func (m model) Init() tea.Cmd {
	utils.CreateFile(constant.ASoulPath,
		constant.InputBasePath+
			constant.VideoBasePath+
			constant.BilibiliPath)
	utils.CreateFile(constant.ASoulPath,
		constant.InputBasePath+
			constant.DanmakuBasePath+
			constant.BilibiliPath)
	utils.CreateFile(constant.ASoulPath,
		constant.InputBasePath+
			constant.SubtitleBasePath+
			constant.BilibiliPath)
	utils.CreateFile(constant.ASoulPath,
		constant.InputBasePath+
			constant.VideoBasePath+
			constant.DouyinPath)

	err := download.DownloadFn(
		constant.ASoulPath,
		constant.VideoBasePath,
		constant.DouyinPath,
		&dl)
	if err != nil {
		log.Fatal(err)
	}
	err = download.DownloadFn(
		constant.ASoulPath,
		constant.VideoBasePath,
		constant.BilibiliPath,
		&dl)
	if err != nil {
		log.Fatal(err)
	}
	err = download.DownloadFn(
		constant.ASoulPath,
		constant.DanmakuBasePath,
		constant.BilibiliPath,
		&dl)
	if err != nil {
		log.Fatal(err)
	}
	err = download.DownloadFn(
		constant.ASoulPath,
		constant.SubtitleBasePath,
		constant.BilibiliPath,
		&dl)
	if err != nil {
		log.Fatal(err)
	}
	return tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case tickMsg:
		dl.StartDownload()
		dl.Wcl.FilterWc()
		// if dl.CheckOver() {
		// 	return m, tea.Quit
		// }
		return m, tick
	}
	return m, nil
}

func ShowProgress(m model) string {
	var s string
	for k, v := range dl.Wcl.List {
		prog := m.ProgressList[k]
		s += fmt.Sprintf("%s: %s\n", v.FilePath, prog.ViewAs(v.Percent))
	}
	return indent.String(s, 1)
}

func (m model) View() string {
	return ShowProgress(m)
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Second / 60)
	return tickMsg{}
}
