# 3 way Quick Sort

- It has 3 steps
  - do partition
  - recursively call quick sort for lo - lt-1
  - recursively call quicksort for gt+1 - hi

- In this the partition maintains three iterators and 1 key
  -  iterator lt := all the elements to the left side of lt is lesser than the pivot element
  - iterator gt := all the elements to the right of gt are greater than the pivot element
  -  iterator i => which just moves across the loop
  -  pivot with which we need to compare a[i]

Starting with the loop we initialise
**lt,gt,i := lo,hi,lo**
pivot will be a[lo]
if the value of a[i] > pivot then move the gt pointer left
if the value of a[i] < pivot then move the lt pointer right
else move i right

What this is doing is making all the elements which are less than pivot to the left hand side and greater to the righ hand side. equal elements will come in middle

