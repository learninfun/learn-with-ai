1. Gated Recurrent Unit (GRU)是一种循环神经网络 (Recurrent Neural Network, RNN) 的一种变种，用于处理序列式资料。

2. 透过 Gate Mechanism 控制每个单元 (unit) 的讯息流动，可以有效地解决梯度消失 (vanishing gradient) 和梯度爆炸 (exploding gradient) 的问题。

3. GRU 的单元包含了更新门 (update gate)、重置门 (reset gate) 和新的候选隐藏状态 (candidate hidden state)。

4. 更新门控制了过去隐藏状态的影响程度，重置门控制了新输入和过去隐藏状态的交互作用，候选隐藏状态决定了当前时刻的隐藏状态。

5. GRU 通常用于语音识别、机器翻译、文本生成等自然语言处理任务中。

6. 可以通过调整 GRU 单元的参数和架构的设计，提高模型的性能。