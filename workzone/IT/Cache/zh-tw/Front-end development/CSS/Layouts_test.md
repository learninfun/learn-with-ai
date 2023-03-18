

1. 給定一個正整數n，輸出一個n*n的矩陣，該矩陣中每個元素的值都是以該元素的行列索引為基準的曼哈頓距離。曼哈頓距離指：在一個二維網格中，從一點到另一點的距離，是該點水平它垂直方向上的距離之和。
例如：
輸入：3
輸出：
0 1 2
1 2 3
2 3 4

2. 給定一個包含n個非負整數的數組，將其排列成一個類似螺旋矩陣的形式。例如，對於數組{1,2,3,4,5,6,7,8,9}，其排列結果如下：
5 4 3
6 1 2
7 8 9

3. 給定一個正整數n，輸出一個n*n的矩陣，該矩陣中每個元素的值都是以該元素的行列索引為基準的切比雪夫距離。切比雪夫距離指：在一個二維網格中，從一點到另一點的距離，是該點水平它垂直方向上距離的最大值。
例如：
輸入：3
輸出：
0 1 2
1 2 3
2 3 4

4. 給定一個包含n個非負整數的數組，將其排列成一個類似旋轉矩陣的形式。例如，對於數組{1,2,3,4,5,6,7,8,9}，其排列結果如下：
7 4 1
8 5 2
9 6 3

5. 給定一個網格圖，該圖由0，1兩類單元組成，0表示障礙物，1表示可以通過的空間。你需要從圖中找到兩個最遠的1之間的曼哈頓距離。例如，對於下圖，最遠1之間的曼哈頓距離為6。

0010
1010
0111
0000

答案：
1. https://www.geeksforgeeks.org/print-a-given-matrix-in-spiral-form/
2. https://leetcode.com/problems/spiral-matrix-ii/
3. https://www.geeksforgeeks.org/find-distance-source-cell-destination-cell-matrix-connected-cells/
4. https://leetcode.com/problems/rotate-image/
5. https://leetcode.com/problems/as-far-from-land-as-possible/