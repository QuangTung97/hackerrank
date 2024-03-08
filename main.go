package main

func checkRegion(
	board [][]byte,
	rowBegin int, rowEnd int,
	colBegin int, colEnd int,
) bool {
	set := map[byte]struct{}{}
	for row := rowBegin; row < rowEnd; row++ {
		for col := colBegin; col < colEnd; col++ {
			e := board[row][col]
			if e == '.' {
				continue
			}
			num := e - '0'
			_, existed := set[num]
			if existed {
				return false
			}
			set[num] = struct{}{}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		if ok := checkRegion(board, row, row+1, 0, 9); !ok {
			return false
		}
	}
	for col := 0; col < 9; col++ {
		if ok := checkRegion(board, 0, 9, col, col+1); !ok {
			return false
		}
	}
	for row := 0; row < 9; row += 3 {
		for col := 0; col < 9; col += 3 {
			if ok := checkRegion(board, row, row+3, col, col+3); !ok {
				return false
			}
		}
	}

	return true
}
