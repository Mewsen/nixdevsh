package ui

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/erikgeiser/promptkit"
	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/mewsen/nixdevsh/logic"
)

const (
	listHeight   = 14
	defaultWidth = 20
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type item struct {
	Title string
}

func (i item) FilterValue() string { return i.Title }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Title)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list             list.Model
	SelectedAnOption bool
	confirmation     *confirmation.Model
}

func (m *model) Init() tea.Cmd {
	keyMap := &confirmation.KeyMap{
		SelectYes: []string{"y", "Y"},
		SelectNo:  []string{"n", "N"},
		Toggle:    []string{"tab"},
		Submit:    []string{"enter"},
		Abort:     []string{"ctrl+c"},
	}

	conf := &confirmation.Confirmation{
		Prompt:         "Initialize Git repository",
		DefaultValue:   confirmation.Yes,
		Template:       confirmation.DefaultTemplate,
		ResultTemplate: confirmation.DefaultResultTemplate,
		KeyMap:         keyMap,
		WrapMode:       promptkit.Truncate,
	}
	m.confirmation = confirmation.NewModel(conf)

	return m.confirmation.Init()
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			if m.SelectedAnOption {
				confirmation, _ := m.confirmation.Value()
				if confirmation {
					err := logic.InitGitRepository()
					if err != nil {
						log.Fatal("Error initializing Git repository", err)
					}
				}

				return m, tea.Quit
			}

			if !m.SelectedAnOption {
				m.SelectedAnOption = true

				err := logic.CreateFlakeFile(m.list.SelectedItem().(item).Title)
				if err != nil {
					log.Fatal("Error creating flake.nix:", err)
				}

				err = logic.CreateEnvRCInCWD()
				if err != nil {
					log.Fatal("Error creating .envrc:", err)
				}
			}

			return m, nil
		}

	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)

		return m, nil
	default:
	}

	var cmd tea.Cmd
	if !m.SelectedAnOption {
		m.list, cmd = m.list.Update(msg)
	} else {
		_, cmd = m.confirmation.Update(msg)
	}

	return m, cmd
}

func (m *model) View() string {
	if m.SelectedAnOption {
		return "\n" + m.confirmation.View()
	}
	return "\n" + m.list.View()
}

func NewListModel() list.Model {
	dirNames := logic.DirNamesFromEmbededDir()

	items := make([]list.Item, len(dirNames))

	for i, dirName := range dirNames {
		items[i] = item{Title: dirName}
	}

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Select a Dev Shell"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	return l
}

func StartUI() {
	p := tea.NewProgram(&model{list: NewListModel()}, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal("Error running program:", err)
	}
}
