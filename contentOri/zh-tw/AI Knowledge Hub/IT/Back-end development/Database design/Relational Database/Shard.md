## 習題預習
{{< ask_chatgpt >}}
給我5題Shard的問題
{{< /ask_chatgpt >}}



1. Shard是什麼？
2. Shard如何幫助分布式系統的擴展性？
3. 如何在Shard之間分配數據？
4. Shard給系統帶來的優勢有哪些？
5. 如何確保Shard之間的數據一致性？   

## 說明知識
{{< ask_chatgpt >}}
說明Shard並舉例
{{< /ask_chatgpt >}}



Shard是指將大型資料庫系統分割成多個較小的子系統，每個子系統都可以獨立地處理資料的存儲和查詢，可以提高資料庫的可伸縮性和效能。

舉例來說，假設有一家銀行的資料庫系統需要處理每日數萬筆交易紀錄，如果全部都存儲在同一個資料庫中，會影響查詢效能並且不利於資料的擴展。因此，該銀行可以將整個系統分割成多個Shard，例如按照地區或客戶類型分割，每個Shard處理自己的資料，並且可以獨立地添加或移除Shard來調整系統的容量。

另一個例子是網路搜尋引擎，例如Google，也使用了Shard來處理海量的網頁索引。Google將數據分割成多個Shard，每個Shard包含一個索引子集，這些Shard可以在不同的伺服器上運行並且可以平行處理多個搜索請求，以提高搜索效能。   

## 彙總重點
{{< ask_chatgpt >}}
條列Shard的重點
{{< /ask_chatgpt >}}



1. 分散式資料庫：Shard 是一種分散式資料庫技術，它將資料庫分成多個分片（Shard），每個分片都存儲部份資料。因此，Shard 可以分散資料庫的負載，提高資料庫的效能和可靠性。

2. 水平切割：Shard 通常是根據資料的某個特定屬性進行水平切割的，例如根據日期、地區、用戶 ID 等。這樣可以使得相關資料存放在同一個分片中，方便進行查詢和管理。

3. 高可靠性：Shard 可以為資料庫提供高可靠性，因為當某個分片出現問題時，其他分片仍然可以正常運作。此外，Shard 還可以為資料庫提供容錯能力和可擴展性。

4. 資料一致性：在使用 Shard 技術時，需要考慮如何維護多個分片之間的資料一致性。通常使用複本（Replica）機制或分布式事務（Distributed Transaction）技術實現。

5. 易於擴展：Shard 技術可以讓資料庫更容易擴展，當資料量增加時，可以輕鬆地添加新的分片來處理更多的資料。此外，分片還可以根據需要進行水平擴展和垂直擴展。

6. 高效查詢：使用 Shard 技術可以實現高效的查詢，因為相關的資料存放在同一個分片中，可以減少跨多個分片的查詢，提高查詢效率。

7. 可靠性：Shard 可靠性高，在某個健康狀態不佳的機器被發現後，它們可以自動停用，直到問題解決為止。這樣可以最小化服務中斷的風險，提高可靠性。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Shard的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1) 有兩個整數列表，找出兩個列表中共同出現的最小值。
   答案：SELECT MIN(shard1.val) FROM shard1 JOIN shard2 ON shard1.val=shard2.val;

2) 在給定的列表中找到前k個最大的元素(假設k小於列表的長度)。
   答案：SELECT val FROM shard ORDER BY val DESC LIMIT k;

3) 有一個包含重複元素的列表，找出存在超過n次的元素。
   答案：SELECT val FROM (SELECT val, count(*) as count FROM shard GROUP BY val) as val_count WHERE count>n;

4) 找出包含最多元素的重複子串(子串不需要連續)。
   答案：SELECT SUBSTRING(shard.str, start, length) as sub_str FROM (SELECT str, SUBSTRING_INDEX(SUBSTRING_INDEX(str, ',', numbers.n), ',', -1) as val, LENGTH(SUBSTRING_INDEX(SUBSTRING_INDEX(str, ',', numbers.n), ',', -1)) as length, LOCATE(SUBSTRING_INDEX(SUBSTRING_INDEX(str, ',', numbers.n), ',', -1), str)+1 as start FROM shard,(SELECT @row := @row + 1 as n FROM (select 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3) t1 CROSS JOIN (select 0 UNION ALL SELECT 1 UNION ALL SELECT 2 UNION ALL SELECT 3) t2 CROSS JOIN (select @row:=0) t3) numbers WHERE CHAR_LENGTH(str)-CHAR_LENGTH(REPLACE(str,',',''))>=numbers.n-1) as shards WHERE LENGTH(sub_str)>=2 GROUP BY sub_str ORDER BY COUNT(*) DESC LIMIT 1;

5) 找出前k個出現最多的單詞。
   答案：SELECT word, COUNT(*) as count FROM shard GROUP BY word ORDER BY count DESC LIMIT k;   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Shard的網路資料
{{< /ask_chatgpt >}}



1. "Introduction to Shard: A Scalable and Secure Blockchain Solution" - https://medium.com/hasher/create-our-own-blockchain-using-sharding-2e441f0a02df

This article provides an introduction to Shard, explaining how it is designed to address the problems of scalability and security that plague traditional blockchains. It outlines the key features of Shard and how they work together to make it a highly effective and efficient blockchain platform.

2. "How Shard can help solve blockchain scalability issues" - https://www.coingecko.com/buzz/how-shard-can-help-solve-blockchain-scalability-issues

This article focuses on the scalability benefits of the Shard blockchain. It explains how Shard's use of sharding technology enables it to handle a much larger volume of transactions than traditional blockchains. The article also explores some potential use cases for Shard, such as in supply chain management and gaming.

3. "Shard: A Brief Overview of the Innovative Blockchain Solution" - https://www.cryptopolitan.com/shard-blockchain/

This article takes a broader look at Shard, providing an overview of its key features and benefits. It explains how Shard is designed to enable greater scalability and security than traditional blockchains, and how it can be used for a variety of applications.

4. "How Sharding Enables Shard to Be Highly Scalable and Efficient" - https://www.cryptonewsz.com/how-sharding-enables-shard-to-be-highly-scalable-and-efficient/8171/

This article delves deeper into how sharding technology enables Shard to achieve its scalability and efficiency gains. It explains how the use of shards enables the network to handle a much larger volume of transactions, with each shard functioning independently of the others.

5. "Shard: The Scalable Blockchain Platform for Secure and Efficient Transactions" - https://www.investinblockchain.com/shard-scalable-blockchain-platform/

This article provides an in-depth analysis of Shard's scalability and security features. It details how Shard's sharding technology enables it to process more transactions per second than traditional blockchains, while maintaining a high level of security. The article also explores some potential applications of Shard beyond finance, such as in healthcare and logistics.   

