卷积神经网络（CNNs）是一种人工神经网络的应用，通常用于图像分类和识别等计算机视觉任务。它的主要特点是采用了卷积层、池化层和全连接层等特征提取和分类技术，能够对图像的局部特征进行提取，并且具备平移不变性和局部关联性等特点。

例如，对于一张狗的图像进行分类，卷积神经网络会先进行卷积操作，通过一个固定大小的视窗，将图像中的每个像素与周围的像素进行运算，提取出图像的局部特征，然后再进行池化操作，将提取出来的特征压缩成一个更小的数组，接着再进行全连接层，将池化后的特征向量和分类器结合，最终得出图像分类的结果。

总之，CNNs通过卷积层、池化层和全连接层等技术，实现了对图像的特征提取和分类，是现今最为先进的图像识别算法之一。