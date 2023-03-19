

Graph traversal algorithms are algorithms that visit each vertex (node) in a graph exactly once. The traversal algorithms are used to search and explore the graph, and they can be used to find specific vertices, determine the connectivity of vertices or to solve different problems that are modeled as graphs.

Here are some examples of graph traversal algorithms:

1. Breadth-first search (BFS) - this algorithm traverses a graph in breadth-first order. It starts at the root node and explores all vertices at the current depth level before moving on to the next depth level. BFS is often used to find the shortest path between two vertices in an unweighted graph.

Example: Imagine you want to find the shortest distance between two cities on a map, BFS algorithm can be used to explore all possible routes from the starting city to the destination city, and find the shortest path.

2. Depth-first search (DFS) - this algorithm traverses a graph in depth-first order. It starts at the root node and explores as far as possible along each branch before backtracking. DFS is often used to find strongly connected components in a directed graph.

Example: Imagine you have a maze and you want to find all the possible paths from one end to another. DFS algorithm can be used to search each possible path, and keep track of all paths that reach the end.

3. Dijkstra's Algorithm - this algorithm is used to find the shortest path between two vertices in a weighted graph. It works by maintaining a priority queue of vertices, where the priority is the current best distance to reach that vertex.

Example: Imagine you have a road network between cities with different distances, Dijkstra's Algorithm can be used to find the shortest distance to travel from one city to another. 

These are just a few examples of graph traversal algorithms, there are other algorithms such as A* algorithm, Kruskal’s algorithm, and Prim's algorithm that can be used based on the problem to be solved.