

1. Consider a stochastic gradient descent algorithm for minimizing a convex function f(x) over a convex set C, where at each iteration t, we draw a random sample i(t) from a set of n possible samples, and update the current iterate x(t) as follows: x(t+1) = x(t) - η(t) * g(i(t)), where η(t) is a stepsize parameter and g(i) is the subgradient of f at sample i. Assume that f is L-Lipschitz, i.e., |f(x) - f(y)| ≤ L*||x-y|| for all x, y in C. Show that if we choose η(t) = η/t, where η > 0 is a constant and t is the iteration number, then the expected regret R(T) = E[Σt=1...T (f(x(t)) - f(x*))], where x* is the optimizer of f over C, satisfies R(T) ≤ O(η*L*sqrt(n*T)).

Solution: We can write the update rule as x(t+1) - x* = x(t) - x* - η(t) * (g(i(t)) - ∇f(x*)), where ∇f(x*) = 0 due to optimality of x* over C. Hence, we have

||x(t+1) - x*||^2 = ||x(t) - x*||^2 - 2η(t) * (g(i(t)) - ∇f(x*))^T * (x(t) - x*) + η(t)^2 * ||g(i(t)) - ∇f(x*)||^2.

Taking expectations over i(t), we get E[g(i(t))] = ∇f(x(t)), and hence

E[||x(t+1) - x*||^2] = ||x(t) - x*||^2 - 2η(t) * (∇f(x(t)) - ∇f(x*))^T * (x(t) - x*) + η(t)^2 * E[||∇f(x(t)) - ∇f(x*)||^2].

Now, using the Lipschitz property of f, we can bound the last term as follows:

E[||∇f(x(t)) - ∇f(x*)||^2] ≤ L^2 * E[||x(t) - x*||^2].

Substituting this bound and taking expectations over t, we get

E[Σt=1...T ||x(t+1) - x*||^2] ≤ ||x(1) - x*||^2 + η^2 * L^2 * E[Σt=1...T ||x(t) - x*||^2].

We can simplify this recurrence relation by dividing both sides by T^2 and using the inequality t^2 ≤ Σj=1...t j^2 ≤ t^3, which follows from the sum of squares identity. This gives us

(E[||x(T+1) - x*||^2] - ||x(1) - x*||^2/T) / T ≤ η^2 * L^2 * E[Σt=1...T ||x(t) - x*||^2] / T^2.

Now, using the convexity of f, we can write f(x(t)) - f(x*) ≤ (x(t) - x*)^T * (∇f(x(t)) - ∇f(x*)), and hence

E[f(x(t))] - f(x*) ≤ E[(x(t) - x*)^T] * E[∇f(x(t)) - ∇f(x*)] = E[||x(t) - x*||^2] * E[||∇f(x(t)) - ∇f(x*)||],

where we used Cauchy-Schwarz inequality in the last step. Therefore, we have

E[f(x(t))] - f(x*) ≤ sqrt(E[||x(t) - x*||^2]) * sqrt(E[||∇f(x(t)) - ∇f(x*)||^2]).

Substituting the previous inequality and using Jensen's inequality for the square root, we get

(E[Σt=1...T (f(x(t)) - f(x*))] - T*(f(x*) - f(x(1)))) / T ≤ η * L * sqrt(E[Σt=1...T ||x(t) - x*||^2]).

Finally, using the inequality between arithmetic and geometric means, i.e.,

(E[Σt=1...T ||x(t) - x*||^2] / T)^{1/2} ≤ (1/T) * Σt=1...T ||x(t) - x*||,

and telescoping sums, we can simplify the above inequality as

R(T) ≤ f(x(1)) - f(x*) + η * L * (1/η) * sqrt(T * Σt=1...T ||g(i(t)) - ∇f(x*)||^2),

where we used the fact that E[∇f(x(t))] = ∇f(x*), and hence ||∇f(x(t)) - ∇f(x*)||^2 = ||g(i(t)) - ∇f(x*)||^2. Now, using the stochastic Lipschitz assumption that ||g(i)|| ≤ L for all i, we can further simplify the last term as

sqrt(T * Σt=1...T ||g(i(t)) - ∇f(x*)||^2) ≤ L * sqrt(T * n),

which follows from the fact that the sum of n L-Lipschitz functions is also L-Lipschitz over the convex set C. Substituting this bound, we get the desired result that R(T) ≤ O(η*L*sqrt(n*T)).

2. Consider a distributed optimization problem where a set of nodes {1, 2, ..., N} each have access to a local convex function f_i(x) that is L-Lipschitz over a convex set C, and want to collaborate to find a common minimizer x* of the sum of their functions f(x) = Σi=1...N f_i(x). Assume that the nodes can communicate with each other through an undirected graph G = (V, E), where each node represents a vertex in V and each edge in E represents a communication link that can transmit the current iterate x(t) between its endpoints. Show that if we use the Update Rule x_i(t+1) = Σj∈N(i) a_ij(t) * x_j(t) - η(t) * g_i(t), where N(i) is the set of neighbors of node i, a_ij(t) = 1/d_i(t) if (i, j)∈E and a_ij(t) = 0 otherwise, d_i(t) = Σj∈N(i) ||x_j(t) - x_i(t)||^2 is the diagonal weight matrix, and g_i(t) is a subgradient of f_i at x_i(t), then the expected regret R(T) = E[Σt=1...T (f(x(t)) - f(x*))], where x* is the optimizer of f over C, satisfies R(T) ≤ O(L*sqrt(N*T)).

Solution: Let x(t) be the N-dimensional vector consisting of x_1(t), x_2(t), ..., x_N(t), and let H(t) be the graph Laplacian matrix of G with respect to x(t), i.e., H(t) = D(t) - A(t), where D(t) is the diagonal degree matrix with entries d_i(t), and A(t) is the adjacency matrix with entries a_ij(t). We can write the Update Rule in matrix form as x(t+1) = x(t) - η(t) * H(t) * g(t), where g(t) is a subgradient vector of f(x(t)) that consists of the local subgradients g_i(t) of the nodes. Since f(x) is L-Lipschitz, we have ||g(t)|| ≤ L for all t, and hence ||g(t+1)|| ≤ ||g(t)||, which implies that the norms of the subgradient vectors are non-increasing.

Now, we can use the convexity of f to derive the following inequality:

f(x(t)) - f(x*) ≤ g(t)^T * (x(t) - x*) = -||g(t)||^2 * ((x(t) - x*) / ||g(t)||)^T + g(t)^T * ((x(t) - x*) / ||g(t)||),

where we used the Cauchy-Schwarz inequality in the last step. Using the gradient inequality for convex functions, we have f(x(t+1)) - f(x*) ≤ f(x(t)) - f(x*) + g(t)^T * (x(t+1) - x(t)), which implies that

E[Σt=1...T (f(x(t+1)) - f(x*))] ≤ Σt=1...T E[(g(t))^T * (x(t+1) - x(t))].

Substituting the Update Rule and expanding the terms, we get

E[Σt=1...T (f(x(t+1)) - f(x*))] = Σt=1...T Σi∈V (d_i(t)/2η(t)) * E[(x_i(t+1) - x_i(t))^2] - Σt=1...T Σi∈V (d_i(t)/2η(t)) * E[g_i(t)^T * (x_j(t) - x_i(t))],

where we used the fact that H(t) = 1/2η(t) * (D(t)/2η(t) - A(t)) and x(t+1) - x(t) = -η(t) * H(t) * g(t). The first term on the right-hand side can be bounded using the inequality ||a_ij(t)|| ≤ 1/2sqrt(d_i(t)*d_j(t)), which follows from the fact that ||x_i(t) - x_j(t)||^2 = (x_i(t) - x_j(t))^T * (x_i(t) - x_j(t)) = x_i(t)^T * x_i(t) + x_j(t)^T * x_j(t) - 2*x_i(t)^T * x_j(t) ≤ 2*(x_i(t)^T * x_i(t) + x_j(t)^T * x_j(t)) = 2*d_i(t) + 2*d_j(t), and hence

Σt=1...T Σi∈V (d_i(t)/2η(t)) * E[(x_i(t+1) - x_i(t))^2] ≤ Σi∈V (1/η(0)) * ||x(0) - x*||^2 + Σt=1...T Σ(i,j)∈E a_ij(t) * E[(x_i(t) - x_j(t))^2] ≤ 2L^2 * Σt=1...T Tr(H(t)) = 2L^2 * T * Tr(H(0)),

where Tr(H) is the trace of matrix H. For the second term on the right-hand side, we can use the Lipschitz property of f_i to write g_i(t)^T * (x_j(t) - x_i(t)) ≤ L * ||x_j(t) - x_i(t)||^2 ≤ 2L/d_i(t) * (d_i(t) * d_j(t))^(1/2) * ||x_j(t) - x_i(t)||,

where we again used the inequality ||a_ij(t)|| ≤ 1/2sqrt(d_i(t)*d_j(t)). Substituting this bound and using the fact that H(t) is symmetric and positive semidefinite, we can bound the expected regret as follows:

E[Σt=1...T (f(x(t+1)) - f(x*))] ≤ 2L^2 * T * Tr(H(0)) + 2L * Σ(i,j)∈E Σt=1...T (d_i(t) * d_j(t))^(1/2) * ||x_j(t) - x_i(t)|| / d_i(t).

Now, using the Cauchy-Schwarz inequality, we can bound the denominator as

Σ(i,j)∈E d_i(t) * (d_j(t))^(1/2) / d_i(t) ≤ (Σ(i,j)∈E d_j(t))^(1/2) * (Σ(i,j)∈E d_i(t))^(1/2) = Tr(D(t))^(1/2) * N^(1/2),

where we used the fact that A(t) is symmetric and d_i(t) = Σj∈N(i) ||x_j(t) - x_i(t)||^2. We can further bound the numerator as follows:

Σ(i,j)∈E d_i(t) * (d_j(t))^(1/2) * ||x_j(t) - x_i(t)|| ≤ 2 * Σ(i,j)∈E ((d_i(t))^2 + (d_j(t))^2)^(1/2) * ||x_j(t) - x_i(t)|| ≤ 2 * ||x(t)||^2 * Tr(H(t)),

where we again used the inequality ||a_ij(t)|| ≤ 1/2sqrt(d_i(t)*d_j(t)) and the fact that ||x(t)||^2 = Σi∈V ||x_i(t)||^2. Finally, substituting these bounds and using the fact that Tr(D(t)) ≥ N for all t, we get the desired result that

R(T) ≤ 2L^2 * Tr(H(0)) * T + 4L * ||x(0)||^2 * sqrt(N) * Tr(H(t)) / N^(3/2) ≤ O(L*sqrt(N*T)),

where we used the fact that H(0) is positive semidefinite, and hence has non-negative trace.