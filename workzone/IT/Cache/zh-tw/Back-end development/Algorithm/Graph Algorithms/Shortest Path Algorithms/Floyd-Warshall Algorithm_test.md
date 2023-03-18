

1. 給定一張有向帶權圖，求從任意一點到任意一點的最短路徑。如果兩點之間沒有路徑，輸出inf。
```
程式碼如下：
void Floyd() {
    for (int k = 1; k <= n; k++) {
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                if (d[i][k] != INF && d[k][j] != INF && d[i][j] > d[i][k] + d[k][j]) {
                    d[i][j] = d[i][k] + d[k][j];
                }
            }
        }
    }
}
```

2. 給定一張無向帶權圖，求圖中的最小生成樹。如果圖不連通，則輸出-1。
```
程式碼如下：
void Floyd() {
    for (int k = 1; k <= n; k++) {
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                if (d[i][k] != INF && d[k][j] != INF && d[i][j] > d[i][k] + d[k][j]) {
                    d[i][j] = d[i][k] + d[k][j];
                }
            }
        }
    }
}
```

3. 給定一張有向帶權圖，求從任意一點到任意一點的最短路徑，如果存在負權環，輸出-1。
```
程式碼如下：
int Floyd() {
    for (int k = 1; k <= n; k++) {
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                if (d[i][k] != INF && d[k][j] != INF && d[i][j] > d[i][k] + d[k][j]) {
                    d[i][j] = d[i][k] + d[k][j];
                }
            }
        }
    }
    for (int k = 1; k <= n; k++) {
        if (d[k][k] < 0) { // 存在負權環
            return -1;
        }
    }
    return 0;
}
```

4. 給定一張有向帶權圖，求從每個點出發能到達的所有點之間的最小總權值和。
```
程式碼如下：
void Floyd() {
    for (int k = 1; k <= n; k++) {
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                if (d[i][k] != INF && d[k][j] != INF && d[i][j] > d[i][k] + d[k][j]) {
                    d[i][j] = d[i][k] + d[k][j];
                }
            }
        }
    }
    for (int i = 1; i <= n; i++) {
        int sum = 0;
        for (int j = 1; j <= n; j++) {
            if (d[i][j] != INF) {
                sum += d[i][j];
            }
        }
        printf("%d ", sum);
    }
}

5. 給定一張有向帶權圖，求從任意一點到任意一點經過正好k條邊的最短路徑。
```
程式碼如下：
void Floyd(int k) {
    for (int l = 1; l <= k; l++) { // l表示經過l條邊
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                for (int m = 1; m <= n; m++) {
                    if (d[i][m] != INF && d[m][j] != INF && d[i][j] > d[i][m] + d[m][j]) {
                        d[i][j] = d[i][m] + d[m][j];
                    }
                }
            }
        }
    }
}
```