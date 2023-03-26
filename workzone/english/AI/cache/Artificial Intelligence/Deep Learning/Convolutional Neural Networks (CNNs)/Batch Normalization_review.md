1. What is Batch Normalization and what problem does it solve?
Answer: Batch Normalization is a technique used in deep learning to standardize input data to a neural network by normalizing it to follow standard normal distribution with zero mean and unit variance. It solves the problem of internal covariate shift, where the distribution of input data to each layer of a deep neural network changes during training. 

2. What are the two parameters that Batch Normalization introduces for each input feature and what is their purpose?
Answer: The two parameters introduced by Batch Normalization are gamma and beta. Gamma is used to scale the normalized input data, while beta is used to shift it. These parameters allow the network to learn the optimal mean and variance for the input data, and improve the model's ability to fit the data. 

3. What are the benefits of using Batch Normalization in deep learning models?
Answer: Batch Normalization can increase both the speed and accuracy of a deep learning model by reducing the internal covariate shift problem, improving model generalization, and enabling higher learning rates. It also helps to regularize the model and reduce overfitting. 

4. How does Batch Normalization work during training and prediction?
Answer: During training, Batch Normalization calculates the mean and variance of each input feature within a mini-batch of data, and uses these values to normalize the input. It then scales and shifts the normalized values using the gamma and beta parameters. During prediction, Batch Normalization uses the running mean and variance of each input feature estimated during training to normalize the input. 

5. Are there any limitations or drawbacks to using Batch Normalization in deep learning models?
Answer: Batch Normalization can increase the computational cost of training a deep learning model due to the additional calculations required. It may also introduce noise into the model due to the stochastic nature of mini-batches. Additionally, it is not always effective for smaller datasets or shallower networks, and may not be suitable for certain types of neural network architectures, such as convolutional neural networks.