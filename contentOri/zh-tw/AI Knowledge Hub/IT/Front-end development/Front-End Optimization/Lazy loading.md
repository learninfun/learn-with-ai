+++
title = "Lazy loading 懶加載"
weight = "5"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Lazy loading的中文問題
{{< /ask_chatgpt >}}



1. 什麼是Lazy loading？ 
2. 哪些網頁元素應該使用Lazy loading？ 
3. Lazy loading對網站速度和性能有什麼影響？ 
4. 如何在網站中實現Lazy loading？ 
5. 什麼是Progressive lazy loading？它如何工作？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Lazy loading並舉例
{{< /ask_chatgpt >}}



Lazy loading是一種網頁加載優化技術，它利用延遲載入的方式，先載入最需要的部分，而讓其他內容在需要時再加載，減少首次頁面載入時間，提高網站速度和用戶體驗。

舉例來說，當加載一個較長的網頁時，圖片是一個常見的資源，如果一次性載入所有的圖片，會導致首次載入時間過長，但透過Lazy loading技術，可以先載入首屏內容，而等到用戶滾動到其他部分或觸發某些事件時再加載其他圖片，這樣可以加速網站頁面載入速度。

另一個例子是當網頁有許多JS或CSS檔案需要載入時，也可以透過Lazy loading技術，只先載入最基本的CSS樣式或JS代碼，等到用戶進入特定頁面或功能，再去載入需要的JS或CSS檔案。

綜而言之，Lazy loading技術可以在不影響用戶體驗的情況下，提高網站速度和性能，特別對於移動端，更是能夠減少用戶等待時間，提高使用體驗。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Lazy loading的重點
{{< /ask_chatgpt >}}



以下是Lazy loading的重點：

1. 延遲載入：Lazy loading是一種技術，它讓網頁上的圖片、影片、音訊及其他資源等，在使用者需要時才載入。

2. 優化網頁速度：Lazy loading能夠減少網頁的載入時間，並且讓使用者在網頁開始載入後，就可以立即開始瀏覽網頁的內容。

3. 標籤及選擇器：開發人員可以使用HTML、CSS及JavaScript等標籤及選擇器完成Lazy loading，並且讓網頁的資源按需而讀取。

4. 支援多種瀏覽器：Lazy loading技術能夠支援大部分的瀏覽器，並且能夠處理所有種類的資源。

5. 改善用戶體驗：Lazy loading能夠改善使用者體驗，提高網站的互動性及使用率，讓使用者能夠更快速、更方便地瀏覽網頁的內容。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Lazy loading的中等難度問題，並在後面列出答案
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

