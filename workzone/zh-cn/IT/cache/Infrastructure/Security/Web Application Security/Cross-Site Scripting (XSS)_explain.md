

跨站脚本攻击（Cross-Site Scripting，简称 XSS）指的是攻击者通过在网站输入框中输入恶意脚本代码，使得网站的其他用户在访问该网站时也受到攻击。攻击者可以通过 XSS 攻击盗取用户的资讯，例如登入密码、Cookie 信息等，甚至可以绕过网站的防火墙，从而掌握网站系统的控制权。

举个例子，假如有一个网站有一个搜索框，当用户在该搜索框中输入一些内容后，网站会显示相关的搜索结果。如果攻击者在搜索框中输入一段 JavaScript 代码，并能够让其被其他用户访问，那么其他用户访问该网站时就会执行攻击者所提供的 JavaScript 代码，从而造成 XSS 攻击。

例如：

当用户在搜索框中输入以下代码：

```
<script>alert('攻击成功！')</script>
```

攻击者就可以成功地绕过网站的防护措施，并在用户访问该网站时弹出一个包含“攻击成功！”字样的提示框，此时攻击者就可以收获用户的关键资讯，造成严重的安全问题。