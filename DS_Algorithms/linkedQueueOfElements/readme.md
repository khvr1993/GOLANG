Queue Implementation using linked list.

Node will have val and next
Queue will have head and tail
When element is inserted
1. In the first element head and tail will be assigned same.
2. For second element head.next will point to tail and tail will be the new element
3. The elements are inserted as stated in 2.
4. While dequeuing the head element is returned and the the head.next will become the head
