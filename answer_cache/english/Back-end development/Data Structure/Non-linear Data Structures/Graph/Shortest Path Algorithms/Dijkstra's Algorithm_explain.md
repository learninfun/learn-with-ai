

Dijkstra's algorithm, also known as the shortest path algorithm, is a famous algorithm for finding the shortest path between two nodes in a graph. To begin the algorithm, a source node is selected, and then the algorithm visits all neighboring nodes of the source node, assigning a tentative distance value to each one. The tentative distance value is the distance from the source node to that node. This process continues until all the nodes in the graph have been visited. Ultimately, Dijkstra's algorithm produces a table of the shortest distances from the source node to every other node in the graph.

Example: Let's find the shortest path between node A and node D given the following directed weighted graph:

![alt text](https://1.bp.blogspot.com/-ht5ik5C5p5E/WGtVJDDtRVI/AAAAAAAABj8/mxE3zMcZbesXJtpxQDLdHjKwGMJYhJLzgCLcB/s640/dijkstra%25E2%2580%2599s%2520algorithm.jpg)

1. First, we start at node A and mark the distance to itself as 0. We also set all the other distances as infinity since we don't know them yet.

![alt text](https://1.bp.blogspot.com/-7PUMG-L9uZk/WGtWfLeOOEI/AAAAAAAABkE/7vJ6W_8f0Q4-SZ9zOKhbvNKhUlDR0EQQwCLcB/s640/dijkstra-1.jpg)

2. From node A, the adjacent nodes are B and C. We update their distances with the weight of the edges between them and A.

![alt text](https://1.bp.blogspot.com/-USZ5v5-j5ng/WGtWfNHwGcI/AAAAAAAABkA/T8e2EI4WSM4AdByvBczLdg00n8Y0NwOwwCLcB/s640/dijkstra-2.jpg)

3. The node with the shortest distance is B. So we move to node B and update the distance of its adjacent nodes, C and D.

![alt text](https://1.bp.blogspot.com/-CWbOxfJno0o/WGtWfDLOgxI/AAAAAAAABj8/jrdBXjp89ZYZluGc5CcFMik5G98ge53SACLcB/s640/dijkstra-3.jpg)

4. We then move to node C, which has a smaller distance compared to D. We update the distance of its neighboring node, D.

![alt text](https://1.bp.blogspot.com/-kcWzoJf7sFc/WGtWfFw_fXI/AAAAAAAABj8/E68zp-rzkeIqbkByA8WGajXLx2-PLLZBwCLcB/s640/dijkstra-4.jpg)

5. Since we have now visited all the nodes, we have obtained the shortest distance to each node from the source node (A). Thus, the shortest path between A and D is A -> C -> D, which has a distance of 4.

![alt text](https://1.bp.blogspot.com/-1anZ9h4lHUQ/WGtWfOvAEHI/AAAAAAAABj8/nhWObgfsaIsuhRA5ZyxRldtwJz5p5c5ngCLcB/s640/dijkstra-5.jpg)