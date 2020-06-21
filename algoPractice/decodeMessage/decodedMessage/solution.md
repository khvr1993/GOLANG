# Solution through recursive approach

Below are some of the cases

for
| encoded message |                               decoded message                               |
| :-------------: | :-------------------------------------------------------------------------: |
|       {1}       |                                      a                                      |
|      {1,1}      |                                  {a,a},{k}                                  |
|     {1,1,1}     |                             {a,a,a},{a,k},{k,a}                             |
|    {1,1,1,2}    |                   {a,a,a,b},{a,k,b},{k,a,b},{a,a,l},{k,l}                   |
|   {1,1,1,2,3}   | {a,a,a,b,c},{a,k,b,c},{k,a,b,c},{a,a,l,c},{k,l,c},{a,a,a,w},{a,k,w},{k,a,w} |

#
If we see the logic we can say that the decodings are similar to a recursion
problem.
# Logic
- if the previous element is less than equal to 2
SUM(N-2)+SUM(N-1)
- If the previous element is greater than 2 then
SUM (N-1)

Call this logic recursively to get the solution