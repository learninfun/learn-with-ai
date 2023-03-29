1.什麼是ELK Stack 6.0的主要升級？ 

2.如何從Kibana API獲取數據？

3.在Logstash中什麼是conditionals？

4.在Elasticsearch中如何使用wildcard query搜索？ 

5.如何設置Elasticsearch的cluster health status為green？

答案：
1.ELK Stack 6.0的主要升級是使用了Apache Lucene 7.0引擎，包括全新的時區支持，全新的分析器和分析器鏈，和新的分詞算法。

2.可以使用Kibana API的Search API獲取數據，或者使用Kibana API的Export API將數據匯出為JSON文件。

3.conditionals是指在Logstash中使用條件語句來執行不同的過濾或輸出操作。 

4.可以在Elasticsearch中使用wildcard query搜索，通過在搜索查詢中使用*或?通配符實現模糊搜索。

5.可以使用curl命令向Elasticsearch REST API發送一個PUT請求，將cluster.routing.allocation.enable參數設置為all即可將cluster health status設置為green。