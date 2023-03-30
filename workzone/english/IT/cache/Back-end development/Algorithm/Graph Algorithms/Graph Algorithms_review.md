1) What is a minimum spanning tree, and how is it found in a graph?
Answer: A minimum spanning tree of a connected weighted graph is a subgraph that is a tree (has no cycles) and includes all the vertices, with the minimum possible total edge weight. It can be found using Prim's or Kruskal's algorithms.

2) What is a topological sort, and when is it used in graph algorithms?
Answer: A topological sort of a directed acyclic graph (DAG) is a linear ordering of its vertices such that for every directed edge (u,v), vertex u comes before vertex v in the ordering. It can be used for scheduling tasks that have dependencies or for evaluating expressions in a programming language.

3) What is a strongly connected component (SCC) of a directed graph, and how can it be found?
Answer: An SCC is a subset of vertices in a directed graph such that there is a path between every pair of vertices in the subset. It can be found using Kosaraju's algorithm, which involves two passes through the graph.

4) What is the shortest path problem in a graph, and what algorithms can be used to solve it?
Answer: The shortest path problem in a graph is finding the shortest path between two vertices, where the cost of each edge is a positive number. Dijkstra's algorithm and Bellman-Ford algorithm can solve this problem.

5) What is a maximum flow problem in a graph, and what algorithms can be used to solve it?
Answer: The maximum flow problem in a network is finding the maximum amount of flow that can be sent from a source vertex to a sink vertex, given capacity constraints on the edges. Ford-Fulkerson algorithm and Edmonds-Karp algorithm can solve this problem.