+++
title = "JSON Web Token"
weight = "8"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題JWT的中文問題
{{< /ask_chatgpt >}}

1. 什麼是JWT？
2. JWT的結構及其包含哪些基本元素？
3. JWT的優點和缺點分別是什麼？
4. 如何生成和驗證JWT？
5. JWT在應用程序中的使用場景及常見應用場景是什麼？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明JWT並舉例
{{< /ask_chatgpt >}}

JWT（JSON Web Token）是一種用於網路傳輸協議的驗證方式，通常被用於 API 驗證、Oauth2 授權等情境。JWT 使用 JSON 格式將驗證資訊加密在 token 中，在用戶每次訪問 API 時都需要將 token 携帶在請求中，以便服務端驗證用戶身份。JWT 主要由三部分組成，分別是 Header、Payload 和 Signature。

Header 包含兩種類型的資訊，分別是 token 使用的加密算法（如 HS256、HS512、RS256）和 token 類型（Bearer Token）。

Payload 包含了一些用戶和驗證相關的資訊，例如用戶的 ID、用戶的權限、token 的過期時間等。

Signature 則是使用了加密算法以 Header、Payload 和一個密鑰組合成鉴权标签。如果 token 可以被正確驗證，那麼可以確定 token 來自於合法身份。

以下是 JWT 的一個簡單範例：

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

上面的範例可以被解析成下面的 JSON 格式：

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022
}
```

其中 alg 表示 JWT 使用的加密算法，typ 表示 token 的類型。sub 表示 JWT 所代表的使用者 ID，name 表示 JWT 所代表的使用者姓名，iat 表示 JWT 發布的時間。這篇範例中，如果 JWT 的 Signature 可以被成功驗證，那麼可以確認使用者 ID 為 1234567890、使用者姓名為 John Doe，並且 token 發布時間為 1516239022。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列JWT的重點
{{< /ask_chatgpt >}}

1. JWT是一種JSON格式的簽名令牌，用於在客戶端和服務器之間進行身份驗證和授權。

2. JWT由三部分組成：頭部、資料載荷和簽名，每個部分都使用BASE64編碼。

3. 頭部包含令牌的類型（即JWT）、加密算法（如HMAC SHA256）和其他元數據。

4. 資料載荷包含附加的身份驗證信息，例如用戶ID、角色、有效期限等。

5. 簽名用於驗證令牌的完整性和真實性，通常使用密鑰進行加密。

6. JWT的主要優勢是可攜帶性和安全性，因為令牌是自我描述的並且包含所有所需的信息。

7. 使用JWT進行身份驗證時，通常由客戶端發送令牌，服務器則驗證令牌並批准或拒絕請求。

8. JWT也支持持久性，因為令牌是自包含的，不需要每次請求時都與服務器進行交互。

9. 如需保護敏感信息，建議將資料載荷加密，以增強安全性。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題JWT的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}

1. JWT 的全名是什麼？它是什麼儲存格式？

2. JWT 有哪些主要的元素？

3. 如何驗證 JWT 是否合法？

4. 在什麼情況下，JWT 會被視為不合法？

5. JWT 的優點是什麼？

答案：
1. JSON Web Token，儲存格式為 Base64 編碼的 JSON 格式。
2. 有三個主要元素：Header、Payload、Signature。
3. 檢查 JWT 的 Signature 是否正確，通常是透過 Secret Key。
4. 如果 JWT 的 Signature 不正確、Expired Time 過期等。
5. JWT 可以包含足夠的用戶信息，無需反覆對數據庫請求驗證。可以跨平臺使用，輕鬆獲得授權。加密和解密操作在客戶端完成，減輕了服務器負擔。結構緊湊，容易傳輸和存儲。   

