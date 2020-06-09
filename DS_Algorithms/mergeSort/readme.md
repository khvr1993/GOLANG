MergeSort

Works recursively
Call the mergeSort itself by passing the array, copy array,low,high points

the low and high points will depend on the mid val.
mid val = (low+high)/2

we have 2 recursive calls
one for the left array and one for the righ array
and merge at the end

Algo {
  if low <= high
  {
    mergeSort(a,copy,low,mid)
    mergeSort(a,copy,mid+1,high)
    merge(a,copy,low,high,mid)
  }
}