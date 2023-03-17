

Selection Sort is a simple sorting algorithm that sorts an array by repeatedly finding the minimum element from an unsorted section of the array and placing it at the beginning of the sorted section. The algorithm divides the array into two sections, one sorted and one unsorted. In each iteration, the smallest element in the unsorted section is found and swapped with the first element in the unsorted section. This process continues until the entire array is sorted.

For example, let us consider an array of numbers: [7, 2, 4, 1, 5, 3].

Step 1: The minimum value is 1. Swap index 0 (7) with index 3 (1). The array becomes [1, 2, 4, 7, 5, 3].
Step 2: The minimum value in the unsorted list (2, 4, 7, 5, 3) is 2. Swap index 1 (2) with index 4 (5). The array becomes [1, 2, 4, 7, 3, 5].
Step 3: The minimum value in the unsorted list (4, 7, 3, 5) is 3. Swap index 2 (4) with index 5 (3). The array becomes [1, 2, 3, 7, 4, 5].
Step 4: The minimum value in the unsorted list (7, 4, 5) is 4. Swap index 3 (7) with index 4 (4). The array becomes [1, 2, 3, 4, 7, 5].
Step 5: The minimum value in the unsorted list (7, 5) is 5. Swap index 4 (7) with index 5 (5). The array becomes [1, 2, 3, 4, 5, 7].

The final sorted array is [1, 2, 3, 4, 5, 7].