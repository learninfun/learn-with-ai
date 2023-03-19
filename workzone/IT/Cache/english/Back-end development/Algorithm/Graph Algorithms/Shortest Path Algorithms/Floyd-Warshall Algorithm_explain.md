

The Floyd-Warshall Algorithm is an efficient algorithm used to find the shortest path between all pairs of vertices in a weighted graph. It works by considering all possible intermediate vertices in a path, along with their combination so as to find the shortest distance path between every pair of vertices.

Let's consider the following example to demonstrate the Floyd-Warshall Algorithm:

Suppose we have the following graph with its edge weights:

```
           3 
    (0)------->(3) 
    |         /|\ 
  2 |          | 
    |          | 1 
   \|/         | 
    (2)<-------(1) 
          10
```

The matrix of the shortest path distances between each pair of vertices can be initialized as follows:

```
   0       1      2        3
0  0       Inf    Inf     3
1  Inf     0      10      Inf
2  Inf     2      0       Inf
3  Inf     Inf    Inf      0
```

Here `Inf` represents an infinite distance between two vertices that are not connected by an edge.

Now, we apply the Floyd-Warshall Algorithm to find the shortest distance between all pairs of vertices in the graph.

We update each entry in this matrix using the following formula:

```
distance[i][j] = min(distance[i][j], distance[i][k] + distance[k][j])
```

where `i`, `j`, and `k` are integers, and `distance[i][j]` represents the shortest distance between vertices `i` and `j`.

After applying the algorithm, we get the resulting matrix:

```
   0       1      2       3 
0  0       11     9       3 
1  Inf     0      10      14 
2  Inf     2      0       12 
3  Inf     Inf    Inf      0
```

As we can see, this matrix now contains the shortest distance between all pairs of vertices in the graph. For example, the shortest distance between vertices `0` and `1` is `11` and the shortest distance between vertices `2` and `3` is `12`.