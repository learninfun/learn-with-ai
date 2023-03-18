

Edmonds-Karp Algorithm is a variation of the Ford-Fulkerson algorithm used to find the maximum flow in a flow network. This algorithm improves the runtime complexity of the Ford-Fulkerson algorithm, which relies on picking an arbitrary path from source to sink, by finding a shortest path in the residual graph using breadth-first search.

The algorithm works as follows:

1. Initialize the flow network with capacities, a source s and a sink t.

2. Initialize the flow f to 0.

3. While there exists a path p from s to t in the residual network:

   * Find the minimum capacity residual edge c(p) along path p.
   
   * Augment flow f along path p by c(p).
   
   * Update the residual capacities of edges on path p.
   
4. Return the maximum flow f.

Here is an example:

Consider a flow network with 4 nodes and 5 edges. The capacities of the edges are given as follows:

* s -> 1: 16
* s -> 2: 13
* 1 -> 2: 10
* 2 -> 1: 4
* 1 -> 3: 12
* 2 -> 4: 14
* 3 -> 2: 9
* 3 -> t: 20
* 4 -> 3: 7
* 4 -> t: 4

The source is node s and the sink is node t. We aim to find the maximum flow in the network.

We begin by initializing the flow to 0. We then find the shortest path from s to t in the residual network using BFS, which is:

s -> 2 -> 4 -> t

The minimum capacity residual edge along this path is 4, so we augment the flow by 4. The updated flow network becomes:

* s -> 1: 16/16
* s -> 2: 13/13
* 1 -> 2: 6/10
* 2 -> 1: 4/4
* 1 -> 3: 12/12
* 2 -> 4: 10/14
* 3 -> 2: 9/9
* 3 -> t: 20/20
* 4 -> 3: 3/7
* 4 -> t: 4/4

We repeat this process until there are no more paths from s to t in the residual network. The final maximum flow is 23.