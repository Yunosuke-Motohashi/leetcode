/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    ans := &ListNode{}
    current := ans
    carryer := 0
    for l1 != nil || l2 != nil || carryer != 0 {
        v1 := 0
        v2 := 0
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }
        tmp := v1 + v2 + carryer
        carryer = tmp / 10
        digit := tmp % 10
        current.Next = &ListNode{Val: digit}
        current = current.Next
    }
    return ans.Next

}