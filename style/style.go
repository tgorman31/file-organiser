package style

import (
	"fmt"
	"os"

	cmds "file-organiser/cmd"

	"github.com/charmbracelet/bubbles/table"
	tbl "github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

var baseStyle = lg.NewStyle().
	BorderStyle(lg.NormalBorder()).
	BorderBackground(lg.Color("240"))

type Table struct {
	Name string
	Size int
	Dir  []cmds.Dir
}

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// var d []cmds.Dir
	// var r tbl.Row

	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		// These keys should quit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "enter" key and the spacebar toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter":

			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[0]),
			)

		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"

}

func Dir_Tbl_Rows(dirs []cmds.Dir) []table.Row {
	var t []tbl.Row
	for _, dir := range dirs {
		t = append(t, table.Row{dir.Name, fmt.Sprint(dir.Size)})
	}
	return t
}

func File_Tbl_Rows(files []cmds.File) []table.Row {
	var t []tbl.Row
	for _, file := range files {
		t = append(t, table.Row{file.Name, fmt.Sprint(file.Size)})
	}
	return t
}

func CreateTable(r []table.Row) {
	columns := []table.Column{
		{Title: "FOLDER", Width: 22},
		{Title: "SIZE", Width: 15},
	}
	t := tbl.New(
		tbl.WithColumns(columns),
		tbl.WithRows(r),
		tbl.WithFocused(true),
		tbl.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lg.NormalBorder()).
		BorderForeground(lg.Color("240")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lg.Color("259")).
		Background(lg.Color("57")).
		Bold(false)

	t.SetStyles(s)

	m := model{t}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
