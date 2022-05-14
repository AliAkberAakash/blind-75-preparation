func updateMatrix(mat [][]int) [][]int {
    
    inf := 999999
    
    rowLim := len(mat)
    colLim := len(mat[0])
    
    for row,v := range mat {
        for col,t := range v{
            if(t==1){
                mat[row][col]=inf
            }
        }
    }
    
    
    for row,v := range mat {
        for col,t := range v{
            if(t==0){
                floodFill(mat,0,row,col, rowLim, colLim)
            }
        }
    }
    
    return mat
}

func floodFill(mat [][]int, currentVal, row, col, rowLim, colLim int){
    if(row<0 || row>=rowLim || col<0 || col>=colLim) {
        return
    }
    
    if(mat[row][col]!=0){
        mat[row][col] = min(currentVal+1, mat[row][col])
        currentVal++;
    }
    
    if row+1 < rowLim && mat[row+1][col]>currentVal+1 {
        floodFill(mat,currentVal,row+1,col,rowLim,colLim)
    }
    
    if row-1 >=0 && mat[row-1][col]>currentVal+1 {
        floodFill(mat,currentVal,row-1,col,rowLim,colLim)
    }
    
    if col+1 < colLim && mat[row][col+1]>currentVal+1 {
        floodFill(mat,currentVal,row,col+1,rowLim,colLim)
    }
    
    if col-1 >=0 && mat[row][col-1]>currentVal+1 {
        floodFill(mat,currentVal,row,col-1,rowLim,colLim)
    }
}

func min(a1,a2 int) int {
    if a1<a2 {
        return a1
    }
    return a2
}
