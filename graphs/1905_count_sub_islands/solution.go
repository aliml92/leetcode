package countsubislands1905

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	rows, cols := len(grid1), len(grid1[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || grid2[i][j] == 0 {
			return
		}
		grid2[i][j] = 0
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid2[i][j] == 1 && grid1[i][j] == 0 {
				dfs(i, j)
			}
		}
	}

	var count int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid2[i][j] == 1 {
				dfs(i, j)
				count++
			}
		}
	}
	return count
}
