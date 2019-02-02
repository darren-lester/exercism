package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type matrix [][]int

func (m *matrix) Rows() [][]int {
	rows := make([][]int, len(*m))
	for i, row := range *m {
		rows[i] = append([]int{}, row...)
	}
	return rows
}

func (m *matrix) Cols() [][]int {
	rows := *m
	cols := make([][]int, len(rows[0]))
	for i := range cols {
		cols[i] = make([]int, len(rows))
		for j := range cols[i] {
			cols[i][j] = rows[j][i]
		}
	}
	return cols
}

func (m *matrix) Set(r, c, v int) bool {
	if r < 0 || c < 0 {
		return false
	}
	if r >= len(*m) || c >= len((*m)[0]) {
		return false
	}
	(*m)[r][c] = v
	return true
}

func New(input string) (matrix, error) {
	rows := strings.Split(input, "\n")
	m := make([][]int, len(rows))

	for i := range rows {
		rowElements := strings.Fields(rows[i])
		if len(rowElements) == 0 {
			return nil, errors.New("empty row")
		}

		if i > 0 && len(rowElements) != len(m[0]) {
			return nil, errors.New("rows are not of equal length")
		}

		row := make([]int, len(rowElements))
		for i, e := range rowElements {
			val, err := strconv.Atoi(e)
			if err != nil {
				return nil, fmt.Errorf("Cannot convert %v to int", e)
			}
			row[i] = val
		}
		m[i] = row
	}

	return m, nil
}
