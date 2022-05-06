/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

    var prev *ListNode
    for head != nil {
        temp := head.Next
        
        head.Next = prev
        prev = head
        head = temp
    }
    
    return prev
}
