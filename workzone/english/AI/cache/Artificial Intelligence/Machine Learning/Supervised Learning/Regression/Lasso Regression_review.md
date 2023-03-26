1. What is the purpose of Lasso Regression?
Answer: The purpose of Lasso Regression is to select a subset of independent variables that are most important in predicting the dependent variable while minimizing overfitting.

2. How does Lasso Regression differ from Ridge Regression?
Answer: Lasso Regression imposes a constraint on the sum of the absolute values of the coefficients, while Ridge Regression imposes a constraint on the sum of the squares of the coefficients. This leads to Lasso Regression having more tendency to set coefficients to zero, thus performing feature selection.

3. What is the impact of increasing the regularization parameter in Lasso Regression?
Answer: Increasing the regularization parameter in Lasso Regression leads to more coefficients being set to zero, thus increasing the amount of feature selection and decreasing the complexity of the model.

4. Can Lasso Regression be used for classification problems?
Answer: Yes, Lasso Regression can be used for classification problems by transforming the predicted values into binary outcomes.

5. Is Lasso Regression suitable for datasets with highly correlated variables?
Answer: Lasso Regression may not be suitable for datasets with highly correlated variables as it may arbitrarily select one of the variables and set the other to zero. In this case, other methods such as Elastic Net may be more appropriate.