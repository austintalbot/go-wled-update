package main

// A simple example that shows how to render a progress bar in a "pure"
// fashion. In this example we bump the progress by 25% every second,
// maintaining the progress state on our top level model using the progress bar
// model's ViewAs method only for rendering.
//
// The signature for ViewAs is:
//
//     func (m Model) ViewAs(percent float64) string
//
// So it takes a float between 0 and 1, and renders the progress bar
// accordingly. When using the progress bar in this "pure" fashion and there's
// no need to call an Update method.
//
// The progress bar is also able to animate itself, however. For details see
// the progress-animated example.

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/austintalbot/go-wled-update/internal/ping"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

func main() {
	prog := progress.New(progress.WithScaledGradient(
		"#FF7CCB",
		"#FDFF8C",
	))

	ip := "192.168.68.200"

	if _, err := tea.NewProgram(model{
		progress: prog,
		ping:     ping.NewModel(ip),
	}).Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}

type tickMsg time.Time

type model struct {
	percent  float64
	progress progress.Model
	ping     ping.Model
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tickCmd(),
		m.ping.Init(),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Let the ping model handle the message first
	var pingCmd tea.Cmd
	var pingModel tea.Model
	pingModel, pingCmd = m.ping.Update(msg)
	m.ping = pingModel.(ping.Model)
	if pingCmd != nil {
		cmds = append(cmds, pingCmd)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = min(msg.Width-padding*2-4, maxWidth)
		return m, nil

	case tickMsg:
		m.percent += 0.01
		if m.percent > 1.0 {
			m.percent = 1.0
			return m, tea.Quit
		}
		cmds = append(cmds, tickCmd())
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.ViewAs(m.percent) + "\n\n" +
		pad + m.ping.View() + "\n" +
		pad + helpStyle("Press any key to quit")
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
