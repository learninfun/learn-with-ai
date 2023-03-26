Data set splitting is the process of dividing a dataset into two or more separate subsets for training, testing, and validation purposes. The primary purpose of this method is to evaluate the performance of machine learning models using previously unseen data, helping to prevent the model from overfitting or underfitting.

For example, imagine a dataset of customer information for an online e-commerce store. The dataset contains multiple features such as age, gender, buying history, and product preferences. The store wants to build a machine learning model that predicts and recommends the products that customers are most likely to buy. 

To accomplish this, the store will divide the dataset into two or more subsets - a training set (used to train the model), a testing set (used to evaluate the trained model), and possibly a validation set (used to fine-tune the model's hyperparameters). The model then trains on the training set before making predictions on the testing set. The performance of the model is evaluated based on its accuracy in predicting the products that customers are most likely to buy, the recall rate for each class, and other standard metrics. 

By splitting the dataset, the e-commerce store can ensure their model generalizes well and performs well on unseen data. It also helps to prevent the model from becoming over-reliant on specific features, which can lead to biased models that don't represent the entire population.