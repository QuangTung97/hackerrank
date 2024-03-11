package main

type position struct {
	row int
	col int
}

type window struct {
	top   int
	bot   int
	left  int
	right int
}

func computeNext(p position, d position, w window) (position, position, window) {
	rotate := func() {
		newDir := position{
			row: d.col,
			col: -d.row,
		}
		d = newDir
	}

	if d.col == 1 {
		if p.col+1 >= w.right {
			w.top++
			rotate()
		}
	}

	if d.row == 1 {
		if p.row+1 >= w.bot {
			w.right--
			rotate()
		}
	}

	if d.col == -1 {
		if p.col-1 <= w.left {
			w.bot--
			rotate()
		}
	}

	if d.row == -1 {
		if p.row-1 <= w.top {
			w.left++
			rotate()
		}
	}

	p.row += d.row
	p.col += d.col
	return p, d, w
}

func spiralOrder(matrix [][]int) []int {
	p := position{
		row: 0,
		col: 0,
	}
	d := position{
		row: 0,
		col: 1,
	}
	w := window{
		top:   -1,
		bot:   len(matrix),
		left:  -1,
		right: len(matrix[0]),
	}

	num := len(matrix) * len(matrix[0])
	result := make([]int, 0, num)
	result = append(result, matrix[0][0])
	for i := 0; i < num-1; i++ {
		p, d, w = computeNext(p, d, w)
		result = append(result, matrix[p.row][p.col])
	}
	return result
}
