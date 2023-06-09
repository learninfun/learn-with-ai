Pooling Layers是一種在神經網路中常用的層級，它有助於減少輸入影像大小，並在一定程度上提高模型的魯棒性。Pooling Layers通常在卷積層後使用，它會把卷積層的輸出進行下採樣，即從一個大的區域中提取出一個小的值，從而減少模型中的參數數量，加速模型訓練速度，並有效預測新數據。

例如，Max Pooling是常用的一種Pooling Layers，它會在網格中選擇最大值，並將其用作下一層的輸入。假設我們有一個2×2的下採樣矩陣，原始輸入矩陣大小為4×4，如下所示：

[4 3 8 1]

[2 6 7 2]

[1 2 5 8]

[9 4 3 6]

當我們使用2×2的Max Pooling進行下採樣時，矩陣會變成2×2的大小，如下所示：

[6 8]

[9 8]

在進行Max Pooling時，我們選擇了每個2×2矩陣中的最大值作為輸出。因此，在此例中，我們選擇了每個2×2矩陣中的最大值6、8、9、8，並將其作為新的下採樣輸出。