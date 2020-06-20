Quick Sort is an inplace sorting algorithm
constant extra space
Depth of recursion is logarithmic

Do a partition to get the inplace element.
use recursion to sort the left part of the inplace element and right place inplace element
Sort(a,lo,hi){
  j := Partition(a,lo,hi);
  sort(a,lo,j-1);
  sort(a,j+1,hi);
}