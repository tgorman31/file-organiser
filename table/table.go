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
	IsFileFiew   bool
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

	return Model{
		dirTable: tbl.New([]tbl.Column{
			tbl.NewColumn(columnKeyName, "Name", 20),
			tbl.NewColumn(columnKeySize, "Size", 15),
		}).WithRows(rows).
			BorderRounded().
			WithBaseStyle(styleBase).
			WithPageSize(20).
			Focused(true),
		IsFileFiew: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.IsFileFiew {
		m.fileTable, cmd = m.fileTable.Update(msg)
	} else {
		m.dirTable, cmd = m.fileTable.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if !m.IsFileFiew {
				selectedFiles := m.dirTable.HighlightedRow().Data[columnkeyFiles].([]ccmd.File)
				m.CurrentFiles = selectedFiles
				m.fileTable = CreateFileTable(selectedFiles)
				m.IsFileFiew = true
			}
		case "esc":
			if m.IsFileFiew {
				m.IsFileFiew = false
			}
		}
	}
	return m, cmd
}

func (m Model) View() string {
	if !m.IsFileFiew {
		return m.fileTable.View()
	}
	return m.dirTable.View()
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

// func updateDir(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
// 	var (
// 		cmd  tea.Cmd
// 		cmds []tea.Cmd
// 	)
// 	// fmt.Println("Made it here")
// 	m.dirTable, cmd = m.dirTable.Update(msg)
// 	cmds = append(cmds, cmd)

// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "q":
// 			cmds = append(cmds, tea.Quit)
// 		case "enter":
// 			m.Files = true
// 			// return updateFile(msg, m)
// 			p := tea.NewProgram(New_File_Model(m.dirTable.HighlightedRow().Data[columnkeyFiles].([]ccmd.File)))
// 			if _, err := p.Run(); err != nil {
// 				fmt.Printf("Alas, ther' be an error: %v", err)
// 				os.Exit(1)
// 			}
// 		}

// 	}

// 	return m, tea.Batch(cmds...)
// }

// func updateFile(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
// 	var (
// 		cmd  tea.Cmd
// 		cmds []tea.Cmd
// 	)

// 	m.dirTable, cmd = m.dirTable.Update(msg)
// 	cmds = append(cmds, cmd)

// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "q":
// 			cmds = append(cmds, tea.Quit)
// 		case "b":
// 			m.Files = false
// 			// return updateDir(msg, m)
// 			// return updateDir(msg, m)
// 			// case "esc":
// 			fmt.Println("you pressed enter!")
// 			p := tea.NewProgram(New_File_Model(m.dirTable.HighlightedRow().Data[columnkeyFiles].([]ccmd.File)))
// 			if _, err := p.Run(); err != nil {
// 				fmt.Printf("Alas, ther' be an error: %v", err)
// 				os.Exit(1)
// 			}
// 		}

// 	}

// 	return m, tea.Batch(cmds...)
// }

// func viewDir(m Model) string {

// 	// Get the metadata back out of the row
// 	selected := m.dirTable.HighlightedRow().Data[columnkeyDir].(TableData)

// 	view := lg.JoinVertical(
// 		lg.Left,
// 		styleSubtle.Render("Highlighted: "+fmt.Sprintf("%s", selected.Name)),
// 		m.dirTable.View(),
// 		styleSubtle.Render("Press q or ctrl+c to quit"),
// 	) + "\n"

// 	return lg.NewStyle().MarginLeft(1).Render(view)
// }

// func viewFile(m Model) string {
// 	// Get the metadata back out of the row
// 	selected := m.dirTable.HighlightedRow().Data[columnkeyDir].(TableData)

// 	view := lg.JoinVertical(
// 		lg.Left,
// 		styleSubtle.Render("Highlighted: "+fmt.Sprintf("%s", selected.Name)),
// 		m.dirTable.View(),
// 		styleSubtle.Render("Press q or ctrl+c to quit"),
// 	) + "\n"

// 	return lg.NewStyle().MarginLeft(1).Render(view)
// }
