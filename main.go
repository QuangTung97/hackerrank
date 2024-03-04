package main

type position struct {
	rows int
	cols int
}

func nextPos(p position, n int) position {
	p2 := position{
		rows: p.rows * 2,
		cols: p.cols * 2,
	}

	mid := position{
		rows: n - 1,
		cols: n - 1,
	}
	v := position{
		rows: p2.rows - mid.rows,
		cols: p2.cols - mid.cols,
	}
	vo := position{
		rows: v.cols,
		cols: -v.rows,
	}
	return position{
		rows: (mid.rows + vo.rows) / 2,
		cols: (mid.cols + vo.cols) / 2,
	}
}

func rotate(matrix [][]int) {
	n := len(matrix[0])
	mid := n / 2
	colMid := (n + 1) / 2
	for row := 0; row < mid; row++ {
		for col := 0; col < colMid; col++ {
			prev := matrix[row][col]
			p := position{
				rows: row,
				cols: col,
			}
			for i := 0; i < 3; i++ {
				p = nextPos(p, n)
				tmp := matrix[p.rows][p.cols]
				matrix[p.rows][p.cols] = prev
				prev = tmp
			}
			matrix[row][col] = prev
		}
	}
}
