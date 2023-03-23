+++
title = "模板方法模式"
weight = "10"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Template Method Pattern的中文問題
{{< /ask_chatgpt >}}



1. 什麼是Template Method Pattern？ 
2. 透過Template Method Pattern如何實現程式碼重複利用及擴展性？ 
3. 與Strategy Pattern有何不同之處？ 
4. 在哪些場景應用Template Method Pattern比較適合？ 
5. Template Method Pattern有哪些優點和缺點？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Template Method Pattern並舉例
{{< /ask_chatgpt >}}



Template Method Pattern是一種行為設計模式，其目的是將一種算法的框架定義在抽像類中，而將算法的具體步驟延遲到子類中實現。這個設計模式可用於在不修改算法結構的情況下，通過子類來擴展算法的特定步驟。

例如，我們可以想像一個製作咖啡和茶的過程。它們都具有相同的基本結構，但是具體步驟略有不同。我們可以使用Template Method Pattern來減少代碼複製並確保系統的分層結構。

首先，我們定義一個抽像類，稱之為"飲料"。這個抽像類有一個名為"prepareBeverage"的方法，它包含類似於下面的一些步驟：

- 燒開水
- 將所需的杯子放在桌子上
- 將一些基本材料加入杯子中
- 調整飲品中的一些特定成分
- 將該飲料倒入杯子中
- 向杯子中加入一些附加物，如糖、奶油或蜂蜜

其中，步驟1、2和6是相同的，我們可以在"飲料"類中定義它們。但是，步驟3、4和5是具體步驟，因此我們需要將它們定義為抽像方法。

然後，我們可以定義兩個具體類，稱之為"咖啡"和"茶"，它們繼承自"飲料"類。這些類可以實現在抽像類中定義的特定步驟，並覆蓋抽像方法以提供其在製作咖啡或茶時的特定實現。

例如，如果我們實現"咖啡"類，我們可以覆蓋"addIngredients"方法以添加咖啡豆和水。如果我們實現"茶"類，我們可以覆蓋"addIngredients"方法以添加茶葉和水。但是，兩種飲料的剩餘步驟仍然是相同的。

這就是Template Method Pattern的作用。通過將一些共同步驟定義在抽像類中，它可以為子類提供代碼復用且易於擴展的基礎。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Template Method Pattern的重點
{{< /ask_chatgpt >}}



- Template Method Pattern是一種行為模式，它定義了一個算法的骨架，將某些步驟延遲到子類實現。
- 模板方法使用了繼承機制，使得子類可以不改變算法骨架的情況下，重新定義某些步驟。
- 模板方法包含一些抽像步驟和具體步驟，並且可以定義一些鉤子方法，這些方法可以由子類實現以改變算法骨架的流程。
- 模板方法的優點包括，提高代碼的重用性、提高代碼的可維護性、封裝算法的實現細節等。
- 但是，使用模板方法也存在一些缺點，如增加了系統的類的數量、涉及到類的層次關係、可能造成性能問題等。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Template Method Pattern的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 在遊戲中，有多種不同的角色可以選擇，每種角色都有不同的技能和能力。使用Template Method Pattern設計一個角色選擇的系統，每個角色都有以下步驟：
- 選擇角色
- 顯示角色資訊
- 顯示角色能力詳情
- 選擇角色技能
- 開始遊戲

答案：
抽像類別Character：
- 選擇角色(selectCharacter方法)
- 顯示角色資訊(displayInformation方法)
- 顯示角色能力詳情(displayAbility方法)
- 選擇角色技能(selectSkill方法)
- 開始遊戲(startGame方法)

具體類別Warrior, Mage, Ranger繼承Character，並實做其方法

2. 設計一個簡單的遊戲關卡系統，遊戲分為多個關卡，每個關卡有以下步驟：
- 進入關卡
- 顯示關卡地圖
- 開始戰鬥
- 顯示戰鬥結果
- 過關

答案：
抽像類別Level：
- 進入關卡(enterLevel方法)
- 顯示關卡地圖(displayMap方法)
- 開始戰鬥(startFight方法)
- 顯示戰鬥結果(displayResult方法)
- 過關(passLevel方法)

具體類別Level1, Level2, Level3繼承Level，並實做其方法

3. 一個文檔編輯器，可以輸入不同格式的文檔，例如txt和pdf。使用Template Method Pattern設計，每次輸入一個新文檔，編輯器會執行以下步驟：
- 開啟文檔
- 讀取文檔
- 顯示文檔內容

答案：
抽像類別Document：
- 開啟文檔(openDocument方法)
- 讀取文檔(readDocument方法)
- 顯示文檔內容(displayDocument方法)

具體類別TxtDocument, PdfDocument繼承Document，並實做其方法

4. 一個網路商店的結帳系統，可以應用Template Method Pattern設計。每次結帳時，系統執行以下步驟：
- 確認訂單明細
- 計算運費
- 確認客戶付款
- 更新庫存

答案：
抽像類別CheckOut：
- 確認訂單明細(confirmOrder方法)
- 計算運費(calculateFreight方法)
- 確認客戶付款(confirmPayment方法)
- 更新庫存(updateInventory方法)

具體類別OnlineCheckOut, CODCheckOut繼承CheckOut，並實做其方法

5. 一個聊天室，使用Template Method Pattern設計。每次聊天時，系統執行以下步驟：
- 連接聊天室
- 發送訊息
- 接收訊息
- 關閉聊天室

答案：
抽像類別ChatRoom：
- 連接聊天室(connectChatRoom方法)
- 發送訊息(sendMessage方法)
- 接收訊息(receiveMessage方法)
- 關閉聊天室(closeChatRoom方法)

具體類別PublicChatRoom, PrivateChatRoom繼承ChatRoom，並實做其方法   

