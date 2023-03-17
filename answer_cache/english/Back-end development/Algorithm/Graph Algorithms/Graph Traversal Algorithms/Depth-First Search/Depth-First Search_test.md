

1. What is the time complexity of Depth-First Search in the worst case scenario?
Answer: The time complexity of Depth-First Search in the worst case scenario is O(|V| + |E|), where |V| is the number of vertices and |E| is the number of edges in the graph.

2. What data structure is used in Depth-First Search to keep track of visited vertices?
Answer: A boolean array or hash set is used in Depth-First Search to keep track of visited vertices.

3. How does Depth-First Search differentiate between a tree edge and a back edge?
Answer: A tree edge is an edge that is visited for the first time during traversal, while a back edge is an edge that leads from a child node to an ancestor node in the traversal tree.

4. Is Depth-First Search an optimal algorithm for finding shortest path in a weighted graph?
Answer: No, Depth-First Search is not an optimal algorithm for finding shortest path in a weighted graph as it does not take into account the weights of the edges.

5. When does Depth-First Search encounter a cycle in a graph?
Answer: Depth-First Search encounters a cycle in a graph when it encounters a back edge during traversal.