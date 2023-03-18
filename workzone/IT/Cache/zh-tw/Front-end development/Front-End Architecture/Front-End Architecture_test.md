

1. 什麼是MVC架構？與MVP或MVVM有何區別？

答：MVC是Model-View-Controller的縮寫，是一種軟體架構模式，將系統分為資料模型（Model）、顯示視圖（View）、和控制器（Controller）三部分。MVP和MVVM則是基於MVC的演化版本，MVVM將控制器改為了ViewModel，MVP則將View和Controller職責進一步拆分。

2. 什麼是單一責任原則（Single Responsibility Principle）？如何在前端應用中運用？

答：單一責任原則是指一個物件或函式只應該擁有一個引起它變化的原因。在前端應用中，可以運用這一原則來拆分Component，每個Component應該只負責一個功能模塊。

3. 什麼是狀態管理庫（State Management）？如何選擇最適合項目的狀態管理庫？

答：狀態管理庫是指管理應用程序狀態的庫，常見的狀態管理庫有Redux、MobX、Vuex等。最適合項目的狀態管理庫取決於項目規模、複雜度、團隊開發經驗等因素。

4. 什麼是Webpack？如何運用Webpack進行模塊打包？

答：Webpack是一個模塊打包器，可將各種類型的檔案（如JavaScript、CSS、圖片等）打包成一個或多個Bundle。運用Webpack打包模塊，首先需要在配置檔設置入口檔、輸出檔路徑和Loader和Plugins等相關配置，然後運行Webpack命令即可進行打包。

5. 什麼是跨域請求（Cross-Origin Request）？如何解決跨域問題？

答：跨域請求是指在網頁中，使用AJAX向不同域名、不同端口或不同協議的服務器發送請求。跨域請求常會受到瀏覽器的同源策略限制，通常可以通過CORS、JSONP、代理服務器等方式來解決跨域問題。