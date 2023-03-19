

1. What is the basic principle behind Dynamic Shortest Path algorithms?
Answer: Dynamic Shortest Path algorithms aim to compute the shortest path between two nodes in a graph that changes over time. These algorithms optimize the path by taking into account both the current graph and the previously computed shortest paths.

2. What is the difference between Dijkstra's algorithm and Bellman-Ford algorithm for Dynamic Shortest Path computation?
Answer: Dijkstra's algorithm is a single-source shortest path algorithm that computes the shortest path only from a single source node to all other nodes in the graph. On the other hand, Bellman-Ford algorithm is a multi-source shortest path algorithm that computes the shortest path from all nodes in the graph to a single destination.

3. How does the concept of relaxation work in Dynamic Shortest Path algorithms?
Answer: Relaxation is the process of updating the distance estimates of the nodes in a graph after each iteration of the algorithm. In Dynamic Shortest Path algorithms, relaxation is done by comparing the distance estimate of each node with the distance estimate of its neighbor nodes. If the neighbor's distance plus the edge weight is less than the current distance estimate, it is updated.

4. How can we handle negative edge weights in Dynamic Shortest Path algorithms?
Answer: Negative edge weights can create issues with the convergence of Dynamic Shortest Path algorithms. One way to handle negative edge weights is to detect negative weight cycles, which involves running the algorithm for an additional cycle and checking for any changes in the shortest path estimate. If there are changes, it indicates the presence of a negative weight cycle.

5. What is the importance of priority queues in Dynamic Shortest Path algorithms?
Answer: Priority queues are used to store and sort the nodes based on their distance estimates during the algorithm's execution. Typically, a Min Heap is used as a priority queue to reduce the time complexity of inserting and deleting nodes. The use of priority queues ensures that the algorithm always chooses the node with the shortest distance estimate first.