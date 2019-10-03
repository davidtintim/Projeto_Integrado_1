package src

import (
	"fmt"
	"strconv"
	"strings"
	)

type Matrix struct{
	rows[][]int
	cols[][]int
}

func (m *Matrix) Set(row, col, val int) (ok bool) {
	mm := *m
	if row < 0 || row >= len(mm) || col < 0 {
		return false
	}
	if cols := len(mm[0]); col >= cols {
		return false
	}
	mm[row][col] = val
	return true
}


func(m *Matrix) Rows() [][]int {
	rows := make([][]int, len(m.rows))
	for i := range m.rows{
		rows[i] = append([]int{}, mr...)
	}
	return rows
}

func (m *Matrix) Cols() [][]int {
	mm := *m
	if len(mm) == 0 {
		return nil
	}
	c := make([][]int, len(mm[0]))
	for i := range c {
		col := make([]int, len(mm))
		for j := range col {
			col[j] = mm[j][i]
		}
		c[i] = col
	}
	return c
}
