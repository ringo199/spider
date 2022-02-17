package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
	"github.com/ringo199/spider/constant"
	"github.com/ringo199/spider/download"
	"github.com/ringo199/spider/utils"
)

var dl download.DownloadMgr = download.DownloadMgr{
	Limit: constant.DownloadLimit,
}

type model struct {
	ProgressList []*progress.Model
	Log          viewport.Model
	IsShowLog    bool
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
		IsShowLog:    false,
	}
}

func (m model) Init() tea.Cmd {
	err := download.InitFileAndWaitDownload(&dl)
	if err != nil {
		utils.SendlogMsg(err.Error())
	}
	//  else {
	// 	m.IsShowLog = false
	// }
	return tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		k := msg.String()
		if k == "t" {
			m.IsShowLog = !m.IsShowLog
			return m, tick
		}
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case tickMsg:
		if err := dl.Update(); err != nil {
			utils.SendlogMsg(err.Error())
		}
		if dl.CheckOver() {
			// utils.SendlogMsg("下载已结束，请点击esc退出")
			m.IsShowLog = true
			return m, tick
		}
		return m, tick
	}

	m.Log, _ = m.Log.Update(msg)
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
		s += fmt.Sprintf("%s: lastSize: %s\n%s curSize: %s size: %s\n",
			v.FilePath, v.Wc.LastTransSizeFormatData, prog.ViewAs(v.Wc.Percent), v.Wc.FormatData, v.Wc.AllFormatData)
	}
	return indent.String(s, 1)
}

func (m model) View() string {
	if m.IsShowLog {
		vp := viewport.Model{Width: 375, Height: 20}

		vp.SetContent(utils.GetLogMsg())
		m.Log = vp
		return m.Log.View()
	}
	return fmt.Sprintf("%s\n\n%s\n", ShowDownloadInfo(m), ShowProgress(m))
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Second / 60)
	return tickMsg{}
}

func Initial() *tea.Program {
	return tea.NewProgram(
		InitialModel(),
		tea.WithMouseCellMotion(),
	)
}
