package views

import (
	"fmt"
	"hinter/hinter/common"
	"math"
	"strings"
)

type TableModel struct {
	columnSize []int
	headers    []string
	rows       [][]string
}

func FromEntries(columnSize []int, entries []common.Entry) TableModel {
	rows := make([][]string, len(entries))
	for row, entry := range entries {
		rows[row] = []string{entry.Key, entry.Value}
	}
	return TableModel{columnSize, []string{"Key", "Value"}, rows}
}

func (m TableModel) View() (v string) {
	assertEqualLenghts(m.headers, m.rows, m.columnSize)
	tableWidth := 1
	for _, width := range m.columnSize {
		tableWidth += width + 1
	}
	divider := strings.Repeat("-", tableWidth) + "\n"
	header := renderRow(m.headers, m.columnSize)
	rows := ""
	for _, currentRow := range m.rows {
		rows += renderRow(currentRow, m.columnSize)
	}
	return header + divider + rows

}

func assertEqualLenghts(headers []string, rows [][]string, columnSizes []int) {
	if len(headers) != len(columnSizes) {
		panic(fmt.Sprintf("headers (%d) and columnSizes (%d) have different length", len(columnSizes), len(headers)))
	}
	if len(rows) > 0 && len(headers) != len(rows[0]) {
		panic(fmt.Sprintf("headers (%d) and rows[0] (%d) have different lengths", len(headers), len(rows[0])))
	}
}

func renderRow(cells []string, columnSizes []int) (r string) {
	necessaryLines := 1
	for col, cell := range cells {
		if len(cell) > 0 {
			linesForThisCell := math.Ceil(float64(len(cell)) / float64(columnSizes[col]))
			necessaryLines = int(math.Max(float64(necessaryLines), linesForThisCell))
		}
	}
	r = ""

	for line := 0; line < necessaryLines; line++ {
		r += "|"
		for col, cell := range cells {
			columnSize := columnSizes[col]
			paddedCell := cell + strings.Repeat(" ", necessaryLines*columnSize-len(cell))
			cellContentForCurrentLine := paddedCell[line*columnSize : (line+1)*columnSize]
			r += padRightWithWhitespace(cellContentForCurrentLine, columnSizes[col]) + "|"
		}
		r += "\n"
	}
	return r
}

func padRightWithWhitespace(value string, columnSize int) string {
	padding := int(math.Max(float64(columnSize-len(value)), 0))
	return value + strings.Repeat(" ", padding)
}
