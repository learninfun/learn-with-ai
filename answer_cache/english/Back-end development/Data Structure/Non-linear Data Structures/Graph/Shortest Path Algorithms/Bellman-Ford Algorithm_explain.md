

Bellman-Ford Algorithm is an algorithm used for finding the shortest path in a weighted graph, from a source vertex to all other vertices. The algorithm computes the distance of the shortest path in a dynamic fashion, i.e. based on the previous computed distance, it computes the next shortest distance until it reaches the correct distance for all vertices. The algorithm can handle both positive and negative weights, but it cannot handle negative weight cycles.

Here are the steps for the Bellman-Ford Algorithm:

1. Initialize the distance of all vertices from the source vertex to infinity, except the source vertex, which is zero.
2. For each vertex, update the distance if a smaller path exists through other vertices.
3. Repeat step 2 V-1 times (where V is the number of vertices).
4. Check for negative weight cycles, and if any are found, report and exit.

Example:

Let's consider the following graph:

![Graph](https://i.imgur.com/E7eCmJ1.jpg)

Starting from vertex 0, we can use Bellman-Ford Algorithm to compute the shortest path to all other vertices:

1. Initialize the distance of all vertices to infinity, except the source vertex, which is 0: 

| Vertex | Dist |
|--------|------|
|   0    |   0  |
|   1    |  Inf |
|   2    |  Inf |
|   3    |  Inf |
|   4    |  Inf |

2. Update distances for all vertices: 

| Vertex |      Dist     |
|--------|---------------|
|   0    |       0       |
|   1    |       10      |
|   2    |       5       |
|   3    |       15      |
|   4    |       20      |

3. Repeat step 2 V-1 times, which in this case is 4 times.
4. Final distance table: 

| Vertex |      Dist     |
|--------|---------------|
|   0    |       0       |
|   1    |       10      |
|   2    |       5       |
|   3    |       15      |
|   4    |       20      |

Thus, the shortest path from vertex 0 to all other vertices are:

- 0 to 0: 0
- 0 to 1: 0 -> 2 -> 1, distance = 5 + 5 = 10
- 0 to 2: 0 -> 2, distance = 5
- 0 to 3: 0 -> 2 -> 1 -> 3, distance = 5 + 5 + 5 = 15
- 0 to 4: 0 -> 2 -> 1 -> 3 -> 4, distance = 5 + 5 + 5 + 5 = 20