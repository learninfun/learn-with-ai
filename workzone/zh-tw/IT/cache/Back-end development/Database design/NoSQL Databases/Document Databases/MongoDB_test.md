

1. 如何使用MongoDB查詢符合特定日期範圍的文檔？

答案： 我們可以使用MongoDB的日期操作符（$lt,$gt）來查詢特定日期範圍的文檔。例如，以下查詢會返回所有creation_date在2018年1月1日至2018年6月30日之間的文檔：

db.collection.find({creation_date: {$lt: new Date("2018-07-01"), $gt: new Date("2018-01-01")}});

2. 如何在MongoDB中使用聚合管道計算平均值？

答案： 我們可以使用MongoDB的$group操作符和$avg操作符來計算平均值。例如，以下聚合管道會計算所有age>20的文檔的年齡平均值：

db.collection.aggregate([{$match: {age: {$gt: 20}}}, {$group: {_id: null, avg_age: {$avg: "$age"}}}]);

3. 如何從MongoDB中查詢嵌套對象的特定值？

答案： 我們可以使用MongoDB的點操作符來查詢嵌套對象的特定值。例如，以下查詢會返回所有address.city字段為"New York"的文檔：

db.collection.find({"address.city": "New York"});

4. 如何使用MongoDB聚合管道匹配嵌套數組中的多個條件？

答案： 我們可以使用MongoDB的$match操作符和$elemMatch操作符來匹配嵌套數組中的多個條件。例如，以下聚合管道會匹配所有有至少一個標籤為"mongodb"且至少一個標籤為"database"的文檔：

db.collection.aggregate([{$match: {tags: {$elemMatch: {$in: ["mongodb", "database"]}}}}]);

5. 如何使用MongoDB的索引來提高查詢性能？

答案： 我們可以使用MongoDB的索引來提高查詢性能。要創建索引，我們可以使用db.collection.createIndex()方法。例如，以下代碼將創建一個按名字升序排列的索引：

db.collection.createIndex({name: 1});

在執行查詢時，MongoDB將使用創建的索引來定位相關文檔，從而提高查詢性能。