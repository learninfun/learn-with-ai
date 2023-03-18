

Radix sort is a non-comparative sorting algorithm that sorts data by grouping elements into buckets based on their radix, which is the number of digits or symbols in the base system being used, and then sorting the elements in each bucket by the next digit, from right to left. This process is continued for all digits, resulting in a fully sorted array.

For example, given the array [170, 45, 75, 90, 802, 24, 2, 66], we can sort by first grouping elements based on their ones digit, resulting in the buckets

Bucket 0: 
Bucket 1: 170 
Bucket 2: 802, 2, 24 
Bucket 3: 
Bucket 4: 45 
Bucket 5: 75 
Bucket 6: 66 
Bucket 7: 
Bucket 8: 90 

We can then combine the buckets in order, resulting in the intermediate array [170, 802, 2, 24, 45, 75, 66, 90].

Next, we repeat the process for the tens digit, resulting in the buckets

Bucket 0: 802, 2, 24, 45, 66 
Bucket 1: 170 
Bucket 2: 
Bucket 3: 75 
Bucket 4: 
Bucket 5: 
Bucket 6: 
Bucket 7: 
Bucket 8: 90 

Combining the buckets in order gives us the sorted array [2, 24, 45, 66, 75, 90, 170, 802].