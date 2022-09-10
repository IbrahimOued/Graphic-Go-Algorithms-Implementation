# Shell sorting

**Shell sorting** : Shell sort is a highly efficient sorting algorithm and is based on insertion sort algorithm. This algorithm avoids large shifts as in case of insertion sort, if the smaller value is to the far right and hase to be moved to the far left.

**Algorithmic result**:
The array is grouped according to a certain increment of subscripts, and the insertion of each group is sorted. As the increment decreases gradually until the increment is 1, the whole data is grouped and sorted.

1. The first sorting
`gap = array.length / 2 = 5`

![1st sorting](1st%20sorting.png)

2. The second sorting
`gap = 5 / 2 = 2`

![2nd sorting](2nd%20sorting.png)

3. The third sorting
`gap = 2/2 = 1`

![](3rd%20sorting.png)
