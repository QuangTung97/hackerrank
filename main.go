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

func generateMatrix(n int) [][]int {
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
		bot:   n,
		left:  -1,
		right: n,
	}

	num := n * n
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	result[0][0] = 1
	for i := 1; i < num; i++ {
		p, d, w = computeNext(p, d, w)
		result[p.row][p.col] = i + 1
	}
	return result
}
