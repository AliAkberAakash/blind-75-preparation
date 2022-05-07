func containsDuplicate(nums []int) bool {
    
    var frequencyMap = make(map[int]bool)
    
    for _,num := range nums {
        
        if frequencyMap[num]{
            return true;
        }
        
        frequencyMap[num]=true;
        
    }
    
    return false;
}
