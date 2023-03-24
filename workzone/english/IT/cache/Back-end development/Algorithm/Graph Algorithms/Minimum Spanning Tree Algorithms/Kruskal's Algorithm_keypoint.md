

1. Kruskal's Algorithm is used to solve minimum spanning tree (MST) problems in a weighted graph.
2. The algorithm starts by considering each vertex as a separate tree, and then it iteratively merges the trees until only one tree remains, which is the MST.
3. To merge the trees, Kruskal's Algorithm initially sorts all the edges in the graph by weight.
4. It then selects the edge with the lowest weight, and checks whether it forms a cycle or not. If the edge forms a cycle, it is discarded, else it is added to the MST.
5. This process is repeated until all the vertices are connected in a single tree, and the MST is obtained.
6. Kruskal's Algorithm is optimal and efficient, with a run-time complexity of O(E log E), where E is the number of edges in the graph.
7. It is widely used in various applications, such as network design, clustering, and machine learning.