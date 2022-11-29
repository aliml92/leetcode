package validsudoku36


func isValidSudoku(board [][]byte) bool {
    if (isRowValid(board) && isColValid(board) && isBoxValid(board)) {
		return true
	}
	return false
}

func isRowValid( board [][]byte ) bool {
	for i := 0; i < 9; i++ {
		m := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			item := board[i][j]
			if item != '.' {
				_, ok := m[item]
				if ok {
					return false
				}
				m[item] = true
			}
		}
	}
	return true
}

func isColValid( board [][]byte ) bool {
	for i := 0; i < 9; i++ {
		m := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			item := board[j][i]
			if item != '.' {
				_, ok := m[item]
				if ok {
					return false
				}
				m[item] = true
			}
		}
	}
	return true
}

func isBoxValid( board [][]byte ) bool {
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			m := make(map[byte]bool)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					item := board[i+k][j+l]
					if item != '.' {
						_, ok := m[item]
						if ok {
							return false
						}
						m[item] = true
					}
				}
			}			
		}
	}
	
	return true
}