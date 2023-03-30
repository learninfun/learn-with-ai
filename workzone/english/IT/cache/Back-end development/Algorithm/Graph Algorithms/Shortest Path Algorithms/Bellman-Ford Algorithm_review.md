1. What is Bellman-Ford Algorithm used for?
Answer: Bellman-Ford Algorithm is used to find the shortest path from one starting node to all other nodes in a weighted graph, even if there are negative weights.

2. How does Bellman-Ford Algorithm handle negative weight edges?
Answer: Bellman-Ford Algorithm can handle negative weight edges, but it must run for the same number of iterations as the number of vertices in the graph to avoid any negative weight cycles.

3. What is the time complexity of Bellman-Ford Algorithm?
Answer: The time complexity of Bellman-Ford Algorithm is O(|V||E|), where |V| is the number of vertices in the graph and |E| is the number of edges.

4. When does Bellman-Ford Algorithm fail to find the shortest path?
Answer: Bellman-Ford Algorithm fails to find the shortest path when there is a negative weight cycle present in the graph. This is because the algorithm cannot converge on a solution due to the negative cycle.

5. How does Bellman-Ford Algorithm differ from Dijkstra's Algorithm?
Answer: Bellman-Ford Algorithm can handle graphs with negative weights and Dijkstra's Algorithm cannot. Dijkstra's Algorithm is faster when there are no negative weights present, whereas Bellman-Ford Algorithm can take longer due to the need for multiple iterations.