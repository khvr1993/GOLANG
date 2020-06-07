Selection Sort.

Pick  the first element i = 0 and assign min = i
compare with all the values in the array from 1 to n-1
when you encounter an element in the array say at index j  which is less than i
make the min = j

at the end of the loop exchange the values
a[i] = a[min]
and a[min] = a[i]
 continue til the last element