

Ford-Fulkerson Algorithm is an algorithm used for finding the maximum flow in a network. It is a greedy algorithm that takes an iterative approach to find the maximum flow by finding the augmenting path in a residual graph until no more path can be found.

The algorithm follows these steps:

1. Initialize the flow f to 0.

2. While there exists an augmenting path in the residual graph:

   a. Find an augmenting path in the residual graph.

   b. Calculate its bottleneck capacity.

   c. Update the flow f with the bottleneck capacity.

   d. Update the residual graph.

3. Return the flow f.

Example:

Consider the following network with source node s and sink node t:

       s ----10--→ A ----5--→ B ----12--→ t
      / \         |        |        |
     8   5        4        9        2 
    /     \      ↓        ↓        ↓ 
   C       D ----7--→ E ----6--→ F
   
The numbers on edges represent their capacities. 

To find the maximum flow from s to t using the Ford-Fulkerson algorithm, we first initialize the flow f to 0. Then, we find an augmenting path in the residual graph:

s → A → B → t, with a bottleneck capacity of 5.

We update the flow f to 5 and update the residual graph by subtracting 5 from the forward edges and adding 5 to the backward edges.

The residual graph becomes:

       s ----5--→ A ----0--→ B ----7--→ t
      / \         ↓        ↓        |
     8   5        4        9        2 
    /     \      ↓        ↓        ↓ 
   C       D ----7--→ E ----6--→ F
   
We continue finding augmenting paths and updating the flow and residual graph until no more augmenting path can be found. 

The final flow is 13, and the residual graph is:

       s ----0--→ A ----0--→ B ----10--→ t
      / \         ↓        ↓        |
     8   2        4        4        2 
    /     \      ↓        ↓        ↓ 
   C       D ----0--→ E ----0--→ F

Therefore, the maximum flow from s to t is 13.