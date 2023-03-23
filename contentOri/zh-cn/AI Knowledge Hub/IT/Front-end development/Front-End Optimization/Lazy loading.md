+++
title = "Lazy loading 懒加载"
weight = "5"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Lazy loading的问题
{{< /ask_chatgpt >}}



1. 什么是Lazy loading？ 
2. 哪些网页元素应该使用Lazy loading？ 
3. Lazy loading对网站速度和性能有什么影响？ 
4. 如何在网站中实现Lazy loading？ 
5. 什么是Progressive lazy loading？它如何工作？   

## 说明知识
{{< ask_chatgpt >}}
说明Lazy loading并举例
{{< /ask_chatgpt >}}



Lazy loading是一种网页加载优化技术，它利用延迟载入的方式，先载入最需要的部分，而让其他内容在需要时再加载，减少首次页面载入时间，提高网站速度和用户体验。

举例来说，当加载一个较长的网页时，图片是一个常见的资源，如果一次性载入所有的图片，会导致首次载入时间过长，但透过Lazy loading技术，可以先载入首屏内容，而等到用户滚动到其他部分或触发某些事件时再加载其他图片，这样可以加速网站页面载入速度。

另一个例子是当网页有许多JS或CSS档案需要载入时，也可以透过Lazy loading技术，只先载入最基本的CSS样式或JS代码，等到用户进入特定页面或功能，再去载入需要的JS或CSS档案。

综而言之，Lazy loading技术可以在不影响用户体验的情况下，提高网站速度和性能，特别对于移动端，更是能够减少用户等待时间，提高使用体验。   

## 汇总重点
{{< ask_chatgpt >}}
条列Lazy loading的重点
{{< /ask_chatgpt >}}



以下是Lazy loading的重點：

1. 延遲載入：Lazy loading是一種技術，它讓網頁上的圖片、影片、音訊及其他資源等，在使用者需要時才載入。

2. 優化網頁速度：Lazy loading能夠減少網頁的載入時間，並且讓使用者在網頁開始載入後，就可以立即開始瀏覽網頁的內容。

3. 標籤及選擇器：開發人員可以使用HTML、CSS及JavaScript等標籤及選擇器完成Lazy loading，並且讓網頁的資源按需而讀取。

4. 支援多種瀏覽器：Lazy loading技術能夠支援大部分的瀏覽器，並且能夠處理所有種類的資源。

5. 改善用戶體驗：Lazy loading能夠改善使用者體驗，提高網站的互動性及使用率，讓使用者能夠更快速、更方便地瀏覽網頁的內容。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Lazy loading的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 前端頁面需載入大量資料，如何實現Lazy Loading？

答案：使用Intersection Observer API監視頁面上的元素，當元素進入視窗範圍內時，才載入該元素所需的資料。可以使用像是react-lazyload和ng-lazyload-image這樣的套件來實現。

2. 如何避免在第一次訪問網站時，將所有頁面資料都加載出來？

答案：使用Lazy Loading機制。將頁面核心資料與功能做好分類，只在用戶訪問時加載當前頁面所需的資料，其他部分透過按需載入實現最佳的性能優化。

3. 如何實現按需載入圖片？

答案：使用Lazy Loading機制。通常有兩種方式實現：以空的圖片檔代替原始圖檔，再透過自訂屬性儲存圖片路徑等資訊，再透過JavaScript進行替換；透過intersection Observer API監視圖片是否進入視窗範圍，並將其載入。

4. 如何確定設計的Lazy Loading是有效果的？

答案：使用Chrome瀏覽器的Dev Tool，啟用Network面板，模擬網路速度，使用Performance面板觀察頁面載入時間、如何加快頁面載入等表現，更可以使用Lighthouse或WebPageTest進行真實測試。

5. 如何實現Lazy Loading時，保證頁面導航的可用性？

答案：加載較大的ressource時，要確保網頁的載入順序是正確的。優先加載頁面的核心資料和功能，再進行圖片和其他資源的加載和渲染，避免出現網頁空白等問題。可以將頁面下滑一半的位置加載接下來的內容，並顯示加載狀態提示，讓用戶了解當前頁面正在進行資源的加載。   

