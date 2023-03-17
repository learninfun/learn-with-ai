## 習題預習
{{< ask_chatgpt >}}
給我5題CSS preprocessors的問題
{{< /ask_chatgpt >}}



1. 什麼是 CSS Preprocessor？它是用來幫助你編寫 CSS 的工具，有哪些好處？
2. 你瞭解哪些 CSS Preprocessor？它們之間有什麼異同？
3. 請舉個例子，你是如何利用 CSS Preprocessor 來實現程式化的 CSS 設計？
4. 在 CSS Preprocessor 中，有哪些常用的功能和語法？
5. CSS Preprocessor 如何處理和編譯成瀏覽器可以讀取的 CSS？   

## 說明知識
{{< ask_chatgpt >}}
說明CSS preprocessors並舉例
{{< /ask_chatgpt >}}



CSS預處理器（CSS preprocessors）是一種可以擴展CSS語言，並且讓CSS更加易於維護和管理複雜的風格表的軟體工具。它們包括像SASS、LESS和Stylus等工具。

舉例來說，SASS可以提供一些類似程式設計的元素，例如變量、條件語句、計算和函數等，這些都可以讓開發人員更輕鬆地編寫CSS代碼。

一個常見的SASS代碼片段如下所示：

```
$primary-color: #4D4D4D;

.navigation {
  background-color: $primary-color;
  font-size: 1.2em;
  a {
    color: white;
    &:hover {
      text-decoration: none;
    }
  }
}
```

這個代碼示例中，$primary-color變量被創建，並且可以在CSS代碼中被使用。此外，它包括了程式設計概念，如嵌套和小括號，以提高代碼的可讀性和編寫速度。   

## 彙總重點
{{< ask_chatgpt >}}
條列CSS preprocessors的重點
{{< /ask_chatgpt >}}

1. 增加可讀性和維護性：CSS預處理器通過較簡單的語法和拓展引入，提高了CSS的可讀性和維護性。
2. 變量和常量：CSS預處理器允許定義變量和常量，方便在代碼中重複使用。這增加了代碼的重用性，並且簡化了更改全局屬性的任務。
3. 嵌套和層級：CSS預處理器允許嵌套選擇器和屬性，使CSS代碼更具層次感和可讀性。這樣可以更容易地設置特定元素的樣式，且與HTML標記的關係更明顯。
4. 擴展：CSS預處理器提供了擴展機制，可通過引入擴展，將各種樣式規則打包到一起。這增加了代碼的普遍性並且簡化了開發任務。
5. 函數和Mixin：CSS預處理器具有函數和Mixin，這些函數和Mixin可在代碼中定義，將代碼結構化，避免代碼的重複。
6. 工具和框架：CSS預處理器具有不同的工具和框架，通過這些工具和框架，開發人員可以更快地構建專業的CSS樣式表。
7. 編譯：CSS預處理器需要編譯才能將預處理器代碼轉換為純CSS代碼，有利於縮小和優化代碼。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題CSS preprocessors的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 在LESS中，如何重複使用定義的變數？
答案：使用@variable語法。

2. 在Sass中，如何創建混合物（mixins）？
答案：使用@mixins語法。

3. 在SCSS中，如何將屬性綁定到嵌套選擇器的父級？
答案：使用&符號。

4. 在Stylus中，如何定義包含多個屬性的函數？
答案：使用“{ }”定義一個對象，然後將其作為函數的返回值傳遞。

5. 在PostCSS中，如何實現自動添加前綴的功能？
答案：使用autoprefixer插件。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇CSS preprocessors的網路資料
{{< /ask_chatgpt >}}



1. "What are CSS preprocessors and why should I use them?" by SitePoint: https://www.sitepoint.com/css-preprocessors-and-why-to-use-them/

2. "Sass vs. LESS vs. Stylus: Preprocessor Shootout" by Smashing Magazine: https://www.smashingmagazine.com/2015/07/death-to-icon-fonts/

3. "Getting Started with Sass, Less, and Stylus" by TutsPlus: https://code.tutsplus.com/tutorials/getting-started-with-css-preprocessors-sass-less-and-stylus--cms-20157

4. "Why You Should Use a CSS Preprocessor (And You Really Should)" by CSS-Tricks: https://css-tricks.com/why-you-should-use-a-css-preprocessor-and-you-really-should/

5. "CSS Preprocessors: A Comparison of Sass, Less and Stylus" by Scotch.io: https://scotch.io/tutorials/css-preprocessors-a-comparison-of-sass-less-and-stylus   

