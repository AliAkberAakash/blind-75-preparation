func orangesRotting(grid [][]int) int {
    
    row := len(grid)
    
    if row == 0 {
        return 0
    }
    
    col := len(grid[0])
    time,fresh := 0,0
    queue := [][]int{}
    
    for i:=0; i<row; i++ {
        for j:=0; j<col; j++ {
            if grid[i][j]==1 {
                fresh++
            }
            if grid[i][j]==2 {
                queue = append(queue,[]int{i,j})
            }
        }
    }
    
    direction := [][]int{
        {0,1},
        {1,0},
        {0,-1},
        {-1,0},
    }
    
    for len(queue)>0 {
        l := len(queue)
        update := 0;
        for k:=0; k<l; k++ {
            top := queue[0]
            queue = queue[1:]
            for i:=0; i<4; i++ {
                dx := top[0]+direction[i][0]
                dy := top[1]+direction[i][1]
                
                if dx<0 || dx>=row || dy<0 || dy>=col || grid[dx][dy] != 1 {
                    continue
                }
                grid[dx][dy] = 2
                queue = append(queue,[]int{dx,dy})
                fresh--
                update = 1
            }
        }
        time += update  
    }

    if fresh == 0 {
        return time
    } 
    
    return -1
}
