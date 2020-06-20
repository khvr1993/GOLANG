For partition the key process is

select a pivot.
set the lo to begining of the array, high to end of the array
Using the pivot as reference, if the elements are smaller than pivot,
increment i(i++), if it is greater move j left(j--)
if at any point a break happens exchange i and j.
after all the swaps exhange the jth value to the key value so that the position of the key is determined.