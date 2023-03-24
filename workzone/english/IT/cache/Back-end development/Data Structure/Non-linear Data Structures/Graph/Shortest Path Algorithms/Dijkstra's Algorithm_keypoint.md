

1. Dijkstra's algorithm is used to find the shortest path in a graph.

2. The algorithm works by assigning a weight to each node in the graph.

3. It then updates the weight of each node based on the distance to adjacent nodes.

4. The algorithm starts at a specified source node and iteratively moves towards the destination node.

5. At each iteration, it selects the node with the lowest weight and calculates the weight for all adjacent nodes.

6. The path is determined by tracing the path of lowest weight from the source to the destination.

7. Dijkstra's algorithm has a time complexity of O(|V|^2), where |V| is the number of nodes in the graph.

8. However, with the use of data structures, such as a priority queue, the time complexity can be reduced to O(|E| + |V|log|V|), where |E| is the number of edges in the graph.

9. Dijkstra's algorithm can be applied to both directed and undirected graphs. 

10. It requires that the graph does not contain any negative weighted edges.