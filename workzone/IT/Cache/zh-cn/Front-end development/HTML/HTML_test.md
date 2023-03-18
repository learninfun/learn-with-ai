

1. 如何在HTML中创建一个表单（form），并添加三个输入框（input）用于收集名字、电子邮件和密码？

答案：

```
<form>
   <label for="name">名字：</label>
   <input type="text" id="name" name="name"><br><br>
   <label for="email">电子邮件：</label>
   <input type="email" id="email" name="email"><br><br>
   <label for="password">密码：</label>
   <input type="password" id="password" name="password"><br><br>
   <input type="submit" value="提交">
</form>
```

2. 怎样在HTML中创建一个可滚动的div（div）？

答案：

```
<div style="overflow:scroll; height:200px;">
   <!--添加滚动内容的HTML代码-->
</div>
```

3. 如何在HTML中添加一个音频（audio）文件，并使其自动播放？

答案：

```
<audio autoplay>
   <source src="example.mp3" type="audio/mp3">
</audio>
```

4. 怎样让超连结（hyperlink）在新窗口（tab）中打开？

答案：

```
<a href="http://example.com" target="_blank">超连结文字</a>
```

5. 如何使用HTML在网页中添加一个背景图片（background image）？

答案：

```
<body style="background-image:url('background.jpg');">
   <!--网页内容-->
</body>
```