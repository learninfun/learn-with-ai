

Best-Case Time Complexity指的是當算法面對最佳情況（即算法的輸入最符合算法所預期的狀況）時，經過該算法的運算所需的最少時間。

舉例來說，假設有一個在已排序的陣列中尋找目標值的算法，當目標值恰好在陣列的中央時，此算法的Best-Case Time Complexity為O(1)，因為只需在中央位置進行一次比較即可找到目標值。

再舉一個例子，假如有一個在大小為n的陣列中尋找最小值的算法。當最小值恰好在陣列的第一個位置時，此算法的Best-Case Time Complexity為O(1)，因為只需要一次比較即可找到最小值。而當最小值位於中間位置或者最後一個位置時，此算法的Best-Case Time Complexity為O(n)，因為需要逐一比較每一個元素。