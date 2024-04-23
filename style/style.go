package style

import (
	lg "github.com/charmbracelet/lipgloss"
	tbl "github.com/charmbracelet/lipgloss/table"
)

var HeaderStyle = lg.NewStyle().
	Bold(true).
	Align(lg.Center).
	Foreground(lg.Color("#7D56F4")).
	Width(22)

var EvenRowStyle = lg.NewStyle().
	Foreground(lg.Color("#FAFAFA")).
	Bold(false).
	Inherit(HeaderStyle)

var EvenRowStyleNum = lg.NewStyle().
	Align(lg.Right).
	Inherit(EvenRowStyle)

var OddRowStyle = lg.NewStyle().
	Background(lg.Color("#2b2a2a")).
	Inherit(EvenRowStyle)

var OddRowStyleNum = lg.NewStyle().
	Align(lg.Right).
	Inherit(OddRowStyle)

var t = tbl.New().
	Border(lg.NormalBorder()).
	BorderStyle(lg.NewStyle().Foreground(lg.Color("99"))).
	StyleFunc(func(row, col int) lg.Style {
		switch {
		case row == 0:
			return HeaderStyle
		case row%2 == 0:
			if col == 1 {
				return EvenRowStyleNum
			}
			return EvenRowStyle
		default:
			if col == 1 {
				return OddRowStyleNum
			}
			return OddRowStyle
		}
	}).
	Headers("FOLDER", "SIZE")

func CreateTable() *tbl.Table {
	return t
}
