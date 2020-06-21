While finding the kth largest element we can think of sorting the
array and returning the k-1th value as the kth largest element.

But we do not need to sort the entire array. Using the quick sort
way first we paritition the element and find the inplace position.
With this we now have 2 sets. One set to the left hand side and
one set to the right hand side.
if the k position is > than the position of the inplace element
repeat the partition process on the right hand side.
if the k position is < than the position of the inplace element
repeat the partition process on the left hand side.
Do this until you find the k th element.