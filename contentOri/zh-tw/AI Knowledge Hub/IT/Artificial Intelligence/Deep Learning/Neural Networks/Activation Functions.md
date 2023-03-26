+++
title = "啟動函數"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Activation Functions的中文問題
{{< /ask_chatgpt >}}

1. 什麼是Sigmoid Activation Function? 它的作用是什麼？
2. ReLU Activation Function是什麼？為什麼它比其他Activation Functions更受歡迎？
3. Tanh Activation Function與Sigmoid Activation Function之間的區別是什麼？TanH的優勢在哪裡？
4. LeakyReLU Activation Function是什麼？ LeakyReLU如何解決基本ReLu的不足？
5. Softmax Activation Function是什麼？ 這個Activation Function有什麼用途？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Activation Functions並舉例
{{< /ask_chatgpt >}}

Activation functions是神經網路中的一個重要元素，它決定了每個神經元的輸出值。Activation functions可以將輸入的信號轉換成不同的形式，在深度學習領域中常常使用以下的幾種Activation functions：

1. Sigmoid Function：將輸入的值經由sigmoid函數轉換，輸出範圍在0到1之間。Sigmoid函數在binary classification中的應用很廣泛。

2. Tanh Function：tanh函數跟sigmoid函數比較相似，但輸出值範圍在-1到1之間，以及加速神經元的收斂速度。

3. ReLU Function：ReLU函數是近年來深度學習中很受歡迎的Activation functions之一，它在輸入大於0的情況下直接輸出，輸入小於0的時候則輸出0。ReLU函數可以加速神經元的收斂速度。

4. Softmax Function：用於多類別判斷的神經網路中。Softmax函數將輸入的多個值轉換成概率分佈，讓神經網路可以對多個類別進行分類。

以上是一些常見的Activation functions，不同的網絡模型可能會選擇不同的Activation functions，根據任務的不同有時可以使用不同的Activation functions 組合。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列Activation Functions的重點
{{< /ask_chatgpt >}}

1. 激活函数可以增加神经网络的非线性能力，提高模型的拟合能力。

2. sigmoid函数是一个常见的激活函数，其输出范围在0到1之间，适用于二分类问题。

3. tanh函数是sigmoid函数的变形，取值范围在-1到1之间，函数的导数在0处最大。

4. ReLU函数非常简单直接，也是近年来最常用的激活函数之一，特点是没有负数输出，训练速度较快，但容易落入"dead ReLU"现象。

5. Leaky ReLU函数是对ReLU函数的改进，对于负数部分输出非零值，避免了"dead ReLU"现象的发生。

6. ELU函数在负数部分输出负指数，可以强制网络学习到更多的特征，远比ReLU函数表现优秀。

7. Softmax函数是一个常用的分类激活函数，用于将输出映射到0到1之间，并使得所有输出的概率和为1。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Activation Functions的中等難度問題，並在後面列出答案
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

