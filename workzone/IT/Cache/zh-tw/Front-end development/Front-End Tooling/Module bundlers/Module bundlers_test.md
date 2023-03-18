

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