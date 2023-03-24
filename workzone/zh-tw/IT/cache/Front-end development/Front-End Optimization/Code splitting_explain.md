

Code splitting 是一種優化Web應用程式性能的技術，它可以幫助開發者降低頁面載入時間，提升用戶體驗。簡單來說，Code splitting 就是將一個大型應用程式拆分成多個小部件，利用懶加載（Lazy loading）技術使每個小部件在需要時再載入，進而達到提升性能並減少頁面載入時間的效果。

例如，假設有一個網站，包含首頁、產品頁面和購物車頁面。如果每個頁面都使用相同的 JavaScript 代碼，就會導致頁面載入時間變長，影響用戶體驗。這時可以使用 Code splitting 技術，將相同的代碼拆分成多個文件，每個文件對應一個頁面。

舉例來說，使用React框架開發網站時，可以使用React.lazy()方法將各個頁面中使用的組件拆分成不同的文件，實現Code splitting。以下是一個示例：

```
import React, { lazy, Suspense } from 'react';

const HomePage = lazy(() => import('./HomePage'));
const ProductPage = lazy(() => import('./ProductPage'));
const CartPage = lazy(() => import('./CartPage'));

function App() {
  return (
    <div>
      <Suspense fallback={<div>Loading...</div>}>
        <HomePage />
        <ProductPage />
        <CartPage />
      </Suspense>
    </div>
  );
}

export default App;
```

在上面的例子中，使用React.lazy()方法將HomePage、ProductPage和CartPage這三個React組件拆分成獨立的文件，並在App組件中以動態載入的方式使用它們。當網頁初次載入時，只加載App組件，而HomePage、ProductPage和CartPage等其他組件會在需要時再進行加載，實現了Code splitting。