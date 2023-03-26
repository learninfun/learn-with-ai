k-Nearest Neighbors (k-NN) is a machine learning algorithm used for classification and regression analysis. It is a non-parametric algorithm that works by finding the k number of closest data points in the input space and then predicts the class or value of the new data point based on the majority class or average value of the k nearest neighbors.

For example, consider a dataset that contains information about different types of fruits and whether they are sweet or sour. The dataset contains attributes such as color, shape, size, and taste. Now, suppose we have a new fruit that we want to classify as sweet or sour based on its attributes.

Using the k-NN algorithm, we first select the value of k, say k=3. Then, we find the three fruits in the dataset that are closest to the new fruit in terms of the attributes. The majority class of these three fruits is then considered as the predicted class for the new fruit.

For instance, if the three closest fruits are all sweet, then the new fruit is predicted to be sweet as well. On the other hand, if two of the three closest fruits are sour, and one is sweet, then the new fruit is predicted to be sour.