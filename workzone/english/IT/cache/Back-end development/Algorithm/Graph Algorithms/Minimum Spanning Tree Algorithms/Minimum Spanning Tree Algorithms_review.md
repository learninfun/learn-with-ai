1) What is the Kruskal's algorithm for finding the minimum spanning tree of a graph?
Answer: In Kruskal's algorithm, edges are added to the minimum spanning tree in ascending order of weight until all vertices are connected.

2) How does Prim's algorithm differ from Kruskal's algorithm?
Answer: Prim's algorithm starts with a single vertex and adds edges to the tree that have minimum weight and connect to unvisited vertices. Kruskal's algorithm sorts all the edges and adds them one-by-one to form a tree.

3) What is the running time complexity of Kruskal's algorithm?
Answer: Kruskal's algorithm has a time complexity of O(E log E), where E is the number of edges in the graph.

4) What is the difference between a minimum spanning tree and a shortest path tree?
Answer: A minimum spanning tree connects all the vertices in a graph with minimum cost, while a shortest path tree connects a single source vertex to all other vertices in the graph with minimum distance.

5) Can a graph have more than one minimum spanning tree?
Answer: Yes, there can be multiple minimum spanning trees if there are edges of equal weight that can be added in different orders to form a tree.