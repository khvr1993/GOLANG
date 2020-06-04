Interpolation Search is a betterment of Binary Search
Works on sorted array.
But be careful as this works great only on uniformly distributed array.
Time complexity is O(log(logN))
Worst case is O(N)

For position =>

pos = low + (target-A[low])*(high-low)/(A[high]-A[low])