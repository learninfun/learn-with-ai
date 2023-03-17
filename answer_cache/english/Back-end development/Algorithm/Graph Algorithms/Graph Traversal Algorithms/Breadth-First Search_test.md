

1. What is Breadth-First Search (BFS)?
Answer: BFS is a graph traversal algorithm that visits all the nodes of a graph in breadthwise manner. It starts from the root (or any arbitrary node) and visits all the neighbor nodes at the current depth before moving on to nodes at the next depth level.

2. What data structures are used in BFS?
Answer: BFS uses a Queue data structure to store the visited nodes and their neighbors. This ensures that the search is prioritized based on the order of discovery rather than the depth.

3. How does BFS guarantee the shortest path?
Answer: BFS guarantees the shortest path by exploring all the nodes at each depth level before moving on to the next level. This ensures that the first path found to a node is the shortest path from the starting node.

4. What is the time complexity of BFS?
Answer: The time complexity of BFS is O(V + E), where V represents the number of vertices or nodes and E represents the number of edges in the graph.

5. Can BFS be used to detect cycles in a graph?
Answer: Yes, BFS can be used to detect cycles in an undirected graph by keeping track of the parent node for each visited node. If a visited node has an already visited neighbor that is not its parent, then there is a cycle in the graph. However, this method does not work for directed graphs.