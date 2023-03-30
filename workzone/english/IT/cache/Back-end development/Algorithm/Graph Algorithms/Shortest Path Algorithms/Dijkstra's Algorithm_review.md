1. What is the time complexity of Dijkstra's Algorithm?
Answer: The time complexity of Dijkstra's Algorithm is O(E + V log V), where E is the number of edges and V is the number of vertices.

2. Can Dijkstra's Algorithm handle negative-weight edges?
Answer: No, Dijkstra's Algorithm cannot handle negative-weight edges. It assumes that all edge weights are non-negative.

3. What is the purpose of the priority queue in Dijkstra's Algorithm?
Answer: The priority queue is used to keep track of the vertices that have been visited and their distances from the source vertex. It ensures that the vertex with the smallest distance is always visited first.

4. What is the difference between Dijkstra's Algorithm and Prim's Algorithm?
Answer: Dijkstra's Algorithm is used to find the shortest path between two vertices in a weighted graph, while Prim's Algorithm is used to find the minimum spanning tree of a weighted graph. Both algorithms are similar in that they use a priority queue to select the next vertex to visit.

5. What is the initialization step in Dijkstra's Algorithm?
Answer: The initialization step sets the distance of the source vertex to 0 and the distances of all other vertices to infinity. It also adds all vertices to the priority queue.