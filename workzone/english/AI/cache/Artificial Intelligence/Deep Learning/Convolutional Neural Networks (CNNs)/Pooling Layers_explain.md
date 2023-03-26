Pooling layers are used in Convolutional Neural Networks (CNN) to reduce the spatial dimensions (width and height) of the input volume or feature map while preserving the depth. Pooling layers are typically used after convolutional layers to downsample and extract the most important or salient features.

A pooling layer works by moving a window across the input volume or feature map and applying a function (such as max, average or sum) to each window to obtain a single output value. The size of the window and the stride (the amount by which the window moves) are hyperparameters that determine the output size of the pooling layer.

An example of a pooling layer in CNN is max pooling, which takes the maximum value within a window, ignoring all other values. For instance, a 2x2 max pooling window with a stride of 2 would divide the input volume or feature map into non-overlapping 2x2 windows and output the maximum value of each window. This results in a feature map that is half the size of the original in both width and height. Max pooling reduces the sensitivity of the network to small shifts and distortions, making the model more robust to spatial variations in the input.