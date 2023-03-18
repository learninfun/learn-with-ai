

Document Databases是一種NoSQL（Not Only SQL）非關聯式資料庫，其儲存資料的方式以「文件」（Document）的形式儲存。

與傳統的關聯式資料庫不同，Document Databases沒有固定的資料表和欄位，而是使用一種稱為「集合」（Collection）的方式儲存，每個集合可以容納多個文件。

每個文件都是一個獨立的單位，可以包含任意數量的欄位，甚至是嵌套的結構。以JavaScript Object Notation（JSON）格式表示，Document Databases能夠很好地處理非結構化或半結構化的資料，如文章、新聞、照片、影片、音樂等。

常見的Document Databases包括：

1. MongoDB：目前最為流行的Document Databases，使用JSON格式與二進制格式儲存資料，支援複雜的查詢和索引。

2. Couchbase：基於CouchDB的Document Databases，支援多種插件和API，提供高可用性和可擴展性。

3. RavenDB：使用C＃開發的Document Databases，支援多種平台和語言，提供全文檢索和實時資料同步等功能。

4. CouchDB：早期的Document Databases，使用JSON格式儲存資料，支援離線同步和MapReduce。

5. Amazon DynamoDB：基於AWS的Document Databases，提供高度可擴展性和靈活的資料模型，支援多種API和SDK。

總體而言，Document Databases能夠快速處理大量的非結構化或半結構化的資料，並且提供高度的可擴展性和可用性。