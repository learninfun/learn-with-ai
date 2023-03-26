Edge Detection是一種用於檢測影像中主要輪廓、邊緣和細節的技術。它通常用於圖像處理、計算機視覺和機器學習等領域。

Edge Detection主要依靠計算影像中像素值的差異來識別邊緣。在邊緣位置，像素值可能會急劇變化，例如由白色轉變為黑色或由黑色轉變為白色。由此可得到邊緣的數據和位置信息。

例如，下圖是一張簡單的黑白影像，利用Sobel算子進行邊緣檢測：
![Sobel算子](https://upload.wikimedia.org/wikipedia/commons/thumb/b/b1/Mirosseni_sobel_operator_state_machine_diagram.svg/2560px-Mirosseni_sobel_operator_state_machine_diagram.svg.png)

透過Sobel算子與邊緣檢測演算法，圖像中的邊緣可以被清晰地標示出來，例如照片中的物體輪廓、建築物的線條等等。這樣的邊緣檢測技術可以用於圖像解析、自動檢測和認知等領域。