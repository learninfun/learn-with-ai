1. On-demand 和 reserved instances 有何不同之處？它們在什麼情況下最為適用？
答案：On-demand instances 是即時使用的，而 reserved instances 是預先購買的；前者十分靈活，後者可以節省成本。當負載變動較大且需要彈性時，選用 on-demand instances；當有較長時間的穩定負載時，可以選用 reserved instances 以獲得成本效益。

2. 購買 reserved instances 的種類有哪些？各種類之間有何不同？
答案：有 Standard、Convertible、Scheduled 三種。Standard 適用於大多數場景，Convertible 可以在期間內改變 instance 的 Family，Scheduled 則是針對預定的時間節點進行購買。

3. 在使用 reserved instances 時，重複使用性質的 instance 可以彼此匹配嗎？如何進行匹配？
答案：可以匹配，AWS 會自動將有效期相同且可彼此匹配的 instance 匹配起來。但如果一個 instance 的有效期只有一部分和另外一些 instance 相同，就無法匹配，此時需額外支付部分費用。

4. 購買 reserved instances 時，何時會進行付款？如何處理付款？
答案：在購買「標準」或「可轉換」的 reserved instances 時，一次性支付全部費用。在購買「計畫型」reserved instances 時，先支付預定金額，之後再支付定期付款。AWS 會在有效期內自動扣除相應費用。

5. 如何判斷購買 on-demand instances 或 reserved instances 更加划算？需要考慮哪些因素？
答案：要考慮幾個因素：預估使用時間、實際使用時間、instance 的類型、使用於哪些場景以及成本壓力等等。大多數情況下，在預估使用時間較長，且具有一定公差時，購買 reserved instances 更加划算。