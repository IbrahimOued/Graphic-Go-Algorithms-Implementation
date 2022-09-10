# Dichotomy binary search

1. Initialize the lowest index `low = 0` The highest index `high = scores.length-1`
2. Find the `searchValue` of the middle index `mid=(low+high)/2` `scores[mid]`
3. Compare the `scores[mid]` with `searchValue`
4. If the `scores[mid] == searchValue` print current mid index
   1. If `scores[mid] > searchValue` that the `searchValue` will be found between `low` and `mid-1`
5. And so on. Repeat step 3 until you find `searchValue` or `low >= high` to terminate the loop

Example 1 : Find the index of `searchValue = 40` in the array that has been sorted below

![Example 1](dichotomy%20binary%20search%201.png)

Example 2 : Find the index of `searchValue = 90` in the array that has been sorted below

![Example 2](dichotomy%20binary%20search%202.png)
