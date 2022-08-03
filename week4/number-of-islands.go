func numIslands(grid [][]byte) int {
    
    row := len(grid)
    col := len(grid[0])
    
    visited := make([][]bool, row)
    
    for i:=0; i<row; i++ {
        for j:=0; j<col; j++ {
            visited[i] = append(visited[i], false)
        }
    }
    
    val := 0
    
    for i:=0; i<row; i++ {
        for j:=0; j<col; j++ {
            val += floodFill(grid,visited,i,j,row,col)
        }
    }
    
    return val
}

func floodFill(grid [][]byte, visited [][]bool, x int, y int, row int, col int) int {
    
    var dx = []int{
        0,0,1,-1,
    }

    var dy = []int{
        1,-1,0,0,
    }
    
    if  x>=0 && y>=0 && x<row && y<col && !visited[x][y] && grid[x][y]=='1' {
        visited[x][y] = true
        
        for i:=0; i<4; i++ {
            floodFill(grid,visited,x+dx[i],y+dy[i],row,col)
        }
        
        return 1
    }
    return 0
}
