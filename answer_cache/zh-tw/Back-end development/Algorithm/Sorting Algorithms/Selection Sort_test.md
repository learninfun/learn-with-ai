

1. 如何將Selection Sort改為遞增排序？ 

答案：只需要在交換兩個元素的地方進行細微更改即可。原本的比較操作應該是if (arr[j] > arr[min_idx])，將其改為if (arr[j] < arr[min_idx])即可。

2. Selection Sort每次交換的次數是多少？ 

答案：Selection Sort每次可以找到當前最小元素並將其放入正確的位置，這需要進行一次交換。因此，它需要進行n-1次交換。

3. 如果Selection Sort的輸入列表只有一個元素，它能正確地排序嗎？ 

答案：是的，輸入列表只有一個元素的情況下，該元素已經處於正確位置，因此Selection Sort不需要執行任何操作。

4. 如何優化Selection Sort的效率？ 

答案：可以將當前最小元素的位置存儲在變量min_idx中，以避免在每次交換時查找最小元素。此外，可以在某些情況下停止排序，例如當列表已經有序時。

5. 與其他排序算法相比，Selection Sort的時間複雜度是什麼？ 

答案：Selection Sort的時間複雜度是O(n^2)，與其他簡單的排序算法如Bubble Sort和Insertion Sort相同。然而，比較複雜的排序算法如Quick Sort和Merge Sort的時間複雜度要小得多。