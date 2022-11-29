package maxareaofisland695


func maxAreaOfIsland(grid [][]int) int {
    var max int
    var dfs func(i int, j int, area *int)
    dfs = func(i int , j int, area *int) {
        if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
            return 
        }
        *area = *area + 1
        grid[i][j] = 0
        dfs(i+1,j, area)
        dfs(i-1,j, area)
        dfs(i,j+1, area)
        dfs(i,j-1, area)
    }
    for i:=0; i < len(grid); i++ {
        for j:=0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                var area int
                dfs(i,j, &area)
                if area > max {
                    max = area
                }
            }
        }   
    }
    return max
}