1. What is a Support Vector Machine (SVM) and how does it work?
Answer: A support vector machine (SVM) is a classification algorithm that attempts to separate data into distinct classes by finding the optimal hyperplane that maximally separates the data. The hyperplane is chosen in such a way that it maximizes the margin between the two classes, with the margin being defined as the distance between the hyperplane and the closest data points.

2. How is the margin calculated in Support Vector Machines?
Answer: The margin is calculated as the shortest distance between the hyperplane and the closest data points of each class. These closest data points are known as support vectors, as they define the boundary between the classes.

3. What is kernel trick in Support Vector Machines?
Answer: Kernel trick is a way to transform non-linearly separable data into a higher-dimensional space where it can be linearly separable. SVMs use different kernel functions such as Polynomial, Radial Basis Function (RBF) and Sigmoid to transform the data according to a specific kernel function which separates the given data without mapping it to the higher dimensional space directly.

4. How is the optimal hyperplane chosen in Support Vector Machines?
Answer: The optimal hyperplane is chosen by maximizing the margin between the two classes. The objective function of the SVM is to find the hyperplane that minimizes the classification error and maximizes the distance between the hyperplane and the closest data points of each class.

5. What are some advantages of Support Vector Machines over other classification algorithms?
Answer: SVMs are often able to generalize well on small training sets because they focus on maximizing the margin between the two classes. They are also able to handle high-dimensional data well and can be used for both classification and regression tasks. SVMs are less prone to overfitting than other algorithms and do not make assumptions about the distribution of the data.