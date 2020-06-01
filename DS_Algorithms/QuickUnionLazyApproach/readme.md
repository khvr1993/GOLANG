Goes by Logic :
Take the first item and make it as child to the seconditem
p is the child to q
IMP: Always the parent nodes go as values in the array
eg:
0 1 2 3 idx
1 3 3 3 val

1 is the parent node to 0
3 is the parent node to 1
3 is the parent node to 2

      3
    2   1
          0

Let us say we have n points

Given (p q)
A[p] = A[q]

idx 0 1 2 3 4
val 0 1 2 3 4
if we have to union (2,3)
we make 3 as the parent of 2
The root of 2 is 3

we replace the array value of index 2 to 3 to say that the root of 2 is 3
---------------
A
idx 0 1 2 3 4
val 0 1 3 3 4
A[2] => 2
A[2] = A[3]

0 1 3 4
    |
    2
---------------
if we have union of 3,4

0 1 4
    |
    3
    |
    2

we make 4 as the parent of 3
and we replace the array value of index 3 to 4 to say that the root of 3 is 4