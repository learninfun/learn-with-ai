

Topological sort is an algorithm that orders the vertices of a directed acyclic graph (DAG) in such a way that for every directed edge (u,v), vertex u comes before vertex v in the ordering. It is used to show dependencies between tasks and to schedule tasks in the correct order.

The algorithm works by selecting a vertex with no incoming edges, adding it to the sorted list, and removing its outgoing edges. This process is repeated until all vertices have been added to the list.

For example, consider the following graph:

```
A -> B -> E
|         ^
v         |
+ -> C -> D
```

A valid topological sort for this graph would be A, C, B, D, E. Starting with A, we remove its outgoing edge to B. Since B now has no incoming edges, it is added to the sort. We then remove B’s outgoing edge to E. Moving on to C, we remove its outgoing edge to D. Finally, D has no incoming edges so it is added to the sort and E is added as the final vertex in the sort.