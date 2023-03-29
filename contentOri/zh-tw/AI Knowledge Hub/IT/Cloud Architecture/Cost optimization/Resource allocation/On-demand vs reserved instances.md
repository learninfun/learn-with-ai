+++
title = "On-demand vs reserved instances"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題On-demand vs reserved instances的中文問題
{{< /ask_chatgpt >}}

1. 什麼是On-demand instances 和reserved instances？
2. 在使用AWS時，什麼時候應該使用On-demand instances？何時用reserved instances比較划算？
3. 在AWS中，如何更改On-demand instances和reserved instances之間的模式？
4. 如果我儲存了幾個On-demand instances和reserved instances的圖像，是否可以將圖像從On-demand instances更改為reserved instances？
5. 在AWS中，如何節省使用On-demand instances的費用？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明On-demand vs reserved instances並舉例
{{< /ask_chatgpt >}}

On-demand和reserved instances都是Amazon Web Services (AWS)提供的虛擬機器租用模式。他們兩者的主要不同是支付方式和價格穩定性。

On-demand指的是每次你需要使用虛擬機器時，你都可以隨時按照幾個小时的計費方式支付相應的費用，不需要長期的合同約束。這種模式非常適用於短期，突發性的需求，並且不需要長時間維護使用。

相反，reserved instances則是預購虛擬機器的一種模式。通過預先簽訂一年或三年合同，你可以在保證一定時間的情況下購買一些虛擬機器，這種方式相對於On-demand來說，能夠節省一些成本。Reserved instances的費用通常是一固定的低價，不會因為佔用時間而產生劃分的分條計價。

舉個例子，假設你需要開發一個多台虛擬機器的應用程序，On-demand方式更適合你，因為你可能會隨時增減購買數量。但是，如果你已經有計劃，并且知道需要整整一年或三年使用固定數量的虛擬機器，那麼reserved instances可節省成本且保持更大的彈性。

總體而言，在選擇虛擬機器租用模式時，你需要考慮你需要多少數量，使用期限，及成本等因素。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列On-demand vs reserved instances的重點
{{< /ask_chatgpt >}}

On-demand instances：

1. 進行即時部署，可以在不同的時間點靈活地啟動或停止
2. 不需要長期簽約，可以按需進行付費
3. 價格較高，但對於不定期使用或突發性需求的情況非常適合
4. 無法做到長期花費節省，適合用於短期的應用案例

Reserved instances：

1. 需要事先預定，簽約長期使用，可減少成本
2. 比On-demand instances更便宜，預定時間越長越便宜
3. 仍可靈活部署，但需預定後按照預定的方式使用
4. 適合長期的穩定使用，例如每天、每週、每月的例行作業


重點比較：

1. On-demand instances更注重靈活性，能夠根據用戶需求及時啟動或停止，但價格較高；而Reserved instances則重視穩定性及成本控制，通常適合長期穩定使用。
2. On-demand instances適合短期的應用案例，例如突發性需求或非定期使用情況；Reserved instances則適合長期的穩定使用，例如每天、每週、每月的例行作業。
3. 兩者都具有靈活部署的特性，但在使用方式上有所不同，需要根據用戶需求選擇合適的方式。
4. 長期成本控制上，Reserved instances更有優勢，可以通過長時間的預訂來達到更佳的成本控制效果。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題On-demand vs reserved instances的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}

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

