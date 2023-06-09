

1. 給定一個有向加權圖，求任意兩點之間的最短距離，並找出這些最短距離中的最大值。

答案：

求出任意兩點之間的最短距離後，遍歷所有距離的值，找出其中的最大值即可。

2. 給定一個有向加權圖，求圖的直徑，即圖中最遠兩點之間的距離。

答案：

運用 Floyd-Warshall Algorithm 求出任意兩點之間的最短距離，然後遍歷所有距離的值，找出其中的最大值即可。

3. 給定一個有向加權圖，該圖邊權值為正數，求圖中所有負環。

答案：

運用 Floyd-Warshall Algorithm 求出任意兩點之間的最短距離，如果任意一個點到自己的距離是負數，則該點在負環上。

4. 給定一個有向加權圖，求任意兩點之間恰好經過一次某條特定邊的最短路徑。

答案：

用 Floyd-Warshall Algorithm 求出任意兩點之間的最短距離，然後將該特定邊設為一個很大的值，再求一次最短路徑，即可得到恰好經過該特定邊的最短路徑。

5. 給定一個無向加權圖，求任意兩點之間的最短路徑，但要求經過的邊權值的和不能超過一個給定值。

答案：

將邊權值視為花費，運用 Floyd-Warshall Algorithm 求出任意兩點之間的最短距離，然後在求最短路徑時限制經過的邊權值之和不超過給定值，可以用 Dijkstra Algorithm 實現。