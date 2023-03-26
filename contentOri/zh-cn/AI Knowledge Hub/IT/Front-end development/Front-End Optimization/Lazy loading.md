+++
title = "Lazy loading 懒加载"
weight = "5"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Lazy loading的中文问题
{{< /ask_chatgpt >}}



1. 什么是Lazy loading？ 
2. 哪些网页元素应该使用Lazy loading？ 
3. Lazy loading对网站速度和性能有什么影响？ 
4. 如何在网站中实现Lazy loading？ 
5. 什么是Progressive lazy loading？它如何工作？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明Lazy loading并举例
{{< /ask_chatgpt >}}



Lazy loading是一种网页加载优化技术，它利用延迟载入的方式，先载入最需要的部分，而让其他内容在需要时再加载，减少首次页面载入时间，提高网站速度和用户体验。

举例来说，当加载一个较长的网页时，图片是一个常见的资源，如果一次性载入所有的图片，会导致首次载入时间过长，但透过Lazy loading技术，可以先载入首屏内容，而等到用户滚动到其他部分或触发某些事件时再加载其他图片，这样可以加速网站页面载入速度。

另一个例子是当网页有许多JS或CSS档案需要载入时，也可以透过Lazy loading技术，只先载入最基本的CSS样式或JS代码，等到用户进入特定页面或功能，再去载入需要的JS或CSS档案。

综而言之，Lazy loading技术可以在不影响用户体验的情况下，提高网站速度和性能，特别对于移动端，更是能够减少用户等待时间，提高使用体验。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列Lazy loading的重点
{{< /ask_chatgpt >}}



以下是Lazy loading的重点：

1. 延迟载入：Lazy loading是一种技术，它让网页上的图片、影片、音讯及其他资源等，在使用者需要时才载入。

2. 优化网页速度：Lazy loading能够减少网页的载入时间，并且让使用者在网页开始载入后，就可以立即开始浏览网页的内容。

3. 标签及选择器：开发人员可以使用HTML、CSS及JavaScript等标签及选择器完成Lazy loading，并且让网页的资源按需而读取。

4. 支援多种浏览器：Lazy loading技术能够支援大部分的浏览器，并且能够处理所有种类的资源。

5. 改善用户体验：Lazy loading能够改善使用者体验，提高网站的互动性及使用率，让使用者能够更快速、更方便地浏览网页的内容。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题Lazy loading的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 前端页面需载入大量资料，如何实现Lazy Loading？

答案：使用Intersection Observer API监视页面上的元素，当元素进入视窗范围内时，才载入该元素所需的资料。可以使用像是react-lazyload和ng-lazyload-image这样的套件来实现。

2. 如何避免在第一次访问网站时，将所有页面资料都加载出来？

答案：使用Lazy Loading机制。将页面核心资料与功能做好分类，只在用户访问时加载当前页面所需的资料，其他部分透过按需载入实现最佳的性能优化。

3. 如何实现按需载入图片？

答案：使用Lazy Loading机制。通常有两种方式实现：以空的图片档代替原始图档，再透过自订属性储存图片路径等资讯，再透过JavaScript进行替换；透过intersection Observer API监视图片是否进入视窗范围，并将其载入。

4. 如何确定设计的Lazy Loading是有效果的？

答案：使用Chrome浏览器的Dev Tool，启用Network面板，模拟网路速度，使用Performance面板观察页面载入时间、如何加快页面载入等表现，更可以使用Lighthouse或WebPageTest进行真实测试。

5. 如何实现Lazy Loading时，保证页面导航的可用性？

答案：加载较大的ressource时，要确保网页的载入顺序是正确的。优先加载页面的核心资料和功能，再进行图片和其他资源的加载和渲染，避免出现网页空白等问题。可以将页面下滑一半的位置加载接下来的内容，并显示加载状态提示，让用户了解当前页面正在进行资源的加载。   

