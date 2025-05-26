package ping

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	probing "github.com/prometheus-community/pro-bing"
)

type pingMsg struct {
	status string
}

type Model struct {
	IP     string
	Status string
}

func NewModel(ip string) Model {
	return Model{IP: ip, Status: "Pinging..."}
}

func (m Model) Init() tea.Cmd {
	return pingCmd(m.IP)
}

func pingCmd(ip string) tea.Cmd {
	return func() tea.Msg {
		pinger, err := probing.NewPinger(ip)
		if err != nil {
			return pingMsg{status: fmt.Sprintf("Failed to initialize pinger: %v", err)}
		}
		pinger.Count = 1
		err = pinger.Run()
		if err != nil {
			return pingMsg{status: fmt.Sprintf("Host %s is unreachable: %v", ip, err)}
		}
		stats := pinger.Statistics()
		if stats.PacketsRecv > 0 {
			return pingMsg{status: fmt.Sprintf("Host %s is reachable", ip)}
		}
		return pingMsg{status: fmt.Sprintf("Host %s is unreachable", ip)}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case pingMsg:
		m.Status = msg.status
		// Schedule another ping after 1 second
		return m, tea.Tick(time.Second, func(time.Time) tea.Msg {
			return pingCmd(m.IP)()
		})
	}
	return m, nil
}

func (m Model) View() string {
	return fmt.Sprintf("Ping status for %s:\n%s\n", m.IP, m.Status)
}
