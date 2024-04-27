package style

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitialModel() model {
	return model{
		choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should quit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up.
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down.
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		}

	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Not that we're not returning a command
	return m, nil
}

func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The fotter
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s

}

// var HeaderStyle = lg.NewStyle().
// 	Bold(true).
// 	Align(lg.Center).
// 	Foreground(lg.Color("#7D56F4")).
// 	Width(22)

// var EvenRowStyle = lg.NewStyle().
// 	Foreground(lg.Color("#FAFAFA")).
// 	Bold(false).
// 	Inherit(HeaderStyle)

// var EvenRowStyleNum = lg.NewStyle().
// 	Align(lg.Right).
// 	Inherit(EvenRowStyle)

// var OddRowStyle = lg.NewStyle().
// 	Background(lg.Color("#2b2a2a")).
// 	Inherit(EvenRowStyle)

// var OddRowStyleNum = lg.NewStyle().
// 	Align(lg.Right).
// 	Inherit(OddRowStyle)

// var t = tbl.New().
// 	Border(lg.NormalBorder()).
// 	BorderStyle(lg.NewStyle().Foreground(lg.Color("99"))).
// 	StyleFunc(func(row, col int) lg.Style {
// 		switch {
// 		case row == 0:
// 			return HeaderStyle
// 		case row%2 == 0:
// 			if col == 1 {
// 				return EvenRowStyleNum
// 			}
// 			return EvenRowStyle
// 		default:
// 			if col == 1 {
// 				return OddRowStyleNum
// 			}
// 			return OddRowStyle
// 		}
// 	}).
// 	Headers("FOLDER", "SIZE")

// func CreateTable() *tbl.Table {
// 	return t
// }
