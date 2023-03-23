

SQL Injection（注入攻擊）是一種常見的網路安全漏洞，攻擊者透過在使用者輸入處添加惡意程式碼，從而取得系統敏感信息或執行惡意操作的一種攻擊方式。

例如，一個具有搜尋功能的網站，正確的查詢語句如下： 

```
SELECT * FROM Products WHERE Name = 'shirt';
```

而攻擊者可在搜尋框輸入以下語句： 

```
' or '1'='1
```

這種攻擊式SQL Injection漏洞可以讓攻擊者改變原本的查詢語句為： 

```
SELECT * FROM Products WHERE Name = '' or '1'='1';
```

這樣攻擊者即可獲取該網站所有產品信息。 

此外，攻擊者還可通過SQL Injection漏洞執行其他危險操作，如插入或刪除數據，甚至取得系統管理權限。因此，應當在開發過程中注意防範SQL Injection漏洞，並及時更新補丁。