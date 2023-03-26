Recurrent Neural Networks (RNNs)是一种神经网路的架构，最初被创造用来处理序列的资料，例如语言、音讯及影像等。

RNNs的特点是有一个额外的网路层，被称为hidden layer。这个hidden layer，跟一般神经网路的hidden layer不一样的地方在于：每个hidden node都会接收上一个时间点的hidden node的讯息。换句话说，hidden layer具有时序性，并且可以在处理序列资料时记忆先前预测的结果。这个过程称为记忆体(memory)或循环性(recurrence)。

以下是一个简单的RNNs例子：我们要用RNNs创建一个模型，可以翻译英文为法文的问候语。假设我们有以下的英语问候语：

"Hello, how are you?"
"Good morning, how's your day going?"
"Hi there, what's up?"

RNNs模型会被训练，输入每个单词的资料，并且在最后输出法文的翻译。在这过程中，RNNs会从前一个word node储存目前的资讯，并且随着时间记录预测翻译的过程。这样，RNNs可以在翻译每个单词时，考虑到整个语句上下文，并进行更准确的翻译。