1. 激活函数可以增加神经网络的非线性能力，提高模型的拟合能力。

2. sigmoid函数是一个常见的激活函数，其输出范围在0到1之间，适用于二分类问题。

3. tanh函数是sigmoid函数的变形，取值范围在-1到1之间，函数的导数在0处最大。

4. ReLU函数非常简单直接，也是近年来最常用的激活函数之一，特点是没有负数输出，训练速度较快，但容易落入"dead ReLU"现象。

5. Leaky ReLU函数是对ReLU函数的改进，对于负数部分输出非零值，避免了"dead ReLU"现象的发生。

6. ELU函数在负数部分输出负指数，可以强制网络学习到更多的特征，远比ReLU函数表现优秀。

7. Softmax函数是一个常用的分类激活函数，用于将输出映射到0到1之间，并使得所有输出的概率和为1。