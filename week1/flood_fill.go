func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    
    oldColor := image[sr][sc]
    image[sr][sc] = newColor;
    
    if oldColor == newColor {
        return image
    }
    
    myFloodFill(image,sr+1,sc,oldColor,newColor)
    myFloodFill(image,sr-1,sc,oldColor,newColor)
    myFloodFill(image,sr,sc+1,oldColor,newColor)
    myFloodFill(image,sr,sc-1,oldColor,newColor)
    
    return image
}

func myFloodFill(image [][]int, sr int, sc int, oldColor int, newColor int) [][]int {
    
    if sr<0 || sr>=len(image) || sc<0 || sc>=len(image[0]) {
        return image
    }
    
    fmt.Println(image[sr][sc])
    if image[sr][sc] != oldColor {
        return image
    }
    
    image[sr][sc] = newColor;
    
    myFloodFill(image,sr+1,sc,oldColor,newColor)
    myFloodFill(image,sr-1,sc,oldColor,newColor)
    myFloodFill(image,sr,sc+1,oldColor,newColor)
    myFloodFill(image,sr,sc-1,oldColor,newColor)
    
    return image
}
