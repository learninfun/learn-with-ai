

1. Floyd-Warshall Algorithm is a dynamic programming algorithm used to find the shortest paths between all pairs of vertices in a graph.

2. It is used for computing the shortest path between all pairs of vertices in a weighted graph.

3. It works by considering all possible intermediate vertices between any two vertices in the graph.

4. It employs a matrix to store the weight of each edge between two vertices, and uses this matrix to compute the shortest distance between the vertices.

5. The algorithm runs in O(n^3) time, where n is the number of vertices in the graph.

6. It can be used to detect negative cycles in a graph.

7. The algorithm is based on the principle of relaxation, where the shortest path is progressively updated by considering all possible paths with intermediate vertices.

8. The output of the algorithm is a matrix that provides the shortest distance between any two vertices in the graph.

9. It is guaranteed to find the shortest path between any two vertices if the graph does not contain any negative cycles.