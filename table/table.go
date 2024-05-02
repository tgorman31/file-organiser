package table

import (
	ccmd "file-organiser/cmd"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	tbl "github.com/evertras/bubble-table/table"
)

type Obj_Type string

const (
	columnKeyName = "name"
	columnKeySize = "size"

	// Hidden data
	columnkeyDir   = "dir"
	columnkeyFiles = "files"

	typeFile Obj_Type = "File"
	typeDir  Obj_Type = "Directory"
)

var (
	styleSubtle = lg.NewStyle().Foreground(lg.Color("#888"))

	styleBase = lg.NewStyle().
			Foreground(lg.Color("#a7a")).
			BorderForeground(lg.Color("#a38")).
			Align(lg.Right)

	objColors = map[Obj_Type]string{
		typeDir:  "#5900ff",
		typeFile: "#d1b9ff",
	}
)

type TableData struct {
	Name string
	Size int
	File []ccmd.File
}

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Esc   key.Binding
	Help  key.Binding
	Quit  key.Binding
}

func New_Table_Data(name string, size int, file []ccmd.File) TableData {
	return TableData{
		Name: name,
		Size: size,
		File: file,
	}
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Enter}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Enter},
		{k.Esc, k.Help, k.Quit},
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter", "enter"),
		key.WithHelp("↵", "enter directory"),
	),
	Esc: key.NewBinding(
		key.WithKeys("esc", "esc"),
		key.WithHelp("esc", "exit directory"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func (t TableData) ToRow() tbl.Row {

	return tbl.NewRow(tbl.RowData{
		columnKeyName: t.Name,
		columnKeySize: t.Size,

		columnkeyFiles: t.File,
		columnkeyDir:   t,
	})
}

type Model struct {
	dirTable     tbl.Model
	fileTable    tbl.Model
	CurrentFiles []ccmd.File
	IsFileView   bool
	keys         keyMap
	help         help.Model
	lastKey      string
}

func Table(dir []ccmd.Dir) {
	p := tea.NewProgram(NewModel(dir))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, ther' be an error: %v", err)
		os.Exit(1)
	}
}

func NewModel(dir []ccmd.Dir) Model {

	rows := []tbl.Row{}

	for _, d := range dir {
		t := New_Table_Data(d.Name, d.Size, d.File)
		rows = append(rows, t.ToRow())
	}

	dirTable := tbl.New([]tbl.Column{
		tbl.NewColumn(columnKeyName, "Name", 20),
		tbl.NewColumn(columnKeySize, "Size", 20),
	}).WithRows(rows).
		BorderRounded().
		WithBaseStyle(styleBase).
		WithPageSize(20).
		Focused(true)

	fileTable := tbl.New([]tbl.Column{
		tbl.NewColumn(columnKeyName, "Name", 20),
		tbl.NewColumn(columnKeySize, "Size", 20),
	}).WithRows([]tbl.Row{}).
		BorderRounded().
		WithBaseStyle(styleBase).
		WithPageSize(20)

	return Model{
		dirTable:   dirTable,
		fileTable:  fileTable,
		IsFileView: false,
		help:       help.New(),
		keys:       keys,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.IsFileView {
		m.fileTable, cmd = m.fileTable.Update(msg)
	} else {
		m.dirTable, cmd = m.dirTable.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			m.lastKey = "↑"
		case key.Matches(msg, m.keys.Down):
			m.lastKey = "↓"
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Esc):
			if m.IsFileView {
				m.IsFileView = false
			}
		case key.Matches(msg, m.keys.Enter):
			if !m.IsFileView {
				selectedFiles := m.dirTable.HighlightedRow().Data[columnkeyFiles].([]ccmd.File)
				m.CurrentFiles = selectedFiles
				m.fileTable = CreateFileTable(selectedFiles)
				m.IsFileView = true
			}
		}
	}
	return m, cmd
}

func (m Model) View() string {
	helpView := m.help.View(m.keys)
	if m.IsFileView {
		view := lg.JoinVertical(
			lg.Left,
			m.fileTable.View(),
			helpView,
			// styleSubtle.Render("Press q or ctrl+c to quit. Press esc to return"),
		) + "\n"
		// view = view + helpView
		return lg.NewStyle().Render(view)
	}
	view := lg.JoinVertical(
		lg.Left,
		m.dirTable.View(),
		helpView,
		// styleSubtle.Render("Press q or ctrl+c to quit"),
	) + "\n"
	// view = view + helpView
	return lg.NewStyle().Render(view)
}

func CreateFileTable(files []ccmd.File) tbl.Model {

	rows := []tbl.Row{}

	for _, f := range files {
		t := New_Table_Data(f.Name, f.Size, nil)
		rows = append(rows, t.ToRow())
	}

	fileTable := tbl.New([]tbl.Column{
		tbl.NewColumn(columnKeyName, "Name", 20),
		tbl.NewColumn(columnKeySize, "Size", 15),
	}).WithRows(rows).
		BorderRounded().
		WithBaseStyle(styleBase).
		WithPageSize(20).
		Focused(true)

	return fileTable
}
