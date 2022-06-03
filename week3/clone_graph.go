func cloneGraph(node *Node) *Node {
    
    if node == nil {
        return nil
    }
    
    visited := make(map[*Node]*Node)
    
    return dfs(node,visited);
}

func dfs(node *Node, visited map[*Node]*Node) *Node{
    
    if _,ok := visited[node]; ok {
        return visited[node]
    }
    
    var newNeighbors []*Node
    newNode := &Node{
        node.Val,
        newNeighbors,
    }
    
    visited[node] = newNode
    
    for _,neighbor := range node.Neighbors {
        newNode.Neighbors = append(newNode.Neighbors, dfs(neighbor,visited))
    }
    
    return newNode
}
