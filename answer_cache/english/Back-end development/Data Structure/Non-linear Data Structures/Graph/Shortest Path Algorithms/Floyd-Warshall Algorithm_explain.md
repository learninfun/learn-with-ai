

The Floyd-Warshall algorithm is a dynamic programming algorithm used to find the shortest path between all pairs of vertices in a weighted graph. It works by building a matrix of distances between each pair of vertices, where the entry (i, j) in the matrix represents the shortest distance between vertex i and vertex j. The algorithm then iteratively updates this matrix by adding intermediate vertices and checking if the resulting path is shorter than the one currently recorded.

Here is an example of the Floyd-Warshall algorithm:

Consider the following directed graph with weights on each edge:

```
      10       3
   0------>1------>2
   |       |      ^ |
   |7      1|5     | | 6
   v       v      | v
   3------>4------>5
      2       9
```

We want to find the shortest path between all pairs of vertices. We can represent the graph as an adjacency matrix, where the value M(i,j) represents the weight of the edge from vertex i to vertex j, or Infinity if there is no edge.

```
   M = [
        [ 0, 10, N,  7, N, N ],
        [ N,  0, 3,  1, N,  5 ],
        [ N, N,  0, N, N,  6 ],
        [ N, N, N,  0, 2, N ],
        [ N, N, N, N,  0, 9 ],
        [ N, N, N, N, N,  0 ]
       ]
```

To find the shortest path between all pairs of vertices, we initialize the matrix with the direct paths between each pair of vertices.

```
   M = [
        [ 0, 10, N,  7, N, N ],
        [ N,  0, 3,  1, N,  5 ],
        [ N, N,  0, N, N,  6 ],
        [ N, N, N,  0, 2, N ],
        [ N, N, N, N,  0, 9 ],
        [ N, N, N, N, N,  0 ]
       ]
```

Then we iterate over all intermediate vertices k and update the matrix to include the path through k if it is shorter than the previously recorded path.

```
   k = 0:
   M = [
        [ 0, 10, N,  7, N, N ],
        [ N,  0, 3,  1, N,  5 ],
        [ N, N,  0,  8, N,  6 ],
        [ N, N, N,  0, 2, N ],
        [ N, N, N, N,  0, 9 ],
        [ N, N, N, N, N,  0 ]
       ]

   k = 1:
   M = [
        [ 0, 10, 13,  7, N, 12 ],
        [ N,  0,  3,  1, N,  5 ],
        [ N, N,  0,  8, N,  6 ],
        [ N, N, N,  0, 2, 11 ],
        [ N, N, N, N,  0,  9 ],
        [ N, N, N, N, N,  0 ]
       ]

   k = 2:
   M = [
        [ 0, 10,  3,  7, N,  5 ],
        [ N,  0,  3,  1, N,  5 ],
        [ N, N,  0,  8, N,  6 ],
        [ N, N, N,  0,  2,  8 ],
        [ N, N, N, N,  0,  9 ],
        [ N, N, N, N, N,  0 ]
       ]

   k = 3:
   M = [
        [ 0, 10,  3,  7, N,  5 ],
        [ 6,  0,  3,  1, 11,  5 ],
        [ N, N,  0,  8, N,  6 ],
        [ N, N, N,  0,  2,  8 ],
        [ N, N, N, N,  0,  9 ],
        [ N, N, N, N, N,  0 ]
       ]

   k = 4:
   M = [
        [ 0, 10,  3,  7,  9,  5 ],
        [ 6,  0,  3,  1,  7,  5 ],
        [ N, N,  0,  8, 10,  6 ],
        [ N, N, N,  0,  2,  8 ],
        [ N, N, N, N,  0,  9 ],
        [ N, N, N, N, N,  0 ]
       ]

   k = 5:
   M = [
        [ 0, 10,  3,  7,  9,  5 ],
        [ 6,  0,  3,  1,  7,  5 ],
        [11, 21,  0,  8, 10,  6 ],
        [ 8, 18, 11,  0,  2,  8 ],
        [17, 27, 20, 12,  0,  9 ],
        [11, 21, 14,  6,  8,  0 ]
       ]
```

The final matrix represents the shortest path between each pair of vertices. For example, the shortest path from vertex 2 to vertex 4 is 8, and the path through vertices 2 and 1 is the shortest path.