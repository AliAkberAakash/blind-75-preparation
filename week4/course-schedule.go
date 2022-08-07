func canFinish(numCourses int, prerequisites [][]int) bool {
    
    matrix := make(map[int][]int)
    visited := make(map[int]bool, numCourses)
    connections := make([]int, numCourses)
    
    fmt.Println(matrix)
    
    for _,p := range prerequisites {
        
        if _,ok:= matrix[p[0]]; ok {
            
        }else{
            matrix[p[0]] = []int{}
        }
        
        matrix[p[0]] = append(matrix[p[0]], p[1])
        
        if _,ok:= matrix[p[1]]; ok {
            
        }else{
            matrix[p[1]] = []int{}
        }
        
        matrix[p[1]] = append(matrix[p[1]], p[0])
        
        connections[p[0]]++        
    }
    
    queue := []int{}
    
    for i,val := range connections {
        if val == 0 {
            queue = append(queue,i)
        }
    }
    
    for len(queue) > 0 {
        top := queue[0]
        queue = queue[1:]
        
        if !visited[top] {
            visited[top] = true

            for _,element := range matrix[top] {
                connections[element]--
                
                if connections[element]==0 {
                    queue = append(queue,element)
                }
            }
        }
        
    }
    
    for _,element := range connections {
        if element > 0 {
            return false
        }
    }
    
    return true
}
