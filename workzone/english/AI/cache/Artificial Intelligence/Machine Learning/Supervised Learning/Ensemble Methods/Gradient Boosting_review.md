1. What is Gradient Boosting and how does it differ from traditional boosting algorithms?
Answer: Gradient Boosting is a machine learning technique that involves building a model by iteratively adding weak learners to an ensemble, with each subsequent model learning from the errors of the previous models. It differs from traditional boosting algorithms in that it uses gradient descent to optimize the loss function, as opposed to manually adjusting weights.

2. What is the role of the learning rate in Gradient Boosting?
Answer: The learning rate determines the impact of each individual tree on the final prediction. A smaller learning rate will result in a more conservative model that takes longer to converge, while a larger learning rate will result in a more aggressive model that may overfit.

3. How do you prevent Gradient Boosting models from overfitting?
Answer: Regularization techniques such as reducing the depth of the tree, increasing the number of training samples or using early stopping can be used to prevent overfitting.

4. How does Gradient Boosting handle missing values and outliers?
Answer: Gradient Boosting imputes missing values and can handle outliers in the training data by constraining tree growth or assigning higher weights to misclassified data points.

5. What are the key hyperparameters in Gradient Boosting and how do you tune them?
Answer: The key hyperparameters in Gradient Boosting include the number of trees, learning rate, maximum depth of the tree and the minimum number of samples required for a split. These can be tuned using cross-validation techniques or grid search to identify the optimal values for each hyperparameter.