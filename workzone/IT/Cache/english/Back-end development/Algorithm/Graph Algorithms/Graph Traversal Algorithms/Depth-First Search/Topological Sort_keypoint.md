

1. Topological sort is a technique for ordering the vertices of a directed acyclic graph (DAG).

2. It provides a linear ordering of the vertices such that for every directed edge from vertex A to vertex B, vertex A comes before vertex B in the ordering.

3. The ordering obtained from topological sort is not unique, but all valid orderings satisfy the same constraints.

4. It can be used to solve a variety of problems such as scheduling tasks, determining dependencies, and finding a feasible order of events.

5. If a DAG contains a cycle, it is impossible to define a topological ordering because the cycle will prevent any vertex from being placed before another vertex in the ordering.

6. Topological sort can be implemented using either depth-first search (DFS) or breadth-first search (BFS) algorithms.

7. The time complexity of topological sort is O(V+E), where V is the number of vertices and E is the number of edges in the graph.