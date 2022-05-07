/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    temp := head
    size := 0
    
    for temp != nil {
        size++
        temp = temp.Next
    }
    
    x := int(size/2)
    
    temp = head
    for x > 0 {
        x --
        temp = temp.Next
    }
    
    return temp
}
