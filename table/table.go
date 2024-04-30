package table

import (
	ccmd "file-organiser/cmd"
	"fmt"
	"os"

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

func New_Table_Data(name string, size int, file []ccmd.File) TableData {
	return TableData{
		Name: name,
		Size: size,
		File: file,
	}
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
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if !m.IsFileView {
				selectedFiles := m.dirTable.HighlightedRow().Data[columnkeyFiles].([]ccmd.File)
				m.CurrentFiles = selectedFiles
				m.fileTable = CreateFileTable(selectedFiles)
				m.IsFileView = true
			}
		case "esc":
			if m.IsFileView {
				m.IsFileView = false
			}
		}
	}
	return m, cmd
}

func (m Model) View() string {
	if m.IsFileView {
		view := lg.JoinVertical(
			lg.Left,
			m.fileTable.View(),
			styleSubtle.Render("Press q or ctrl+c to quit"),
		) + "\n"
		return lg.NewStyle().Render(view)
	}
	view := lg.JoinVertical(
		lg.Left,
		m.dirTable.View(),
		styleSubtle.Render("Press q or ctrl+c to quit"),
	) + "\n"
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