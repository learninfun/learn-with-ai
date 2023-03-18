+++
title = "命令模式"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Command Pattern的問題
{{< /ask_chatgpt >}}



1. Command模式的主要目的是什麼，它通常在什麼情況下使用？
2. Command模式中的什麼元素包括Command、Receiver、Invoker和Client？
3. 在Command模式中，如何將命令封裝成對象，並在適當的時候執行它們？
4. Command模式中的Undo和Redo如何實現，這些功能可以如何幫助系統開發人員？
5. Command模式何時適合使用，什麼情況下可以使用其他設計模式來替代它？   

## 說明知識
{{< ask_chatgpt >}}
說明Command Pattern並舉例
{{< /ask_chatgpt >}}



Command Pattern是一種行為型設計模式，它提供了一種將動作封裝成物件的方式，從而允許操作與其執行的物件解耦。這種模式的核心思想是將一個操作的相關數據以及對象封裝在一起，形成一個命令對象，該對象可以用於執行操作或撤銷該操作。

例如，一個文本編輯器可以用Command Pattern來實現「撤銷」和「重做」功能。當用戶在編輯器中進行某些操作（如刪除一個字母），該操作被封裝成一個命令對象。隨後，該命令對象被存儲在一個命令歷史記錄中。如果用戶想要撤銷進行的操作，編輯器可以簡單地從命令歷史記錄中拿出最新的命令對象並執行它。如果用戶想要重做撤銷的操作，編輯器可以從命令歷史記錄中拿出上一個命令對象並執行它。

另一個例子是一個遙控器，它可以用Command Pattern來實現不同的遙控操作。例如，一個遙控器可能包含控制音量的按鈕和控制頻道的按鈕。當按下音量按鈕時，遙控器將創建一個專門用於增加音量的命令對象。隨後，該命令對象將被儲存到命令歷史紀錄中。同樣地，當按下控制頻道的按鈕時，遙控器將創建一個專門用於更改頻道的命令對象。這些命令對象可以儲存在命令歷史紀錄中，以供後續執行或撤銷。   

## 彙總重點
{{< ask_chatgpt >}}
條列Command Pattern的重點
{{< /ask_chatgpt >}}



1. Command Pattern是一種行為型設計模式。
2. Command Pattern讓你能夠將特定操作的信息從其執行中分離出來，並封裝成一個獨立的物件中。
3. Command Pattern讓你可以將特定的操作序列化、日誌記錄、取消或延遲其執行。
4. Command Pattern中的關鍵角色有Command、Invoker、Receiver和Client。
5. Command是行為請求的選擇接口，Invoker引用並調用命令，Receiver實現命令和最終操作，Client則創建一個具體的Command對象並將其傳遞給Invoker。
6. Command Pattern的優點包括解耦程式碼、易於修改、簡化操作層級和支援撤銷和恢復功能等。
7. Command Pattern的缺點包括生成大量命令物件可能會影響效能、需要額外實現的模式如果實現不好可能會產生更多的問題。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Command Pattern的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 情境描述：建立一個遙控器，它可以控制不同品牌的音響、電視和DVD播放器。使用Command Pattern設計一個遙控器，並可以任意添加、刪除和調整每個設備的指令。

答案：創建一個遙控器的Command介面，用於定義execute()方法。讓每個設備都實現Command介面，同時實現其各自的方法。最終，遙控器內部將有一個可迭代的命令列表，可以設置、刪除和調整這些命令。

2. 情境描述：設計一個遊戲，其中每個角色都有不同的技能和攻擊方法。使用Command Pattern設計一個角色控制面板，以便玩家可以輕鬆地使用這些技能和攻擊。

答案：創建一個Command介面，其中定義execute()和undo()方法。每個技能和攻擊都是一個具體的Command對象，每個角色都實現Command介面並實現其各自的方法。然後，遊戲方面可以使用這些命令在角色控制面板上創建一個命令列表。

3. 情境描述：設計一個文件操作系統，其中有一些操作如複製、貼上和刪除。使用Command Pattern設計此操作系統。

答案：創建一個Command介面，其中定義execute()和undo()方法。然後，每個命令如複製、貼上和刪除都是一個具體的Command對象。內部命令模式使用命令模式，其中具體命令可以使用遞歸檢查特定文件或文件夾中的所有文件。

4. 情境描述：如何使用Command Pattern設計一個餐廳菜單，讓服務員、廚師和收銀員可以更好地通信？

答案：為每種菜品創建一個Command介面，其中定義execute()方法。服務員將菜單項目映射到每個Command對象，並且能夠添加和刪除訂單。每個Command對象都存儲了菜品的名稱和數量，廚師可以查看這個列表並開始準備菜品。最終，收銀員將運行整個訂單列表並計算總價格。

5. 情境描述：設計一個電子商務網站，用於訂購產品和處理退貨。使用Command Pattern設計此網站。

答案：為每個操作（訂單、付款、發貨、退貨）創建一個Command介面，其中定義execute()和undo()方法。控制器可以將這些命令添加到一個可迭代列表中，並且可以反過來運行列表以取消訂單。為了處理退貨流程，可以使用不同的Command介面，其中還包括額外的方法如validate()和approve()。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Command Pattern的網路資料
{{< /ask_chatgpt >}}



1. 維基百科的Command Pattern介紹：https://zh.wikipedia.org/wiki/%E5%91%BD%E4%BB%A4%E6%A8%A1%E5%BC%8F

該頁面為中文維基百科上的Command Pattern條目，簡單介紹了Command Pattern的應用場景、結構和優缺點。

2. Refactoring Guru上的Command Pattern詳細解讀：https://refactoring.guru/design-patterns/command

該頁面為Refactoring Guru網站上的Command Pattern解讀，從多個方面、多個角度進行了詳細解讀，並給出了實際的使用案例。

3. Java Tutorials上的Command Pattern應用實例：https://docs.oracle.com/javase/tutorial/uiswing/examples/events/CommandDemoProject/src/events/CommandDemo.java

該頁面為Oracle官方Java Tutorials網站上的Command Pattern應用實例，通過一個事件操作的範例展示了Command Pattern的使用。

4. Design Patterns in Python上的Command Pattern範例代碼：https://www.giacomodebidda.com/command-pattern-python/

該頁面為Design Patterns in Python網站上的Command Pattern範例代碼，介紹了如何在Python中實現Command Pattern。

5. BitDegree上的Command Pattern入門指南：https://www.bitdegree.org/programming-interview-questions/command-pattern

該頁面為BitDegree網站上的Command Pattern入門指南，從問題設定、解決方案、代碼示例等方面進行了詳細介紹。   

