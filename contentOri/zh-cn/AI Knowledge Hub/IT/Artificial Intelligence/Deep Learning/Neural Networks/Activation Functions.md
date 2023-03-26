+++
title = "激活函数"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Activation Functions的中文问题
{{< /ask_chatgpt >}}

1. 什么是Sigmoid Activation Function? 它的作用是什么？
2. ReLU Activation Function是什么？为什么它比其他Activation Functions更受欢迎？
3. Tanh Activation Function与Sigmoid Activation Function之间的区别是什么？TanH的优势在哪里？
4. LeakyReLU Activation Function是什么？ LeakyReLU如何解决基本ReLu的不足？
5. Softmax Activation Function是什么？ 这个Activation Function有什么用途？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明Activation Functions并举例
{{< /ask_chatgpt >}}

Activation functions是神经网路中的一个重要元素，它决定了每个神经元的输出值。Activation functions可以将输入的信号转换成不同的形式，在深度学习领域中常常使用以下的几种Activation functions：

1. Sigmoid Function：将输入的值经由sigmoid函数转换，输出范围在0到1之间。Sigmoid函数在binary classification中的应用很广泛。

2. Tanh Function：tanh函数跟sigmoid函数比较相似，但输出值范围在-1到1之间，以及加速神经元的收敛速度。

3. ReLU Function：ReLU函数是近年来深度学习中很受欢迎的Activation functions之一，它在输入大于0的情况下直接输出，输入小于0的时候则输出0。ReLU函数可以加速神经元的收敛速度。

4. Softmax Function：用于多类别判断的神经网路中。Softmax函数将输入的多个值转换成概率分布，让神经网路可以对多个类别进行分类。

以上是一些常见的Activation functions，不同的网络模型可能会选择不同的Activation functions，根据任务的不同有时可以使用不同的Activation functions 组合。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列Activation Functions的重点
{{< /ask_chatgpt >}}

1. 激活函数可以增加神经网络的非线性能力，提高模型的拟合能力。

2. sigmoid函数是一个常见的激活函数，其输出范围在0到1之间，适用于二分类问题。

3. tanh函数是sigmoid函数的变形，取值范围在-1到1之间，函数的导数在0处最大。

4. ReLU函数非常简单直接，也是近年来最常用的激活函数之一，特点是没有负数输出，训练速度较快，但容易落入"dead ReLU"现象。

5. Leaky ReLU函数是对ReLU函数的改进，对于负数部分输出非零值，避免了"dead ReLU"现象的发生。

6. ELU函数在负数部分输出负指数，可以强制网络学习到更多的特征，远比ReLU函数表现优秀。

7. Softmax函数是一个常用的分类激活函数，用于将输出映射到0到1之间，并使得所有输出的概率和为1。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题Activation Functions的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}

1. ReLU(Rectified Linear Unit)是一種常用的Activation Function，請問在什麼情況下會使得輸出為0？
答案：當輸入小於0時，ReLU會輸出0。

2. Sigmoid是一種Activation Function，它可以將輸入轉化為0到1之間的輸出值，請問當輸入是正無限大或負無限大時，Sigmoid的輸出值為多少？
答案：當輸入是正無限大時，Sigmoid的輸出值接近於1；當輸入是負無限大時，Sigmoid的輸出值接近於0。

3. Tanh是一種常用的Activation Function，它可以將輸入轉化為-1到1之間的輸出值，請問當輸入是正無限大或負無限大時，Tanh的輸出值為多少？
答案：當輸入是正無限大時，Tanh的輸出值接近於1；當輸入是負無限大時，Tanh的輸出值接近於-1。

4. LeakyReLU是一種Activation Function，它與ReLU相似但在輸入小於0時會保持一個較小的斜率，請問LeakyReLU的斜率通常為多少？
答案：LeakyReLU的斜率通常為0.01。

5. Softmax是一種Activation Function，它通常用於多分類問題，請問Softmax的輸出值是什麼意思？
答案：Softmax的輸出值表示每個分類的概率，所有分類的概率之和等於1。   

