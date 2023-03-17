## 習題預習
{{< ask_chatgpt >}}
給我5題Module bundlers的問題
{{< /ask_chatgpt >}}



1. 什麼是Module Bundler，它如何工作？
2. Module Bundling的目的是什麼？它解決了哪些問題？
3. Webpack和Parcel有什麼不同？它們各自有哪些優點和局限性？
4. 如何在Webpack中進行Code Splitting？它有什麼好處？
5. 如何在Parcel中使用自定義的Babel設置？   

## 說明知識
{{< ask_chatgpt >}}
說明Module bundlers並舉例
{{< /ask_chatgpt >}}



Module bundlers是一種將多個模組(Module)合併成單一檔案的工具，用於簡化前端開發及減少網頁請求量。

舉例來說，Webpack是一個常使用的模組打包工具，它支援多種模組系統包括CommonJS, AMD及ES6模組系統，可以將不同模組系統的模組打包成單一檔案供網頁使用。

另外，Rollup是一個以ES模組為主的模組打包工具，它將類似的模組進行Tree Shaking偵測，去除未使用到的變數及函數，最後編譯輸出最小化的JavaScript程式碼。這有助於提高性能及減少檔案大小。

其他常見的模組打包工具還包括Parcel及SystemJS等。透過使用這些工具，開發者可以更容易地管理複雜的前端程式碼，並減少網頁請求量，提高效能。   

## 彙總重點
{{< ask_chatgpt >}}
條列Module bundlers的重點
{{< /ask_chatgpt >}}



1. 為什麼需要Module bundlers：Module bundlers可以將多個檔案合併為一個文件，減少網頁請求的次數，提高性能。

2. 模組系統：Module bundlers支援各種前端框架與庫，同時可以維護依賴性，確保載入的檔案順序正確，減少錯誤。

3. 優化：Module bundlers還可以進行代碼壓縮、刪除未使用的代碼等優化操作，讓載入速度更快、文件更小。

4. 熱模組重載：Module bundlers支援熱模組重載，即只載入需要更新的模組代碼，提高開發效率。

5. 社區支援：Module bundlers有強大的社區支援，提供了許多開發者友好的功能和擴充功能。

6. 常用的Module bundlers：Webpack、Rollup、Parcel等是常用的Module bundlers。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Module bundlers的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 請問什麼是 Tree Shaking？

答：Tree Shaking 是一個在模組打包時，移除沒有被使用的程式碼的技術。

2. 請問什麼是 Code Splitting？

答：Code Splitting 是將一個大的 JavaScript 模組分成多個小的模組，並且只加載需要的部分，從而減少整體頁面載入時間。

3. 請問 Webpack 的 entry、output、loader 與 plugin 分別有什麼作用？

答：entry 為定義進入點，output 為定義輸出文件的位置與文件名稱，loader 為處理非 JavaScript 文件，將其轉換成 JavaScript 可以執行的形式，plugin 則是擴展 Webpack 的功能，並進行更進階的設定。

4. 請問 Webpack 中常用的優化方式有哪些？

答：常用的優化方式有：使用 Tree Shaking、Code Splitting、使用緩存、使用懶加載（Lazy Loading）等。

5. 請問 Rollup 與 Webpack 有哪些不同點？

答：Rollup 的優點是產生的 bundle 更小、更快，並且更適合用於打包開源庫等程式庫，而 Webpack 則更適合用於打包應用程式，並且支援更多的優化方式，如 Tree Shaking、Code Splitting 等。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Module bundlers的網路資料
{{< /ask_chatgpt >}}



1. "Webpack" from the official website of Webpack
https://webpack.js.org/

2. "An Introduction to JavaScript Module Bundling with Webpack" by Tania Rascia on Scotch.io
https://scotch.io/tutorials/javascript-module-bundling-an-introduction-with-webpack

3. "Beginner's guide to Webpack" by Sumit Kumar on Medium
https://medium.com/@souravdey/beginners-guide-to-webpack-b63d18bc7db#.emlt24sv4

4. "Getting Started with Browserify and Gulp" by Blake Embrey on SitePoint
https://www.sitepoint.com/getting-started-browserify-gulp/

5. "Understanding JavaScript Modules: Bundling & Transpiling" by T.J. VanToll on Telerik Developer Network
http://developer.telerik.com/content-types/tutorials/understanding-javascript-modules-bundling-transpiling/   

