

線性資料結構是指其資料元素按照一定的順序排列，並且每個資料元素都只有一個前驅元素（第一個資料元素除外）和一個後繼元素（最後一個資料元素除外），即資料元素之間存在一對一的前後關係。常見的線性資料結構包括陣列，鏈表，佇列，堆棧等。

以下是幾種常見的線性資料結構：

1. 陣列：陣列是一個在內存中分配連續記憶體的資料結構，它通過索引來訪問和操作元素。陣列的特點是能夠快速訪問元素，但在插入和刪除操作時需要移動陣列中其他元素的位置。例如，int nums [5] = {1, 2, 3, 4, 5}。

2. 鏈表：鏈表是由節點組成的資料結構，每個節點包含資料和指向下一個節點的指針。鏈表的特點是在插入和刪除記錄時不需要移動其他元素位置，但是訪問元素時需要遍歷整個鏈表。例如，單鏈表、雙向鏈表、循環鏈表等。

3. 佇列：佇列是具有先進先出（FIFO）特點的資料結構，類似於排隊。在佇列的一端添加元素，在另一端刪除元素。例如，等待列。

4. 堆棧：堆棧是具有後進先出（LIFO）特點的資料結構，類似於一疊盤子，最後放上去的最先被取下來。在堆棧中添加元素的操作稱為推入(push)，刪除元素的操作稱為彈出(pop)。例如，瀏覽器的後退按鈕。