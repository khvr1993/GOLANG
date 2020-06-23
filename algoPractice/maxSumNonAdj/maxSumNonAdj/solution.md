We solve this based on the concept that once we see the old elements max and add it to the current max we again get max

## Approach

- Array = [5,1,1,0,5]

inclusion = previous excl + currval
exclusion will be the maximum of included or excluded value

# Solution
| Array indx | CurrentValue | Maximum | Exclusion |
| :--------- | :----------- | :------ | :-------- |
| i = 0      | 5            | 5       | 0         |
| i = 1      | 1            | 5       | 1         |
| i = 2      | 1            | 6       | 1         |
| i = 3      | 0            | 6       | 1         |
| i = 4      | 5            | 11      | 1         |