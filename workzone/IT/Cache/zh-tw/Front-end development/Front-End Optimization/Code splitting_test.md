

1. 創建一個代碼分割點，當用戶點擊按鈕時，加載一個包含一些複雜邏輯的模組。模組的名字應該是 `lazyModule.js`。

```javascript
// App.js

import React, { useState } from 'react';

function App() {
  const [showModule, setShowModule] = useState(false);

  const handleClick = () => {
    import('./lazyModule.js').then((module) => {
      setShowModule(true);
    });
  };

  return (
    <div>
      <button onClick={handleClick}>Load module</button>
      {showModule && <LazyModule />}
    </div>
  );
}

export default App;

// lazyModule.js

import React from 'react';

function LazyModule() {
  return <div>This module has some complex logic...</div>;
}

export default LazyModule;
```

2. 創建一個路由器，在每個路由之間動態切換代碼分割點。當用戶訪問 `/` 時，顯示一個顯示選擇器的頁面，當用戶選擇項目並提交時，切換到對應的代碼分割點，加載不同的組件。

```javascript
// Router.js

import React, { lazy, useState } from 'react';
import { BrowserRouter, Link, Route, Switch } from 'react-router-dom';

const CategoryA = lazy(() => import('./CategoryA'));
const CategoryB = lazy(() => import('./CategoryB'));

function Router() {
  const [selectedCategory, setSelectedCategory] = useState(null);

  const handleSelectCategory = (category) => {
    setSelectedCategory(category);
  };

  return (
    <BrowserRouter>
      <nav>
        <Link to="/">Home</Link>
        <Link to="/category-a">Category A</Link>
        <Link to="/category-b">Category B</Link>
      </nav>
      <Switch>
        <Route exact path="/">
          <div>
            <p>Select a category:</p>
            <button onClick={() => handleSelectCategory('a')}>Category A</button>
            <button onClick={() => handleSelectCategory('b')}>Category B</button>
          </div>
        </Route>
        <Route path="/category-a">
          {selectedCategory === 'a' && <CategoryA />}
        </Route>
        <Route path="/category-b">
          {selectedCategory === 'b' && <CategoryB />}
        </Route>
      </Switch>
    </BrowserRouter>
  );
}

export default Router;

// CategoryA.js

import React from 'react';

function CategoryA() {
  return <div>This is Category A.</div>;
}

export default CategoryA;

// CategoryB.js

import React from 'react';

function CategoryB() {
  return <div>This is Category B.</div>;
}

export default CategoryB;
```

3. 通過組合使用 `React.lazy` 和 `React.Suspense` 實現按需加載幾個組件。

```javascript
// App.js

import React, { lazy, Suspense } from 'react';

const ComponentA = lazy(() => import('./ComponentA'));
const ComponentB = lazy(() => import('./ComponentB'));
const ComponentC = lazy(() => import('./ComponentC'));

function App() {
  return (
    <div>
      <Suspense fallback={<div>Loading...</div>}>
        <ComponentA />
        <ComponentB />
        <ComponentC />
      </Suspense>
    </div>
  );
}

export default App;

// ComponentA.js

import React from 'react';

function ComponentA() {
  return <div>This is Component A.</div>;
}

export default ComponentA;

// ComponentB.js

import React from 'react';

function ComponentB() {
  return <div>This is Component B.</div>;
}

export default ComponentB;

// ComponentC.js

import React from 'react';

function ComponentC() {
  return <div>This is Component C.</div>;
}

export default ComponentC;
```

4. 通過按需加載不同的組件實現動態切換視圖。當用戶點擊不同的選項時，加載對應的組件，顯示在頁面上。

```javascript
// App.js

import React, { useState } from 'react';

function App() {
  const [view, setView] = useState(null);

  const handleSetView = (viewName) => {
    import(`./${viewName}`).then((module) => {
      setView(module.default);
    });
  };

  return (
    <div>
      <button onClick={() => handleSetView('ViewA')}>View A</button>
      <button onClick={() => handleSetView('ViewB')}>View B</button>
      {view && <view />}
    </div>
  );
}

export default App;

// ViewA.js

import React from 'react';

function ViewA() {
  return <div>This is View A.</div>;
}

export default ViewA;

// ViewB.js

import React from 'react';

function ViewB() {
  return <div>This is View B.</div>;
}

export default ViewB;
```

5. 實現按需加載第三方庫，如 Lodash。

```javascript
// App.js

import React, { useState } from 'react';

function App() {
  const [lodash, setLodash] = useState(false);

  const handleLoadLodash = () => {
    import('lodash').then((module) => {
      setLodash(module);
    });
  };

  return (
    <div>
      <button onClick={handleLoadLodash}>Load Lodash</button>
      {lodash && <div>{lodash.VERSION}</div>}
    </div>
  );
}

export default App;
```

答案僅供參考，並不是唯一正確的答案。