Bitonic Array is an Array where the elements are
in ascending order immediatly followed by elements in descending order

Not Handling Duplicate cases here

eg :

4, 78, 90, 45, 23

To find the max we use the binary search approach with little modification.
get the middle element
check the previous element and the next element.

if previous element is smaller and the next element is smaller then the middle element is the greatest element

if the previous element is greater and next element is smaller move the high point to pos -1
repeat the process. (Overshot the max element pos)

if the previous element is smaller and next element is greater then move the low position to pos+1 (Under shot the max element position)


Result will be the index of the maximum element