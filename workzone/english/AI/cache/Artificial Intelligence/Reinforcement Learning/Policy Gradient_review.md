1. What is the primary advantage of using policy gradient methods compared to value-based methods?
Answer: Policy gradient methods can handle continuous action spaces, whereas value-based methods require discretization of the action space.

2. What is the role of the baseline in policy gradient methods? 
Answer: The baseline is subtracted from the estimated advantage to reduce variance and improve the performance of the algorithm.

3. What is the motivation behind using the entropy term in the policy gradient objective function?
Answer: The entropy term encourages exploration and prevents the policy from becoming too deterministic and getting stuck in suboptimal local maxima.

4. How do we compute the gradient of the policy with respect to the policy parameters in policy gradient methods?
Answer: The gradient is usually computed using the likelihood ratio trick, which involves multiplying the gradient of the log-probability of the action taken with the estimated advantage.

5. What are some common methods of optimizing the policy gradient objective function? 
Answer: Some common methods include gradient ascent using stochastic gradient descent, trust region policy optimization, and natural policy gradient methods.