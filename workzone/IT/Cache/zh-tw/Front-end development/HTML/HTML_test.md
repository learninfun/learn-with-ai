

1. 如何在HTML中創建一個表單（form），並添加三個輸入框（input）用於收集名字、電子郵件和密碼？

答案：

```
<form>
   <label for="name">名字：</label>
   <input type="text" id="name" name="name"><br><br>
   <label for="email">電子郵件：</label>
   <input type="email" id="email" name="email"><br><br>
   <label for="password">密碼：</label>
   <input type="password" id="password" name="password"><br><br>
   <input type="submit" value="提交">
</form>
```

2. 怎樣在HTML中創建一個可滾動的div（div）？

答案：

```
<div style="overflow:scroll; height:200px;">
   <!--添加滾動內容的HTML代碼-->
</div>
```

3. 如何在HTML中添加一個音頻（audio）文件，並使其自動播放？

答案：

```
<audio autoplay>
   <source src="example.mp3" type="audio/mp3">
</audio>
```

4. 怎樣讓超連結（hyperlink）在新窗口（tab）中打開？

答案：

```
<a href="http://example.com" target="_blank">超連結文字</a>
```

5. 如何使用HTML在網頁中添加一個背景圖片（background image）？

答案：

```
<body style="background-image:url('background.jpg');">
   <!--網頁內容-->
</body>
```