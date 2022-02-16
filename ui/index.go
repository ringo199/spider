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
)

var dl download.DownloadMgr = download.DownloadMgr{
	Limit: constant.DownloadLimit,
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
	err := download.InitFileAndWaitDownload(&dl)
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
		dl.Update()
		if dl.CheckOver() {
			return m, tea.Quit
		}
		return m, tick
	}
	return m, nil
}

func ShowDownloadInfo(m model) string {
	wls := dl.GetWaitListSize()
	dls := dl.GetDownloadingListSize()
	if wls == 0 && dls == 0 {
		return "当前没有需要下载的文件，请点击esc退出"
	}
	return fmt.Sprintf("当前待下载:%d，当前正在下载:%d",
		wls, dls)
}

func ShowProgress(m model) string {
	var s string
	for k, v := range dl.DownloadingList {
		prog := m.ProgressList[k]
		s += fmt.Sprintf("%s:\n%s size：%s\n",
			v.FilePath, prog.ViewAs(v.Wc.Percent), v.Wc.AllFormatData)
	}
	return indent.String(s, 1)
}

func (m model) View() string {
	return fmt.Sprintf("%s\n\n%s\n", ShowDownloadInfo(m), ShowProgress(m))
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Second / 60)
	return tickMsg{}
}
