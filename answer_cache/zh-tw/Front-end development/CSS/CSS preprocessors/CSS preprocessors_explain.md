

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