package getIntersectionNode

/*
https://leetcode.com/problems/intersection-of-two-linked-lists/

The current aproach uses extra space.
 To avoid space iterate through one list and attach the tail it to the head.
 Now check for the starting point of the circular portion of List 2

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	hashMap := make(map[*ListNode]int)

	for headA != nil {
		hashMap[headA] = 1
		headA = headA.Next
	}

	for headB != nil {
		_, ok := hashMap[headB]
		if ok {
			return headB
		}

		headB = headB.Next
	}

	return nil

}
