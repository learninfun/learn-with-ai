

Trie，也稱為字典樹或前綴樹，是一種數據結構，用於有效地存儲和檢索字符串。 

Trie樹特別適合於應用場景，如搜尋輸入框中的自動完成，單詞應用程序的拼寫檢查或基因組學中的基因組匹配等。

Trie樹由一組節點構成，其中每個節點代表一個字符，節點之間的連接代表字符之間的關係。每條從根節點到葉節點的路徑代表一個字符串。

下面是一個例子：

假設要存儲以下字典：apple，applaud，application，apply，book，boxing

首先，創建一個空的Trie樹，如下圖所示：

![image.png](attachment:image.png)

然後，將字典中的單詞添加到Trie樹中。每個字母都在Trie樹中表示為一個節點。下面是修改後的Trie樹：

![image-2.png](attachment:image-2.png)

Trie樹允許更快地查找字符串。以「appl」為例，Trie樹可以從根節點開始，按照「a」，「p」，「p」順序遍歷它的三個子節點。當它到達第三個節點時，它會發現它是一個單詞的結尾，因此它可以確定「appl」字符串在字典中存在。

總體來說，Trie樹是一種效率高且易於實現的數據結構，適用於許多應用程序場景。