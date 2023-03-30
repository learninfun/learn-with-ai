1. What is the maximum flow theorem in network flow algorithms?
Answer: The maximum flow theorem states that the maximum flow in a network equals the minimum cut in the network.

2. What is Ford-Fulkerson algorithm and how does it work?
Answer: Ford-Fulkerson algorithm is a basic algorithm for computing maximum flow in a network. It works by repeatedly finding augmenting paths from the source to the sink and adding the flow along those paths until no more augmenting paths can be found.

3. What is the difference between residual capacity and residual graph in network flow algorithms?
Answer: Residual capacity is the amount of additional flow that can be sent along an edge in a residual graph, while the residual graph is a modified version of the original network that shows the remaining capacity on each edge after flow has already been sent through it.

4. What is a minimum cut in a network flow problem?
Answer: A minimum cut in a network flow problem is an edge set whose removal disconnects the source and sink nodes and has the minimum total capacity needed to do so.

5. How do you find the maximum flow using the Edmonds-Karp Algorithm?
Answer: The Edmonds-Karp Algorithm is a variation of the Ford-Fulkerson Algorithm that uses breadth-first search to find the shortest augmenting path. It works by repeatedly finding the shortest path from the source to the sink in the residual graph and increasing flow along that path until no more augmenting paths can be found.