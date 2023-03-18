+++
title = "模组打包器"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Module bundlers的问题
{{< /ask_chatgpt >}}



1. 什麼是Module Bundler，它如何工作？
2. Module Bundling的目的是什麼？它解決了哪些問題？
3. Webpack和Parcel有什麼不同？它們各自有哪些優點和局限性？
4. 如何在Webpack中進行Code Splitting？它有什麼好處？
5. 如何在Parcel中使用自定義的Babel設置？   

## 说明知识
{{< ask_chatgpt >}}
说明Module bundlers并举例
{{< /ask_chatgpt >}}



Module bundlers是一種將多個模組(Module)合併成單一檔案的工具，用於簡化前端開發及減少網頁請求量。

舉例來說，Webpack是一個常使用的模組打包工具，它支援多種模組系統包括CommonJS, AMD及ES6模組系統，可以將不同模組系統的模組打包成單一檔案供網頁使用。

另外，Rollup是一個以ES模組為主的模組打包工具，它將類似的模組進行Tree Shaking偵測，去除未使用到的變數及函數，最後編譯輸出最小化的JavaScript程式碼。這有助於提高性能及減少檔案大小。

其他常見的模組打包工具還包括Parcel及SystemJS等。透過使用這些工具，開發者可以更容易地管理複雜的前端程式碼，並減少網頁請求量，提高效能。   

## 汇总重点
{{< ask_chatgpt >}}
条列Module bundlers的重点
{{< /ask_chatgpt >}}



1. 为什么需要Module bundlers：Module bundlers可以将多个档案合并为一个文件，减少网页请求的次数，提高性能。

2. 模组系统：Module bundlers支援各种前端框架与库，同时可以维护依赖性，确保载入的档案顺序正确，减少错误。

3. 优化：Module bundlers还可以进行代码压缩、删除未使用的代码等优化操作，让载入速度更快、文件更小。

4. 热模组重载：Module bundlers支援热模组重载，即只载入需要更新的模组代码，提高开发效率。

5. 社区支援：Module bundlers有强大的社区支援，提供了许多开发者友好的功能和扩充功能。

6. 常用的Module bundlers：Webpack、Rollup、Parcel等是常用的Module bundlers。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Module bundlers的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 请问什么是 Tree Shaking？

答：Tree Shaking 是一个在模组打包时，移除没有被使用的程式码的技术。

2. 请问什么是 Code Splitting？

答：Code Splitting 是将一个大的 JavaScript 模组分成多个小的模组，并且只加载需要的部分，从而减少整体页面载入时间。

3. 请问 Webpack 的 entry、output、loader 与 plugin 分别有什么作用？

答：entry 为定义进入点，output 为定义输出文件的位置与文件名称，loader 为处理非 JavaScript 文件，将其转换成 JavaScript 可以执行的形式，plugin 则是扩展 Webpack 的功能，并进行更进阶的设定。

4. 请问 Webpack 中常用的优化方式有哪些？

答：常用的优化方式有：使用 Tree Shaking、Code Splitting、使用缓存、使用懒加载（Lazy Loading）等。

5. 请问 Rollup 与 Webpack 有哪些不同点？

答：Rollup 的优点是产生的 bundle 更小、更快，并且更适合用于打包开源库等程式库，而 Webpack 则更适合用于打包应用程式，并且支援更多的优化方式，如 Tree Shaking、Code Splitting 等。   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Module bundlers的网络数据
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

