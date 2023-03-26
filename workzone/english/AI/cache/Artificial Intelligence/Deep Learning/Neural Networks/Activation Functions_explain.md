Activation functions are mathematical equations that determine the output of a neural network given input values. They are mostly used for learning purposes during the training of a neural network model. An activation function takes a linear combination of inputs from the previous layer and applies a non-linear transformation to generate the outputs of a neural network.

There are different types of activation functions, but some of the most common are:

1. Sigmoid Function: It's a non-linear function that maps any input value to a value between 0 and 1. It's mainly used in binary classification models, and the output ranges from 0 to 1.

Formula: $f(x) = \frac{1}{1 + e^{-x}}$

2. ReLU (Rectified Linear Unit): It's a non-linear function that outputs the input if it's positive, and 0 if negative. ReLU is widely used in deep learning neural networks, as it speeds the training process.

Formula: $f(x) = max(0,x)$

3. TanH (Hyperbolic Tangent): It depends on the inputs and produces an output between -1 and 1. The TanH function tends to be mostly used in the hidden layer of a neural network.

Formula: $f(x) = \frac{e^x - e^{-x}}{e^x + e^{-x}}$

Activation functions play a crucial role in the performance of deep learning models, and selecting the right activation functions can lead to improved accuracy and faster training times.