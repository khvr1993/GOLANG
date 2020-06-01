Logic is Take the second item value and replace the values in the first item in the array. Meaning you traverse through the array and whenever
you find the matching value you replace it with the second item

A[p] = A[q]


idx 0 1 2 3 4 5
val 0 1 2 3 4 5

Union(1,2)
idx 0 1 2 3 4 5
val 0 2 2 3 4 5

Union (3,2)
idx 0 1 2 3 4 5
val 0 2 2 2 4 5

Union (2,0)
idx 0 1 2 3 4 5
val 0 0 0 0 4 5




IsConnected -> Check if both p and q points have same value
Union -> Convert the values in p[i] to q[i] only if it is not already converted
