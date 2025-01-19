package main 


type ListNode struct {
    Val int
    Next *ListNode
}
 
func deleteDuplicates(head *ListNode) *ListNode {

	for node := head; node.Next != nil ; node = node.Next{
		if node.Val == node.Next.Val{
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

    return head
}

func main() {

}