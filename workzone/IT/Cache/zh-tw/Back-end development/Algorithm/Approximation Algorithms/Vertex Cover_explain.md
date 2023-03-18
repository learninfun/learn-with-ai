

Vertex Cover是一種圖論中的問題，目的是找到一個最小的集合，可以覆蓋一張無向圖中所有的邊。換句話說，就是找到一些點，讓這些點所相連的邊涵蓋了整個圖。

例如，下圖中，有一個六個節點的無向圖，其中的所有邊都用虛線標記。如果要找到一個Vertex Cover，可以選擇以下三個點：1、3和5。這三個點所連接的邊（用實線表示）可以涵蓋整個圖中的所有邊。

![Vertex Cover example](https://miro.medium.com/max/3248/1*FovPhPTOG64NKGpbv1zExQ.png)

在這個例子中，這個Vertex Cover的大小為3，因為我們只需要三個節點就可以完全涵蓋整個圖了。Vertex Cover問題是一個NP完全問題，因此通常需要使用近似算法進行求解。