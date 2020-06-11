Can we use bottom up merge sort here ?

Bottom Up merge Sort uses Merge functionality. But Here instead of recursion we call
merge multiple times.

Algo can be similar to


for i = 1; i < n;i=2*i{'
  for (j = 0;j <n;j=j+2*i)
  merge()
}

On first iteration we merge array of length 1
merging elements 0 and 1
(0,1)(2,3) ....

On first iteration we merge array of length 2
merging elements 0,1 & 2,3
(0,3)(4,7)

second iteration array of lengths 4


Application of counting Inversions
how far (or close) the array is from being sorted. If array is already sorted then inversion count is 0. If array is sorted in reverse order that inversion count is the maximum.