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