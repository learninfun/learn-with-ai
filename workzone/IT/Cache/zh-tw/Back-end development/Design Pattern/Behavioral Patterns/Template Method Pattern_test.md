

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