Naive Bayes is a classification algorithm that uses probabilities to predict the classification of new data. It is based on Bayes' theorem, which states that the probability of a hypothesis (in this case, a classification) is proportional to the probability of the evidence (in this case, the data) given the hypothesis.

The "naive" part of Naive Bayes refers to the assumption that the predictors are independent of each other, which simplifies the calculation of probabilities. 

An example of Naive Bayes is spam filtering. The algorithm can be trained on a dataset of emails that are labeled as either spam or not spam (ham). The algorithm uses the words in the email as predictors and calculates the probability that a new email containing those words is spam or ham. For example, if an email contains the words "viagra," "money," and "free," the algorithm may predict that it is spam with a high probability. Conversely, if an email contains the words "meeting," "schedule," and "agenda," the algorithm may predict that it is ham with a high probability. Thus, Naive Bayes is an effective tool for automatically classifying new data based on patterns in the training data.