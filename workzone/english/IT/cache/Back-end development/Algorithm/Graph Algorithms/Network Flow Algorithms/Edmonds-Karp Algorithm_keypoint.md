

1. Edmonds-Karp algorithm is a variation of Ford-Fulkerson algorithm to find maximum flow in a flow network.
2. It uses Breadth-First Search (BFS) to find the shortest augmenting path from source to sink with available capacity.
3. The augmenting path is used to update the flow along the edges in the path.
4. The process is repeated until there are no more augmenting paths available in the network.
5. The time complexity of the Edmonds-Karp algorithm is O(VE^2), where V is the number of vertices and E is the number of edges in the network.
6. The algorithm guarantees finding the maximum flow in polynomial time, and it is also efficient for dense graphs.
7. Edmonds-Karp algorithm can be used to solve various optimization problems such as bipartite matching, assignment problem, and image segmentation.