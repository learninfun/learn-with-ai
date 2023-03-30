1. What is a Minimum Spanning Tree (MST)?
Answer: A MST is a tree in a weighted graph that connects all nodes with the minimum possible total edge weight.

2. How is a MST different from a Shortest Path Tree?
Answer: A MST connects all nodes with the minimum edge weight, while a Shortest Path Tree connects a single source node to all other nodes with the minimum path weight.

3. What is the Prim's Algorithm for finding a MST?
Answer: Prim's Algorithm is a greedy algorithm that starts with an arbitrary node and repeatedly adds the lowest-weight edge that connects to an unexplored node until all nodes are connected.

4. What is the Kruskal's Algorithm for finding a MST?
Answer: Kruskal's Algorithm is a greedy algorithm that sorts edges by weight and gradually adds edges with the lowest weight that don't create a cycle until all nodes are connected.

5. Can a graph have multiple MSTs?
Answer: Yes, if there are multiple edges with the same minimum weight or if the graph has a symmetry, there may be multiple MSTs.