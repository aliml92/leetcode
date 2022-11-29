package islandperimeter463


func islandPerimeter(grid [][]int) int {
    var p int
    visit := make(map[[2]int]struct{})
    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 || grid[i][j] == 0 {
            p += 1
            return
        }
        _, ok := visit[[2]int{i,j}]
        if ok {
            return
        }
        visit[[2]int{i,j}] = struct{}{}
        dfs(i, j+1)
        dfs(i+1, j)
        dfs(i, j-1)
        dfs(i-1, j)
        
    }
    
    for i:= 0; i < len(grid); i++ {
        for j:= 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                dfs(i,j)
                return p
            }
        }
    }
    
    return p
}