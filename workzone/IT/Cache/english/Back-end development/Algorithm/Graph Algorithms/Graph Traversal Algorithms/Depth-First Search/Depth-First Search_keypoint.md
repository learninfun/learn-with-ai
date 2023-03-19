

1. DFS is a graph traversal technique that visits all the nodes of a graph in a depthward motion until there is no more node left to visit.

2. This algorithm starts the visit from the root node and explores as far as possible along each branch before backtracking to the previous node.

3. In DFS, nodes are marked as visited when they are being explored.

4. If a node has multiple unvisited neighbors, the algorithm proceeds with any one of them.

5. The algorithm maintains a stack to keep track of the visited nodes during the search.

6. The time complexity of the DFS algorithm is O(V+E) where V is the number of vertices and E is the number of edges.

7. DFS can be used to detect cycles in a graph and to solve problems such as topological sorting and strongly connected components.

8. DFS is typically implemented using recursion, but it can also be implemented using a stack data structure.