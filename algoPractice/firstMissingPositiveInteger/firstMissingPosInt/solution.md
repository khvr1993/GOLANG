1. Check for the existance of 1. If it is not present then we can safely return 1
2. Convert all the 0's, the negative numbers, the positive numbers which are greater than
   n to 1.
3. Create a hashmap in the array itself. Loop through array and when you flip the sign of the
  element at a[i]-1. After completion loop through the array again and return the j+1 th value
  where you encounter a positive number
