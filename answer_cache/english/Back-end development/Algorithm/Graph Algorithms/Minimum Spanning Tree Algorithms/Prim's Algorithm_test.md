

1. What is the main objective of Prim's Algorithm? 

Ans: The main objective of Prim's Algorithm is to find the Minimum Spanning Tree (MST) for a given connected and weighted undirected graph. 

2. How does Prim's Algorithm work? 

Ans: Prim's Algorithm works by starting with an arbitrary vertex and adding the edge with minimum weight to the existing MST in each step until all vertices are included. 

3. What is the time complexity of Prim's Algorithm? 

Ans: The time complexity of Prim's Algorithm is O(E log V), where E is the number of edges and V is the number of vertices in the graph. 

4. Can Prim's Algorithm be used for directed graphs? 

Ans: No, Prim's Algorithm is specifically designed for undirected graphs. For directed graphs, other algorithms like Dijkstra's Algorithm or Kruskal's Algorithm can be used. 

5. What is the difference between Prim's Algorithm and Kruskal's Algorithm? 

Ans: Prim's Algorithm starts with an arbitrary vertex and gradually builds the MST by adding the edge with the minimum weight at each step. Kruskal's Algorithm, on the other hand, starts with the edges of the graph arranged in increasing order of their weight, and adds them to the MST one by one as long as they do not create a cycle.