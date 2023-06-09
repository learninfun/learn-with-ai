

Observer Pattern 的重點包括：

1. 定義主題（Subject）和觀察者（Observer）的關係，主題維護所有已註冊的觀察者，當主題狀態發生變化時，通知所有觀察者。

2. 主題和觀察者可分別定義一個介面，讓實現類別繼承並實現介面方法以定義主題和觀察者的行為。

3. 觀察者可註冊於多個主題，當某一主題狀態發生變化時，觀察者即可同時收到相應的通知。

4. Observer Pattern 可以實現鬆耦合關係，主題和觀察者只需維護各自的狀態，彼此無需知道對方的實現細節。

5. 主題和觀察者之間的通信可以使用多種方式實現，如事件通知、回呼方法等。