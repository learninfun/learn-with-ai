

1. What is a weighted graph? 
Answer: A weighted graph is a graph in which each edge is assigned a weight/value.

2. What is the purpose of assigning weights to edges in a graph? 
Answer: Assigning weights to edges in a graph allows the graph algorithms to calculate the shortest path, minimum spanning tree, and other useful calculations that are not feasible with unweighted graphs.

3. What is the difference between a weighted graph and a directed graph? 
Answer: A weighted graph can be either directed or undirected, but each edge has an assigned weight. In contrast, a directed graph (also known as a digraph) assigns a direction to each edge.

4. Can a weighted graph have negative edge weights? 
Answer: Yes, a weighted graph can have negative edge weights. Negative edge weights are usually used to represent scenarios where the weight of an edge represents a cost, and negative costs can occur in certain situations.

5. What is Dijkstra's algorithm? 
Answer: Dijkstra's algorithm is an algorithm used to find the shortest path between two nodes in a weighted graph. It works by maintaining a set of unvisited nodes and assigning tentative distances to each node in the graph, then selecting the node with the smallest tentative distance and updating its neighbors' tentative distances. This process is repeated until the desired endpoint is reached, and the shortest path is calculated based on the distances assigned to each node.