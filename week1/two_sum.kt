class Solution {
    fun twoSum(nums: IntArray, target: Int): IntArray {
        
        var lookupTable = mutableMapOf<Int,Int>()
        
        nums.forEachIndexed { index, value ->
            lookupTable[value] = index
        }
        
        var ans = IntArray(2)
        nums.forEachIndexed { index, value ->
            var tmp = target-value;
            if(lookupTable.containsKey(tmp)){
                if(index!=lookupTable[tmp]){
                  ans[0]=index
                  ans[1]=lookupTable[tmp]!!
                }
            }
        }
        
        return ans
    }
}
