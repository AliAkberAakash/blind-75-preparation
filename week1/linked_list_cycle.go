/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    lookupTable := make(map[int]int)
    
    temp := head
    
    for temp != nil {
        lookupTable[temp.Val]++;
        
        if lookupTable[temp.Val]>1 {
            return true;
        }
        temp = temp.Next
    }
    
    return false;
}
