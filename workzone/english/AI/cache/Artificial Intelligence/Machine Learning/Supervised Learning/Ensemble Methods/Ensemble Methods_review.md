1. What is Bagging and how does it work in Ensemble Methods?
Answer: Bagging is Bootstrap Aggregating, which involves taking multiple bootstrap samples of the original dataset, building decision trees on each sample, and then combining-predictions of all trees to make a final prediction. It reduces the variance and overfitting, and it is used in random forests.

2. How is Boosting different from Bagging in Ensemble Methods?
Answer: Boosting is another ensemble method that improves the accuracy of weak learners by training them sequentially, and combining them to create a strong learner. It focuses on the training instances that were misclassified in the previous iterations, and tries to improve their weights in the following iterations.

3. What is the difference between Adaboost and Gradient Boosting in Ensemble Methods?
Answer: Adaboost is an iterative algorithm that takes weak learners and modifies the data distribution associated with each iteration based on the misclassification error of the previous iteration. Gradient Boosting, on the other hand, fits a regression model to the input data and then fits additional models to the residual errors.

4. What is the concept of Stacking in Ensemble Methods?
Answer: Stacking is a model aggregation technique, where the outputs of multiple base-level models are combined or stacked to train a meta-model or an ensemble-level model. It involves using outputs of various models as input features to the meta-model, which combines the strengths of all models and improves overall performance.

5. Can Ensemble Methods be used for unsupervised learning tasks?
Answer: Yes, Ensemble Methods can be used for unsupervised learning tasks as well. For example, clustering algorithms like K-means, can be used to create multiple clusterings, which can be combined using Ensemble Methods to improve the clustering accuracy. It is also used in anomaly detection and dimensionality reduction tasks.