+++
title = "前端性能預算"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Front-End Performance Budgets的問題
{{< /ask_chatgpt >}}



1. 什麼是Front-End Performance Budgets? 
2. 為什麼需要Front-End Performance Budgets? 
3. 使用Front-End Performance Budgets可以發現哪些問題？ 
4. 如何設置Front-End Performance Budgets？ 
5. 怎樣量測Front-End Performance Budgets？   

## 說明知識
{{< ask_chatgpt >}}
說明Front-End Performance Budgets並舉例
{{< /ask_chatgpt >}}



前端響應式設計是通過將設計網頁從不同的設備和屏幕大小進行優化以響應用戶需求，成為Web設計的標準。而Front-End Performance Budgets（前端性能預算）就是通過對前端性能進行調整，來提高網站響應速度和用戶體驗。所以，一個前端性能預算是指為了避免出現不良後果而建立的預算。它列出了為網站應該採用一些性能範圍的具體措施。前端性能預算可以包含許多方面。例如:

1. 頁面加載時間
在加載時間方面的性能預算只是一個數字。它可以是3秒或5秒等。如果你的網站在規定的時間內不能夠加載，那麼你需要修改你的性能預算。

2. 壓縮圖像
在性能預算中，必須明確地考慮為網站壓縮圖片的大小和壓縮格式。

3. 最大化網絡請求
在性能預算的方面，我們必須明確地平衡每個資源的請求大小。每個資源應該有一個確定的大小或者是請求次數的限制。在這種情況下，我們可以限制每個頁面的請求數量，並在內置本地緩存後處理數據。

舉個例子來說，假設你的公司的前端性能預算限制為500KB。這就意味著你的網站的所有靜態內容，包括CSS、JavaScript、圖片和其他資源都不能超過這個限制。假設你的CSS文件大小為80KB，JavaScript文件大小為150KB，圖片和其他資源總計大小為270KB。那麼，這意味著你還有剩餘的大小176KB在你的性能預算之內，你可以用於更多的圖片和資源，或者用於更多的JavaScript交互和動畫。通過這種方式，可以確保你網站的速度和性能表現始終如一，提高了用戶的網站體驗。   

## 彙總重點
{{< ask_chatgpt >}}
條列Front-End Performance Budgets的重點
{{< /ask_chatgpt >}}



1.約束預算：Front-end Performance Budgets是一種設計和開發網站的方法，它強制開發人員在特定的性能操作上設置截止日期。

2. 固定大小和響應性：通過設置網站元素的固定大小和響應性，可以確保網站的頁面元素在各種設備上都能夠完美地呈現。

3. 網站性能分析：通過使用工具進行網站性能測試和分析，可以確定網站上哪些元素需要進行優化。

4. 應用程式性能優化：通過最佳化網站的HTML、CSS和JavaScript代碼，可以大大提高網站的性能，同時減少頁面載入時間。

5. 及時更新：Front-end Performance Budgets需要定期更新，以反映網站上更常見的流量模式、瀏覽器和設備。

6. 最佳實踐：Front-end Performance Budgets設置和執行需要遵循最佳實踐，以最大程度地提高網站的性能和流暢性。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Front-End Performance Budgets的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 如果要求每個頁面的加載時間必須在3秒內完成，請問您會如何限制JavaScript文件的大小？
答案：限制JavaScript文件的大小不超過200KB。

2. 如果您希望每個頁面的可見部分（above-the-fold content）的加載時間應該在1.5秒內完成，請問應該如何優化圖像的加載時間？
答案：針對可見部分和其他區域，分別優化圖像的加載時間。採用WebP格式的圖像格式，適當地縮小圖像大小，減少需要下載的像素數量。

3. 如果您目標是每個頁面的DOM大小不能超過1MB，請問應該如何減少CSS文件的大小？
答案：減少CSS文件的大小是通過優化CSS樣式表的過程實現的，其中包括：移除不必要的樣式代碼，包括不使用的樣式，以及唯一樣式的副本；採用可壓縮的格式，如GZIP或BROTLI，優化樣式表的結構。

4. 如果您想要確保網站的平均加載時間不超過5秒，您應該如何進行JavaScript代碼優化？
答案：JavaScript代碼優化可以通過移除不必要的代碼，使用簡單的算法和查詢，並壓縮代碼。此外，避免使用較慢的JavaScript函數和複雜的查詢操作。

5. 如果您想要減少網頁加載時間並保證高效的性能，您應該如何使用Web Worker？
答案：Web Worker可以在後臺線程中執行複雜的JavaScript代碼，從而減少主線程的工作量。這樣可以提高網頁的反應速度，同時確保高效的性能。Web Worker通常在處理大量數據，計算密集型操作和耗時操作方面非常有用。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Front-End Performance Budgets的網路資料
{{< /ask_chatgpt >}}



1. Front-End Performance Budgets: A Complete Guide | SitePoint
- https://www.sitepoint.com/front-end-performance-budgets-complete-guide/

2. Crafting Your Front-End Performance Budget – Smashing Magazine
- https://www.smashingmagazine.com/2015/06/crafting-a-future-proof-responsive-design-strategy/

3. Frontend Performance Budgets – Performance Calendar
- https://calendar.perfplanet.com/2014/frontend-performance-budgets/

4. Implementing a Front-End Performance Budget with Google Analytics
- https://mattzeunert.com/2019/05/24/implementing-a-front-end-performance-budget-with-google-analytics.html

5. Measuring Front-end Performance with SpeedCurve
- https://speedcurve.com/blog/measuring-frontend-performance-with-speedcurve/   
