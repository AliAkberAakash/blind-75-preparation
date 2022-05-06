func search(nums []int, target int) int {
    
    start := 0
    end := len(nums)-1
    
    var mid int
    for start <= end {
        mid = int((start+end)/2)
        
        if(nums[mid]==target){
            return mid
        }
        
        if(target > nums[mid]) {
            start = mid+1
        }
        
        if(target <nums[mid]) {
            end = mid-1
        }
    }
    
    return -1
}
