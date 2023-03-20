1. 題目：有一張有向帶權圖，每條邊有一個開關，表示這條邊是否斷開。每個時刻，你可以選擇打開一個開關，使得這條邊變得通路可走；或者關閉一個開關，使得這條邊變得不可行走。求從起點到終點的最短距離。
答案：該問題可以用動態規劃和 Dijkstra 算法求解。

2. 題目：有一個有向帶權圖，每個頂點有一個權值，每個時刻你可以花費一定的代價額外增加一個頂點，或者刪除其中一個頂點，求從起點到終點的最短距離。
答案：該問題可以用 Dijkstra 算法和 prims 算法求解。

3. 題目：有一個有向帶權圖，每個頂點有一個權值，每個頂點有一條出邊指向另一個頂點，但這條邊的權值隨時間變化。在每個時刻，你可以選擇任意一個頂點，將其出邊的權值加上一個固定值，求從起點到終點的最短距離。
答案：該問題可以用 Dijkstra 算法和 Bellman-Ford 算法求解。

4. 題目：有一個有向帶權圖，每個頂點都有一條指向另一個頂點的出邊，邊權隨時間變化。在每個時刻，你可以選擇任意一個頂點，將其出邊的權值乘以一個固定值，或者除以一個固定值。求從起點到終點的最短距離。
答案：該問題可以用 Dijkstra 算法和 Bellman-Ford 算法求解。

5. 題目：給定一個起點和一些中轉點以及一個終點，求從起點到終點的一條最短路徑，並且這條路徑經過的中轉點是固定的。中轉點的位置不會改變，但其之間的距離會隨時間變化。
答案：該問題可以用 Dijkstra 算法和 Floyd 算法求解。