func threeSum(nums []int) [][]int {
    
    var res [][]int
    
    sort.Ints(nums)
    
    for i,num := range nums {
        left := i+1
        right := len(nums)-1
        
        for left < right {
            target := nums[left]+nums[right]+num
            if target == 0 {
                var temp []int
                temp = append(temp,nums[left],nums[right],num)
                sort.Ints(temp)
                if !exists(res,temp) {
                    res = append(res, temp)
                }
                left++
                right--
            }else if target > 0 {
                right--
            }else{
                left++
            }
        }
        
    }
    
    return res
}

func exists(source [][]int, target []int) bool {
    for _,num := range source {
        if num[0]==target[0] && num[1]==target[1] && num[2]==target[2] {
            return true
        }
    }
    return false
}
