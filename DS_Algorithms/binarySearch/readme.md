Binary Search operates as O(Logn)

take the sorted Array.
Mark Low and High points. Initially there are 0 and n-1

get the middle index.
if the middle index is > the target element make the low point to middle index +1
else make the high point middle index -1
if equal return true

if none found return false