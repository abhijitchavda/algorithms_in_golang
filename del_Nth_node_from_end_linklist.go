/*
remove Nth node from the end of a singly linked list
*/
 type ListNode struct {
     Val int
     Next *ListNode
 }
 
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    storage:=make([]*ListNode,0)
    count:=0
    for head!=nil {
        storage=append(storage,head)
        head=head.Next
        count++
    }
    k:=count-n-1
    if k==(-1) {
        return storage[0].Next
    }
    if k<(-1) {
        return nil
    }
    storage[k].Next=storage[k].Next.Next
    return storage[0]
}