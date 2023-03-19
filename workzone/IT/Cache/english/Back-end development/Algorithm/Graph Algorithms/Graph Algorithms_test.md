

1. What is the shortest path between vertices A and C in the graph below, if the edges have weights as marked?

       A---1---B----4-----C
           |           /   |  
           2         2    4
           |       /      |
           D---5---E---3--F

Answer: The shortest path from A to C is A-B-E-C, with a total weight of 6.

2. What is the minimum number of colors required to color the vertices of the graph below in such a way that no two adjacent vertices have the same color?

            A------B
            |\    /|
            | \  / |
            |  \/  |
            |  /\  |
            | /  \ |
            |/    \|
            C------D

Answer: Three colors are sufficient to color the vertices in such a way that no two adjacent vertices have the same color. One possible coloring is A-red, B-green, C-blue, and D-red.

3. What is the maximum flow that can be sent from vertex S to vertex T in the network below?

            S -------2------ A -------3------ B -------4------T
             |\                  |\                  |\ 
             | \                 | \                 | \
             |  \                |  \                |  \ 
             1   4               2   1               3   2
             |   \              |   \              |   \
             |    \             |    \             |    \
             v     v            v     v            v     v
             D------3------- E -------1------ F -------2------ G

Answer: The maximum flow that can be sent from vertex S to vertex T is 5. One possible flow configuration is: 
S-A: 2 
S-D: 1 
A-B: 1 
A-E: 1 
B-T: 4 
D-E: 2 
E-F: 1 
F-G: 2 
G-T: 2 

4. In the graph below, what is the minimum cut that separates vertex S from vertex T?

            S-------3------A-------2------B
            |\                 |\                |\
            | \                | \               | \
            |  \               |  \              |  \
            |   \              |   \             |   \ 
            1  | 4              2  | 1             4  | 1
            |   /              |   /              |   /
            |  /               |  /               |  /
            | /                | /                | /
            v/                 |/                 |/
            C-------2------D-------3------T 

Answer: The minimum cut that separates vertex S from vertex T is {A,B,C,D}, with a capacity of 6. This cut separates S from T because no path from S to T crosses it. 

5. Given the following input graph, perform a topological sort and output the vertices in the order they would be visited in a depth-first search starting from vertex A. 

            A-------B      
            |       |     
            C-------D      
            |       |     
            E-------F      

Answer: One possible topological sort for this graph is: A, C, E, D, F, B. In a depth-first search starting from A, one possible visit order is: A, C, E, F, D, B.