Bag of Words Model，簡稱BOW模型，是一種常用的文本表示法。這種模型的主要想法是將文本中的單詞提取出來，然後建立一個稱為文本詞袋的容器，將單詞及其出現的次數存放在裡面。該模型表示文本時沒有考慮單詞出現的順序，因此可以有效地解決語言文本的問題。

例如以下表格就是一段文字轉換成BOW模型的結果：

|單詞|出現次數|
|---|---|
|I|1|
|love|2|
|you|1|
|and|1|
|he|1|
|like|1|

可以看到，單詞"love"出現了2次，而其他單詞都只出現了1次。這樣，我們就可以用這些數字來表示這段文本，而這種表示方法可以用於語言處理中的許多任務，如文本分類、語言識別、自然語言生成等。