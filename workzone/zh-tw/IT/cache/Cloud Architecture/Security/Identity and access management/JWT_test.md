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