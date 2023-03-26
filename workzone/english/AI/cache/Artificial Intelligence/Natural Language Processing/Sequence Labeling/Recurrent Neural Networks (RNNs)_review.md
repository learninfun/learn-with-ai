1. What is the key difference between RNNs and regular feedforward neural networks?
Answer: RNNs have feedback connections that allow information to flow from one time step to another, allowing them to work with sequential data such as time series or natural language.

2. How are the weights and biases updated during training in an RNN?
Answer: In an RNN, backpropagation through time (BPTT) is used to compute gradients of the loss function with respect to the parameters at each time step, which are then used to update the weights and biases through gradient descent.

3. What is the vanishing gradient problem, and how does it affect training of RNNs?
Answer: The vanishing gradient problem occurs when the gradients used to update the weights of an RNN are small, causing the network to converge slowly or not at all. This is a common issue in deep RNNs, where information from earlier time steps can get diluted as it passes through many layers. 

4. What are some common activation functions used in RNNs, and what are their properties?
Answer: Common activation functions used in RNNs include sigmoid, tanh, and ReLU. Sigmoid and tanh are often used for gating and activation functions because they take inputs in the range (-1, 1) and (0, 1), respectively, which can be useful for controlling the flow of information. ReLU is often used for its computational efficiency and ability to handle large-scale data. 

5. What are some applications of RNNs?
Answer: RNNs have been used in a wide range of applications including speech recognition, language modeling, machine translation, sentiment analysis, and anomaly detection. They are particularly well-suited for tasks where input data is sequential and contextual information is important.