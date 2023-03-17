## 習題預習
{{< ask_chatgpt >}}
給我5題RESTful API的問題
{{< /ask_chatgpt >}}



1. 什麼是RESTful API？

2. RESTful API與傳統API有何區別？

3. 什麼是URI？URI在RESTful API中扮演什麼角色？

4. RESTful API中的CRUD是什麼意思？

5. 如何保證RESTful API的安全性和可靠性？   

## 說明知識
{{< ask_chatgpt >}}
說明RESTful API並舉例
{{< /ask_chatgpt >}}



RESTful API（Representational State Transfer API）是一种基于HTTP协议、遵循REST原则的Web API，可以用来让不同的软件应用程序之间通讯，提供资源和操作。

REST原则包括：使用URI描述资源、使用HTTP动词描述操作、使用HTTP状态码返回结果、使用Hypermedia作为应用程序状态转移的引擎等。

以下是一个简单的例子，说明如何创建一个RESTful API：

1. 为资源设计一个唯一的URI
例如：https://www.example.com/products/12345，其中12345是产品的唯一标识符。

2. 使用HTTP动词表示操作
例如，使用GET方法检索产品，使用PUT方法更新产品，使用DELETE方法删除产品。

3. 使用HTTP状态码返回结果
例如，当成功检索到资源时，返回200 OK状态码。如果请求的资源不存在，返回404 Not Found状态码。

4. 使用Hypermedia引擎
Hypermedia是一种用于描述应用程序状态的格式，可以表达所有可能从当前状态进行的状态转移。例如，其他可用的资源可以通过链接提供，从而帮助客户端进行导航和发现。

总的来说，RESTful API是一种非常灵活和可扩展的API设计模式，可以与各种不同的客户端和服务器技术一起使用。其优点包括相对简单易用、易于扩展、面向资源等特点。常见的RESTful API包括Twitter API、GitHub API等。   

## 彙總重點
{{< ask_chatgpt >}}
條列RESTful API的重點
{{< /ask_chatgpt >}}



1. 基於資源：RESTful API是基於資源的設計，每一個資源都有一個獨立的URI，表示用戶可以通過URI訪問資源。

2. 請求方法：使用HTTP請求方法（GET, POST, PUT, DELETE）來執行操作。

3. 無狀態：RESTful API是無狀態的，每一個請求都是獨立的，伺服器不會保留任何用戶的狀態。

4. 資源表示：用戶可以通過請求URI獲取資源的資訊，通常使用JSON等格式來表示資源。

5. HATEOAS：Hypermedia as the Engine of Application State，RESTful API通常提供超媒體鏈接，讓用戶可以通過連結來訪問相關資源。

6. 安全性：RESTful API需要實現相應的安全機制，包括身份驗證、授權、加密等。

7. 可測試性：RESTful API需要提供測試工具和文檔，方便使用者對API進行測試和開發。

8. 可擴展性：RESTful API需要支持擴展，可以通过增加資源和URI等方式擴展API的功能。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題RESTful API的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 如何在 RESTful API 中实现身份验证和授权？在 API 响应中包含哪些信息？

答：RESTful API 的身份验证和授权可通过 JWT（JSON Web Token）进行实现。在 API 响应中，通常包含以下信息：
- 状态代码（状态代码指示请求成功与否，例如：200 表示请求成功，401 表示未授权等）
- 消息（即响应的文本信息，例如：成功响应、错误消息等）

此外，还可以在响应头部中包含一些特定的信息，例如认证令牌、缓存控制当前用户等信息。

2. 如何在 RESTful API 中支持分页？

答：RESTful API 中支持分页通常需要在 API 调用中包含以下信息：
- limit：返回的记录数量限制；
- offset：返回记录的偏移量。

例如，下面的 URL 可用于获得返回前 20 条记录：
`https://example.com/api/products?limit=20&offset=0`

3. 如何在 RESTful API 中处理文件上传？

答：文件上传通常使用 POST 方法，并将文件作为路径参数或表单数据进行提交。在许多情况下，建议使用 POST 方法，因为它支持更大的文件上传。

4. 如何在 RESTful API 中实现缓存控制？

答：缓存控制可以通过 HTTP 响应标头中的 ETag 和 Last-Modified 字段来实现。当客户端发出请求时，服务器将比较这些字段，以确定是否需要更新缓存数据或返回缓存数据。

例如，下面的示例中，服务器可向客户端提供 ETag 标头：
`ETag: "686897696a7c876b7e"`

当客户端再次发起请求时，它可以在 If-None-Match 标头中提供 ETag 标识符，以便服务器只在标识符不匹配时返回新数据：
`If-None-Match: "686897696a7c876b7e"`

5. 如何在 RESTful API 中实现搜索功能？

答：搜索功能可以使用 GET 方法和查询参数来实现。查询参数指的是在 URL 中添加搜索参数，例如：
`https://example.com/api/products?search=keyword`

在服务器端中，可以使用这些参数进行搜索查询，并返回匹配结果。同时，还可以使用其他查询参数，例如 limit 和 offset 点进行分页操作。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇RESTful API的網路資料
{{< /ask_chatgpt >}}



1. RESTful API設計指南: https://blog.gtwang.org/programming/restful-api-design-guidelines/
2. RESTful API設計原則: https://ithelp.ithome.com.tw/articles/10245385
3. RESTful API介紹與實作: https://dotblogs.com.tw/explooosion/2017/05/04/011300
4. RESTful API介紹與實作 - 建立Web API: https://www.dotblogs.com.tw/kingon2008/2017/05/09/192047
5. RESTful API範例教學: https://cloud.tencent.com/developer/article/1335585   

