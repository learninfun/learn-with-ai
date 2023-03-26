1) What is AdaBoost?
Answer: AdaBoost (Adaptive Boosting) is a machine learning algorithm that combines multiple weak learners to form a strong learner.

2) What are weak learners in AdaBoost?
Answer: Weak learners are models that perform slightly better than chance (50% accuracy). Examples include decision stumps (decision trees with only one split) or logistic regression models.

3) How does AdaBoost work?
Answer: AdaBoost assigns weights to each training instance and trains a weak learner on the weighted data. The algorithm then adjusts the weights of the incorrectly classified instances and reruns the training process with the updated weights. This process is repeated for a fixed number of iterations, and the resulting weak learners are combined to make predictions on new data.

4) What is the role of the learning rate in AdaBoost?
Answer: The learning rate controls the contribution of each weak learner to the final prediction. A higher learning rate gives greater weight to each weak learner, which can lead to overfitting, while a lower learning rate gives more importance to the majority of correctly classified instances.

5) Can AdaBoost be used for both classification and regression?
Answer: Yes, AdaBoost can be used for both classification and regression problems. However, it is more commonly used for classification tasks.