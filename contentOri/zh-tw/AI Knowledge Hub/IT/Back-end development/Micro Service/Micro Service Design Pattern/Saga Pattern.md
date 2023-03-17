## 習題預習
{{< ask_chatgpt >}}
給我5題Saga Pattern的問題
{{< /ask_chatgpt >}}



1. 什麼是Saga Pattern？它主要用於哪些場景？
2. Saga Pattern是如何實現分佈式事務管理的？它有哪些優點和缺點？
3. 在Saga Pattern中，什麼是Compensating Action？它和Rollback有什麼區別？
4. Saga Pattern和傳統的ACID事務有什麼區別？Saga Pattern是如何保證數據一致性的？
5. 在Saga Pattern中，如果某個步驟出現故障，整個事務會如何回滾？如何處理已經執行成功的步驟？   

## 說明知識
{{< ask_chatgpt >}}
說明Saga Pattern並舉例
{{< /ask_chatgpt >}}



Saga Pattern是一種在分佈式系統中實現長流程事務的設計模式。該模式將一個複雜的事務分解成多個短期事務，這些短期事務將在一系列步驟中執行，以達到原始複雜事務的目標。

具體來說，Saga Pattern通常包括以下組件：

1. Saga：整個流程的控制器，負責協調和管理各個步驟，以確保事務的最終一致性。
2. Choreography：各個步驟之間的協作和通信，以確保數據和狀態的傳遞和同步。
3. Compensation：每個步驟的回滾機制，以確保在異常情況下能夠撤消已經執行的操作。

下面以訂單服務為例說明Saga Pattern的應用：

1. 訂單服務收到訂單請求。這是整個流程的開始。

2. 訂單服務啟動Saga，並向庫存服務發送「減少庫存」請求。此時庫存服務會將商品庫存減少相應的數量。

3. 如果庫存服務成功執行，訂單服務會向支付服務發送「扣款」請求，並將訂單狀態設置為「付款中」。

4. 如果支付服務成功執行，訂單服務將訂單狀態設置為「已付款」，向物流服務發送「發貨」請求。

5. 如果發貨成功，訂單服務將訂單狀態設置為「已完成」，完成整個流程。

但是，如果某個步驟執行失敗，整個流程將出現異常。為了確保事務的一致性，Saga Pattern會啟動補償機制。例如：

1. 如果庫存服務失敗，訂單服務會向庫存服務發送「增加庫存」請求，以恢復庫存數量。

2. 如果支付服務失敗，訂單服務會向支付服務發送「退款」請求，以撤銷扣款操作。

3. 如果發貨失敗，訂單服務會向物流服務發送「撤銷發貨」請求，以恢復商品的庫存和狀態。

綜上所述，Saga Pattern通過拆分事務，協調各個步驟，以及實現補償機制，可以應對分佈式系統中的各種異常情況，確保事務的最終一致性。   

## 彙總重點
{{< ask_chatgpt >}}
條列Saga Pattern的重點
{{< /ask_chatgpt >}}



- Saga Pattern可以用來協調分散的事務，保證系統在失敗的情況下能夠恢復正常運作
- 單獨的每個事務都是冪等的，即在重複執行時不會改變系統的狀態
- Saga是一個長時間的交互過程，它管理一系列相關的事務
- Saga Pattern具有不可撤銷性，一旦一個事務完成，它就無法被撤回
- Saga模式的實現方式有兩種： Choreography-based saga和Orchestration-based saga   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Saga Pattern的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 您正在開發一個電子商務網站，需要實現一個交付系統，該系統具有出色的回滾能力。實現一個符合Saga模式的交付系統。

答案：在實現Saga模式的交付系統中，可以首先執行所有必要的檢查和驗證。 如果一切都運行正常，則可以從訂單服務器發送交付命令。 接下來，系統會執行商品準備，配送和最終交付。 如果某個步驟遇到問題（例如商品缺貨或配送失敗），系統將回滾到之前的步驟，直到所有步驟都成功完成或者所有步驟都回滾完成。

2. 假設您正在開發一個大型分佈式應用程序，其中包含多個微服務，您需要實現一種可靠且可伸縮性的Saga模式，以確保所有微服務之間的事務一致性。該怎麼實現？

答案：為了實現可靠且可伸縮性的Saga模式，您需要劃分事務的業務邏輯，將處理相關的操作放在同一個Saga內。同時，在Saga模式中，您需要實現有關如何回滾事務的邏輯，以確保當有錯誤時，系統能夠回滾到之前的狀態。使用一些消息隊列（MQ）作為Saga模式中交流的介面，這樣可以實現微服務之間的解耦。

3. 您開發了一個訂單管理系統，該系統可以向購買商品後的用戶發送感謝信。但是，由於網路故障等因素，有些感謝信並未成功發送。請項使用Saga模式實現重複發送這些感謝信。

答案：在實現Saga模式的重複發送感謝信系統中，您需要實現一個記錄器，記錄感謝信是否成功發送。 當系統發現某個感謝信發送失敗時，它將啟動一個Saga來進行自動重試。 Saga將使用之前收集的資料並重新嘗試以發送此感謝信。 一旦再次成功後，該Saga將終止，否則它將不斷進行重新嘗試，直到成功。

4. 您正在開發一個物流系統，該系統將車輛分派給不同的司機。當有司機接收任務後，系統需要實現對車輛和司機任務的修改能力，還需要能夠回滾操作以解決可能出現的任何問題。請使用Saga模式來實現該系統。

答案：在實現Saga模式的物流系統中，您可以首先執行所需的驗證和檢查。如果一切正常，系統將創建一個Saga用於相關的車輛和司機操作。在Saga中，系統將執行車輛分配，司機任務修改和固定/解除固定車輛等操作。 如果發生任何錯誤，系統將回滾到之前的狀態，直到所有操作都成功完成，或者回滾完成。

5. 假設您正在構建一個購物車系統，其中含有多個商品。您需要實現一種Saga模式，以確保在提交購物車訂單時，所有商品都有庫存可用。如果有任何商品庫存不足，系統將回滾到之前的狀態。請列出用於實現該Saga模式的步驟。

答案：實現Saga模式的購物車系統包括以下步驟：

1. 檢查所有商品的庫存量是否充足。

2. 如果庫存量充足，則在購物車中添加所有商品。

3. 執行付款操作。

4. 訂單成功後，從庫存中扣除所有商品數量。

5. 如果任何商品缺貨，回滾操作至第一步。

6. 確認所有商品都成功下單後，將訂單提交。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Saga Pattern的網路資料
{{< /ask_chatgpt >}}



1. "Design Patterns: Saga" by Chris Richardson, InfoQ, October 31, 2018.
https://www.infoq.com/articles/saga-design-pattern/
This article provides an overview of the Saga Pattern and its implementation in a microservices architecture. It includes a code example and discusses pros and cons of the pattern.

2. "Sagas: How to Maintain Data Consistency in a Microservice Architecture" by Daniel Bryant, DZone, January 25, 2017.
https://dzone.com/articles/sagas-how-to-maintain-data-consistency-in-a-microse
This article defines the Saga Pattern and its role in maintaining data consistency in a microservice architecture. It includes a code example in Java and discusses how to handle failures in a saga.

3. "Using the Saga pattern with AWS Step Functions" by Stefano Buliani, AWS Compute Blog, February 19, 2019. 
https://aws.amazon.com/blogs/compute/using-the-saga-pattern-with-aws-step-functions/
This article explains how to use the AWS Step Functions service to implement the Saga Pattern. It includes a code example and discusses how to handle compensation logic in a saga.

4. "The Saga Pattern in Distributed Systems" by Edvin Syse, Medium, May 27, 2019.
https://medium.com/@edvinsyse/the-saga-pattern-in-distributed-systems-cc413cbb2d38
This article provides an introduction to the Saga Pattern and its role in maintaining consistency in distributed systems. It includes a code example in Python and discusses how to handle failures in a saga.

5. "Implementing Sagas in Microservices Architecture" by Abdennour Toumi, Towards Data Science, October 19, 2020.
https://towardsdatascience.com/implementing-sagas-in-microservices-architecture-a073c974a2bd
This article provides a detailed explanation of the Saga Pattern, including its advantages and disadvantages. It includes an example in Node.js and discusses best practices for implementing sagas in a microservices architecture.   

