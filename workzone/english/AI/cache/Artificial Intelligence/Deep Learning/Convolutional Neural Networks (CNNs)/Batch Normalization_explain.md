Batch Normalization is a technique used to stabilize the distribution of inputs to each layer of a neural network. It involves normalizing each batch of inputs, meaning that the mean of the inputs is subtracted and the variance is divided by a standard deviation that is learned during training. This technique helps to prevent overfitting, reduce the gradient vanishing and exploding problem, and improve the efficiency of training.

For example, let's consider a convolutional neural network (CNN) for image recognition. At each layer of the CNN, Batch Normalization can be applied to normalize the inputs. Let's say that the first layer of the CNN receives an input image of size (100x100x3). The Batch Normalization layer would calculate the mean and variance of the input values across a batch and normalize each channel across that batch. This normalization process would make the input values more comparable across different batches, making the neural network more stable and effective in learning the underlying patterns of the image data.