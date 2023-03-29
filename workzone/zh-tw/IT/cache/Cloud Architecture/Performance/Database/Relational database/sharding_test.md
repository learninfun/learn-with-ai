1. 什麼是sharding？請簡要解釋其用途和優點。
2. 在資料庫sharding架構中，如何實現垂直分割和水平分割？有何不同？
3. 如何設計一個sharding架構，以確保資料的一致性和可靠性？可舉例說明。
4. 應該怎麼決定shard key，並解釋如何考慮相關因素？
5. 如果shard server失效怎麼辦？請簡述應對策略。

答案：
1. sharding指將資料分散在多個節點上，以提高資料處理的效能和可擴展性。其優點包括提高系統的縱向和橫向擴展能力、提高查詢效能和吞吐量、減少單點故障等。
2. 垂直分割（vertical partitioning）是根據資料的不同屬性將資料分割成不同的表，水平分割（horizontal partitioning）是將同一類型的資料按照某個shard key值分割到不同的節點上。垂直分割強調的是資料的歸屬，水平分割強調的是資料的分散。
3. 設計一個sharding架構，需要考慮諸多因素，包括shard key的設計、shard server的部署、資料的分配和移動、資料的一致性和可靠性等。一種可行的方案是採用一致性哈希算法，將資料根據shard key映射到不同的節點上。在節點失效時，可以通過備份和自動故障轉移等方式保證資料的可靠性。
4. 決定shard key的具體策略因情況而異，可以根據資料的特點、查詢的頻率、分佈的規律等因素考慮。一般來說，shard key應具有良好的分散性、均勻性和唯一性，同時要盡量避免熱點問題和資料倾斜等問題。
5. 如果shard server失效，可以通過備份和自動故障轉移等方式來處理。在節點失效時，系統可以將該節點上的資料自動轉移至其他節點上，以確保資料的可靠性和一致性。此外，還可以通過實時監控和自動調整等手段，提高系統的穩定性和可靠性。