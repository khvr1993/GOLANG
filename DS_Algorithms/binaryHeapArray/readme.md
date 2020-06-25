# Binary Heap

## Proposition :
- a[1] is the largest key which is the root of the binary tree

## Proposition
- Can use array indices to move through the tree
- Parent of node k is at k/2
- Children of node at k are at 2k and 2k-1

## Scenario : Childs key becomes larger than parents key (SWIM)

To eliminate the violation
- Exhange key in child with key in parent.
- Repeat until heap oder is restored


## Insert into a heap (Insert)
- Add Node at end and move it up

## Parents key becomes smaller than one or both of the children (SINK)
- Exchange the key in parent with key in larger child
- repeat till heap is restored

## Delete/get  the maximum
- Exchange the root with node at end and then sink it down

