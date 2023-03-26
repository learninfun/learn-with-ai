

1. 给定一张有向带权图，求从任意一点到任意一点的最短路径。如果两点之间没有路径，输出inf。
```
程式码如下：
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

2. 给定一张无向带权图，求图中的最小生成树。如果图不连通，则输出-1。
```
程式码如下：
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

3. 给定一张有向带权图，求从任意一点到任意一点的最短路径，如果存在负权环，输出-1。
```
程式码如下：
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
        if (d[k][k] < 0) { // 存在负权环
            return -1;
        }
    }
    return 0;
}
```

4. 给定一张有向带权图，求从每个点出发能到达的所有点之间的最小总权值和。
```
程式码如下：
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

5. 给定一张有向带权图，求从任意一点到任意一点经过正好k条边的最短路径。
```
程式码如下：
void Floyd(int k) {
    for (int l = 1; l <= k; l++) { // l表示经过l条边
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